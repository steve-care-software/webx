package interpreters

import (
	"fmt"
	"strconv"
)

const (
	// NZero represents a zero
	NZero (uint8) = iota

	// NOne represents a one
	NOne

	// NTwo represents a two
	NTwo

	// NThree represents a tree
	NThree

	// NFour represents a four
	NFour

	// NFive represents a five
	NFive

	// NSix represents a six
	NSix

	// NSeven represents a seven
	NSeven

	// NHeight represents an height
	NHeight

	// NNine represents a nine
	NNine

	// LLA represents lowercase letter a
	LLA

	// LLB represents lowercase letter b
	LLB

	// LLC represents lowercase letter c
	LLC

	// LLD represents lowercase letter d
	LLD

	// LLE represents lowercase letter e
	LLE

	// LLF represents lowercase letter f
	LLF

	// LLG represents lowercase letter g
	LLG

	// LLH represents lowercase letter h
	LLH

	// LLI represents lowercase letter i
	LLI

	// LLJ represents lowercase letter j
	LLJ

	// LLK represents lowercase letter k
	LLK

	// LLL represents lowercase letter l
	LLL

	// LLM represents lowercase letter m
	LLM

	// LLN represents lowercase letter n
	LLN

	// LLO represents lowercase letter o
	LLO

	// LLP represents lowercase letter p
	LLP

	// LLQ represents lowercase letter q
	LLQ

	// LLR represents lowercase letter r
	LLR

	// LLS represents lowercase letter s
	LLS

	// LLT represents lowercase letter t
	LLT

	// LLU represents lowercase letter u
	LLU

	// LLV represents lowercase letter v
	LLV

	// LLW represents lowercase letter w
	LLW

	// LLX represents lowercase letter x
	LLX

	// LLY represents lowercase letter y
	LLY

	// LLZ represents lowercase letter z
	LLZ

	// ULA represents uppercase letter a
	ULA

	// ULB represents uppercase letter b
	ULB

	// ULC represents uppercase letter c
	ULC

	// ULD represents uppercase letter d
	ULD

	// ULE represents uppercase letter e
	ULE

	// ULF represents uppercase letter f
	ULF

	// ULG represents uppercase letter g
	ULG

	// ULH represents uppercase letter h
	ULH

	// ULI represents uppercase letter i
	ULI

	// ULJ represents uppercase letter j
	ULJ

	// ULK represents uppercase letter k
	ULK

	// ULL represents uppercase letter l
	ULL

	// ULM represents uppercase letter m
	ULM

	// ULN represents uppercase letter n
	ULN

	// ULO represents uppercase letter o
	ULO

	// ULP represents uppercase letter p
	ULP

	// ULQ represents uppercase letter q
	ULQ

	// ULR represents uppercase letter r
	ULR

	// ULS represents uppercase letter s
	ULS

	// ULT represents uppercase letter t
	ULT

	// ULU represents uppercase letter u
	ULU

	// ULV represents uppercase letter v
	ULV

	// ULW represents uppercase letter w
	ULW

	// ULX represents uppercase letter x
	ULX

	// ULY represents uppercase letter y
	ULY

	// ULZ represents uppercase letter z
	ULZ

	// SPlus represents a plus sign
	SPlus

	// SMinus represents a minus sign
	SMinus

	// SEqual represents an equal sign
	SEqual

	// SSemiColon represents a semi-colon sign
	SSemiColon
)

type lineExecuteFn func(map[string]string) (any, error)
type blockExecuteFn func(map[string]string, *value) error

