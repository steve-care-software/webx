package elements

import "github.com/steve-care-software/webx/domain/databases/entities"

type element struct {
	entity      entities.Entity
	cardinality entities.Identifier
	content     Content
}

func createElement(
	entity entities.Entity,
	cardinality entities.Identifier,
	content Content,
) Element {
	out := element{
		entity:      entity,
		cardinality: cardinality,
		content:     content,
	}

	return &out
}

// Entity returns the entity
func (obj *element) Entity() entities.Entity {
	return obj.entity
}

// Cardinality returns the cardinality
func (obj *element) Cardinality() entities.Identifier {
	return obj.cardinality
}

// Content returns the content
func (obj *element) Content() Content {
	return obj.content
}
