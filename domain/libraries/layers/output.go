package layers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type output struct {
	hash     hash.Hash
	variable string
	kind     Kind
	execute  string
}

func createOutput(
	hash hash.Hash,
	variable string,
	kind Kind,
) Output {
	return createOutputInternally(hash, variable, kind, "")
}

func createOutputWithExecute(
	hash hash.Hash,
	variable string,
	kind Kind,
	execute string,
) Output {
	return createOutputInternally(hash, variable, kind, execute)
}

func createOutputInternally(
	hash hash.Hash,
	variable string,
	kind Kind,
	execute string,
) Output {
	out := output{
		hash:     hash,
		variable: variable,
		kind:     kind,
		execute:  execute,
	}

	return &out
}

// Hash returns the hash
func (obj *output) Hash() hash.Hash {
	return obj.hash
}

// Variable returns the variable
func (obj *output) Variable() string {
	return obj.variable
}

// Kind returns the kind
func (obj *output) Kind() Kind {
	return obj.kind
}

// HasExecute returns true if there is an execute, false otherwise
func (obj *output) HasExecute() bool {
	return obj.execute != ""
}

// Execute returns the execute, if any
func (obj *output) Execute() string {
	return obj.execute
}
