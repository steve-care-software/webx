package modifications

import (
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/cryptography/signatures"
)

type content struct {
	name string
	sig  signatures.PrivateKey
	enc  keys.PrivateKey
}

func createContentWithName(
	name string,
) Content {
	return createContentInternally(name, nil, nil)
}

func createContentWithSignature(
	sig signatures.PrivateKey,
) Content {
	return createContentInternally("", sig, nil)
}

func createContentWithEncryption(
	enc keys.PrivateKey,
) Content {
	return createContentInternally("", nil, enc)
}

func createContentWithNameAndSignature(
	name string,
	sig signatures.PrivateKey,
) Content {
	return createContentInternally(name, sig, nil)
}

func createContentWithNameAndEncryption(
	name string,
	enc keys.PrivateKey,
) Content {
	return createContentInternally(name, nil, enc)
}

func createContentWithSignatureAndEncryption(
	sig signatures.PrivateKey,
	enc keys.PrivateKey,
) Content {
	return createContentInternally("", sig, enc)
}

func createContentWithNameAndSignatureAndEncryption(
	name string,
	sig signatures.PrivateKey,
	enc keys.PrivateKey,
) Content {
	return createContentInternally(name, sig, enc)
}

func createContentInternally(
	name string,
	sig signatures.PrivateKey,
	enc keys.PrivateKey,
) Content {
	out := content{
		name: name,
		sig:  sig,
		enc:  enc,
	}

	return &out
}

// HasName returns true if there is a name, false otherwise
func (obj *content) HasName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *content) Name() string {
	return obj.name
}

// HasSignature returns true if there is a signature pk, false otherwise
func (obj *content) HasSignature() bool {
	return obj.sig != nil
}

// Signature returns the signature pk, if any
func (obj *content) Signature() signatures.PrivateKey {
	return obj.sig
}

// HasEncryption returns true if there is an encryption pk, false otherwise
func (obj *content) HasEncryption() bool {
	return obj.enc != nil
}

// Encryption returns the encryption pk, if any
func (obj *content) Encryption() keys.PrivateKey {
	return obj.enc
}
