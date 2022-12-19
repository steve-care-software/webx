package references

type reference struct {
	content ContentKeys
	commits Commits
}

func createReference(
	content ContentKeys,
	commits Commits,
) Reference {
	out := reference{
		content: content,
		commits: commits,
	}

	return &out
}

// Next returns the next beginning index for a pointer
func (obj *reference) Next() int64 {
	return 0
}

// ContentKeys returns the contentKeys
func (obj *reference) ContentKeys() ContentKeys {
	return obj.content
}

// Commits returns the commits, if any
func (obj *reference) Commits() Commits {
	return obj.commits
}
