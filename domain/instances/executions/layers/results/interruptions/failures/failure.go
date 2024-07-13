package failures

import "github.com/steve-care-software/historydb/domain/hash"

type failure struct {
	hash            hash.Hash
	index           uint
	code            uint
	isRaisedInLayer bool
	message         string
}

func createFailure(
	hash hash.Hash,
	index uint,
	code uint,
	isRaisedInLayer bool,
) Failure {
	return createFailureInternally(hash, index, code, isRaisedInLayer, "")
}

func createFailureWithMessage(
	hash hash.Hash,
	index uint,
	code uint,
	isRaisedInLayer bool,
	message string,
) Failure {
	return createFailureInternally(hash, index, code, isRaisedInLayer, message)
}

func createFailureInternally(
	hash hash.Hash,
	index uint,
	code uint,
	isRaisedInLayer bool,
	message string,
) Failure {
	out := failure{
		hash:            hash,
		index:           index,
		code:            code,
		isRaisedInLayer: isRaisedInLayer,
		message:         message,
	}

	return &out
}

// Hash returns the hash
func (obj *failure) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *failure) Index() uint {
	return obj.index
}

// Code returns the code
func (obj *failure) Code() uint {
	return obj.code
}

// IsRaisedInLayer returns true if raisedInLayer, false otherwise
func (obj *failure) IsRaisedInLayer() bool {
	return obj.isRaisedInLayer
}

// HasMessage returns true if there is a message, false otherwise
func (obj *failure) HasMessage() bool {
	return obj.message != ""
}

// Message returns the message, if any
func (obj *failure) Message() string {
	return obj.message
}
