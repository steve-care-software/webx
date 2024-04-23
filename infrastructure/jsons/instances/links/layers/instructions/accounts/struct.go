package accounts

import (
	json_inserts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/accounts/inserts"
	json_updates "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/layers/instructions/accounts/updates"
)

// Account represents an account
type Account struct {
	Insert *json_inserts.Insert `json:"insert"`
	Update *json_updates.Update `json:"update"`
	Delete string               `json:"delete"`
}
