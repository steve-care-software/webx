package selectors

import "github.com/steve-care-software/webx/domain/databases/entities"

type selector struct {
	entity entities.Entity
	token  entities.Identifier
	inside entities.Identifier
	fn     entities.Identifier
}

func createSelector(
	entity entities.Entity,
	token entities.Identifier,
	inside entities.Identifier,
	fn entities.Identifier,
) Selector {
	out := selector{
		entity: entity,
		token:  token,
		inside: inside,
		fn:     fn,
	}

	return &out
}

// Entity returns the entity
func (obj *selector) Entity() entities.Entity {
	return obj.entity
}

// Token returns the token
func (obj *selector) Token() entities.Identifier {
	return obj.token
}

// Inside returns the inside
func (obj *selector) Inside() entities.Identifier {
	return obj.inside
}

// Func returns the func
func (obj *selector) Func() entities.Identifier {
	return obj.fn
}
