package communications

import (
	json_signs "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/signs"
	json_votes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/votes"
)

// Communication represents a communication
type Communication struct {
	Sign         *json_signs.Sign `json:"sign"`
	Vote         *json_votes.Vote `json:"vote"`
	GenerateRing string           `json:"generate_ring"`
}
