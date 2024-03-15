package resources

import "github.com/steve-care-software/datastencil/domain/hash"

type kind struct {
	hash       hash.Hash
	native     Native
	reference  []string
	connection string
}

func createKindWithNative(
	hash hash.Hash,
	native Native,
) Kind {
	return createKindInternally(hash, native, nil, "")
}

func createKindWithReference(
	hash hash.Hash,
	reference []string,
) Kind {
	return createKindInternally(hash, nil, reference, "")
}

func createKindWithConnection(
	hash hash.Hash,
	connection string,
) Kind {
	return createKindInternally(hash, nil, nil, connection)
}

func createKindInternally(
	hash hash.Hash,
	native Native,
	reference []string,
	connection string,
) Kind {
	out := kind{
		hash:       hash,
		native:     native,
		reference:  reference,
		connection: connection,
	}

	return &out
}

// Hash returns the hash
func (obj *kind) Hash() hash.Hash {
	return obj.hash
}

// IsNative returns true if there is a native, false otherwise
func (obj *kind) IsNative() bool {
	return obj.native != nil
}

// Native returns the native, if any
func (obj *kind) Native() Native {
	return obj.native
}

// IsReference returns true if there is a reference, false otherwise
func (obj *kind) IsReference() bool {
	return obj.reference != nil
}

// Reference returns the reference, if any
func (obj *kind) Reference() []string {
	return obj.reference
}

// IsConnection returns true if there is a connection, false otherwise
func (obj *kind) IsConnection() bool {
	return obj.connection != ""
}

// Connection returns the connection, if any
func (obj *kind) Connection() string {
	return obj.connection
}
