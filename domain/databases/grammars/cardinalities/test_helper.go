package cardinalities

import (
	"math/rand"
	"time"
)

// NewCardinalityForTests creates a new cardinality for tests
func NewCardinalityForTests(hasMax bool) Cardinality {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	min := uint(r1.Int())

	builder := NewBuilder().Create().WithMin(min)
	if hasMax {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		max := uint(r1.Int()/3) + min
		builder.WithMax(max)
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
