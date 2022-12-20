package contents

import (
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// NewContentForTests creates a new content for tests
func NewContentForTests(kind uint, data []byte) Content {
	pHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is a content hash: %d", kind)))
	if err != nil {
		panic(err)
	}

	ins, err := NewBuilder().Create().WithHash(*pHash).WithData(data).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
