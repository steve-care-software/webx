package grammars

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/cryptography/hashtrees"
)

type adapter struct {
	hashAdapter     hash.Adapter
	hashTreeAdapter hashtrees.Adapter
	builder         Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	hashTreeAdapter hashtrees.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter:     hashAdapter,
		hashTreeAdapter: hashTreeAdapter,
		builder:         builder,
	}

	return &out
}

// ToGrammar returns the content to a grammar instance
func (app *adapter) ToGrammar(content []byte) (Grammar, error) {
	return nil, nil
}

// ToContent returns the grammar to content
func (app *adapter) ToContent(ins Grammar) ([]byte, error) {
	return nil, nil
}
