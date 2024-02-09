package results

import "github.com/steve-care-software/identity/domain/hash"

type result struct {
	hash    hash.Hash
	success Success
	failure Failure
}

func createResultWithSuccess(
	hash hash.Hash,
	success Success,
) Result {
	return createResultInternally(hash, success, nil)
}

func createResultWithFailure(
	hash hash.Hash,
	failure Failure,
) Result {
	return createResultInternally(hash, nil, failure)
}

func createResultInternally(
	hash hash.Hash,
	success Success,
	failure Failure,
) Result {
	out := result{
		hash:    hash,
		success: success,
		failure: failure,
	}

	return &out
}

// Hash returns the hash
func (obj *result) Hash() hash.Hash {
	return obj.hash
}

// IsSuccess returns true if there is a success, false otherwise
func (obj *result) IsSuccess() bool {
	return obj.success != nil
}

// Success returns the success, if any
func (obj *result) Success() Success {
	return obj.success
}

// IsFailure returns true if there is a failure, false otherwise
func (obj *result) IsFailure() bool {
	return obj.failure != nil
}

// Failure returns the failure, if any
func (obj *result) Failure() Failure {
	return obj.failure
}
