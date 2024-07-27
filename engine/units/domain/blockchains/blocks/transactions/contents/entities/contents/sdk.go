package contents

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions/contents/entities/contents/votes"
	"github.com/steve-care-software/webx/engine/units/domain/governances"
	"github.com/steve-care-software/webx/engine/units/domain/identities/profiles"
	"github.com/steve-care-software/webx/engine/units/domain/units/transfers"
)

// Content represents an entity content
type Content interface {
	Hash() hash.Hash
	IsTransfer() bool
	Transfer() transfers.Transfer
	IsVote() bool
	Vote() votes.Vote
	IsProfile() bool
	Profile() profiles.Profile
	IsGovernance() bool
	Governance() governances.Governance
}
