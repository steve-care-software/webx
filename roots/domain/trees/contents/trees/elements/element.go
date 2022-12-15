package elements

import "github.com/steve-care-software/webx/domain/databases/entities"

type element struct {
	entity   entities.Entity
	grammar  entities.Identifier
	contents entities.Identifiers
}

func createElement(
	entity entities.Entity,
	grammar entities.Identifier,
	contents entities.Identifiers,
) Element {
	out := element{
		entity:   entity,
		grammar:  grammar,
		contents: contents,
	}

	return &out
}

// Entity returns the entity
func (obj *element) Entity() entities.Entity {
	return obj.entity
}

// Grammar returns the grammar
func (obj *element) Grammar() entities.Identifier {
	return obj.grammar
}

// Contents returns the contents
func (obj *element) Contents() entities.Identifiers {
	return obj.contents
}
