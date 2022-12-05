package transfers

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/cryptography/signatures"
	"github.com/steve-care-software/webx/licenses/domain/licenses"
)

// Transfer represents a license transfer
type Transfer interface {
	Hash() hash.Hash
	Content() Content
	Signature() signatures.RingSignature
}

// Content represents a transfer content
type Content interface {
	Hash() hash.Hash
	License() licenses.License
	Owner() []hash.Hash
}
