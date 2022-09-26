package wallets

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/wallets/assets"
	"github.com/steve-care-software/syntax/domain/identity/wallets/transactions"
)

// Builder represents a wallets builder
type Builder interface {
	Create() Builder
	WithList(list []Wallet) Builder
	Now() (Wallets, error)
}

// Wallets represents wallets
type Wallets interface {
	List() []Wallet
	FetchByUnits(unitHashes []hash.Hash) (Wallet, error)
	FetchListExceptID(id uuid.UUID) ([]Wallet, error)
	FetchByID(id uuid.UUID) (Wallet, error)
}

// WalletBuilder represents a wallet builder
type WalletBuilder interface {
	Create() WalletBuilder
	WithID(id uuid.UUID) WalletBuilder
	WithName(name string) WalletBuilder
	WithDescription(description string) WalletBuilder
	WithIncoming(incoming assets.Assets) WalletBuilder
	WithOutgoing(outgoing transactions.Transactions) WalletBuilder
	Now() (Wallet, error)
}

// Wallet represents a wallet
type Wallet interface {
	ID() uuid.UUID
	Name() string
	Description() string
	HasIncoming() bool
	Incoming() assets.Assets
	HasOutgoing() bool
	Outgoing() transactions.Transactions
}
