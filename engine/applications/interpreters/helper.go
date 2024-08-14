package interpreters

func newGrammar(
	blockEntry string,
	blocks []*block,
	elements []*element,
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
		blocks:        mpBlock,
		elements:      mElements,
		tokens:        mTokens,
		cardinalities: mCardinalities,
		bytesMapping:  bytesMapping,
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

func newElementWithToken(name string, token string, cardinality string) *element {
	return &element{
		name:        name,
		token:       token,
		block:       "",
		cardinality: cardinality,
	}
}

func newElementWithBlock(name string, block string, cardinality string) *element {
	return &element{
		name:        name,
		token:       "",
		block:       block,
		cardinality: cardinality,
	}
}

func newToken(name string, characters []byte) *token {
	return &token{
		name:       name,
		characters: characters,
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
