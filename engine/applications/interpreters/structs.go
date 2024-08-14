package interpreters

type grammar struct {
	blockEntry    string
	bytesMapping  map[uint8]string
	blocks        map[string]*block
	elements      map[string]*element
	tokens        map[string]*token
	cardinalities map[string]*cardinality
}

type block struct {
	name  string
	lines []*line
}

type line struct {
	elements map[string]string
	execFn   executeFn
}

type element struct {
	name        string
	token       string
	block       string
	cardinality string
}

type token struct {
	name       string
	characters []byte
}

type cardinality struct {
	name    string
	min     uint
	pAmount *uint
}
