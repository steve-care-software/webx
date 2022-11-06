package grammars

import "github.com/steve-care-software/webx/domain/databases/entities"

type grammar struct {
	entity   entities.Entity
	root     entities.Identifier
	channels entities.Identifiers
}

func createGrammar(
	entity entities.Entity,
	root entities.Identifier,
) Grammar {
	return createGrammarInternally(entity, root, nil)
}

func createGrammarWithChannels(
	entity entities.Entity,
	root entities.Identifier,
	channels entities.Identifiers,
) Grammar {
	return createGrammarInternally(entity, root, channels)
}

func createGrammarInternally(
	entity entities.Entity,
	root entities.Identifier,
	channels entities.Identifiers,
) Grammar {
	out := grammar{
		entity:   entity,
		root:     root,
		channels: channels,
	}

	return &out
}

// Entity returns the entity
func (obj *grammar) Entity() entities.Entity {
	return obj.entity
}

// Root returns the root
func (obj *grammar) Root() entities.Identifier {
	return obj.root
}

// HasChannels returns true if there is channels, false otherwise
func (obj *grammar) HasChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels, if any
func (obj *grammar) Channels() entities.Identifiers {
	return obj.channels
}
