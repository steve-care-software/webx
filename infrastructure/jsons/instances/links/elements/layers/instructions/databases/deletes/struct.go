package deletes

// Delete represents a delete
type Delete struct {
	Context    string `json:"context"`
	Path       string `json:"path"`
	Identifier string `json:"ientifier"`
}
