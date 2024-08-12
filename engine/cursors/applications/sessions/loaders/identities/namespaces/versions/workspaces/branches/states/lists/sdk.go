package lists

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states/singles/lists/creates"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/deletes"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/updates"
)

// Application represents a resource application
type Application interface {
	Create(input resources.Resource, create creates.Create) (resources.Resource, error) // creates a new list
	Load(input resources.Resource, name string) (resources.Resource, error)
	Select(input resources.Resource, name string) (resources.Resource, error)
	Insert(input resources.Resource, values []string) (resources.Resource, error)        // insert elements in the current list
	Remove(input resources.Resource, values []string) (resources.Resource, error)        // remove elements from the current list
	Fetch(input resources.Resource, index uint64, amount uint64) (singles.Single, error) // fetch elements from the current list
	Union(input resources.Resource, listName string) (singles.Single, error)             // union the specified list with the current one
	Intersect(input resources.Resource, listName string) (singles.Single, error)         // intersect the specified list with the current one

	Delete(input resources.Resource, delete deletes.Delete) (resources.Resource, error) // delete the current list
	Retrieve(input resources.Resource) (singles.Single, error)                          // retrieve the current list
	Update(input resources.Resource, update updates.Update) (resources.Resource, error)
	Commit(input resources.Resource) (transactions.Transactions, error)
	Transact(input resources.Resource, trx transactions.Transactions) error
}