// NewApplication creates a new application
func NewApplication() Application {
	out := application{newGrammar(
		"bRoot",
		[]*block{
			newBlockWithExecution(
				"bRoot",
				[]*line{
					newLine([]*variable{
						newVariable("variableName", "eVariableName"),
						newVariable("equal", "eEqual"),
						newVariable("value", "eOperation"),
					}),
				},
				func(variables map[string]string, input *value) error {
					fmt.Printf("\n_______%s,%d\n", variables["variableName"], input.variables["value"].block[0].retExecLineValue)
					return nil
				},
			),
			newBlock(
				"bOperations",
				[]*line{
					newLineWithExecution([]*variable{
						newVariable("first", "eAnyNumber"),
						newVariable("plus", "ePlus"),
						newVariable("second", "eAnyNumber"),
					},
						func(variables map[string]string) (any, error) {
							first, err := strconv.Atoi(variables["first"])
							if err != nil {
								return nil, err
							}

							second, err := strconv.Atoi(variables["second"])
							if err != nil {
								return nil, err
							}

							result := first + second
							return result, nil
						},
					),
					newLineWithExecution(
						[]*variable{
							newVariable("first", "eAnyNumber"),
							newVariable("minus", "eMinus"),
							newVariable("second", "eAnyNumber"),
						},
						func(variables map[string]string) (any, error) {
							first, err := strconv.Atoi(variables["first"])
							if err != nil {
								return nil, err
							}

							second, err := strconv.Atoi(variables["second"])
							if err != nil {
								return nil, err
							}

							result := first - second
							return result, nil
						},
					),
				},
			),
			newBlock(
				"bVariableName",
				[]*line{
					newLine([]*variable{
						newVariable("firstLetter", "eOneLowercaseLetter"),
						newVariable("otherLeters", "eAnyLetters"),
					}),
				},
			),
			newBlock(
				"bLowercaseOrUppercaseLetters",
				[]*line{
					newLine([]*variable{
						newVariable("lowercase", "eAnyLowercaseLetter"),
					}),
					newLine([]*variable{
						newVariable("uppercase", "eAnyUppercaseLetter"),
					}),
				},
			),
		},
		[]*element{
			newElementWithBlock("eOperation", "bOperations", "cOnce"),
			newElementWithBlock("eVariableName", "bVariableName", "cOnce"),
			newElementWithBlock("eAnyLetters", "bLowercaseOrUppercaseLetters", "cOnceOrMore"),
			newElementWithToken("eAnyNumber", "tOneNumber", "cOnceOrMore"),
			newElementWithToken("eOneLowercaseLetter", "tOneLowercaseLetter", "cOnce"),
			newElementWithToken("eAnyLowercaseLetter", "tOneLowercaseLetter", "cOnceOrMore"),
			newElementWithToken("eAnyUppercaseLetter", "tOneUppercaseLetter", "cOnceOrMore"),
			newElementWithToken("ePlus", "tPlus", "cOnce"),
			newElementWithToken("eMinus", "tMinus", "cOnce"),
			newElementWithToken("eEqual", "tEqual", "cOnce"),
		},
		[]*token{
			newToken("tPlus", []byte{SPlus}),
			newToken("tMinus", []byte{SMinus}),
			newToken("tEqual", []byte{SEqual}),
			newToken("tSemiColon", []byte{SSemiColon}),
			newToken("tOneNumber", []byte{NZero, NOne, NTwo, NThree, NFour, NFive, NSix, NSeven, NHeight, NNine}),
			newToken("tOneLowercaseLetter", []byte{LLA, LLB, LLC, LLD, LLE, LLF, LLG, LLH, LLI, LLJ, LLK, LLL, LLM, LLN, LLO, LLP, LLQ, LLR, LLS, LLT, LLU, LLV, LLW, LLX, LLY, LLZ}),
			newToken("tOneUppercaseLetter", []byte{ULA, ULB, ULC, ULD, ULE, ULF, ULG, ULH, ULI, ULJ, ULK, ULL, ULM, ULN, ULO, ULP, ULQ, ULR, ULS, ULT, ULU, ULV, ULW, ULX, ULY, ULZ}),
		},
		[]*cardinality{
			newCardinalityWithAmount("cOnce", 1, 1),
			newCardinality("cAny", 0),
			newCardinality("cOnceOrMore", 1),
		},
		map[uint8]string{
			NZero:      "0",
			NOne:       "1",
			NTwo:       "2",
			NThree:     "3",
			NFour:      "4",
			NFive:      "5",
			NSix:       "6",
			NSeven:     "7",
			NHeight:    "8",
			NNine:      "9",
			SPlus:      "+",
			SMinus:     "-",
			SEqual:     "=",
			SSemiColon: ";",
			LLA:        "a",
			LLB:        "b",
			LLC:        "c",
			LLD:        "d",
			LLE:        "e",
			LLF:        "f",
			LLG:        "g",
			LLH:        "h",
			LLI:        "i",
			LLJ:        "j",
			LLK:        "k",
			LLL:        "l",
			LLM:        "m",
			LLN:        "n",
			LLO:        "o",
			LLP:        "p",
			LLQ:        "q",
			LLR:        "r",
			LLS:        "s",
			LLT:        "t",
			LLU:        "u",
			LLV:        "v",
			LLW:        "w",
			LLX:        "x",
			LLY:        "y",
			LLZ:        "z",
			ULA:        "A",
			ULB:        "B",
			ULC:        "C",
			ULD:        "D",
			ULE:        "E",
			ULF:        "F",
			ULG:        "G",
			ULH:        "H",
			ULI:        "I",
			ULJ:        "J",
			ULK:        "K",
			ULL:        "L",
			ULM:        "M",
			ULN:        "N",
			ULO:        "O",
			ULP:        "P",
			ULQ:        "Q",
			ULR:        "R",
			ULS:        "S",
			ULT:        "T",
			ULU:        "U",
			ULV:        "V",
			ULW:        "W",
			ULX:        "X",
			ULY:        "Y",
			ULZ:        "Z",
		},
	)}
	return &out
}

// Application represents the interpreter application
type Application interface {
	Execute(input []byte) ([]byte, error)
}
