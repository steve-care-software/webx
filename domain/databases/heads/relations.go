package heads

type relations struct {
	list []Relation
}

func createRelations(
	list []Relation,
) Relations {
	out := relations{
		list: list,
	}

	return &out
}

// List returns the relations
func (obj *relations) List() []Relation {
	return obj.list
}
