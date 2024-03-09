package executions

// Execution represents an execution
type Execution interface {
	Input() string
	Layer() string
	Library() string
	Context() string
}
