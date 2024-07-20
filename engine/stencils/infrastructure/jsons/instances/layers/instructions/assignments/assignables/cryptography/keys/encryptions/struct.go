package encryptions

import (
	json_decrypts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	json_encrypts "github.com/steve-care-software/webx/engine/stencils/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

// Encryption represents the encryption
type Encryption struct {
	IsGeneratePrivateKey bool                   `json:"generate_pk"`
	FetchPublicKey       string                 `json:"fetch_pubkey"`
	Decrypt              *json_decrypts.Decrypt `json:"decrypt"`
	Encrypt              *json_encrypts.Encrypt `json:"encrypt"`
}
