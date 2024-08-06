package cursors

// Cursors represents cursors
type Cursors interface {
	LIst() []Cursor
}

// CursorFactory represents a cursor factory
type CursorFactory interface {
	Create() (Cursor, error)
}

// Cursor represents a cursor
type Cursor interface {
	HasIdentity() bool
	Identity() string
	HasNamespace() bool
	Namespace() string
	HasVersion() bool
	Version() string
	HasBranches() bool
	Branches() []string
}
