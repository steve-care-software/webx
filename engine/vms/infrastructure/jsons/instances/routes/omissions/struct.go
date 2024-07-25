package omissions

import (
	json_elements "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/elements"
)

// Omission represents the omission
type Omission struct {
	Prefix *json_elements.Element `json:"prefix"`
	Suffix *json_elements.Element `json:"suffix"`
}
