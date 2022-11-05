package fns

import (
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a func builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithProgram(program entities.Identifier) Builder
	WithParam(param uint) Builder
	IsSingle() Builder
	IsContent() Builder
	Now() (Fn, error)
}

// Fn represents a func
type Fn interface {
	Entity() entities.Entity
	IsSingle() bool
	IsContent() bool
	Program() entities.Identifier
	Param() uint
}
