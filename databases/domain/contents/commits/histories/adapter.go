package histories

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type adapter struct {
	hashAdapter hash.Adapter
	builder     Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter: hashAdapter,
		builder:     builder,
	}

	return &out
}

// ToContent converts an history instance to content
func (app *adapter) ToContent(ins History) ([]byte, error) {
	commitBytes := ins.Commit().Bytes()

	scoreBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(scoreBytes, uint64(ins.Score()))

	output := []byte{}
	output = append(output, commitBytes...)
	output = append(output, scoreBytes...)

	return output, nil
}

// ToHistory converts content to an history instance
func (app *adapter) ToHistory(content []byte) (History, error) {
	contentLength := len(content)
	if contentLength != historySize {
		str := fmt.Sprintf("the content was expected to contain %d bytes in order to convert to an History instance, %d provided", historySize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	score := binary.LittleEndian.Uint64(content[hash.Size:])
	return app.builder.Create().
		WithCommit(*pHash).
		WithScore(uint(score)).
		Now()
}
