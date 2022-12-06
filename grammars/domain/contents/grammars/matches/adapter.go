package matches

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
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

// ToContent converts a Match instance to content
func (app *adapter) ToContent(ins Match) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	tokenBytes := ins.Token().Bytes()

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, tokenBytes...)

	suiteHashes := ins.Suites()
	for _, oneSuiteHash := range suiteHashes {
		output = append(output, oneSuiteHash...)
	}

	return output, nil
}

// ToMatch converts content to a Match instance
func (app *adapter) ToMatch(content []byte) (Match, error) {
	contentLength := len(content)
	if contentLength < minMatchSize {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a Match instance, %d provided", minMatchSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	tokenDelimiter := hash.Size + hash.Size
	pTokenHash, err := app.hashAdapter.FromBytes(content[hash.Size:tokenDelimiter])
	if err != nil {
		return nil, err
	}

	remaining := content[tokenDelimiter:]
	remainingLength := len(remaining)
	builder := app.builder.Create().WithHash(*pHash).WithToken(*pTokenHash)
	if remainingLength > 0 {
		if remainingLength%hash.Size != 0 {
			str := fmt.Sprintf("the content's remaining length (%d) was expected to be a multiple of %d in order to convert it to a Match instance (suite's hash list)", remainingLength, hash.Size)
			return nil, errors.New(str)
		}

		suiteHashes := []hash.Hash{}
		amount := len(remaining) / hash.Size
		for i := 0; i < amount; i++ {
			startsOn := i * hash.Size
			endsOn := startsOn + hash.Size
			pHash, err := app.hashAdapter.FromBytes(remaining[startsOn:endsOn])
			if err != nil {
				return nil, err
			}

			suiteHashes = append(suiteHashes, *pHash)
		}

		builder.WithSuites(suiteHashes)
	}

	return builder.Now()
}
