package interpreters

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

// NewApplication creates a new application
func NewApplication() Application {
	out := application{newGrammar(
		"bRoot",
		[]*block{
			newBlock("bRoot", [][]string{
				{"vAnyNumber", "vPlus", "vAnyNumber"},
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
	)}
	return &out
}

// Application represents the interpreter application
type Application interface {
	Execute(input []byte) ([]byte, error)
}
