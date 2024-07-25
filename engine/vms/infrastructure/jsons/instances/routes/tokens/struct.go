package tokens

import (
	json_elements "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/elements"
	json_omissions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/omissions"
	json_cardinalities "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/tokens/cardinalities"
)

// Token represents a token
type Token struct {
	Elements    []json_elements.Element        `json:"elements"`
	Cardinality json_cardinalities.Cardinality `json:"cardinality"`
	Omission    *json_omissions.Omission       `json:"omission"`
}
