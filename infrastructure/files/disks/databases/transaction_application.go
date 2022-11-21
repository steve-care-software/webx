package databases

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/steve-care-software/webx/applications/databases/transactions"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type transactionApplication struct {
	absFilePath string
	contexts    map[uint][]*transactionCommand
}

type transactionCommand struct {
	insert  []byte
	delete  *hash.Hash
	approve *transactionCommandApprove
}

type transactionCommandApprove struct {
	hash  hash.Hash
	proof big.Int
}

func createTransactionApplication(
	absFilePath string,
) transactions.Application {
	out := transactionApplication{
		absFilePath: absFilePath,
		contexts:    map[uint][]*transactionCommand{},
	}

	return &out
}

// Begin begins a context
func (app *transactionApplication) Begin() (*uint, error) {
	context := app.generateContext()
	app.contexts[context] = []*transactionCommand{}
	return &context, nil
}

// Insert inserts content to a context
func (app *transactionApplication) Insert(context uint, content []byte) error {
	if _, ok := app.contexts[context]; !ok {
		str := fmt.Sprintf("the context (%d) is invalid and therefore cannot be used to insert content", context)
		return errors.New(str)
	}

	// hash the content:

	// make sure the key referenced by the hashed content does not already exists in the reference:

	app.contexts[context] = append(app.contexts[context], &transactionCommand{
		insert: content,
	})

	return nil
}

// Delete deletes content from a context
func (app *transactionApplication) Delete(context uint, hash hash.Hash) error {
	if _, ok := app.contexts[context]; !ok {
		str := fmt.Sprintf("the context (%d) is invalid and therefore cannot be used to delete content", context)
		return errors.New(str)
	}

	// make sure the key referenced by the hash exists in the reference:

	app.contexts[context] = append(app.contexts[context], &transactionCommand{
		delete: &hash,
	})

	return nil
}

// Approve approves a content to a context
func (app *transactionApplication) Approve(context uint, hash hash.Hash, proof big.Int) error {
	if _, ok := app.contexts[context]; !ok {
		str := fmt.Sprintf("the context (%d) is invalid and therefore cannot be used to approve content", context)
		return errors.New(str)
	}

	// make sure the key referenced by the hash is currently pending:

	app.contexts[context] = append(app.contexts[context], &transactionCommand{
		approve: &transactionCommandApprove{
			hash:  hash,
			proof: proof,
		},
	})

	return nil
}

// Push pushes a context
func (app *transactionApplication) Push(context uint) error {
	return nil
}

// Cancel cancels a context
func (app *transactionApplication) Cancel(context uint) error {
	if _, ok := app.contexts[context]; !ok {
		str := fmt.Sprintf("the context (%d) is invalid and therefore cannot be used to cancel a series of operations", context)
		return errors.New(str)
	}

	delete(app.contexts, context)
	return nil
}

func (app *transactionApplication) generateContext() uint {
	randNumber := func() uint {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		return uint(r1.Int())
	}

	context := randNumber()
	if _, ok := app.contexts[context]; ok {
		return app.generateContext()
	}

	return context
}
