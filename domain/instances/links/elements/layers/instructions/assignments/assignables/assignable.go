package assignables

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography"
)

type assignable struct {
	hash     hash.Hash
	bytes    bytes.Bytes
	constant constants.Constant
	crypto   cryptography.Cryptography
	compiler compilers.Compiler
	query    string
}

func createAssignableWithBytes(
	hash hash.Hash,
	bytes bytes.Bytes,
) Assignable {
	return createAssignableInternally(hash, bytes, nil, nil, nil, "")
}

func createAssignableWithConstant(
	hash hash.Hash,
	constant constants.Constant,
) Assignable {
	return createAssignableInternally(hash, nil, constant, nil, nil, "")
}

func createAssignableWithCryptography(
	hash hash.Hash,
	crypto cryptography.Cryptography,
) Assignable {
	return createAssignableInternally(hash, nil, nil, crypto, nil, "")
}

func createAssignableWithCompiler(
	hash hash.Hash,
	compiler compilers.Compiler,
) Assignable {
	return createAssignableInternally(hash, nil, nil, nil, compiler, "")
}

func createAssignableWithQuery(
	hash hash.Hash,
	query string,
) Assignable {
	return createAssignableInternally(hash, nil, nil, nil, nil, query)
}

func createAssignableInternally(
	hash hash.Hash,
	bytes bytes.Bytes,
	constant constants.Constant,
	crypto cryptography.Cryptography,
	compiler compilers.Compiler,
	query string,
) Assignable {
	out := assignable{
		hash:     hash,
		bytes:    bytes,
		constant: constant,
		crypto:   crypto,
		compiler: compiler,
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

// IsCryptography returns true if there is a cryptography, false otherwise
func (obj *assignable) IsCryptography() bool {
	return obj.crypto != nil
}

// Cryptography returns the cryptography, if any
func (obj *assignable) Cryptography() cryptography.Cryptography {
	return obj.crypto
}

// IsCompiler returns true if there is a compiler, false otherwise
func (obj *assignable) IsCompiler() bool {
	return obj.compiler != nil
}

// Compiler returns the compiler, if any
func (obj *assignable) Compiler() compilers.Compiler {
	return obj.compiler
}

// IsQuery returns true if there is a query, false otherwise
func (obj *assignable) IsQuery() bool {
	return obj.query != ""
}

// Query returns the query, if any
func (obj *assignable) Query() string {
	return obj.query
}
