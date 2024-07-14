package retrieveats

// RetrieveAt represents a retrieve at
type RetrieveAt interface {
	Context() string
	Index() string
	Execution() string
}
