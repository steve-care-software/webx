package jsons

// MetaData represents a metadata
type MetaData struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Path        []string `json:"path"`
}
