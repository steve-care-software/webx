package links

import (
	json_elements "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/links/elements"
	json_references "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/pointers/resources/logics/links/references"
)

// Link represents the link
type Link struct {
	Elements   []json_elements.Element     `json:"elements"`
	References []json_references.Reference `json:"references"`
}
