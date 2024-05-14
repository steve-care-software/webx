package databases

import (
	json_commits "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases/commits"
)

// Database represents a database
type Database struct {
	Path        []string            `json:"path"`
	Description string              `json:"description"`
	IsActive    bool                `json:"is_active"`
	Head        json_commits.Commit `json:"head"`
}
