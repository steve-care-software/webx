package bytes

import (
	"errors"
	"fmt"
)

func fetchAmountReturnRemaining(data []byte) (int, []byte, error) {
	expectation := AmountOfBytesIntUint64
	if len(data) < expectation {
		str := fmt.Sprintf("the data was expected to contain at least %d bytes, %d provided", expectation, len(data))
		return 0, nil, errors.New(str)
	}

	return int(BytesToUint64(data[:expectation])), data[expectation:], nil
}
