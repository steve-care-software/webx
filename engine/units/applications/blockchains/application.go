package blockchains

import (
	"errors"
	"fmt"

	entity_applications "github.com/steve-care-software/webx/engine/databases/entities/applications"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions"
)

type application struct {
	entityApplication   entity_applications.Application
	transactionsBuilder transactions.Builder
	queue               map[string][]transactions.Transaction
}

func createApplication(
	entityApplication entity_applications.Application,
	transactionsBuilder transactions.Builder,
) Application {
	out := application{
		entityApplication:   entityApplication,
		transactionsBuilder: transactionsBuilder,
		queue:               map[string][]transactions.Transaction{},
	}

	return &out
}

// List lists the blockchain identifiers
func (app *application) List() ([]string, error) {
	return nil, nil
}

// Insert inserts a blockchain
func (app *application) Insert(blockchain blockchains.Blockchain) error {
	return nil
}

// Update updates a blockchain identifier
func (app *application) Update(currentIdentifier string, newIdentifier string) error {
	return nil
}

// Delete deletes a blockchain
func (app *application) Delete(hash hash.Hash) error {
	return nil
}

// Retrieve retrieves a blockchain by identifier
func (app *application) Retrieve(identifier string) (blockchains.Blockchain, error) {
	return nil, nil
}

// Transact executes a transaction
func (app *application) Transact(identifier string, trx transactions.Transactions) error {
	// retrieve the current queue, then append the new ones to the queue:
	newTrx := []transactions.Transaction{}
	if currentTrx, ok := app.queue[identifier]; ok {
		newTrx = currentTrx
	}

	newTrx = append(newTrx, trx.List()...)
	app.queue[identifier] = newTrx

	return nil
}

// Queue returns the transaction queue
func (app *application) Queue(identifier string) (transactions.Transactions, error) {
	if currentTrx, ok := app.queue[identifier]; ok {
		return app.transactionsBuilder.Create().
			WithList(currentTrx).
			Now()
	}

	str := fmt.Sprintf("there is no transaction queue for the provided blockchain identifier: %s", identifier)
	return nil, errors.New(str)
}

// Mine mines a block on the blockchain
func (app *application) Mine(identifier string) error {
	return nil
}

// Sync sync a blockchain with peers
func (app *application) Sync(identifier string) error {
	return nil
}
