package everythings

import "github.com/steve-care-software/webx/domain/databases/entities"

type everyting struct {
	entity    entities.Entity
	exception entities.Identifier
	escape    entities.Identifier
}

func createEverything(
	entity entities.Entity,
	exception entities.Identifier,
) Everything {
	return createEverythingInternally(entity, exception, nil)
}

func createEverythingWithEscape(
	entity entities.Entity,
	exception entities.Identifier,
	escape entities.Identifier,
) Everything {
	return createEverythingInternally(entity, exception, escape)
}

func createEverythingInternally(
	entity entities.Entity,
	exception entities.Identifier,
	escape entities.Identifier,
) Everything {
	out := everyting{
		entity:    entity,
		exception: exception,
		escape:    escape,
	}

	return &out
}

// Entity returns the entity
func (obj *everyting) Entity() entities.Entity {
	return obj.entity
}

// Exception returns the exception
func (obj *everyting) Exception() entities.Identifier {
	return obj.exception
}

// HasEscape returns true if there is an escape, false otherwise
func (obj *everyting) HasEscape() bool {
	return obj.escape != nil
}

// Escape returns the escape, if any
func (obj *everyting) Escape() entities.Identifier {
	return obj.escape
}
