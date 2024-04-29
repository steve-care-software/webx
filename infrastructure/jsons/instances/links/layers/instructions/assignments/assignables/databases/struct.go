package databases

import (
	json_repositories "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/databases/repositories"
	json_services "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/databases/services"
)

// Database represents a database
type Database struct {
	Repository *json_repositories.Repository `json:"repository"`
	Service    *json_services.Service        `json:"service"`
}
