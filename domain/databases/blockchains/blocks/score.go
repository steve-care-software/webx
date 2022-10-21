package blocks

import "math/big"

type score struct {
	pNext    *big.Int
	pPending *big.Int
	pTotal   *big.Int
}

func createScore(
	pNext *big.Int,
	pPending *big.Int,
	pTotal *big.Int,
) Score {
	out := score{
		pNext:    pNext,
		pPending: pPending,
		pTotal:   pTotal,
	}

	return &out
}

// Next returns the next
func (obj *score) Next() *big.Int {
	return obj.pNext
}

// Pending returns the pending
func (obj *score) Pending() *big.Int {
	return obj.pPending
}

// Total returns the total
func (obj *score) Total() *big.Int {
	return obj.pTotal
}
