package databases

import (
	"errors"
	"time"
)

type headBuilder struct {
	name           string
	sections       Sections
	pBlockInterval *time.Duration
	pSyncInterval  *time.Duration
	migration      Migration
}

func createHeadBuilder() HeadBuilder {
	out := headBuilder{
		name:           "",
		sections:       nil,
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

// WithSections adds a sections to the builder
func (app *headBuilder) WithSections(sections Sections) HeadBuilder {
	app.sections = sections
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

	if app.sections == nil {
		return nil, errors.New("the sections is mandatory in order to build a Head instance")
	}

	if app.pBlockInterval == nil {
		return nil, errors.New("the blockInterval is mandatory in order to build a Head instance")
	}

	if app.pSyncInterval == nil {
		return nil, errors.New("the syncInterval is mandatory in order to build a Head instance")
	}

	if app.migration != nil {
		return createHeadWithMigration(app.name, app.sections, *app.pBlockInterval, *app.pSyncInterval, app.migration), nil
	}

	return createHead(app.name, app.sections, *app.pBlockInterval, *app.pSyncInterval), nil
}
