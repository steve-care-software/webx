package files

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/applications"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
	infra_bytes "github.com/steve-care-software/webx/engine/databases/bytes/infrastructure/bytes"
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

const amountOfBytesIntUint64 = 8

const contextIdentifierUndefinedPattern = "the context identifier (%d) does not exists"

// NewApplicationBuilder creates a new application builder
func NewApplicationBuilder() applications.Builder {
	hashAdapter := hash.NewAdapter()
	statesAdapter := infra_bytes.NewStateAdapter()
	statesBuilder := states.NewBuilder()
	stateBuilder := states.NewStateBuilder()
	pointersBuilder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	entriesBuilder := entries.NewBuilder()
	entryBuilder := entries.NewEntryBuilder()
	delimitersBuilder := delimiters.NewBuilder()
	delimiterBuilder := delimiters.NewDelimiterBuilder()
	return createApplicationBuilder(
		hashAdapter,
		statesAdapter,
		statesBuilder,
		stateBuilder,
		pointersBuilder,
		pointerBuilder,
		entriesBuilder,
		entryBuilder,
		delimitersBuilder,
		delimiterBuilder,
	)
}
