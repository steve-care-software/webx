package inserts

// Insert represents an insert
type Insert struct {
	Context  string `json:"context"`
	Path     string `json:"path"`
	Instance string `json:"instance"`
}
