package transactions

import (
	"bytes"
	"math/big"
	"testing"
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

func TestCalculateScore_Success(t *testing.T) {
	pMaxScore := FetchMaxScore()
	data := map[uint8]big.Int{
		0:  *big.NewInt(0),
		1:  *big.NewInt(16),
		2:  *big.NewInt(256),
		64: *pMaxScore,
		65: *pMaxScore,
	}

	for input, output := range data {
		retValue := CalculateScore(input)
		if bytes.Compare(retValue.Bytes(), output.Bytes()) != 0 {
			t.Errorf("when the input is: %d, the output was expected to be '%s', '%s' returned", input, output.String(), retValue.String())
			return
		}
	}
}

func TestDiscoverMiningProof_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	difficulty := uint8(3)
	pAsset, _ := hashAdapter.FromBytes([]byte("this is an asset"))
	timeOut := time.Duration(20 * time.Second)
	pProof, err := DiscoverMiningProof(hashAdapter, difficulty, *pAsset, timeOut)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	mine, _ := ExecuteMiner(hashAdapter, *pAsset, *pProof)
	retAmount := FetchMinedAmount(*mine)
	if retAmount < difficulty {
		t.Errorf("the returned amount was expected to be at least of the required difficulty (%d), %d returned", difficulty, retAmount)
		return
	}
}

func TestDiscoverMiningProof_timeOut_returnsError(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pAsset, _ := hashAdapter.FromBytes([]byte("this is an asset"))
	timeOut := time.Duration(1)
	_, err := DiscoverMiningProof(hashAdapter, MaxAmount, *pAsset, timeOut)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
