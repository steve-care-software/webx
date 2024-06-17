package references

import (
	"github.com/steve-care-software/datastencil/infrastructure/jsons/instances"
)

// Reference represents a reference
type Reference struct {
	Variable string              `json:"variable"`
	Instance *instances.Instance `json:"instance"`
}
