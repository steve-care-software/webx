package references

type weightedRelation struct {
	from uint
	to   WeightedElements
}

func createWeightedRelation(
	from uint,
	to WeightedElements,
) WeightedRelation {
	out := weightedRelation{
		from: from,
		to:   to,
	}

	return &out
}

// From returns the from index
func (obj *weightedRelation) From() uint {
	return obj.from
}

// To returns the to elements
func (obj *weightedRelation) To() WeightedElements {
	return obj.to
}
