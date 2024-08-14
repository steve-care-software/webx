package interpreters

type grammar struct {
	blockEntry    string
	bytesMapping  map[uint8]string
	blockPointers map[string]*blockPointer
	blocks        map[string]*block
	values        map[string]*value
	tokenPointers map[string]*tokenPointer
	tokens        map[string]*token
	cardinalities map[string]*cardinality
}

type blockPointer struct {
	name        string
	block       string
	cardinality string
}

type block struct {
	name  string
	lines []*line
}

type line struct {
	values map[string]string
	execFn executeFn
}

type value struct {
	name         string
	token        string
	tokenPointer string
	blockPointer string
}

type tokenPointer struct {
	name        string
	token       string
	cardinality string
}

type token struct {
	name        string
	characters  []byte
	cardinality string
}

type cardinality struct {
	name    string
	min     uint
	pAmount *uint
}
