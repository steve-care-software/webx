package conditions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions/operators"
)

type comparison struct {
	hash      hash.Hash
	operator  operators.Operator
	condition Condition
}

func createComparison(
	hash hash.Hash,
	operator operators.Operator,
	condition Condition,
) Comparison {
	out := comparison{
		hash:      hash,
		operator:  operator,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *comparison) Hash() hash.Hash {
	return obj.hash
}

// Operator returns the operator
func (obj *comparison) Operator() operators.Operator {
	return obj.operator
}

// Condition returns the condition
func (obj *comparison) Condition() Condition {
	return obj.condition
}
