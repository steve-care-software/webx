package queries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions"
)

type query struct {
	hash      hash.Hash
	entity    string
	condition conditions.Condition
	fields    []string
}

func createQuery(
	hash hash.Hash,
	entity string,
	condition conditions.Condition,
) Query {
	return createQueryInternally(
		hash,
		entity,
		condition,
		nil,
	)
}

func createQueryWithFields(
	hash hash.Hash,
	entity string,
	condition conditions.Condition,
	fields []string,
) Query {
	return createQueryInternally(
		hash,
		entity,
		condition,
		fields,
	)
}

func createQueryInternally(
	hash hash.Hash,
	entity string,
	condition conditions.Condition,
	fields []string,
) Query {
	out := query{
		hash:      hash,
		entity:    entity,
		condition: condition,
		fields:    fields,
	}

	return &out
}

// Hash returns the hash
func (obj *query) Hash() hash.Hash {
	return obj.hash
}

// Entity returns the entity
func (obj *query) Entity() string {
	return obj.entity
}

// Condition returns the condition
func (obj *query) Condition() conditions.Condition {
	return obj.condition
}

// HasFields returns true if there is fields, false otherwise
func (obj *query) HasFields() bool {
	return obj.fields != nil
}

// Fields returns the fields, if any
func (obj *query) Fields() []string {
	return obj.fields
}
