package repositories

// Repository represents a repository
type Repository struct {
	IsSkeleton bool   `json:"is_skeleton"`
	IsHeight   bool   `json:"is_height"`
	List       string `json:"list"`
	Retrieve   string `json:"retrieve"`
}
