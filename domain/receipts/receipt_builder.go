package receipts

import (
	"errors"

	"github.com/steve-care-software/identity/domain/accounts/signers"
	"github.com/steve-care-software/identity/domain/hash"
	"github.com/steve-care-software/datastencil/domain/receipts/commands"
)

type receiptBuilder struct {
	hashAdapter hash.Adapter
	commands    commands.Commands
	signature   signers.Signature
}

func createReceiptBuilder(
	hashAdapter hash.Adapter,
) ReceiptBuilder {
	out := receiptBuilder{
		hashAdapter: hashAdapter,
		commands:    nil,
		signature:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *receiptBuilder) Create() ReceiptBuilder {
	return createReceiptBuilder(
		app.hashAdapter,
	)
}

// WithCommands add commands to the builder
func (app *receiptBuilder) WithCommands(commands commands.Commands) ReceiptBuilder {
	app.commands = commands
	return app
}

// WithSignature add signature to the builder
func (app *receiptBuilder) WithSignature(signature signers.Signature) ReceiptBuilder {
	app.signature = signature
	return app
}

// Now builds a new Receipt instance
func (app *receiptBuilder) Now() (Receipt, error) {
	if app.commands == nil {
		return nil, errors.New("the commands is mandatory in order to build a Receipt instance")
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Receipt instance")
	}

	sigBytes, err := app.signature.Bytes()
	if err != nil {
		return nil, err
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.commands.Hash().Bytes(),
		sigBytes,
	})

	if err != nil {
		return nil, err
	}

	return createReceipt(*pHash, app.commands, app.signature), nil
}
