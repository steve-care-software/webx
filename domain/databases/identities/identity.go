package identities

import "github.com/steve-care-software/webx/domain/databases/entities"

type identity struct {
	entity        entities.Entity
	modifications entities.Identifiers
}

func createIdentity(
	entity entities.Entity,
	modifications entities.Identifiers,
) Identity {
	out := identity{
		entity:        entity,
		modifications: modifications,
	}

	return &out
}

// Entity retruns the entity
func (obj *identity) Entity() entities.Entity {
	return obj.entity
}

// Modifications retruns the modifications
func (obj *identity) Modifications() entities.Identifiers {
	return obj.modifications
}
