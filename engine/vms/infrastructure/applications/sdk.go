package applications

import (
	stencil_applications "github.com/steve-care-software/webx/engine/stencils/applications"
	vm_layers "github.com/steve-care-software/webx/engine/vms/applications/layers"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/binaries"
)

const keyEncryptionBitrate = 4096

// NewApplicationFactory creates a new application factory
func NewApplicationFactory(
	localApplicationBuilder stencil_applications.LocalBuilder,
	remoteApplicationBuilder stencil_applications.RemoteBuilder,
) vm_layers.Factory {
	return createFactory(
		localApplicationBuilder,
		remoteApplicationBuilder,
	)
}

// NewLayerBinaryApplication creates a new layer binary application
func NewLayerBinaryApplication() binaries.Application {
	return createLayerBinaryApplication()
}
