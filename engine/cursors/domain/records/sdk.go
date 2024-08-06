package records

import "github.com/steve-care-software/webx/engine/cursors/domain/cursors"

// Builder represents a records builder
type Builder interface {
	Create() Builder
	WithList(list []Record) Builder
	Now() (Records, error)
}

// Records represents records
type Records interface {
	List() []Record
	FetchByName(name string) (Record, error)
	FetchExceptName(name string) ([]Record, error)
}

// RecordBuilder represents a record builder
type RecordBuilder interface {
	Create() RecordBuilder
	WithName(name string) RecordBuilder
	WithCursor(cursor cursors.Cursor) RecordBuilder
	Now() (Record, error)
}

// Record represents a record
type Record interface {
	Name() string
	Cursor() cursors.Cursor
}
