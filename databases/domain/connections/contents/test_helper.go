package contents

import (
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// NewContentsForTests creates a new contents for tests
func NewContentsForTests(kind uint, data [][]byte) Contents {
	list := []Content{}
	for _, oneData := range data {
		ins := NewContentForTests(kind, oneData)
		list = append(list, ins)
	}

	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentForTests creates a new content for tests
func NewContentForTests(kind uint, data []byte) Content {
	pHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is a content hash: %d", kind)))
	if err != nil {
		panic(err)
	}

	ins, err := NewContentBuilder().Create().WithHash(*pHash).WithData(data).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
