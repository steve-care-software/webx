package references

type content struct {
	active   ContentKeys
	pendings ContentKeys
	deleted  ContentKeys
}

func createContent(
	active ContentKeys,
	pendings ContentKeys,
	deleted ContentKeys,
) Content {
	out := content{
		active:   active,
		pendings: pendings,
		deleted:  deleted,
	}

	return &out
}

// HasActive returns true if active, false otherwise
func (obj *content) HasActive() bool {
	return obj.active != nil
}

// Active returns the active keys, if any
func (obj *content) Active() ContentKeys {
	return obj.active
}

// HasPendings returns true if pendings, false otherwise
func (obj *content) HasPendings() bool {
	return obj.pendings != nil
}

// Pendings returns the pendings, if any
func (obj *content) Pendings() ContentKeys {
	return obj.pendings
}

// HasDeleted returns true if deleted, false otherwise
func (obj *content) HasDeleted() bool {
	return obj.deleted != nil
}

// Deleted returns the deleted keys, if any
func (obj *content) Deleted() ContentKeys {
	return obj.deleted
}
