package resources

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/deletes"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/inserts"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/updates"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type application struct {
	dbApp               databases.Application
	builder             resources.Builder
	storageBuilder      storages.StorageBuilder
	switchersBuilder    switchers.Builder
	switcherBuilder     switchers.SwitcherBuilder
	singleBuiler        singles.Builder
	delimiterBuilder    delimiters.DelimiterBuilder
	transactionsBuilder transactions.Builder
	transactionBuilder  transactions.TransactionBuilder
	voteAdapter         signers.VoteAdapter
	hashAdapter         hash.Adapter
}

func createApplication(
	dbApp databases.Application,
	builder resources.Builder,
	storageBuilder storages.StorageBuilder,
	switchersBuilder switchers.Builder,
	switcherBuilder switchers.SwitcherBuilder,
	singleBuiler singles.Builder,
	delimiterBuilder delimiters.DelimiterBuilder,
	transactionsBuilder transactions.Builder,
	transactionBuilder transactions.TransactionBuilder,
	voteAdapter signers.VoteAdapter,
	hashAdapter hash.Adapter,
) Application {
	out := application{
		dbApp:               dbApp,
		builder:             builder,
		storageBuilder:      storageBuilder,
		switchersBuilder:    switchersBuilder,
		switcherBuilder:     switcherBuilder,
		singleBuiler:        singleBuiler,
		delimiterBuilder:    delimiterBuilder,
		transactionsBuilder: transactionsBuilder,
		transactionBuilder:  transactionBuilder,
		voteAdapter:         voteAdapter,
		hashAdapter:         hashAdapter,
	}

	return &out
}

// Insert inserts a resource
func (app *application) Insert(input resources.Resource, insert inserts.Insert) (resources.Resource, error) {
	switcher, err := app.switcherBuilder.Create().
		WithInsert(insert).
		Now()

	if err != nil {
		return nil, err
	}

	return app.loadSwitcherInResource(input, switcher)
}

// Load loads a resource
func (app *application) Load(input resources.Resource, name string) (resources.Resource, error) {
	storage, err := input.All().FetchByName(name)
	if err != nil {
		return nil, err
	}

	delimiter := storage.Delimiter()
	retBytes, err := app.dbApp.Read(delimiter)
	if err != nil {
		return nil, err
	}

	single, err := app.singleBuiler.Create().
		WithStorage(storage).
		WithBytes(retBytes).
		Now()

	if err != nil {
		return nil, err
	}

	switcher, err := app.switcherBuilder.Create().
		WithOriginal(single).
		Now()

	if err != nil {
		return nil, err
	}

	return app.loadSwitcherInResource(input, switcher)
}

// Select selects a loaded resource
func (app *application) Select(input resources.Resource, name string) (resources.Resource, error) {
	if !input.HasLoaded() {
		return nil, errors.New(noLoadedResourceErr)
	}

	loaded := input.Loaded()
	current, err := loaded.FetchByName(name)
	if err != nil {
		return nil, err
	}

	all := input.All()
	return app.builder.Create().
		WithAll(all).
		WithCurrent(current).
		WithLoaded(loaded).
		Now()
}

// Delete deletes the selected resource
func (app *application) Delete(input resources.Resource, delete deletes.Delete) (resources.Resource, error) {
	switcher, err := app.switcherBuilder.Create().
		WithDelete(delete).
		Now()

	if err != nil {
		return nil, err
	}

	return app.loadSwitcherInResource(input, switcher)
}

// Retrieve retrieves the selected resource
func (app *application) Retrieve(input resources.Resource) (singles.Single, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noSelectedResourceErr)
	}

	return input.Current().Current(), nil
}

// Update updates the blacklist and whitelist of our resource
func (app *application) Update(input resources.Resource, update updates.Update) (resources.Resource, error) {
	switcher, err := app.switcherBuilder.Create().WithUpdate(update).Now()
	if err != nil {
		return nil, err
	}

	return app.updateStorageInResource(input, switcher)
}

