package contexts

// Context represents a context
type Context struct {
	Identifier uint     `json:"identifier"`
	Head       string   `json:"head"`
	Executions []string `json:"executions"`
}
