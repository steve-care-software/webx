package lines

import "github.com/steve-care-software/webx/domain/databases/entities"

type line struct {
	entity   entities.Entity
	grammar  entities.Identifier
	elements entities.Identifiers
}

func createLine(
	entity entities.Entity,
	grammar entities.Identifier,
	elements entities.Identifiers,
) Line {
	out := line{
		entity:   entity,
		grammar:  grammar,
		elements: elements,
	}

	return &out
}

// Entity returns the entity
func (obj *line) Entity() entities.Entity {
	return obj.entity
}

// Grammar returns the grammar
func (obj *line) Grammar() entities.Identifier {
	return obj.grammar
}

// Elements returns the elements
func (obj *line) Elements() entities.Identifiers {
	return obj.elements
}