// Commit commits the resource and returns the transactions
func (app *application) Commit(input resources.Resource) (transactions.Transactions, error) {
	if !input.HasLoaded() {
		return nil, nil // nothing to commit
	}

	trxList := []transactions.Transaction{}
	singlesList := []singles.Single{}
	loaded := input.Loaded().List()
	for _, oneLoaded := range loaded {
		trxBuilder := app.transactionBuilder.Create()
		if oneLoaded.HasInsert() {
			insert := oneLoaded.Insert()
			retSingle, err := app.insert(input, insert)
			if err != nil {
				return nil, err
			}

			singlesList = append(singlesList, retSingle)
			trxBuilder.WithInsert(insert)
		}

		if oneLoaded.HasUpdate() {
			update := oneLoaded.Update()
			retSingle, err := app.update(input, update)
			if err != nil {
				return nil, err
			}

			singlesList = append(singlesList, retSingle)
			trxBuilder.WithUpdate(update)
		}

		if oneLoaded.HasDelete() {
			delete := oneLoaded.Delete()
			retSingle, err := app.delete(input, delete)
			if err != nil {
				return nil, err
			}

			singlesList = append(singlesList, retSingle)
			trxBuilder.WithDelete(delete)
		}

		trx, err := trxBuilder.Now()
		if err != nil {
			return nil, err
		}

		trxList = append(trxList, trx)
	}

	nextIndex := input.All().NextIndex()
	err := app.write(nextIndex, singlesList)
	if err != nil {
		return nil, err
	}

	return app.transactionsBuilder.Create().
		WithList(trxList).
		Now()
}

func (app *application) write(startAtIndex uint64, singles []singles.Single) error {
	cpyFromIndex := startAtIndex
	for _, oneResource := range singles {
		storage := oneResource.Storage()
		if storage.IsDeleted() {
			continue
		}

		delimiter := storage.Delimiter()
		bytes := oneResource.Bytes()
		index := storage.Delimiter().Index()
		err := app.dbApp.CopyBeforeThenWrite(cpyFromIndex, index, bytes)
		if err != nil {
			return nil
		}

		cpyFromIndex = delimiter.Index() + delimiter.Length()
	}

	return nil
}

// Transact execute transactions
func (app *application) Transact(input resources.Resource, trx transactions.Transactions) error {
	lastRet := input
	list := trx.List()
	for _, oneTrx := range list {
		if oneTrx.IsInsert() {
			insert := oneTrx.Insert()
			retRes, err := app.Insert(lastRet, insert)
			if err != nil {
				return err
			}

			lastRet = retRes
			continue
		}

		if oneTrx.IsUpdate() {
			update := oneTrx.Update()
			retRes, err := app.Update(lastRet, update)
			if err != nil {
				return err
			}

			lastRet = retRes
			continue
		}

		delete := oneTrx.Delete()
		retRes, err := app.Delete(lastRet, delete)
		if err != nil {
			return err
		}

		lastRet = retRes
	}

	return nil
}

func (app *application) insert(input resources.Resource, insert inserts.Insert) (singles.Single, error) {
	data := insert.Bytes()
	nextIndex := input.All().NextIndex()
	length := uint64(len(data))
	delimiter, err := app.delimiterBuilder.Create().
		WithIndex(nextIndex).
		WithLength(length).
		Now()

	if err != nil {
		return nil, err
	}

	storageBuilder := app.storageBuilder.Create().
		WithDelimiter(delimiter)

	if insert.HasBlacklist() {
		storageBuilder.WithBlacklist(insert.Blacklist())
	}

	if insert.HasWhitelist() {
		storageBuilder.WithBlacklist(insert.Whitelist())
	}

	storage, err := storageBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.singleBuiler.Create().
		WithBytes(data).
		WithStorage(storage).
		Now()
}

func (app *application) update(input resources.Resource, update updates.Update) (singles.Single, error) {
	content := update.Content()
	retResource, err := app.Select(input, content.Name())
	if err != nil {
		return nil, err
	}

	single := retResource.Current().Current()
	storage := single.Storage()
	err = app.validate(
		content.Hash(),
		update.Vote(),
		storage.Blacklist(),
		storage.Whitelist(),
	)

	if err != nil {
		return nil, err
	}

	blacklist := app.mergeLists(
		storage.Blacklist(),
		content.BlacklistAddition(),
		content.BlacklistRemoval(),
	)

	whitelist := app.mergeLists(
		storage.Whitelist(),
		content.WhitelistAddition(),
		content.WhitelistRemoval(),
	)

	delimiter := storage.Delimiter()
	storageBuilder := app.storageBuilder.Create().WithDelimiter(delimiter)
	if len(blacklist) > 0 {
		storageBuilder.WithBlacklist(blacklist)
	}

	if len(blacklist) > 0 {
		storageBuilder.WithBlacklist(blacklist)
	}

	if len(whitelist) > 0 {
		storageBuilder.WithWhitelist(whitelist)
	}

	if storage.IsDeleted() {
		storageBuilder.IsDeleted()
	}

	updatedStorage, err := storageBuilder.Now()
	if err != nil {
		return nil, err
	}

	data := single.Bytes()
	return app.singleBuiler.Create().
		WithBytes(data).
		WithStorage(updatedStorage).
		Now()
}

