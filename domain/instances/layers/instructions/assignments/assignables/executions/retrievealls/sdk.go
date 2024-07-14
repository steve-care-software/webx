package retrievealls

// RetrieveAll represents a retrieve all
type RetrieveAll interface {
	Context() string
	Index() string
	Length() string
	Executions() string
}
