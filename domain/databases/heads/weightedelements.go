package heads

type weightedElements struct {
	list []WeightedElement
}

func createWeightedElements(
	list []WeightedElement,
) WeightedElements {
	out := weightedElements{
		list: list,
	}

	return &out
}

// List returns the weightedElements
func (obj *weightedElements) List() []WeightedElement {
	return obj.list
}
