package databases

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/programs"
)

type migrationBuilder struct {
	hashAdapter hash.Adapter
	previous    Content
	pHeight     *uint
	description string
	program     programs.Program
}

func createMigrationBuilder(
	hashAdapter hash.Adapter,
) MigrationBuilder {
	out := migrationBuilder{
		hashAdapter: hashAdapter,
		previous:    nil,
		pHeight:     nil,
		description: "",
		program:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *migrationBuilder) Create() MigrationBuilder {
	return createMigrationBuilder(app.hashAdapter)
}

// WithPrevious adds a previous content to the builder
func (app *migrationBuilder) WithPrevious(previous Content) MigrationBuilder {
	app.previous = previous
	return app
}

// WithHeight adds an height  to the builder
func (app *migrationBuilder) WithHeight(height uint) MigrationBuilder {
	app.pHeight = &height
	return app
}

// WithDescription adds a description  to the builder
func (app *migrationBuilder) WithDescription(description string) MigrationBuilder {
	app.description = description
	return app
}

// WithProgram adds a program  to the builder
func (app *migrationBuilder) WithProgram(program programs.Program) MigrationBuilder {
	app.program = program
	return app
}

// Now builds a new Migration instance
func (app *migrationBuilder) Now() (Migration, error) {
	if app.previous == nil {
		return nil, errors.New("the previous content is mandatory in order to build a Migration instance")
	}

	if app.pHeight == nil {
		return nil, errors.New("the height is mandatory in order to build a Migration instance")
	}

	data := [][]byte{
		app.previous.Hash().Bytes(),
		[]byte(strconv.Itoa(int(*app.pHeight))),
	}

	if app.description != "" {
		data = append(data, []byte(app.description))
	}

	if app.program != nil {
		data = append(data, app.program.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.program != nil {
		return createMigrationWithProgram(*pHash, app.previous, *app.pHeight, app.description, app.program), nil
	}

	return createMigration(*pHash, app.previous, *app.pHeight, app.description), nil
}
