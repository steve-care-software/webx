package interruptions

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results/interruptions/failures"
	"github.com/steve-care-software/historydb/domain/hash"
)

type interruption struct {
	hash        hash.Hash
	pStopAtLine *uint
	failure     failures.Failure
}

func createInterruptionWithStop(
	hash hash.Hash,
	pStopAtLine *uint,
) Interruption {
	return createInterruptionInternally(hash, pStopAtLine, nil)
}

func createInterruptionWithFailure(
	hash hash.Hash,
	failure failures.Failure,
) Interruption {
	return createInterruptionInternally(hash, nil, failure)
}

func createInterruptionInternally(
	hash hash.Hash,
	pStopAtLine *uint,
	failure failures.Failure,
) Interruption {
	out := interruption{
		hash:        hash,
		pStopAtLine: pStopAtLine,
		failure:     failure,
	}

	return &out
}

// Hash returns the hash
func (obj *interruption) Hash() hash.Hash {
	return obj.hash
}

// IsStop returns true if stop, false otherwise
func (obj *interruption) IsStop() bool {
	return obj.pStopAtLine != nil
}

// Stop returns the stop line, if any
func (obj *interruption) Stop() *uint {
	return obj.pStopAtLine
}

// IsFailure returns true if failure, false otherwise
func (obj *interruption) IsFailure() bool {
	return obj.failure != nil
}

// Failure returns the failure, if any
func (obj *interruption) Failure() failures.Failure {
	return obj.failure
}
