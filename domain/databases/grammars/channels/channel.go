package channels

import "github.com/steve-care-software/webx/domain/databases/entities"

type channel struct {
	entity entities.Entity
	token  entities.Identifier
	prev   entities.Identifier
	next   entities.Identifier
}

func createChannel(
	entity entities.Entity,
	token entities.Identifier,
) Channel {
	return createChannelInternally(entity, token, nil, nil)
}

func createChannelWithPrevious(
	entity entities.Entity,
	token entities.Identifier,
	prev entities.Identifier,
) Channel {
	return createChannelInternally(entity, token, prev, nil)
}

func createChannelWithNext(
	entity entities.Entity,
	token entities.Identifier,
	next entities.Identifier,
) Channel {
	return createChannelInternally(entity, token, nil, next)
}

func createChannelWithPreviousAndNext(
	entity entities.Entity,
	token entities.Identifier,
	prev entities.Identifier,
	next entities.Identifier,
) Channel {
	return createChannelInternally(entity, token, prev, next)
}

func createChannelInternally(
	entity entities.Entity,
	token entities.Identifier,
	prev entities.Identifier,
	next entities.Identifier,
) Channel {
	out := channel{
		entity: entity,
		token:  token,
		prev:   prev,
		next:   next,
	}

	return &out
}

// Entity returns the entity
func (obj *channel) Entity() entities.Entity {
	return obj.entity
}

// Token returns the token
func (obj *channel) Token() entities.Identifier {
	return obj.token
}

// HasPrevious returns true if there is a previous, false otherwise
func (obj *channel) HasPrevious() bool {
	return obj.prev != nil
}

// Previous returns the previous, if any
func (obj *channel) Previous() entities.Identifier {
	return obj.prev
}

// HasNext returns true if there is a next, false otherwise
func (obj *channel) HasNext() bool {
	return obj.next != nil
}

// Next returns the next, if any
func (obj *channel) Next() entities.Identifier {
	return obj.next
}
