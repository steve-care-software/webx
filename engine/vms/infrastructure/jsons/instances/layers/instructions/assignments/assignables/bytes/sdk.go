package bytes

import "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/bytes"

// NewAdapter creates a new adapter
func NewAdapter() bytes.Adapter {
	builder := bytes.NewBuilder()
	return createAdapter(
		builder,
	)
}
