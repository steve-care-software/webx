package heads

type weightedRelations struct {
	list []WeightedRelation
}

func createWeightedRelations(
	list []WeightedRelation,
) WeightedRelations {
	out := weightedRelations{
		list: list,
	}

	return &out
}

// List returns the weightedRelations
func (obj *weightedRelations) List() []WeightedRelation {
	return obj.list
}
