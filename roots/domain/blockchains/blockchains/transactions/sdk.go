package transactions

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

// MiningCharacter represents the mining character
const MiningCharacter = "0"

// MaxAmount returns the maximum amount of leading mining value a transaction can have
const MaxAmount = 64

// ValuePerMiningCharacter represents the value per mining character
const ValuePerMiningCharacter = 16

// NewBuilder creates a new transactions builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewTransactionBuilder creates a new transaction builder
func NewTransactionBuilder() TransactionBuilder {
	hashAdapter := hash.NewAdapter()
	return createTransactionBuilder(hashAdapter)
}

// FetchMaxScore returns the max score a transaction can have
func FetchMaxScore() *big.Int {
	x := big.NewInt(ValuePerMiningCharacter)
	y := big.NewInt(MaxAmount)
	return x.Exp(x, y, nil)
}

// CalculateScore returns the score of a transaction
func CalculateScore(amount uint8) *big.Int {
	output := big.NewInt(0)
	if amount <= 0 {
		return output
	}

	if amount > MaxAmount {
		amount = MaxAmount
	}

	castedAmount := big.NewInt(int64(amount))
	castedValuePerCharacter := big.NewInt(int64(ValuePerMiningCharacter))

	output = output.Exp(castedValuePerCharacter, castedAmount, nil)
	return output
}

// DiscoverMiningProof discovers a mining proof that matches the required difficulty
func DiscoverMiningProof(hashAdapter hash.Adapter, difficulty uint8, asset hash.Hash, timeOut time.Duration) (*big.Int, error) {
	increment := big.NewInt(1)
	pProof := big.NewInt(0)
	beginsOn := time.Now().UTC()
	for {
		mine, err := ExecuteMiner(hashAdapter, asset, *pProof)
		if err != nil {
			return nil, err
		}

		amount := FetchMinedAmount(*mine)
		if amount >= difficulty {
			return pProof, nil
		}

		pProof = pProof.Add(pProof, increment)
		current := time.Now().UTC()
		elapsed := current.Sub(beginsOn)
		if elapsed >= timeOut {
			break
		}
	}

	str := fmt.Sprintf("the mining proof (required difficulty: %d) could not be found within the timeout value (%s), the last non-found proof was: %s", difficulty, timeOut.String(), pProof.String())
	return nil, errors.New(str)
}

// FetchMinedAmount fetches the mined amount from the mine hash
func FetchMinedAmount(mine hash.Hash) uint8 {
	cpt := uint8(0)
	str := mine.String()
	for {
		if !strings.HasPrefix(str, MiningCharacter) {
			break
		}

		str = strings.TrimPrefix(str, MiningCharacter)
		cpt++
	}

	return cpt
}

// ExecuteMiner executes the miner
func ExecuteMiner(hashAdapter hash.Adapter, asset hash.Hash, proof big.Int) (*hash.Hash, error) {
	return hashAdapter.FromMultiBytes([][]byte{
		asset.Bytes(),
		proof.Bytes(),
	})
}

// Builder represents a transactions builder
type Builder interface {
	Create() Builder
	WithList(list []Transaction) Builder
	Now() (Transactions, error)
}

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	Fetch(hash hash.Hash) (Transaction, error)
	List() []Transaction
	Score() *big.Int
}

// TransactionBuilder represents a transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithAssset(asset hash.Hash) TransactionBuilder
	WithProof(proof big.Int) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Asset() hash.Hash
	Proof() big.Int
	Mine() hash.Hash
	Score() *big.Int
}
