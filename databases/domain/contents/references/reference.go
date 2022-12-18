package references

type reference struct {
	content Content
	commits Commits
}

func createReference(
	content Content,
) Reference {
	return createReferenceInternally(content, nil)
}

func createReferenceWithCommits(
	content Content,
	commits Commits,
) Reference {
	return createReferenceInternally(content, commits)
}

func createReferenceInternally(
	content Content,
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

// Content returns the content
func (obj *reference) Content() Content {
	return obj.content
}

// HasCommits returns true if there is commits, false otherwise
func (obj *reference) HasCommits() bool {
	return obj.commits != nil
}

// Commits returns the commits, if any
func (obj *reference) Commits() Commits {
	return obj.commits
}
