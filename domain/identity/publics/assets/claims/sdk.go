package claims

import (
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets/units/genesis"
	"github.com/steve-care-software/syntax/domain/syntax/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
)

// Claim represents a claim
type Claim interface {
	Hash() hash.Hash
	Value() Value
	Signature() signatures.RingSignature
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithGenesis(genesis genesis.Genesis) ValueBuilder
	WithGramar(grammar grammars.Grammar) ValueBuilder
	WithCriteria(criteria criterias.Criteria) ValueBuilder
	WithCompiler(compiler compilers.Compiler) ValueBuilder
	Now() (Value, error)
}

// Value represents a claim value
type Value interface {
	Genesis() genesis.Genesis
	Content() Content
}

// Content represents the content of a claim
type Content interface {
	IsGrammar() bool
	Grammar() grammars.Grammar
	IsCriteria() bool
	Criteria() criterias.Criteria
	IsCompiler() bool
	Compiler() compilers.Compiler
}
