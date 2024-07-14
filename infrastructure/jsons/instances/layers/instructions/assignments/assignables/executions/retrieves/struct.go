package retrieves

// Retrieve represents a retrieve
type Retrieve struct {
	Context string `json:"context"`
	Index   string `json:"index"`
	Return  string `json:"return"`
	Length  string `json:"length"`
}
