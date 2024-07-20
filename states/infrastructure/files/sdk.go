package files

import "github.com/steve-care-software/datastencil/states/domain/files"

// NewRepositoryBuilder creates a new repository builder
func NewRepositoryBuilder(
	innerPath []string,
) files.RepositoryBuilder {
	return createFileRepositoryBuilder(
		innerPath,
	)
}

// NewServiceBuilder creates a new service builder
func NewServiceBuilder(
	innerPath []string,
) files.ServiceBuilder {
	return createFileServiceBuilder(
		NewRepositoryBuilder(innerPath),
		innerPath,
	)
}
