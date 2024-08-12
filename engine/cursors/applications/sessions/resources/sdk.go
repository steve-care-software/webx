package resources

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/databases"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/deletes"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/inserts"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/updates"
)

const noLoadedResourceErr = "There is no loaded resource"
const noSelectedResourceErr = "There is no selected resource and therefore the requested action is impossible"
const cannotAlterNeverCommittedErr = "The current resource cannot be altered because it has never been comitted"

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithDatabase(dbApp databases.Application) Builder
	Now() (Application, error)
}

// Application represents a resource application
type Application interface {
	Insert(input resources.Resource, insert inserts.Insert) (resources.Resource, error)
	Load(input resources.Resource, delimiterIndex uint64) (resources.Resource, error)
	Select(input resources.Resource, delimiterIndex uint64) (resources.Resource, error)
	Delete(input resources.Resource, delete deletes.Delete) (resources.Resource, error)
	Retrieve(input resources.Resource) (singles.Single, error)
	Update(input resources.Resource, update updates.Update) (resources.Resource, error)
	Commit(input resources.Resource) error
	Transact(input resources.Resource, trx transactions.Transactions) error
}
