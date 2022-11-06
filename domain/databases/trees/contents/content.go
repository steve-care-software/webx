package contents

import "github.com/steve-care-software/webx/domain/databases/entities"

type content struct {
	entity entities.Entity
	value  Value
	prefix entities.Identifier
}

func createContent(
	entity entities.Entity,
	value Value,
) Content {
	return createContentInternally(entity, value, nil)
}

func createContentWithPrefix(
	entity entities.Entity,
	value Value,
	prefix entities.Identifier,
) Content {
	return createContentInternally(entity, value, prefix)
}

func createContentInternally(
	entity entities.Entity,
	value Value,
	prefix entities.Identifier,
) Content {
	out := content{
		entity: entity,
		value:  value,
		prefix: prefix,
	}

	return &out
}

// Entity returns the entity
func (obj *content) Entity() entities.Entity {
	return obj.entity
}

// Value returns the value
func (obj *content) Value() Value {
	return obj.value
}

// HasPrefix returns true if there is a prefix, false otherwise
func (obj *content) HasPrefix() bool {
	return obj.prefix != nil
}

// Prefix returns the prefix, if any
func (obj *content) Prefix() entities.Identifier {
	return obj.prefix
}
