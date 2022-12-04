package elements

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

type adapter struct {
	hashAdapter        hash.Adapter
	cardinalityAdapter CardinalityAdapter
	builder            Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	cardinalityAdapter CardinalityAdapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter:        hashAdapter,
		cardinalityAdapter: cardinalityAdapter,
		builder:            builder,
	}

	return &out
}

// ToContent converts an element to bytes
func (app *adapter) ToContent(ins Element) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()
	contentBytes := app.contentToBytes(ins.Content())

	contentLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(contentLengthBytes, uint64(len(contentBytes)))

	cardinalityBytes, err := app.cardinalityAdapter.ToContent(ins.Cardinality())
	if err != nil {
		return nil, err
	}

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, contentLengthBytes...)
	output = append(output, contentBytes...)
	output = append(output, cardinalityBytes...)
	return output, nil
}

func (app *adapter) contentToBytes(ins Content) []byte {
	valueBytes := []byte{}
	output := []byte{}
	if ins.IsValue() {
		pValue := ins.Value()
		valueBytes = []byte{
			*pValue,
		}

		output = []byte{
			valueFlag,
		}
	}

	if ins.IsExternal() {
		valueBytes = ins.External().Bytes()
		output = []byte{
			externalFlag,
		}
	}

	if ins.IsToken() {
		valueBytes = ins.Token().Bytes()
		output = []byte{
			tokenFlag,
		}
	}

	if ins.IsEverything() {
		valueBytes = ins.Everything().Bytes()
		output = []byte{
			everythingFlag,
		}
	}

	if ins.IsRecursive() {
		valueBytes = ins.Recursive().Bytes()
		output = []byte{
			recursiveFlag,
		}
	}

	return append(output, valueBytes...)
}

// ToElement converts bytes to an Element instance
func (app *adapter) ToElement(content []byte) (Element, error) {
	contentLength := len(content)
	if contentLength < minElementLength {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to an Element instance, %d provided", minElementLength, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	elementContentLengthDelimiter := hash.Size + 8
	elementContentLength := binary.LittleEndian.Uint64(content[hash.Size:elementContentLengthDelimiter])

	contentDelimiter := elementContentLengthDelimiter + int(elementContentLength)
	if contentLength < contentDelimiter {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to an Element instance, %d provided", contentDelimiter, contentLength)
		return nil, errors.New(str)
	}

	pValue, pExternal, pToken, pEverything, pRecursive, err := app.bytesToContent(content[elementContentLengthDelimiter:contentDelimiter])
	if err != nil {
		return nil, err
	}

	cardinality, err := app.cardinalityAdapter.ToCardinality(content[contentDelimiter:])
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithHash(*pHash).WithCardinality(cardinality)
	if pValue != nil {
		builder.WithValue(*pValue)
	}

	if pExternal != nil {
		builder.WithExternal(*pExternal)
	}

	if pToken != nil {
		builder.WithToken(*pToken)
	}

	if pEverything != nil {
		builder.WithEverything(*pEverything)
	}

	if pRecursive != nil {
		builder.WithRecursive(*pRecursive)
	}

	return builder.Now()
}

func (app *adapter) bytesToContent(bytes []byte) (*uint8, *hash.Hash, *hash.Hash, *hash.Hash, *hash.Hash, error) {
	contentLength := len(bytes)
	if contentLength < 1 {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a element's Content instance, %d provided", 1, contentLength)
		return nil, nil, nil, nil, nil, errors.New(str)
	}

	flag := bytes[:1][0]
	if flag == valueFlag {
		if contentLength < 2 {
			str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a element's Content instance, %d provided", 2, contentLength)
			return nil, nil, nil, nil, nil, errors.New(str)
		}

		value := bytes[1:2][0]
		return &value, nil, nil, nil, nil, nil
	}

	if contentLength < hash.Size+1 {
		str := fmt.Sprintf("the content was expected to contain at least %d bytes in order to convert data to a element's Content instance, %d provided", hash.Size+1, contentLength)
		return nil, nil, nil, nil, nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(bytes[1:])
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	if flag == externalFlag {
		return nil, pHash, nil, nil, nil, nil
	}

	if flag == tokenFlag {
		return nil, nil, pHash, nil, nil, nil
	}

	if flag == everythingFlag {
		return nil, nil, nil, pHash, nil, nil
	}

	if flag == recursiveFlag {
		return nil, nil, nil, nil, pHash, nil
	}

	str := fmt.Sprintf("the element's content flag (%d) is invalid and therefore the provided data cannot be converted to an element's Content instance", flag)
	return nil, nil, nil, nil, nil, errors.New(str)
}
