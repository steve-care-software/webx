package references

type reference struct {
	active            Keys
	pendings          Keys
	deleted           Keys
	links             Links
	relations         Relations
	weightedRelations WeightedRelations
}

func createReference(
	active Keys,
	pendings Keys,
	deleted Keys,
	links Links,
	relations Relations,
	weightedRelations WeightedRelations,
) Reference {
	out := reference{
		active:            active,
		pendings:          pendings,
		deleted:           deleted,
		links:             links,
		relations:         relations,
		weightedRelations: weightedRelations,
	}

	return &out
}

// HasActive returns true if active, false otherwise
func (obj *reference) HasActive() bool {
	return obj.active != nil
}

// Active returns the active keys, if any
func (obj *reference) Active() Keys {
	return obj.active
}

// HasPendings returns true if pendings, false otherwise
func (obj *reference) HasPendings() bool {
	return obj.pendings != nil
}

// Pendings returns the pendings, if any
func (obj *reference) Pendings() Keys {
	return obj.pendings
}

// HasDeleted returns true if deleted, false otherwise
func (obj *reference) HasDeleted() bool {
	return obj.deleted != nil
}

// Deleted returns the deleted keys, if any
func (obj *reference) Deleted() Keys {
	return obj.deleted
}

// HasLinks returns true if links, false otherwise
func (obj *reference) HasLinks() bool {
	return obj.links != nil
}

// Links returns the links, if any
func (obj *reference) Links() Links {
	return obj.links
}

// HasRelations returns true if relations, false otherwise
func (obj *reference) HasRelations() bool {
	return obj.relations != nil
}

// Relations returns the relations, if any
func (obj *reference) Relations() Relations {
	return obj.relations
}

// HasWeightedRelations returns true if weightedRelations, false otherwise
func (obj *reference) HasWeightedRelations() bool {
	return obj.weightedRelations != nil
}

// WeightedRelations returns the weightedRelations, if any
func (obj *reference) WeightedRelations() WeightedRelations {
	return obj.weightedRelations
}
