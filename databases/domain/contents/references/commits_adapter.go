package references

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type commitsAdapter struct {
	adapter CommitAdapter
	builder CommitsBuilder
}

func createCommitsAdapter(
	adapter CommitAdapter,
	builder CommitsBuilder,
) CommitsAdapter {
	out := commitsAdapter{
		adapter: adapter,
		builder: builder,
	}

	return &out
}

// ToContent converts Commits to bytes
func (app *commitsAdapter) ToContent(ins Commits) ([]byte, error) {
	list := ins.List()
	output := []byte{}
	for _, oneCommit := range list {
		content, err := app.adapter.ToContent(oneCommit)
		if err != nil {
			return nil, err
		}

		lengthBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(lengthBytes, uint64(len(content)))

		output = append(output, lengthBytes...)
		output = append(output, content...)
	}

	return output, nil
}

// ToCommits converts bytes to Commits
func (app *commitsAdapter) ToCommits(content []byte) (Commits, error) {
	smallest := 8 + commitMinSize
	if len(content) < smallest {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert to a Commit instance, %d provided", smallest, len(content))
		return nil, errors.New(str)
	}

	list := []Commit{}
	index := 0
	for {
		amount := len(content[index:])
		if amount <= 0 {
			break
		}

		lengthDelimiter := index + 8
		length := int(binary.LittleEndian.Uint64(content[index:lengthDelimiter]))
		contentDelimiter := lengthDelimiter + length
		ins, err := app.adapter.ToCommit(content[lengthDelimiter:contentDelimiter])
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
		index = contentDelimiter
	}

	return app.builder.Create().WithList(list).Now()
}
