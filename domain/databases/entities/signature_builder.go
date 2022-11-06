package entities

import (
	"errors"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"go.dedis.ch/kyber/v3"
)

type signatureBuilder struct {
	ring []hash.Hash
	s    kyber.Scalar
	e    kyber.Scalar
}

func createSignatureBuilder() SignatureBuilder {
	out := signatureBuilder{
		ring: nil,
		s:    nil,
		e:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *signatureBuilder) Create() SignatureBuilder {
	return createSignatureBuilder()
}

// WithRing adds a ring to the builder
func (app *signatureBuilder) WithRing(ring []hash.Hash) SignatureBuilder {
	app.ring = ring
	return app
}

// WithS adds a scalar s to the builder
func (app *signatureBuilder) WithS(s kyber.Scalar) SignatureBuilder {
	app.s = s
	return app
}

// WithE adds a scalar e to the builder
func (app *signatureBuilder) WithE(e kyber.Scalar) SignatureBuilder {
	app.e = e
	return app
}

// Now builds a new Signature instance
func (app *signatureBuilder) Now() (Signature, error) {
	if app.s == nil {
		return nil, errors.New("the scalar s is mandatory in order to build a Signature instance")
	}

	if app.e == nil {
		return nil, errors.New("the scalar e is mandatory in order to build a Signature instance")
	}

	if app.ring != nil && len(app.ring) <= 0 {
		app.ring = nil
	}

	if app.ring == nil {
		return nil, errors.New("the ring is mandatory in order to build a Signature instance")
	}

	return createSignature(app.ring, app.s, app.e), nil
}
