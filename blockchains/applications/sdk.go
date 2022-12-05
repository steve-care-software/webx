package applications

import "github.com/steve-care-software/webx/blockchains/domain/contents/references"

// Application represents the read application
type Application interface {
	Open(name string) (*uint, error)
	Read(context uint, pointer references.Pointer) ([]byte, error)
	ReadAll(context uint, pointers []references.Pointer) ([][]byte, error)
	Write(data []byte) error
	WriteAll(data [][]byte) error
	Cancel(context uint) error
	Commit(context uint) error
	Push(context uint) error
	Close(context uint) error
}
