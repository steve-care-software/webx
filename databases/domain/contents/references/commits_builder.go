package references

import "errors"

type commitsBuilder struct {
	list []Commit
}

func createCommitsBuilder() CommitsBuilder {
	out := commitsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *commitsBuilder) Create() CommitsBuilder {
	return createCommitsBuilder()
}

// WithList add commits to the builder
func (app *commitsBuilder) WithList(list []Commit) CommitsBuilder {
	app.list = list
	return app
}

// Now builds Commits instance
func (app *commitsBuilder) Now() (Commits, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Commit in order to build a Commits instance")
	}

	mp := map[string]Commit{}
	for _, oneCommit := range app.list {
		commitname := oneCommit.Hash().String()
		mp[commitname] = oneCommit
	}

	return createCommits(mp, app.list), nil
}
