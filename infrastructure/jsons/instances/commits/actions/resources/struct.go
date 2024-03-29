package resources

// Resource represents a resource
type Resource struct {
	Path     []string `json:"path"`
	Instance string   `json:"instance"`
}
