package applications

import (
	stencil_applications "github.com/steve-care-software/webx/engine/stencils/applications"
	"github.com/steve-care-software/webx/engine/vms/applications"
	"github.com/steve-care-software/webx/engine/vms/applications/binaries"
)

const keyEncryptionBitrate = 4096

// NewApplicationFactory creates a new application factory
func NewApplicationFactory(
	localApplicationBuilder stencil_applications.LocalBuilder,
	remoteApplicationBuilder stencil_applications.RemoteBuilder,
) applications.Factory {
	return createFactory(
		localApplicationBuilder,
		remoteApplicationBuilder,
	)
}

// NewLayerBinaryApplication creates a new layer binary application
func NewLayerBinaryApplication() binaries.Application {
	return createLayerBinaryApplication()
}
