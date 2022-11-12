package references

type weightedElement struct {
	index  uint
	weight uint
}

func createWeightedElement(
	index uint,
	weight uint,
) WeightedElement {
	out := weightedElement{
		index:  index,
		weight: weight,
	}

	return &out
}

// Index returns the index
func (obj *weightedElement) Index() uint {
	return obj.index
}

// Weight returns the weight
func (obj *weightedElement) Weight() uint {
	return obj.weight
}
