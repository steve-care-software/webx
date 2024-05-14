package commits

// Commit represents a commit
type Commit struct {
	Description string `json:"description"`
	Actions     string `json:"actions"`
	Parent      string `json:"parent"`
}
