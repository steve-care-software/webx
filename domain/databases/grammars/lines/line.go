package lines

import "github.com/steve-care-software/webx/domain/databases/entities"

type line struct {
	entity   entities.Entity
	elements entities.Identifiers
}

func createLine(
	entity entities.Entity,
	elements entities.Identifiers,
) Line {
	out := line{
		entity:   entity,
		elements: elements,
	}

	return &out
}

// Entity returns the entity
func (obj *line) Entity() entities.Entity {
	return obj.entity
}

// Elements returns the elements
func (obj *line) Elements() entities.Identifiers {
	return obj.elements
}
