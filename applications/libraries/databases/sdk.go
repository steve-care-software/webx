package databases

import (
	"github.com/steve-care-software/datastencil/applications/libraries/databases/actions"
	"github.com/steve-care-software/datastencil/applications/libraries/databases/transactions"
	"github.com/steve-care-software/datastencil/domain/commits/actions/resources/instances"
	"github.com/steve-care-software/datastencil/domain/commits/actions/resources/instances/skeletons"
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Application represents the database application
type Application interface {
	Skeleton() skeletons.Skeleton
	List(path []string) []string
	Retrieve(path []string, hash hash.Hash) (instances.Instance, error)
	Begin() transactions.Application
	Commit(trx transactions.Application) (uint, error)
	Actions() (actions.Application, error)
}
