package interpreters

func newGrammar(
	blockEntry string,
	blocks []*block,
	elements []*element,
	blockPointers []*blockPointer,
	tokenPointers []*tokenPointer,
	tokens []*token,
	cardinalities []*cardinality,
	bytesMapping map[uint8]string,
) *grammar {
	mpBlock := map[string]*block{}
	for _, oneBlock := range blocks {
		mpBlock[oneBlock.name] = oneBlock
	}

	mElements := map[string]*element{}
	for _, oneElement := range elements {
		mElements[oneElement.name] = oneElement
	}

	mpBlockPointers := map[string]*blockPointer{}
	for _, oneBlockPointer := range blockPointers {
		mpBlockPointers[oneBlockPointer.name] = oneBlockPointer
	}

	mTokenPointers := map[string]*tokenPointer{}
	for _, oneTokenPointer := range tokenPointers {
		mTokenPointers[oneTokenPointer.name] = oneTokenPointer
	}

	mTokens := map[string]*token{}
	for _, oneToken := range tokens {
		mTokens[oneToken.name] = oneToken
	}

	mCardinalities := map[string]*cardinality{}
	for _, oneCardinality := range cardinalities {
		mCardinalities[oneCardinality.name] = oneCardinality
	}

	return &grammar{
		blockEntry:    blockEntry,
		blockPointers: mpBlockPointers,
		blocks:        mpBlock,
		elements:      mElements,
		tokenPointers: mTokenPointers,
		tokens:        mTokens,
		cardinalities: mCardinalities,
		bytesMapping:  bytesMapping,
	}
}

func newBlockPointer(name string, block string, cardinality string) *blockPointer {
	return &blockPointer{
		name:        name,
		block:       block,
		cardinality: cardinality,
	}
}

func newBlock(name string, lines []*line) *block {
	return &block{
		name:  name,
		lines: lines,
	}
}

func newLineWithExecution(elements map[string]string, execFn executeFn) *line {
	return &line{
		elements: elements,
		execFn:   execFn,
	}
}

func newLine(elements map[string]string) *line {
	return &line{
		elements: elements,
		execFn:   nil,
	}
}

func newElementWithTokenPointer(name string, tokenPointer string) *element {
	return &element{
		name:         name,
		token:        "",
		tokenPointer: tokenPointer,
		blockPointer: "",
	}
}

func newElementWithToken(name string, token string) *element {
	return &element{
		name:         name,
		token:        token,
		tokenPointer: "",
		blockPointer: "",
	}
}

func newTokenPointer(name string, token string, cardinality string) *tokenPointer {
	return &tokenPointer{
		name:        name,
		token:       token,
		cardinality: cardinality,
	}
}

func newToken(name string, characters []byte, cardinality string) *token {
	return &token{
		name:        name,
		characters:  characters,
		cardinality: cardinality,
	}
}

func newCardinality(name string, min uint) *cardinality {
	return &cardinality{
		name:    name,
		min:     min,
		pAmount: nil,
	}
}

func newCardinalityWithAmount(name string, min uint, amount uint) *cardinality {
	return &cardinality{
		name:    name,
		min:     min,
		pAmount: &amount,
	}
}
