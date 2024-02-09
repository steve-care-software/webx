package receipts

import (
	"github.com/steve-care-software/identity/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/receipts/commands"
)

// NewReceiptForTests creates a new receipt for tests
func NewReceiptForTests(commands commands.Commands, sig signers.Signature) Receipt {
	ins, err := NewReceiptBuilder().Create().WithCommands(commands).WithSignature(sig).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
