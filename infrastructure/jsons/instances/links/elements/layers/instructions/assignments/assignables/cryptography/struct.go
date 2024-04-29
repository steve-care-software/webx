package cryptography

import (
	json_decrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/decrypts"
	json_encrypts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/cryptography/encrypts"
)

// Cryptography represents a cryptography
type Cryptography struct {
	Encrypt *json_encrypts.Encrypt `json:"encrypt"`
	Decrypt *json_decrypts.Decrypt `json:"decrypt"`
}
