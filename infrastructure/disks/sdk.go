package disks

import (
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/passwords"
	"github.com/steve-care-software/webx/domain/identities"
)

// NewIdentityRepository creates a new identity repository
func NewIdentityRepository(
	basePath string,
	delimiter string,
	extension string,
) identities.Repository {
	encryptionBuilder := passwords.NewBuilder()
	return createIdentityRepository(
		encryptionBuilder,
		basePath,
		delimiter,
		extension,
	)
}

// NewIdentityService creates a new identity service
func NewIdentityService(
	repository identities.Repository,
	basePath string,
	delimiter string,
	extension string,
) identities.Service {
	encryptionBuilder := passwords.NewBuilder()
	return createIdentityService(
		encryptionBuilder,
		repository,
		basePath,
		delimiter,
		extension,
	)
}
