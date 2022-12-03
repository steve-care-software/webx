package blocks

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/cryptography/hashtrees"
)

type adapter struct {
	hashAdapter     hash.Adapter
	hashTreeAdapter hashtrees.Adapter
	builder         Builder
}

func createAdapter(
	hashAdapter hash.Adapter,
	hashTreeAdapter hashtrees.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		hashAdapter:     hashAdapter,
		hashTreeAdapter: hashTreeAdapter,
		builder:         builder,
	}

	return &out
}

// ToContent converts block to bytes
func (app *adapter) ToContent(ins Block) ([]byte, error) {
	hashBytes := ins.Hash().Bytes()

	heightBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(heightBytes, uint64(ins.Height()))

	nextScoreBytes := ins.NextScore().Bytes()
	nextScoreLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nextScoreLengthBytes, uint64(len(nextScoreBytes)))

	pendingScoreBytes := ins.PendingScore().Bytes()
	pendingScoreLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(pendingScoreLengthBytes, uint64(len(pendingScoreBytes)))

	trxBytes, err := app.hashTreeAdapter.ToContent(ins.Transactions())
	if err != nil {
		return nil, err
	}

	trxLengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(trxLengthBytes, uint64(len(trxBytes)))

	output := []byte{}
	output = append(output, hashBytes...)
	output = append(output, heightBytes...)
	output = append(output, nextScoreLengthBytes...)
	output = append(output, nextScoreBytes...)
	output = append(output, pendingScoreLengthBytes...)
	output = append(output, pendingScoreBytes...)
	output = append(output, trxLengthBytes...)
	output = append(output, trxBytes...)

	if ins.HasPrevious() {
		previousBytes := ins.Previous().Bytes()
		output = append(output, previousBytes...)
	}

	return output, nil
}

// ToBlock converts bytes to block
func (app *adapter) ToBlock(content []byte) (Block, error) {
	contentLength := len(content)
	if contentLength < minBlockSize {
		str := fmt.Sprintf(errorStr, minBlockSize, contentLength)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromBytes(content[:hash.Size])
	if err != nil {
		return nil, err
	}

	heightDelimiter := hash.Size + 8
	height := binary.LittleEndian.Uint64(content[hash.Size:heightDelimiter])

	nextScoreLengthDelimiter := heightDelimiter + 8
	nextScoreLength := binary.LittleEndian.Uint64(content[heightDelimiter:nextScoreLengthDelimiter])
	nextScoreDelimiter := nextScoreLengthDelimiter + int(nextScoreLength)
	if contentLength < nextScoreDelimiter {
		str := fmt.Sprintf(errorStr, nextScoreDelimiter, contentLength)
		return nil, errors.New(str)
	}

	pNextScore := big.NewInt(0).SetBytes(content[nextScoreLengthDelimiter:nextScoreDelimiter])

	pendingScoreLengthDelimiter := nextScoreDelimiter + 8
	pendingScoreLength := binary.LittleEndian.Uint64(content[nextScoreDelimiter:pendingScoreLengthDelimiter])
	pendingScoreDelimiter := pendingScoreLengthDelimiter + int(pendingScoreLength)
	if contentLength < pendingScoreDelimiter {
		str := fmt.Sprintf(errorStr, pendingScoreDelimiter, contentLength)
		return nil, errors.New(str)
	}

	pPendingScore := big.NewInt(0).SetBytes(content[pendingScoreLengthDelimiter:pendingScoreDelimiter])

	trxLengthDelimiter := pendingScoreDelimiter + 8
	trxLength := binary.LittleEndian.Uint64(content[pendingScoreDelimiter:trxLengthDelimiter])
	trxDelimiter := trxLengthDelimiter + int(trxLength)
	if contentLength < trxDelimiter {
		str := fmt.Sprintf(errorStr, trxDelimiter, contentLength)
		return nil, errors.New(str)
	}

	trx, err := app.hashTreeAdapter.ToHashTree(content[trxLengthDelimiter:trxDelimiter])
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithHash(*pHash).WithHeight(uint(height)).WithNextScore(*pNextScore).WithPendingScope(*pPendingScore).WithTransactions(trx)
	if len(content[trxDelimiter:]) > 0 {
		pPrevHash, err := app.hashAdapter.FromBytes(content[trxDelimiter:])
		if err != nil {
			return nil, err
		}

		builder.WithPrevious(*pPrevHash)
	}

	return builder.Now()
}
