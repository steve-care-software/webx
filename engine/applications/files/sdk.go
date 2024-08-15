package files

// LocalBuilder represents a local file builder
type LocalBuilder interface {
	Create() LocalBuilder
	WithBasePath(basePath string) LocalBuilder
	Now() (Application, error)
}

// Application represents a file application
type Application interface {
	Dir(relativePath string) ([]string, error)     // returns the list of directory the current directory contains
	Files(relativePath string) ([]string, error)   // returns the list of files the current directory contains
	Open(relativePath string) (*uint, error)       // opens a file
	Length(context uint) error                     // returns the length of the file
	Seek(context uint, index int) error            // seek a file to this index
	Read(context uint, length int) ([]byte, error) // read a file
	Write(context uint, data []byte) error         // write to a file
	Close(context uint) error                      // close a file
}
