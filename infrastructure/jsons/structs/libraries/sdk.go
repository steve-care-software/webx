package libraries

import (
	"github.com/steve-care-software/datastencil/infrastructure/jsons/structs/libraries/layers"
	"github.com/steve-care-software/datastencil/infrastructure/jsons/structs/libraries/layers/links"
)

// Library represents a library
type Library struct {
	Layers []layers.Layer `json:"layers"`
	Links  []links.Link   `json:"links"`
}
