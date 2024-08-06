package cursors

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/blockchains"
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/identities"
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/namespaces"
)

// Builder represents a cursors builder
type Builder interface {
	Create() Builder
	WithList(list []Cursor) Builder
	Now() (Cursors, error)
}

// Cursors represents cursors
type Cursors interface {
	List() []Cursor
}

// CursorFactory represents a cursor factory
type CursorFactory interface {
	Create() (Cursor, error)
}

// CursorBuilder represents a cursor builder
type CursorBuilder interface {
	Create() CursorBuilder
	WithIdentity(identity identities.Identity) CursorBuilder
	WIthNamespace(namespace namespaces.Namespace) CursorBuilder
	WithBlockchain(blockchain blockchains.Blockchain) CursorBuilder
	WithInitialCursor(cursor Cursor) CursorBuilder
	Now() (Cursor, error)
}

// Cursor represents a cursor
type Cursor interface {
	HasIdentity() bool
	Identity() identities.Identity
	HasNamespace() bool
	Namespace() namespaces.Namespace
	HasBlockchain() bool
	Blockchain() blockchains.Blockchain
}
