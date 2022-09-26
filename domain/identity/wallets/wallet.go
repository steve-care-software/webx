package wallets

import (
	"github.com/steve-care-software/syntax/domain/identity/wallets/assets"
	"github.com/steve-care-software/syntax/domain/identity/wallets/transactions"
)

type wallet struct {
	incoming assets.Assets
	outgoing transactions.Transactions
}
