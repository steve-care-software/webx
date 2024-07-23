package signs

import (
	json_creates "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	json_validates "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

// Sign represents a sign
type Sign struct {
	Create   *json_creates.Create     `json:"create"`
	Validate *json_validates.Validate `json:"validate"`
}
