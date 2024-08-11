package applications

import (
	"fmt"
	"testing"

	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/loaders/identities"
	applications_loaders "github.com/steve-care-software/webx/engine/cursors/applications/sessions/pointers"
	"github.com/steve-care-software/webx/engine/cursors/infrastructure/edwards25519"
)

func TestIdentity_Success(t *testing.T) {
	encryptionApp := edwards25519.NewEncryptionApplication()
	storagePointerApp := applications_loaders.NewApplication()
	identityApp := identities.NewApplication(
		encryptionApp,
		storagePointerApp,
		nil,
		4096,
	)

	fmt.Printf("\n%v\n", identityApp)
}
