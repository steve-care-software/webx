package interruptions

import (
	json_failures "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commands/results/interruptions/failures"
)

// Interruption represents an interruption
type Interruption struct {
	Stop    *uint                  `json:"stop"`
	Failure *json_failures.Failure `json:"failure"`
}
