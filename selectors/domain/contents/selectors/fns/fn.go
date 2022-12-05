package fns

import "github.com/steve-care-software/webx/domain/databases/entities"

type fn struct {
	entity    entities.Entity
	isSingle  bool
	isContent bool
	program   entities.Identifier
	param     uint
}

func createFn(
	entity entities.Entity,
	isSingle bool,
	isContent bool,
	program entities.Identifier,
	param uint,
) Fn {
	out := fn{
		entity:    entity,
		isSingle:  isSingle,
		isContent: isContent,
		program:   program,
		param:     param,
	}

	return &out
}

// Entity returns the entity
func (obj *fn) Entity() entities.Entity {
	return obj.entity
}

// IsSingle returns true if single, false otherwise
func (obj *fn) IsSingle() bool {
	return obj.isSingle
}

// IsContent returns true if content, false otherwise
func (obj *fn) IsContent() bool {
	return obj.isContent
}

// Program returns the program
func (obj *fn) Program() entities.Identifier {
	return obj.program
}

// Param returns the param
func (obj *fn) Param() uint {
	return obj.param
}
