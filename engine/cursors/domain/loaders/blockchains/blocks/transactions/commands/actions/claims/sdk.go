package claims

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands/actions/claims/contents"
	"github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
)

// Claim represents a claim
type Claim interface {
	Hash() hash.Hash
	Content() contents.Content
	Signature() signatures.Signature
}
