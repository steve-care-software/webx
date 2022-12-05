package tokens

import "github.com/steve-care-software/webx/domain/databases/entities"

type token struct {
	entity   entities.Entity
	reverse  entities.Identifier
	element  Element
	pContent *uint
}

func createToken(
	entity entities.Entity,
	reverse entities.Identifier,
	element Element,
) Token {
	return createTokenInternally(entity, reverse, element, nil)
}

func createTokenWithContentIndex(
	entity entities.Entity,
	reverse entities.Identifier,
	element Element,
	pContent *uint,
) Token {
	return createTokenInternally(entity, reverse, element, pContent)
}

func createTokenInternally(
	entity entities.Entity,
	reverse entities.Identifier,
	element Element,
	pContent *uint,
) Token {
	out := token{
		entity:   entity,
		reverse:  reverse,
		element:  element,
		pContent: pContent,
	}

	return &out
}

// Entity returns the entity
func (obj *token) Entity() entities.Entity {
	return obj.entity
}

// Reverse returns the reverse
func (obj *token) Reverse() entities.Identifier {
	return obj.reverse
}

// Element returns the element
func (obj *token) Element() Element {
	return obj.element
}

// HasContent returns true if there is content, false otherwise
func (obj *token) HasContent() bool {
	return obj.pContent != nil
}

// Content returns the content, if any
func (obj *token) Content() *uint {
	return obj.pContent
}
