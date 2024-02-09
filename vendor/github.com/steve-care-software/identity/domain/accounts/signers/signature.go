package signers

import (
	"errors"
	"math/rand"
	"strconv"

	kyber "go.dedis.ch/kyber/v3"
)

type signature struct {
	r PublicKey
	s kyber.Scalar
}

func createSignature(r PublicKey, s kyber.Scalar) (Signature, error) {
	out := signature{r: r, s: s}
	if !out.Verify() {
		return nil, errors.New("the signature could not verify")
	}

	return &out, nil
}

// PublicKey returns the public key of the signature
func (app *signature) PublicKey(msg []byte) (PublicKey, error) {
	// Create a generator.
	g := curve.Point().Base()

	// e = Hash(m || r)
	pubKeyBytes, err := app.r.Bytes()
	if err != nil {
		return nil, err
	}

	combine := []byte{}
	combine = append(combine, []byte(msg)...)
	combine = append(combine, pubKeyBytes...)
	e := createHash(combine)

	// y = (r - s * G) * (1 / e)
	y := curve.Point().Sub(app.r.Point(), curve.Point().Mul(app.s, g))
	y = curve.Point().Mul(curve.Scalar().Div(curve.Scalar().One(), e), y)

	return createPublicKey(y), nil
}

// Verify verifies if the signature has been made by the given public key
func (app *signature) Verify() bool {

	// generate a message:
	msg := []byte(strconv.Itoa(rand.Int()))

	// retrieve pubKey:
	p, err := app.PublicKey(msg)
	if err != nil {
		return false
	}

	// Create a generator.
	g := curve.Point().Base()

	// e = Hash(m || r)
	pubKeyBytes, err := app.r.Bytes()
	if err != nil {
		return false
	}

	combine := []byte{}
	combine = append(combine, []byte(msg)...)
	combine = append(combine, pubKeyBytes...)
	e := createHash(combine)

	// Attempt to reconstruct 's * G' with a provided signature; s * G = r - e * p
	sGv := curve.Point().Sub(app.r.Point(), curve.Point().Mul(e, p.Point()))

	// Construct the actual 's * G'
	sG := curve.Point().Mul(app.s, g)

	// Equality check; ensure signature and public key outputs to s * G.
	return sG.Equal(sGv)
}

// Bytes returns the bytes representation of the signature
func (app *signature) Bytes() ([]byte, error) {
	pubKeyBytes, err := app.r.Bytes()
	if err != nil {
		return nil, err
	}

	sBytes, err := app.s.MarshalBinary()
	if err != nil {
		return nil, err
	}

	output := []byte{}
	output = append(output, pubKeyBytes...)
	output = append(output, []byte(delimiter)...)
	output = append(output, sBytes...)
	return output, nil
}
