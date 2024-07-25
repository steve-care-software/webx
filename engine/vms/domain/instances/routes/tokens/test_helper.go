package tokens

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens/cardinalities"
)

// NewTokensForTests creates tokens for tests
func NewTokensForTests(list []Token) Tokens {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTokenWithOmissionForTests creates token with omission for tests
func NewTokenWithOmissionForTests(elements elements.Elements, cardinality cardinalities.Cardinality, omission omissions.Omission) Token {
	ins, err := NewTokenBuilder().Create().WithElements(elements).WithCardinality(cardinality).WithOmission(omission).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTokenForTests creates token for tests
func NewTokenForTests(elements elements.Elements, cardinality cardinalities.Cardinality) Token {
	ins, err := NewTokenBuilder().Create().WithElements(elements).WithCardinality(cardinality).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
