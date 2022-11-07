package routers

import "github.com/steve-care-software/webx/domain/databases/entities"

type route struct {
	entity    entities.Entity
	grammar   entities.Identifier
	selectors entities.Identifiers
	program   entities.Identifier
}

func createRoute(
	entity entities.Entity,
	grammar entities.Identifier,
	selectors entities.Identifiers,
	program entities.Identifier,
) Route {
	out := route{
		entity:    entity,
		grammar:   grammar,
		selectors: selectors,
		program:   program,
	}

	return &out
}

// Entity returns the entity
func (obj *route) Entity() entities.Entity {
	return obj.entity
}

// Grammar returns the grammar
func (obj *route) Grammar() entities.Identifier {
	return obj.grammar
}

// Selectors returns the selectors
func (obj *route) Selectors() entities.Identifiers {
	return obj.selectors
}

// Program returns the program
func (obj *route) Program() entities.Identifier {
	return obj.program
}
