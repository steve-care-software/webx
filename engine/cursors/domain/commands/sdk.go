package commands

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/commands/contents"
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/signers"
)

// Command represents a command
type Command interface {
	Hash() hash.Hash
	Content() contents.Content
	Signature() signers.Signature
}
