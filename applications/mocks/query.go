package mocks

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/queries"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/queries/conditions"
)

type query struct {
}

func createQuery() queries.Query {
	out := query{}
	return &out
}

// Hash returns the hash
func (obj *query) Hash() hash.Hash {
	return nil
}

// Entity returns the entity
func (obj *query) Entity() string {
	return ""
}

// Condition returns the condition
func (obj *query) Condition() conditions.Condition {
	return nil
}

// HasFields returns true if there is fields, false otherwise
func (obj *query) HasFields() bool {
	return false
}

// Fields returns fields, if any
func (obj *query) Fields() []string {
	return nil
}