func (app *application) delete(input resources.Resource, delete deletes.Delete) (singles.Single, error) {
	retResource, err := app.Select(input, delete.Name())
	if err != nil {
		return nil, err
	}

	single := retResource.Current().Original()
	storage := single.Storage()
	delimiter := storage.Delimiter()
	delimiterIndex := delimiter.Index()
	pHash, err := app.hashAdapter.FromBytes([]byte(strconv.Itoa(int(delimiterIndex))))
	if err != nil {
		return nil, err
	}

	vote := delete.Vote()
	err = app.validate(*pHash, vote, storage.Blacklist(), storage.Whitelist())
	if err != nil {
		return nil, err
	}

	updatedStorageBuilder := app.storageBuilder.Create().
		IsDeleted().
		WithDelimiter(delimiter)

	if storage.HasWhitelist() {
		whitelist := storage.Whitelist()
		updatedStorageBuilder.WithWhitelist(whitelist)
	}

	if storage.HasBlacklist() {
		blacklist := storage.Blacklist()
		updatedStorageBuilder.WithBlacklist(blacklist)
	}

	updatedStorage, err := updatedStorageBuilder.Now()
	if err != nil {
		return nil, err
	}

	data := single.Bytes()
	return app.singleBuiler.Create().
		WithBytes(data).
		WithStorage(updatedStorage).
		Now()
}

func (app *application) validate(
	msg hash.Hash,
	vote signers.Vote,
	blacklist []hash.Hash,
	whitelist []hash.Hash,
) error {
	if blacklist == nil {
		blacklist = []hash.Hash{}
	}

	if whitelist == nil {
		whitelist = []hash.Hash{}
	}

	if len(whitelist) > 0 {
		isApproved, err := app.voteAdapter.ToVerification(vote, msg.String(), whitelist)
		if err != nil {
			return err
		}

		if !isApproved {
			return errors.New("the delete request could not be approved because the resource contains a whitelistt, which the voter is NOT a member of")
		}
	}

	if len(blacklist) > 0 {
		isApproved, err := app.voteAdapter.ToVerification(vote, msg.String(), blacklist)
		if err != nil {
			return err
		}

		if !isApproved {
			return errors.New("the delete request could not be approved because the resource contains a blacklist, which the voter is a member of")
		}
	}

	return nil
}

func (app *application) mergeLists(original []hash.Hash, addition []hash.Hash, removal []hash.Hash) []hash.Hash {
	if original == nil {
		original = []hash.Hash{}
	}

	if addition == nil {
		addition = []hash.Hash{}
	}

	if removal == nil {
		removal = []hash.Hash{}
	}

	list := original
	if len(addition) > 0 {
		list = append(list, addition...)
	}

	if len(removal) > 0 {
		updated := []hash.Hash{}
		for _, oneHash := range removal {
			isRemoved := false
			for _, oneToRemoveHash := range removal {
				if oneHash.Compare(oneToRemoveHash) {
					isRemoved = true
					break
				}
			}

			if !isRemoved {
				updated = append(updated, oneHash)
			}
		}

		list = updated
	}

	return list
}

func (app *application) updateStorageInResource(input resources.Resource, switcher switchers.Switcher) (resources.Resource, error) {
	loadedList := []switchers.Switcher{}
	if input.HasLoaded() {
		loadedList = input.Loaded().List()
	}

	loadedList = append(loadedList, switcher)
	loaded, err := app.switchersBuilder.Create().
		WithList(loadedList).
		Now()

	if err != nil {
		return nil, err
	}

	all := input.All()
	builder := app.builder.Create().
		WithAll(all).
		WithLoaded(loaded)

	if input.HasCurrent() {
		current := input.Current()
		builder.WithCurrent(current)
	}

	return builder.Now()
}

func (app *application) loadSwitcherInResource(input resources.Resource, switcher switchers.Switcher) (resources.Resource, error) {
	loadedList := []switchers.Switcher{}
	if input.HasLoaded() {
		loadedList = input.Loaded().List()
	}

	loadedList = append(loadedList, switcher)
	loaded, err := app.switchersBuilder.Create().
		WithList(loadedList).
		Now()

	if err != nil {
		return nil, err
	}

	all := input.All()
	builder := app.builder.Create().
		WithAll(all).
		WithLoaded(loaded)

	if input.HasCurrent() {
		current := input.Current()
		builder.WithCurrent(current)
	}

	return builder.Now()
}
