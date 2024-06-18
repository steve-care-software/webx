package heads

// Head represents an head
type Head struct {
	Path        []string `json:"path"`
	Description string   `json:"description"`
	IsActive    bool     `json:"is_active"`
}
