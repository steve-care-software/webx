package references

// Reference represents a reference
type Reference struct {
	Variable string   `json:"variable"`
	Path     []string `json:"path"`
}
