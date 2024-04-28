package encryptions

import (
	json_decrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	json_encrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
)

// Encryption represents an encryption
type Encryption struct {
	Encrypt *json_encrypts.Encrypt `json:"encrypt"`
	Decrypt *json_decrypts.Decrypt `json:"decrypt"`
}
