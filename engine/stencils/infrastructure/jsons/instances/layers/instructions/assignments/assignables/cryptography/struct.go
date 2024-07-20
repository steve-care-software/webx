package cryptography

import (
	json_decrypts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	json_encrypts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	json_keys "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys"
)

// Cryptography represents a cryptography
type Cryptography struct {
	Encrypt *json_encrypts.Encrypt `json:"encrypt"`
	Decrypt *json_decrypts.Decrypt `json:"decrypt"`
	Key     *json_keys.Key         `json:"key"`
}
