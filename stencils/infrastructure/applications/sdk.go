package applications

import (
	applications_databases "github.com/steve-care-software/datastencil/states/infrastructure/applications"
	"github.com/steve-care-software/datastencil/stencils/applications"
	"github.com/steve-care-software/datastencil/stencils/applications/layers/binaries"
)

const invalidPatternErr = "the provided context (%d) does not exists"

const keyEncryptionBitrate = 4096

// NewRemoteApplicationBuilder creates a new remote application builder
func NewRemoteApplicationBuilder() applications.RemoteBuilder {
	return createRemoteApplicationBuilder()
}

// NewLayerBinaryApplication creates a new layer binary application
func NewLayerBinaryApplication() binaries.Application {
	return createLayerBinaryApplication()
}

// NewLocalApplicationBuilder creates a new local application builder
func NewLocalApplicationBuilder() applications.LocalBuilder {
	dbAppBuilder := applications_databases.NewBuilder()
	commitInnerPath := []string{"commits"}
	chunksInnerPath := []string{"chunks"}
	sizeToChunk := uint(1024)
	splitHashInThisAmount := uint(16)
	return createLocalApplicationBuilder(
		dbAppBuilder,
		commitInnerPath,
		chunksInnerPath,
		sizeToChunk,
		splitHashInThisAmount,
	)
}
