package executions

import (
	json_databases "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/databases"
	json_links "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links"
)

// Execution represents an execution
type Execution struct {
	Link     json_links.Link         `json:"link"`
	Database json_databases.Database `json:"database"`
}
