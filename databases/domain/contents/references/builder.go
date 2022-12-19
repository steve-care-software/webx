package references

import (
	"errors"
	"net/url"
)

type builder struct {
	contentKeys ContentKeys
	commits     Commits
	peers       []*url.URL
}

func createBuilder() Builder {
	out := builder{
		contentKeys: nil,
		commits:     nil,
		peers:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithContentKeys adds a contentKeys to the builder
func (app *builder) WithContentKeys(contentKeys ContentKeys) Builder {
	app.contentKeys = contentKeys
	return app
}

// WithCommits add commits to the builder
func (app *builder) WithCommits(commits Commits) Builder {
	app.commits = commits
	return app
}

// WithPeers add peers to the builder
func (app *builder) WithPeers(peers []*url.URL) Builder {
	app.peers = peers
	return app
}

// Now builds a new Reference instance
func (app *builder) Now() (Reference, error) {
	if app.contentKeys == nil {
		return nil, errors.New("the ContentKeys is mandatory in order to build a Reference instance")
	}

	if app.commits == nil {
		return nil, errors.New("the Commits is mandatory in order to build a Reference instance")
	}

	if app.peers != nil && len(app.peers) <= 0 {
		app.peers = nil
	}

	if app.peers != nil {
		return createReferenceWithPeers(app.contentKeys, app.commits, app.peers), nil
	}

	return createReference(app.contentKeys, app.commits), nil
}
