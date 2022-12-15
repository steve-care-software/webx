package elements

import (
	"math/rand"
	"time"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// NewElementWithValueForTests creates a new element with value for tests
func NewElementWithValueForTests() Element {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	cardinality := NewCardinalityForTests(true)
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithCardinality(cardinality).WithValue(uint8(45)).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithExternalForTests creates a new element with external for tests
func NewElementWithExternalForTests() Element {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	cardinality := NewCardinalityForTests(false)
	if err != nil {
		panic(err)
	}

	pExternal, err := hash.NewAdapter().FromBytes([]byte("this is an external hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithCardinality(cardinality).WithExternal(*pExternal).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithTokenForTests creates a new element with token for tests
func NewElementWithTokenForTests() Element {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	cardinality := NewCardinalityForTests(true)
	if err != nil {
		panic(err)
	}

	pToken, err := hash.NewAdapter().FromBytes([]byte("this is a token hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithCardinality(cardinality).WithToken(*pToken).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithEverythingForTests creates a new element with everything for tests
func NewElementWithEverythingForTests() Element {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	cardinality := NewCardinalityForTests(false)
	if err != nil {
		panic(err)
	}

	pEverything, err := hash.NewAdapter().FromBytes([]byte("this is an everything hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithCardinality(cardinality).WithEverything(*pEverything).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithRecursiveForTests creates a new element with recursive for tests
func NewElementWithRecursiveForTests() Element {
	pHash, err := hash.NewAdapter().FromBytes([]byte("this is an hash"))
	if err != nil {
		panic(err)
	}

	cardinality := NewCardinalityForTests(true)
	if err != nil {
		panic(err)
	}

	pRecursive, err := hash.NewAdapter().FromBytes([]byte("this is a recursive hash"))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithCardinality(cardinality).WithRecursive(*pRecursive).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCardinalityForTests creates a new cardinality for tests
func NewCardinalityForTests(hasMax bool) Cardinality {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	min := uint(r1.Int())

	builder := NewCardinalityBuilder().Create().WithMin(min)
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
