package selects

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/steve-care-software/syntax/domain/blockchains"
	"github.com/steve-care-software/syntax/domain/blockchains/blocks"
	"github.com/steve-care-software/syntax/domain/blockchains/transactions"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

type application struct {
	builder             blockchains.Builder
	repository          blockchains.Repository
	service             blockchains.Service
	blockRepository     blocks.Repository
	transactionsBuilder transactions.Builder
	reference           hash.Hash
}

func createApplication(
	builder blockchains.Builder,
	repository blockchains.Repository,
	service blockchains.Service,
	blockRepository blocks.Repository,
	transactionsBuilder transactions.Builder,
	reference hash.Hash,
) Application {
	out := application{
		builder:             builder,
		repository:          repository,
		service:             service,
		blockRepository:     blockRepository,
		transactionsBuilder: transactionsBuilder,
		reference:           reference,
	}

	return &out
}

// Retrieve retrieves the current blockchain
func (app *application) Retrieve() (blockchains.Blockchain, error) {
	return app.repository.Retrieve(app.reference)
}

// Transact adds a transaction to the blockchain
func (app *application) Transact(trx transactions.Transaction) error {
	blockchain, err := app.Retrieve()
	if err != nil {
		return err
	}

	return app.updateBlockchain(blockchain, trx, nil)
}

// BlockByPreviousHash returns a block by previous hash
func (app *application) BlockByPreviousHash(prev hash.Hash) (blocks.Block, error) {
	blockchain, err := app.Retrieve()
	if err != nil {
		return nil, err
	}

	reference := blockchain.Reference()
	return app.blockRepository.RetrieveByPreviousHash(reference, prev)
}

// BlockByHash returns a block by hash
func (app *application) BlockByHash(hash hash.Hash) (blocks.Block, error) {
	blockchain, err := app.Retrieve()
	if err != nil {
		return nil, err
	}

	reference := blockchain.Reference()
	return app.blockRepository.RetrieveByHash(reference, hash)
}

// BlockByHeight returns a block by height
func (app *application) BlockByHeight(height uint) (blocks.Block, error) {
	blockchain, err := app.Retrieve()
	if err != nil {
		return nil, err
	}

	reference := blockchain.Reference()
	return app.blockRepository.RetrieveByHeight(reference, height)
}

// Search searches a transaction by hash and returns the blocks where the transaction was found, or the pending transactions
func (app *application) Search(trx hash.Hash) (blocks.Block, transactions.Transactions, error) {
	blockchain, err := app.Retrieve()
	if err != nil {
		return nil, nil, err
	}

	pendings := blockchain.Pendings()
	_, err = pendings.Fetch(trx)
	if err == nil {
		return nil, pendings, nil
	}

	head := blockchain.Head()
	found, err := app.searchInBlock(trx, head)
	if err != nil {
		return nil, nil, err
	}

	if found == nil {
		str := fmt.Sprintf("the transaction (hash: %s) could not be found in the current blockchain (reference: %s)", trx.String(), blockchain.Reference().String())
		return nil, nil, errors.New(str)
	}

	return found, nil, nil
}

// ConnectList adds the connections to the blockchain
func (app *application) ConnectList(conns []*url.URL) error {
	blockchain, err := app.Retrieve()
	if err != nil {
		return err
	}

	return app.updateBlockchain(blockchain, nil, conns)
}

// Connect adds the connection to the blockchain
func (app *application) Connect(conn *url.URL) error {
	blockchain, err := app.Retrieve()
	if err != nil {
		return err
	}

	return app.updateBlockchain(blockchain, nil, []*url.URL{
		conn,
	})
}

func (app *application) searchInBlock(trx hash.Hash, head blocks.Block) (blocks.Block, error) {
	_, err := head.Transactions().Fetch(trx)
	if err == nil {
		return head, nil
	}

	if !head.HasPrevious() {
		return nil, nil
	}

	prevHash := head.Previous()
	prevBlock, err := app.BlockByHash(*prevHash)
	if err != nil {
		return nil, err
	}

	return app.searchInBlock(trx, prevBlock)
}

func (app *application) updateBlockchain(blockchain blockchains.Blockchain, trx transactions.Transaction, connections []*url.URL) error {
	reference := blockchain.Reference()
	head := blockchain.Head()
	createdOn := blockchain.CreatedOn()
	builder := app.builder.Create().WithReference(reference).WithHead(head).CreatedOn(createdOn)

	// connections:
	conns := []*url.URL{}
	if blockchain.HasConnections() {
		conns = blockchain.Connections()
	}

	if connections != nil {
		for _, oneConnection := range connections {
			conns = append(conns, oneConnection)
		}
	}

	mp := map[string]*url.URL{}
	for _, oneConnection := range conns {
		keyname := oneConnection.String()
		mp[keyname] = oneConnection
	}

	uniqueConnections := []*url.URL{}
	for _, oneConnection := range mp {
		uniqueConnections = append(uniqueConnections, oneConnection)
	}

	if len(uniqueConnections) > 0 {
		builder.WithConnections(uniqueConnections)
	}

	// transactions:
	trxList := []transactions.Transaction{}
	if blockchain.HasPendings() {
		trxHash := trx.Hash()
		pendings := blockchain.Pendings()
		_, err := pendings.Fetch(trxHash)
		if err == nil {
			str := fmt.Sprintf("the transaction (hash: %s) is already saved in the pending transactions of the current blockchain (reference: %s)", trxHash.String(), reference.String())
			return errors.New(str)
		}

		trxList = pendings.List()
	}

	trxList = append(trxList, trx)
	trxIns, err := app.transactionsBuilder.Create().WithList(trxList).Now()
	if err != nil {
		return err
	}

	updated, err := builder.WithPendings(trxIns).Now()
	if err != nil {
		return err
	}

	return app.service.Save(updated)
}
