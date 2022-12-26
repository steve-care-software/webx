package references

import "net/url"

type reference struct {
	contentKeys ContentKeys
	commits     Commits
	peers       []*url.URL
}

func createReference(
	contentKeys ContentKeys,
	commits Commits,
) Reference {
	return createReferenceInternally(contentKeys, commits, nil)
}

func createReferenceWithPeers(
	contentKeys ContentKeys,
	commits Commits,
	peers []*url.URL,
) Reference {
	return createReferenceInternally(contentKeys, commits, peers)
}

func createReferenceInternally(
	contentKeys ContentKeys,
	commits Commits,
	peers []*url.URL,
) Reference {
	out := reference{
		contentKeys: contentKeys,
		commits:     commits,
		peers:       peers,
	}

	return &out
}

// ContentKeys returns the contentKeys
func (obj *reference) ContentKeys() ContentKeys {
	return obj.contentKeys
}

// Commits returns the commits, if any
func (obj *reference) Commits() Commits {
	return obj.commits
}

// HasPeers returns true if there is peers, false otherwise
func (obj *reference) HasPeers() bool {
	return obj.peers != nil
}

// Peers returns the peers, if any
func (obj *reference) Peers() []*url.URL {
	return obj.peers
}
