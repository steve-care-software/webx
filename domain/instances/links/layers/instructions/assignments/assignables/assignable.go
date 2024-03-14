package assignables

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries"
)

type assignable struct {
	hash     hash.Hash
	bytes    bytes.Bytes
	constant constants.Constant
	account  accounts.Account
	crypto   cryptography.Cryptography
	library  libraries.Library
	query    string
}

func createAssignableWithBytes(
	hash hash.Hash,
	bytes bytes.Bytes,
) Assignable {
	return createAssignableInternally(hash, bytes, nil, nil, nil, nil, "")
}

func createAssignableWithConstant(
	hash hash.Hash,
	constant constants.Constant,
) Assignable {
	return createAssignableInternally(hash, nil, constant, nil, nil, nil, "")
}

func createAssignableWithAccount(
	hash hash.Hash,
	account accounts.Account,
) Assignable {
	return createAssignableInternally(hash, nil, nil, account, nil, nil, "")
}

func createAssignableWithCryptography(
	hash hash.Hash,
	crypto cryptography.Cryptography,
) Assignable {
	return createAssignableInternally(hash, nil, nil, nil, crypto, nil, "")
}

func createAssignableWithLibrary(
	hash hash.Hash,
	library libraries.Library,
) Assignable {
	return createAssignableInternally(hash, nil, nil, nil, nil, library, "")
}

func createAssignableWithQuery(
	hash hash.Hash,
	query string,
) Assignable {
	return createAssignableInternally(hash, nil, nil, nil, nil, nil, query)
}

func createAssignableInternally(
	hash hash.Hash,
	bytes bytes.Bytes,
	constant constants.Constant,
	account accounts.Account,
	crypto cryptography.Cryptography,
	library libraries.Library,
	query string,
) Assignable {
	out := assignable{
		hash:     hash,
		bytes:    bytes,
		constant: constant,
		account:  account,
		crypto:   crypto,
		library:  library,
		query:    query,
	}

	return &out
}

// Hash returns the hash
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *assignable) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *assignable) Bytes() bytes.Bytes {
	return obj.bytes
}

// IsConstant returns true if there is constant, false otherwise
func (obj *assignable) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *assignable) Constant() constants.Constant {
	return obj.constant
}

// IsAccount returns true if there is an account, false otherwise
func (obj *assignable) IsAccount() bool {
	return obj.account != nil
}

// Account returns the account, if any
func (obj *assignable) Account() accounts.Account {
	return obj.account
}

// IsCryptography returns true if there is a cryptography, false otherwise
func (obj *assignable) IsCryptography() bool {
	return obj.crypto != nil
}

// Cryptography returns the cryptography, if any
func (obj *assignable) Cryptography() cryptography.Cryptography {
	return obj.crypto
}

// IsLibrary returns true if there is a library, false otherwise
func (obj *assignable) IsLibrary() bool {
	return obj.library != nil
}

// Library returns the library, if any
func (obj *assignable) Library() libraries.Library {
	return obj.library
}

// IsQuery returns true if there is a query, false otherwise
func (obj *assignable) IsQuery() bool {
	return obj.query != ""
}

// Query returns the query, if any
func (obj *assignable) Query() string {
	return obj.query
}
