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

	// SPlus represents a plus sign
	SPlus
)

type executeFn func(map[string]string) error

// NewApplication creates a new application
func NewApplication() Application {
	out := application{newGrammar(
		"bRoot",
		[]*block{
			newBlock("bRoot", []*line{
				newLineWithExecution(map[string]string{
					"first":  "vAnyNumber",
					"plus":   "vPlus",
					"second": "vAnyNumber",
				}, func(input map[string]string) error {
					first, err := strconv.Atoi(input["first"])
					if err != nil {
						return err
					}

					second, err := strconv.Atoi(input["second"])
					if err != nil {
						return err
					}

					result := first + second
					fmt.Printf("!!! yes - works %s + %s = %d!!!!", input["first"], input["second"], result)
					return nil
				}),
			}),
		},
		[]*value{
			newValueWithTokenPointer("vAnyNumber", "tpAnyNumber"),
			newValueWithToken("vPlus", "tPlus"),
		},
		[]*blockPointer{},
		[]*tokenPointer{
			newTokenPointer("tpAnyNumber", "tOneNumber", "cOnceOrMore"),
		},
		[]*token{
			newToken("tPlus", []byte{SPlus}, "cOnce"),
			newToken("tOneNumber", []byte{NZero, NOne, NTwo, NThree, NFour, NFive, NSix, NSeven, NHeight, NNine}, "cOnce"),
		},
		[]*cardinality{
			newCardinalityWithAmount("cOnce", 1, 1),
			newCardinality("cAny", 0),
			newCardinality("cOnceOrMore", 1),
		},
		map[uint8]string{
			NZero:   "0",
			NOne:    "1",
			NTwo:    "2",
			NThree:  "3",
			NFour:   "4",
			NFive:   "5",
			NSix:    "6",
			NSeven:  "7",
			NHeight: "8",
			NNine:   "9",
			SPlus:   "+",
		},
	)}
	return &out
}

// Application represents the interpreter application
type Application interface {
	Execute(input []byte) ([]byte, error)
}
