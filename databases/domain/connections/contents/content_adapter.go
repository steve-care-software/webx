package contents

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type contentAdapter struct {
	hashContentAdapter hash.Adapter
	builder            ContentBuilder
}

func createContentAdapter(
	hashContentAdapter hash.Adapter,
	builder ContentBuilder,
) ContentAdapter {
	out := contentAdapter{
		hashContentAdapter: hashContentAdapter,
		builder:            builder,
	}

	return &out
}

// ToContent converts content to bytes
func (app *contentAdapter) ToContent(ins Content) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()

	kindBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(kindBytes, uint64(ins.Kind()))

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, kindBytes...)
	output = append(output, ins.Data()...)
	return output, nil
}

// ToInstance converts bytes to a content instance
func (app *contentAdapter) ToInstance(content []byte) (Content, error) {
	contentLength := len(content)
	if contentLength < minContentSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Content instance, %d provided", minContentSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashContentAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	kindDelimiter := hash.Size + 8
	kind := binary.LittleEndian.Uint64(content[hash.Size:kindDelimiter])

	return app.builder.Create().
		WithHash(*pHash).
		WithKind(uint(kind)).
		WithData(content[kindDelimiter:]).
		Now()
}
