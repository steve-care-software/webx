package links

import (
	json_layers "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions/links/layers"
	json_links "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links"
)

// Link represents a link
type Link struct {
	Input  string              `json:"input"`
	Source json_links.Link     `json:"source"`
	Layers []json_layers.Layer `json:"layers"`
}
