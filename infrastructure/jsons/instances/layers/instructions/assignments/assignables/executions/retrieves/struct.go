package retrieves

// Retrieve represents a retrieve
type Retrieve struct {
	Context string `json:"context"`
	Index   string `json:"index"`
	Length  string `json:"length"`
}
