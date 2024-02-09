package layers

import "github.com/steve-care-software/identity/domain/hash"

type identity struct {
	hash      hash.Hash
	signer    Signer
	encryptor Encryptor
}

func createIdentityWithSigner(
	hash hash.Hash,
	signer Signer,
) Identity {
	return createIdentityInternally(hash, signer, nil)
}

func createIdentityWithEncryptor(
	hash hash.Hash,
	encryptor Encryptor,
) Identity {
	return createIdentityInternally(hash, nil, encryptor)
}

func createIdentityInternally(
	hash hash.Hash,
	signer Signer,
	encryptor Encryptor,
) Identity {
	out := identity{
		hash:      hash,
		signer:    signer,
		encryptor: encryptor,
	}

	return &out
}

// Hash returns the hash
func (obj *identity) Hash() hash.Hash {
	return obj.hash
}

// IsSigner returns true if there is a signer, false otherwise
func (obj *identity) IsSigner() bool {
	return obj.signer != nil
}

// Signer returns the signer, if any
func (obj *identity) Signer() Signer {
	return obj.signer
}

// IsEncryptor returns true if there is an encryptor, false otherwise
func (obj *identity) IsEncryptor() bool {
	return obj.encryptor != nil
}

// Encryptor returns the encryptor, if any
func (obj *identity) Encryptor() Encryptor {
	return obj.encryptor
}
