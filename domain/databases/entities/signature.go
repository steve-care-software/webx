package entities

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"go.dedis.ch/kyber/v3"
)

type signature struct {
	ring []hash.Hash
	s    kyber.Scalar
	e    kyber.Scalar
}

func createSignature(
	ring []hash.Hash,
	s kyber.Scalar,
	e kyber.Scalar,
) Signature {
	out := signature{
		ring: ring,
		s:    s,
		e:    e,
	}

	return &out
}

// Ring returns the ring
func (obj *signature) Ring() []hash.Hash {
	return obj.ring
}

// S returns the s scalar
func (obj *signature) S() kyber.Scalar {
	return obj.s
}

// E returns the e scalar
func (obj *signature) E() kyber.Scalar {
	return obj.e
}
