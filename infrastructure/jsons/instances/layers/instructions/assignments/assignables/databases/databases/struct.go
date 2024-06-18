package databases

// Database represents a database
type Database struct {
	Path        string `json:"path"`
	Description string `json:"description"`
	Head        string `json:"head"`
	IsActive    string `json:"is_active"`
}
