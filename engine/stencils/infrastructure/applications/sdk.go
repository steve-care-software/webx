package applications

import (
	applications_databases "github.com/steve-care-software/webx/engine/states/infrastructure/applications"
	"github.com/steve-care-software/webx/engine/stencils/applications"
	vm_infrastructure_applications "github.com/steve-care-software/webx/engine/vms/infrastructure/applications"
)

const invalidPatternErr = "the provided context (%d) does not exists"

// NewRemoteApplicationBuilder creates a new remote application builder
func NewRemoteApplicationBuilder() applications.RemoteBuilder {
	return createRemoteApplicationBuilder()
}

// NewLocalApplicationBuilder creates a new local application builder
func NewLocalApplicationBuilder() applications.LocalBuilder {
	dbAppBuilder := applications_databases.NewBuilder()
	contextEndPath := []string{"context"}
	commitInnerPath := []string{"commits"}
	chunksInnerPath := []string{"chunks"}
	sizeToChunk := uint(1024)
	splitHashInThisAmount := uint(16)
	ins := createLocalApplicationBuilder(
		dbAppBuilder,
		contextEndPath,
		commitInnerPath,
		chunksInnerPath,
		sizeToChunk,
		splitHashInThisAmount,
	)

	vmAppFactory := vm_infrastructure_applications.NewApplicationFactory(
		ins,
		NewRemoteApplicationBuilder(),
	)

	return ins.(*localApplicationBuilder).init(vmAppFactory)
}
