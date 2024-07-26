package routes

import (
	json_omissions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/omissions"
	json_tokens "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/tokens"
)

// Route represents a route
type Route struct {
	Layer         string                   `json:"layer"`
	Tokens        []json_tokens.Token      `json:"tokens"`
	Global        *json_omissions.Omission `json:"global"`
	TokenOmission *json_omissions.Omission `json:"token_omission"`
}
