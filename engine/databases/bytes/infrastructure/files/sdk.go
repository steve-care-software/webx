package files

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/applications"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/modifications"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"
	infra_bytes "github.com/steve-care-software/webx/engine/databases/bytes/infrastructure/bytes"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

const amountOfBytesIntUint64 = 8

const contentIdentifierUndefinedPattern = "the context identifier (%d) does not exists"

// NewApplicationBuilder creates a new application builder
func NewApplicationBuilder() applications.Builder {
	hashAdapter := hash.NewAdapter()
	statesAdapter := infra_bytes.NewStateAdapter()
	statesBuilder := states.NewBuilder()
	stateBuilder := states.NewStateBuilder()
	containersBuilder := containers.NewBuilder()
	containerBuilder := containers.NewContainerBuilder()
	pointersBuilder := pointers.NewBuilder()
	pointerBuilder := pointers.NewPointerBuilder()
	modificationBuilder := modifications.NewBuilder()
	entriesBuilder := entries.NewBuilder()
	deletesBuilder := deletes.NewBuilder()
	retrievalsBuilder := retrievals.NewBuilder()
	return createApplicationBuilder(
		hashAdapter,
		statesAdapter,
		statesBuilder,
		stateBuilder,
		containersBuilder,
		containerBuilder,
		pointersBuilder,
		pointerBuilder,
		modificationBuilder,
		entriesBuilder,
		deletesBuilder,
		retrievalsBuilder,
	)
}
