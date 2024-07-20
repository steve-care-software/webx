package applications

import "github.com/steve-care-software/webx/engine/stencils/applications/layers/binaries"

type layerBinaryApplication struct {
}

func createLayerBinaryApplication() binaries.Application {
	out := layerBinaryApplication{}
	return &out
}

// Execute executes the application
func (app *layerBinaryApplication) Execute(binaries []byte, cmd []string) ([]byte, error) {
	return nil, nil
}
