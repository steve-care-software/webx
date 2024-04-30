package links

import (
	json_elements "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements"
	json_origins "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/origins"
	json_references "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/references"
)

// Link represents the link
type Link struct {
	Origin     json_origins.Origin         `json:"origin"`
	Elements   []json_elements.Element     `json:"elements"`
	References []json_references.Reference `json:"references"`
}
