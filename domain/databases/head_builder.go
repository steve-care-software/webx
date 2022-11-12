package databases

import (
	"errors"
	"time"

	"github.com/steve-care-software/webx/domain/databases/references"
)

type headBuilder struct {
	name           string
	reference      references.Reference
	pBlockInterval *time.Duration
	pSyncInterval  *time.Duration
	migration      Migration
}

func createHeadBuilder() HeadBuilder {
	out := headBuilder{
		name:           "",
		reference:      nil,
		pBlockInterval: nil,
		pSyncInterval:  nil,
		migration:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *headBuilder) Create() HeadBuilder {
	return createHeadBuilder()
}

// WithName adds a name to the builder
func (app *headBuilder) WithName(name string) HeadBuilder {
	app.name = name
	return app
}

// WithReference adds a reference to the builder
func (app *headBuilder) WithReference(reference references.Reference) HeadBuilder {
	app.reference = reference
	return app
}

// WithBlockInterval adds a blockInterval to the builder
func (app *headBuilder) WithBlockInterval(blockInterval time.Duration) HeadBuilder {
	app.pBlockInterval = &blockInterval
	return app
}

// WithSyncInterval adds a syncInterval to the builder
func (app *headBuilder) WithSyncInterval(syncInterval time.Duration) HeadBuilder {
	app.pSyncInterval = &syncInterval
	return app
}

// WithMigration adds a migration to the builder
func (app *headBuilder) WithMigration(migration Migration) HeadBuilder {
	app.migration = migration
	return app
}

// Now builds a new Head instance
func (app *headBuilder) Now() (Head, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Head instance")
	}

	if app.pBlockInterval == nil {
		return nil, errors.New("the blockInterval is mandatory in order to build a Head instance")
	}

	if app.pSyncInterval == nil {
		return nil, errors.New("the syncInterval is mandatory in order to build a Head instance")
	}

	if app.reference == nil {
		return nil, errors.New("the reference is mandatory in order to build a Database instance")
	}

	if app.migration != nil {
		return createHeadWithMigration(app.name, app.reference, *app.pBlockInterval, *app.pSyncInterval, app.migration), nil
	}

	return createHead(app.name, app.reference, *app.pBlockInterval, *app.pSyncInterval), nil
}
