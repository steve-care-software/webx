package accounts

import (
	json_communications "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/communications"
	json_credentials "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/credentials"
	json_encryptions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/encryptions"
	json_retrieves "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/assignments/assignables/accounts/retrieves"
)

// Account represents an account
type Account struct {
	List          string                             `json:"list"`
	Credentials   *json_credentials.Credentials      `json:"credentials"`
	Retrieve      *json_retrieves.Retrieve           `json:"retrieve"`
	Communication *json_communications.Communication `json:"communication"`
	Encryption    *json_encryptions.Encryption       `json:"encryption"`
}
