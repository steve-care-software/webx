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
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type application struct {
	dbApp            databases.Application
	builder          resources.Builder
	storageBuilder   storages.StorageBuilder
	switchersBuilder switchers.Builder
	switcherBuilder  switchers.SwitcherBuilder
	singleBuiler     singles.Builder
	delimiterBuilder delimiters.DelimiterBuilder
	voteAdapter      signers.VoteAdapter
	hashAdapter      hash.Adapter
}

func createApplication(
	dbApp databases.Application,
	builder resources.Builder,
	storageBuilder storages.StorageBuilder,
	switchersBuilder switchers.Builder,
	switcherBuilder switchers.SwitcherBuilder,
	singleBuiler singles.Builder,
	delimiterBuilder delimiters.DelimiterBuilder,
	voteAdapter signers.VoteAdapter,
	hashAdapter hash.Adapter,
) Application {
	out := application{
		dbApp:            dbApp,
		builder:          builder,
		storageBuilder:   storageBuilder,
		switchersBuilder: switchersBuilder,
		switcherBuilder:  switcherBuilder,
		singleBuiler:     singleBuiler,
		delimiterBuilder: delimiterBuilder,
		voteAdapter:      voteAdapter,
		hashAdapter:      hashAdapter,
	}

	return &out
}

// Insert inserts a resource
func (app *application) Insert(input resources.Resource, data []byte, blacklist []hash.Hash, whitelist []hash.Hash) (resources.Resource, error) {
	nextIndex := input.All().NextIndex()
	length := uint64(len(data))
	delimiter, err := app.delimiterBuilder.Create().WithIndex(nextIndex).WithLength(length).Now()
	if err != nil {
		return nil, err
	}

	storageBuilder := app.storageBuilder.Create().WithDelimiter(delimiter)
	if blacklist != nil {
		storageBuilder.WithBlacklist(blacklist)
	}

	if whitelist != nil {
		storageBuilder.WithBlacklist(whitelist)
	}

	storage, err := storageBuilder.Now()
	if err != nil {
		return nil, err
	}

	switcher, err := app.switcherBuilder.Create().WithUpdated(storage).Now()
	if err != nil {
		return nil, err
	}

	return app.loadSwitcherInResource(input, switcher)
}

// Load loads a resource
func (app *application) Load(input resources.Resource, delimiterIndex uint64) (resources.Resource, error) {
	storage, err := input.All().FetchByDelimiterIndex(delimiterIndex)
	if err != nil {
		return nil, err
	}

	delimiter := storage.Delimiter()
	retBytes, err := app.dbApp.Read(delimiter)
	if err != nil {
		return nil, err
	}

	single, err := app.singleBuiler.Create().WithStorage(storage).WithBytes(retBytes).Now()
	if err != nil {
		return nil, err
	}

	switcher, err := app.switcherBuilder.Create().WithOriginal(single).Now()
	if err != nil {
		return nil, err
	}

	return app.loadSwitcherInResource(input, switcher)
}

// Select selects a loaded resource
func (app *application) Select(input resources.Resource, delimiterIndex uint64) (resources.Resource, error) {
	if !input.HasLoaded() {
		return nil, errors.New(noLoadedResourceErr)
	}

	loaded := input.Loaded()
	current, err := loaded.FetchByDelimiterIndex(delimiterIndex)
	if err != nil {
		return nil, err
	}

	all := input.All()
	return app.builder.Create().WithAll(all).WithCurrent(current).WithLoaded(loaded).Now()
}

// Delete deletes the selected resource
func (app *application) Delete(input resources.Resource, vote signers.Vote) (resources.Resource, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noSelectedResourceErr)
	}

	current := input.Current()
	if !current.HasOriginal() {
		return nil, errors.New(cannotAlterNeverCommittedErr)
	}

	storage := current.Original().Storage()
	delimiter := storage.Delimiter()
	delimiterIndex := delimiter.Index()
	pHash, err := app.hashAdapter.FromBytes([]byte(strconv.Itoa(int(delimiterIndex))))
	if err != nil {
		return nil, err
	}

	updatedStorageBuilder := app.storageBuilder.Create().IsDeleted().WithDelimiter(delimiter)
	if storage.HasWhitelist() {
		whitelist := storage.Whitelist()
		isApproved, err := app.voteAdapter.ToVerification(vote, pHash.String(), whitelist)
		if err != nil {
			return nil, err
		}

		if !isApproved {
			return nil, errors.New("the delete request could not be approved because the resource contains a whitelistt, which the voter is NOT a member of")
		}

		updatedStorageBuilder.WithWhitelist(whitelist)
	}

	if storage.HasBlacklist() {
		blacklist := storage.Blacklist()
		isApproved, err := app.voteAdapter.ToVerification(vote, pHash.String(), storage.Whitelist())
		if err != nil {
			return nil, err
		}

		if isApproved {
			return nil, errors.New("the delete request could not be approved because the resource contains a blacklist, which the voter is a member of")
		}

		updatedStorageBuilder.WithBlacklist(blacklist)
	}

	updatedStorage, err := updatedStorageBuilder.Now()
	if err != nil {
		return nil, err
	}

	switcher, err := app.switcherBuilder.Create().WithDeleted(updatedStorage).Now()
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
func (app *application) Update(
	input resources.Resource,
	addToBlacklist []hash.Hash,
	removeFromBlacklist []hash.Hash,
	addToWhitelist []hash.Hash,
	removeFromWhitelist []hash.Hash,
) (resources.Resource, error) {
	if !input.HasCurrent() {
		return nil, errors.New(noSelectedResourceErr)
	}

	current := input.Current()
	if !current.HasOriginal() {
		return nil, errors.New(cannotAlterNeverCommittedErr)
	}

	storage := current.Current().Storage()
	blacklist := app.mergeLists(
		storage.Blacklist(),
		addToBlacklist,
		removeFromBlacklist,
	)

	whitelist := app.mergeLists(
		storage.Whitelist(),
		addToWhitelist,
		removeFromWhitelist,
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

	switcher, err := app.switcherBuilder.Create().WithUpdated(updatedStorage).Now()
	if err != nil {
		return nil, err
	}

	return app.updateStorageInResource(input, switcher)
}

// Commit commits the resource
func (app *application) Commit(input resources.Resource) error {
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
	loaded, err := app.switchersBuilder.Create().WithList(loadedList).Now()
	if err != nil {
		return nil, err
	}

	all := input.All()
	builder := app.builder.Create().WithAll(all).WithLoaded(loaded)
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
	loaded, err := app.switchersBuilder.Create().WithList(loadedList).Now()
	if err != nil {
		return nil, err
	}

	all := input.All()
	builder := app.builder.Create().WithAll(all).WithLoaded(loaded)
	if input.HasCurrent() {
		current := input.Current()
		builder.WithCurrent(current)
	}

	return builder.Now()
}
