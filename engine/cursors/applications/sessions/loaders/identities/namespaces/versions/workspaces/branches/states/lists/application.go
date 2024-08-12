package lists

import (
	"errors"

	resource_applications "github.com/steve-care-software/webx/engine/cursors/applications/sessions/resources"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states/singles/lists"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states/singles/lists/creates"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/deletes"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/inserts"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/transactions/updates"
)

type application struct {
	resourceApp   resource_applications.Application
	insertBuilder inserts.Builder
	updateBuilder updates.Builder
	deleteBuilder deletes.Builder
	adapter       lists.Adapter
	builder       lists.Builder
	signer        signers.Signer
}

func createApplication(
	resourceApp resource_applications.Application,
	adapter lists.Adapter,
	builder lists.Builder,
) Application {
	out := application{
		resourceApp: resourceApp,
		adapter:     adapter,
		builder:     builder,
	}

	return &out
}

// Create creates a new list
func (app *application) Create(input resources.Resource, create creates.Create) (resources.Resource, error) {
	list := create.List()
	data, err := app.adapter.ToBytes(list)
	if err != nil {
		return nil, err
	}

	name := list.Name()
	whitelist := create.Whitelist()
	insert, err := app.insertBuilder.Create().
		WithName(name).
		WithBytes(data).
		WithWhitelist(whitelist).
		Now()

	if err != nil {
		return nil, err
	}

	return app.resourceApp.Insert(input, insert)
}

// Load loads a new list
func (app *application) Load(input resources.Resource, name string) (resources.Resource, error) {
	return app.resourceApp.Load(input, name)
}

// Select selects s list and put it as its current
func (app *application) Select(input resources.Resource, name string) (resources.Resource, error) {
	return app.resourceApp.Load(input, name)
}

// Insert  elements in the current list
func (app *application) Insert(input resources.Resource, values []string) (resources.Resource, error) {
	if !input.HasCurrent() {
		return nil, errors.New("there is no current list selected")
	}

	retSingle, err := app.resourceApp.Retrieve(input)
	if err != nil {
		return nil, err
	}

	data := retSingle.Bytes()
	list, err := app.adapter.ToInstance(data)
	if err != nil {
		return nil, err
	}

	resources := []string{}
	if list.HasResources() {
		resources = list.Resources()
	}

	name := list.Name()
	builder := app.builder.
		WithName(name)

	if list.IsUnique() {
		builder.IsUnique()
	}

	resources = append(resources, values...)
	updatedList, err := builder.
		WithResources(resources).
		Now()

	if err != nil {
		return nil, err
	}

	updatedBytes, err := app.adapter.ToBytes(updatedList)
	if err != nil {
		return nil, err
	}

	update, err := app.updateBuilder.Create().
		WithName(name).
		WithSigner(app.signer).
		WithData(updatedBytes).
		Now()

	if err != nil {
		return nil, err
	}

	return app.resourceApp.Update(input, update)
}

// Remove  elements from the current list
func (app *application) Remove(input resources.Resource, values []string) (resources.Resource, error) {
	return nil, nil
}

// Fetch  elements from the current list
func (app *application) Fetch(input resources.Resource, index uint64, amount uint64) (singles.Single, error) {
	return nil, nil
}

// Union the specified list with the current one
func (app *application) Union(input resources.Resource, listName string) (singles.Single, error) {
	return nil, nil
}

// Intersect intersects the specified list with the current one
func (app *application) Intersect(input resources.Resource, listName string) (singles.Single, error) {
	return nil, nil
}

// Delete deletes the current list
func (app *application) Delete(input resources.Resource, delete deletes.Delete) (resources.Resource, error) {
	return nil, nil
}

// Retrieve retrieves the current list
func (app *application) Retrieve(input resources.Resource) (singles.Single, error) {
	return nil, nil
}

// Update rthe current list
func (app *application) Update(input resources.Resource, update updates.Update) (resources.Resource, error) {
	return nil, nil
}

// Commit executes a commit
func (app *application) Commit(input resources.Resource) (transactions.Transactions, error) {
	return nil, nil
}

// Transact replay transactions
func (app *application) Transact(input resources.Resource, trx transactions.Transactions) error {
	return nil
}
