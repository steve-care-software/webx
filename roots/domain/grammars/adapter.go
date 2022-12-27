package grammars

import (
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

// ToContent converts a grammar to content
func (app *adapter) ToContent(ins Grammar) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	rootBytes := ins.Root().Bytes()

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, rootBytes...)

	if ins.HasChannels() {
		channels := ins.Channels()
		for _, oneChannel := range channels {
			channelBytes := oneChannel.Bytes()
			output = append(output, channelBytes...)
		}
	}

	return output, nil
}

// ToGrammar converts content to a Grammar instance
func (app *adapter) ToGrammar(content []byte) (Grammar, error) {
	contentLength := len(content)
	if contentLength < minGrammarSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Grammar instance, %d provided", minGrammarSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	rootDelimiter := hash.Size + hash.Size
	pRoot, err := app.hashAdapter.FromBytes(content[hash.Size:rootDelimiter])
	if err != nil {
		return nil, err
	}

	remaining := content[rootDelimiter:]
	remainingLength := len(remaining)
	builder := app.builder.Create().WithHash(*pHash).WithRoot(*pRoot)
	if remainingLength > 0 {
		if remainingLength%hash.Size != 0 {
			str := fmt.Sprintf("the content's remaining length (%d) was expected to be a multiple of %d in order to convert it to a Grammar instance (channel hash list)", remainingLength, hash.Size)
			return nil, errors.New(str)
		}

		channels := []hash.Hash{}
		amount := len(remaining) / hash.Size
		for i := 0; i < amount; i++ {
			startsOn := i * hash.Size
			endsOn := startsOn + hash.Size
			pHash, err := app.hashAdapter.FromBytes(remaining[startsOn:endsOn])
			if err != nil {
				return nil, err
			}

			channels = append(channels, *pHash)
		}

		builder.WithChannels(channels)
	}

	return builder.Now()
}
