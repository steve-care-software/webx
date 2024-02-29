package sqllites

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
	"github.com/steve-care-software/datastencil/domain/orms"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"
)

const resourceNameDelimiter = "_"
const endOfLine = "\n"
const connectionNameDelimiter = "$"

// NewOrmRepository creates a new orm repository
func NewOrmRepository(
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) orms.Repository {
	hashAdapter := hash.NewAdapter()
	builders := map[string]interface{}{
		"libraries_layers_links": links.NewOperatorBuilder(),
	}

	return createOrmRepository(
		hashAdapter,
		builders,
		skeleton,
		dbPtr,
	)
}

// NewOrmService creates a new orm service
func NewOrmService(
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
	txPtr *sql.Tx,
) orms.Service {
	hashAdapter := hash.NewAdapter()
	return createOrmService(
		hashAdapter,
		skeleton,
		dbPtr,
		txPtr,
	)
}

// NewSkeletonFactory creates a new skeleton factory
func NewSkeletonFactory() skeletons.Factory {
	builder := skeletons.NewBuilder()
	resourcesBuilder := resources.NewBuilder()
	resourceBuilder := resources.NewResourceBuilder()
	fieldsBuilder := resources.NewFieldsBuilder()
	fieldBuilder := resources.NewFieldBuilder()
	kindBuilder := resources.NewKindBuilder()
	connectionsBuilder := connections.NewBuilder()
	connectionBuilder := connections.NewConnectionBuilder()
	connectionFieldBuilder := connections.NewFieldBuilder()
	return createSkeletonFactory(
		builder,
		resourcesBuilder,
		resourceBuilder,
		fieldsBuilder,
		fieldBuilder,
		kindBuilder,
		connectionsBuilder,
		connectionBuilder,
		connectionFieldBuilder,
	)
}