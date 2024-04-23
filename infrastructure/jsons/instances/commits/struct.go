package commits

import "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commits/actions"

// Commit represents a commit
type Commit struct {
	Content   Content `json:"content"`
	Signature string  `json:"signature"`
}

// Content represents content
type Content struct {
	Actions  []actions.Action `json:"actions"`
	Previous *Commit          `json:"commit"`
}
