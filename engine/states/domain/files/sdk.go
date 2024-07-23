package files

import "os"

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithBasePath(basePath []string) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a repository
type Repository interface {
	Open(path []string, permission uint) (*os.File, error)
	Exists(path []string) bool
	RetrieveChunk(identifier *os.File, index uint, length uint) ([]byte, error)
	RetrieveFrom(identifier *os.File, index uint) ([]byte, error)
	RetrieveAll(identifier *os.File) ([]byte, error)
	RetrieveFromPath(path []string) ([]byte, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithBasePath(basePath []string) ServiceBuilder
	Now() (Service, error)
}

// Service represents a file service
type Service interface {
	Init(path []string) error
	Lock(path []string) error
	Unlock(path []string) error
	Save(path []string, bytes []byte) error
	Transact(path []string, bytes []byte) error
	Delete(path []string) error
}
