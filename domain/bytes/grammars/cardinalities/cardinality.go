package cardinalities

type cardinality struct {
	min  uint
	pMax *uint
}

func createCardinality(
	min uint,
) Cardinality {
	return createCardinalityInternally(min, nil)
}

func createCardinalityWithMax(
	min uint,
	pMax *uint,
) Cardinality {
	return createCardinalityInternally(min, pMax)
}

func createCardinalityInternally(
	min uint,
	pMax *uint,
) Cardinality {
	out := cardinality{
		min:  min,
		pMax: pMax,
	}

	return &out
}

// Min returns the minimum
func (obj *cardinality) Min() uint {
	return obj.min
}

// HasMax returns true if there is a max, false otherwise
func (obj *cardinality) HasMax() bool {
	return obj.pMax != nil
}

// Max returns the max, if any
func (obj *cardinality) Max() *uint {
	return obj.pMax
}
