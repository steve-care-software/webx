package receipts

import (
	"github.com/steve-care-software/identity/domain/accounts/signers"
	"github.com/steve-care-software/identity/domain/hash"
	"github.com/steve-care-software/datastencil/domain/receipts/commands"
)

// NewReceiptBuilder creates a new builder instance
func NewReceiptBuilder() ReceiptBuilder {
	hashAdapter := hash.NewAdapter()
	return createReceiptBuilder(
		hashAdapter,
	)
}

// ReceiptBuilder represents a receipt builder
type ReceiptBuilder interface {
	Create() ReceiptBuilder
	WithCommands(commands commands.Commands) ReceiptBuilder
	WithSignature(signature signers.Signature) ReceiptBuilder
	Now() (Receipt, error)
}

// Receipt represents a receipt
type Receipt interface {
	Hash() hash.Hash
	Commands() commands.Commands
	Signature() signers.Signature
}
