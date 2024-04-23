package updates

import json_criterias "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/accounts/updates/criterias"

// Update represents an update
type Update struct {
	Credentials string                  `json:"credentials"`
	Criteria    json_criterias.Criteria `json:"criteria"`
}
