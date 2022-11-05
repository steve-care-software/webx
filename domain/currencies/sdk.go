package currencies

import (
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/cryptography/signatures"
	"github.com/steve-care-software/webx/domain/programs"
)

// Builder represents the currency builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSupply(supply uint) Builder
	Now() (Currency, error)
}

// Currency represents a currency
type Currency interface {
	Hash() hash.Hash
	Name() string
	Supply() uint
}

// UnitBuilder represents a unit builder
type UnitBuilder interface {
	Create() UnitBuilder
	WithCurrency(currency Currency) UnitBuilder
	WithAmount(amount hash.Hash) UnitBuilder
	WithNonce(nonce uint) UnitBuilder
	WithOwner(owner []hash.Hash) UnitBuilder
	ActivatedOn(activatedOn time.Time) UnitBuilder
	Now() (Unit, error)
}

// Unit represents a currency unit
type Unit interface {
	Hash() hash.Hash
	Currency() Currency
	Amount() hash.Hash
	Nonce() uint
	Owner() []hash.Hash
	ActivatedOn() time.Time
}

// TransactionBuilder represents a transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithValue(value TransactionValue) TransactionBuilder
	WithSignature(signature signatures.RingSignature) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Value() TransactionValue
	Signature() signatures.RingSignature
}

// TransactionValueBuilder represents a transaction value builder
type TransactionValueBuilder interface {
	Create() TransactionValueBuilder
	WithTransfer(transfer Transfer) TransactionValueBuilder
	WithMerge(merge Merge) TransactionValueBuilder
	WithSplit(split Split) TransactionValueBuilder
	WithCondition(condition programs.Program) TransactionValueBuilder
	Now() (TransactionValue, error)
}

// TransactionValue represents a transaction value
type TransactionValue interface {
	Hash() hash.Hash
	Content() TransactionContent
	Condition() programs.Program
}

// TransactionContent represents a transaction content
type TransactionContent interface {
	Hash() hash.Hash
	IsTransfer() bool
	Transfer() Transfer
	IsMerge() bool
	Merge() Merge
	IsSplit() bool
	Split() Split
}

// TransfersBuilder represents a transfers builder
type TransfersBuilder interface {
	Create() TransfersBuilder
	WithList(list []Transfer) TransfersBuilder
	Now() (Transfers, error)
}

// Transfers represents transfers
type Transfers interface {
	Hash() hash.Hash
	List() []Transfer
}

// TransferBuilder represents a transfer builder
type TransferBuilder interface {
	Create() TransferBuilder
	WithOrigin(origin Origin) TransferBuilder
	WithAmount(amount hash.Hash) TransferBuilder
	WithNonce(nonce uint) TransferBuilder
	Now() (Transfer, error)
}

// Transfer represents a transfer
type Transfer interface {
	Hash() hash.Hash
	Origin() Origin
	Amount() hash.Hash
	Nonce() uint
}

// MergeBuilder represents a merge builder
type MergeBuilder interface {
	Create() MergeBuilder
	From(from Origins) MergeBuilder
	To(to Transfer) MergeBuilder
	Now() (Merge, error)
}

// Merge represents a merge
type Merge interface {
	Hash() hash.Hash
	From() Origins
	To() Transfer
}

// SplitBuilder represents a split builder
type SplitBuilder interface {
	Create() SplitBuilder
	From(from Origin) SplitBuilder
	To(to Transfers) SplitBuilder
	Now() (Split, error)
}

// Split represents a split
type Split interface {
	Hash() hash.Hash
	From() Origin
	To() Transfers
}

// OriginsBuilder represents an origins builder
type OriginsBuilder interface {
	Create() OriginsBuilder
	WithList(list []Origin) OriginsBuilder
	Now() (Origins, error)
}

// Origins represents origins
type Origins interface {
	Hash() hash.Hash
	List() []Origin
}

// OriginBuilder represents an origin builder
type OriginBuilder interface {
	Create() OriginBuilder
	WithUnit(unit Unit) OriginBuilder
	WithOwner(owner Owner) OriginBuilder
	Now() (Origin, error)
}

// Origin represents a unit origin
type Origin interface {
	Hash() hash.Hash
	IsUnit() bool
	Unit() Unit
	IsOwner() bool
	Owner() Owner
}

// OwnerBuilder represents an owner builder
type OwnerBuilder interface {
	Create() OwnerBuilder
	WithSource(souce Source) OwnerBuilder
	WithAmount(amount uint) OwnerBuilder
	WithNonce(nonce uint) OwnerBuilder
	Now() (Owner, error)
}

// Owner represents an owner
type Owner interface {
	Hash() hash.Hash
	Source() Source
	Amount() uint
	Nonce() uint
}

// SourceBuilder represents the source builder
type SourceBuilder interface {
	Create() SourceBuilder
	WithUnit(unit Unit) SourceBuilder
	WithTransfer(transfer Transfer) SourceBuilder
	Now() (Source, error)
}

// Source represents a source
type Source interface {
	Hash() hash.Hash
	IsUnit() bool
	Unit() Unit
	IsTransfer() bool
	Transfer() Transfer
}
