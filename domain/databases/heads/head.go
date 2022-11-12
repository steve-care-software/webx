package heads

type head struct {
	active            Keys
	deleted           Keys
	links             Links
	relations         Relations
	weightedRelations WeightedRelations
}

func createHead(
	active Keys,
	deleted Keys,
	links Links,
	relations Relations,
	weightedRelations WeightedRelations,
) Head {
	out := head{
		active:            active,
		deleted:           deleted,
		links:             links,
		relations:         relations,
		weightedRelations: weightedRelations,
	}

	return &out
}

// HasActive returns true if active, false otherwise
func (obj *head) HasActive() bool {
	return obj.active != nil
}

// Active returns the active keys, if any
func (obj *head) Active() Keys {
	return obj.active
}

// HasDeleted returns true if deleted, false otherwise
func (obj *head) HasDeleted() bool {
	return obj.deleted != nil
}

// Deleted returns the deleted keys, if any
func (obj *head) Deleted() Keys {
	return obj.deleted
}

// HasLinks returns true if links, false otherwise
func (obj *head) HasLinks() bool {
	return obj.links != nil
}

// Links returns the links, if any
func (obj *head) Links() Links {
	return obj.links
}

// HasRelations returns true if relations, false otherwise
func (obj *head) HasRelations() bool {
	return obj.relations != nil
}

// Relations returns the relations, if any
func (obj *head) Relations() Relations {
	return obj.relations
}

// HasWeightedRelations returns true if weightedRelations, false otherwise
func (obj *head) HasWeightedRelations() bool {
	return obj.weightedRelations != nil
}

// WeightedRelations returns the weightedRelations, if any
func (obj *head) WeightedRelations() WeightedRelations {
	return obj.weightedRelations
}
