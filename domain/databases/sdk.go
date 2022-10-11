package databases

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
)

// Database represents a database
type Database interface {
	ID() uuid.UUID
	Name() string
	Grammar() grammars.Grammar
}

// Repository represents a database repository
type Repository interface {
	List() []string
	RetrieveByID(id uuid.UUID) (Database, error)
	RetrieveByName(name string) (Database, error)
}

// Service represents a database service
type Service interface {
	Save(db Database) error
	Delete(db Database) error
}
