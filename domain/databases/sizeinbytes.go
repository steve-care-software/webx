package databases

type sizeInBytes struct {
	maxAmount uint
	amount    uint
}

func createSizeInBytes(
	maxAmount uint,
	amount uint,
) SizeInBytes {
	out := sizeInBytes{
		maxAmount: maxAmount,
		amount:    amount,
	}

	return &out
}

// MaxAmount returns the max amount
func (obj *sizeInBytes) MaxAmount() uint {
	return obj.maxAmount
}

// Amount returns the amount
func (obj *sizeInBytes) Amount() uint {
	return obj.amount
}

// IsZero returns true if the size is zero, false otherwise
func (obj *sizeInBytes) IsZero() bool {
	return obj.maxAmount == 0 && obj.amount == 0
}
