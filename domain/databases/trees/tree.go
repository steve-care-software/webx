package trees

import "github.com/steve-care-software/webx/domain/databases/entities"

type tree struct {
	entity  entities.Entity
	grammar entities.Identifier
	line    entities.Identifier
	suffix  entities.Identifiers
}

func createTree(
	entity entities.Entity,
	grammar entities.Identifier,
	line entities.Identifier,
) Tree {
	return createTreeInternally(entity, grammar, line, nil)
}

func createTreeWithSuffix(
	entity entities.Entity,
	grammar entities.Identifier,
	line entities.Identifier,
	suffix entities.Identifiers,
) Tree {
	return createTreeInternally(entity, grammar, line, suffix)
}

func createTreeInternally(
	entity entities.Entity,
	grammar entities.Identifier,
	line entities.Identifier,
	suffix entities.Identifiers,
) Tree {
	out := tree{
		entity:  entity,
		grammar: grammar,
		line:    line,
		suffix:  suffix,
	}

	return &out
}

// Entity returns the entity
func (obj *tree) Entity() entities.Entity {
	return obj.entity
}

// Grammar returns the grammar
func (obj *tree) Grammar() entities.Identifier {
	return obj.grammar
}

// Line returns the line
func (obj *tree) Line() entities.Identifier {
	return obj.line
}

// HasSuffix returns true if there is a suffix, false otherwise
func (obj *tree) HasSuffix() bool {
	return obj.suffix != nil
}

// Suffix returns the suffix, if any
func (obj *tree) Suffix() entities.Identifiers {
	return obj.suffix
}
