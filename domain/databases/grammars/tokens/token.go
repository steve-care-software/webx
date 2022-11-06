package tokens

import "github.com/steve-care-software/webx/domain/databases/entities"

type token struct {
	entity entities.Entity
	lines  entities.Identifiers
	suites entities.Identifiers
}

func createToken(
	entity entities.Entity,
	lines entities.Identifiers,
) Token {
	return createTokenInternally(entity, lines, nil)
}

func createTokenWithSuites(
	entity entities.Entity,
	lines entities.Identifiers,
	suites entities.Identifiers,
) Token {
	return createTokenInternally(entity, lines, suites)
}

func createTokenInternally(
	entity entities.Entity,
	lines entities.Identifiers,
	suites entities.Identifiers,
) Token {
	out := token{
		entity: entity,
		lines:  lines,
		suites: suites,
	}

	return &out
}

// Entity returns the entity
func (obj *token) Entity() entities.Entity {
	return obj.entity
}

// Lines returns the lines
func (obj *token) Lines() entities.Identifiers {
	return obj.lines
}

// HasSuites returns true if there is suites, false otherwise
func (obj *token) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *token) Suites() entities.Identifiers {
	return obj.suites
}
