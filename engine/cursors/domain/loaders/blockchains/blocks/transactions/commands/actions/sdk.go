package actions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands/actions/claims"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands/actions/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands/actions/transfers"
)

// Actions represents actions
type Actions interface {
	Hash() hash.Hash
	List() []Action
}

// Action represents an action
type Action interface {
	Hash() hash.Hash
	IsClaim() bool
	Claim() claims.Claim
	IsDatabase() bool
	Database() databases.Database
	IsTransfer() bool
	Transfer() transfers.Transfer
}
