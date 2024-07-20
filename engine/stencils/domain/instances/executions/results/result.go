package results

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/interruptions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/executions/results/success"
)

type result struct {
	hash         hash.Hash
	success      success.Success
	interruption interruptions.Interruption
}

func createResultWithSuccess(
	hash hash.Hash,
	success success.Success,
) Result {
	return createResultInternally(hash, success, nil)
}

func createResultWithInterruption(
	hash hash.Hash,
	interruption interruptions.Interruption,
) Result {
	return createResultInternally(hash, nil, interruption)
}

func createResultInternally(
	hash hash.Hash,
	success success.Success,
	interruption interruptions.Interruption,
) Result {
	out := result{
		hash:         hash,
		success:      success,
		interruption: interruption,
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
func (obj *result) Success() success.Success {
	return obj.success
}

// IsInterruption returns true if there is an interruption, false otherwise
func (obj *result) IsInterruption() bool {
	return obj.interruption != nil
}

// Interruption returns the interruption, if any
func (obj *result) Interruption() interruptions.Interruption {
	return obj.interruption
}
