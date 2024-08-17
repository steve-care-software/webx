package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/asts"
	"github.com/steve-care-software/webx/engine/domain/grammars"
	"github.com/steve-care-software/webx/engine/domain/grammars/rules"
)

const ulA = "A"
const ulB = "B"
const ulC = "C"
const ulD = "D"
const ulE = "E"
const ulF = "F"
const ulG = "G"
const ulH = "H"
const ulI = "I"
const ulJ = "J"
const ulK = "K"
const ulL = "L"
const ulM = "M"
const ulN = "N"
const ulO = "O"
const ulP = "P"
const ulQ = "Q"
const ulR = "R"
const ulS = "S"
const ulT = "T"
const ulU = "U"
const ulV = "V"
const ulW = "W"
const ulX = "X"
const ulY = "Y"
const ulZ = "Z"

const ruleValueEscape = "\\"
const ruleValuePrefix = "\""
const ruleValueSuffix = "\""
const ruleNameSeparator = "_"
const ruleNameValueSeparator = ":"

// NewApplication creates a new application
func NewApplication() Application {
	ruleBuilder := rules.NewRuleBuilder()
	possibleRuleNameCharacters := createPossibleRuleNameCharactersList()
	return createApplication(
		ruleBuilder,
		[]byte(ruleNameValueSeparator)[0],
		possibleRuleNameCharacters,
		[]byte(ruleNameSeparator)[0],
		[]byte(ruleValuePrefix)[0],
		[]byte(ruleValueSuffix)[0],
		[]byte(ruleValueEscape)[0],
	)
}

// Application represents the grammar application
type Application interface {
	Parse(lexedInput []byte) (grammars.Grammar, error)
	Compile(grammar grammars.Grammar) (asts.AST, error)
	Decompile(ast asts.AST) (grammars.Grammar, error)
	Compose(grammar grammars.Grammar) ([]byte, error)
}
