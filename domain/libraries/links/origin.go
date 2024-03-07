package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins/resources"
)

type origin struct {
	hash     hash.Hash
	resource resources.Resource
	operator operators.Operator
	next     OriginValue
}

func createOrigin(
	hash hash.Hash,
	resource resources.Resource,
	operator operators.Operator,
	next OriginValue,
) Origin {
	out := origin{
		hash:     hash,
		resource: resource,
		operator: operator,
		next:     next,
	}

	return &out
}

// Hash returns the hash
func (obj *origin) Hash() hash.Hash {
	return obj.hash
}

// Resource returns the resource
func (obj *origin) Resource() resources.Resource {
	return obj.resource
}

// Operator returns the operator
func (obj *origin) Operator() operators.Operator {
	return obj.operator
}

// Next returns the next value
func (obj *origin) Next() OriginValue {
	return obj.next
}
