package scopes

import (
	"strings"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type scope struct {
	hash   hash.Hash
	prefix []string
}

func createScope(
	hash hash.Hash,
	prefix []string,
) Scope {
	out := scope{
		hash:   hash,
		prefix: prefix,
	}

	return &out
}

// Hash returns the hash
func (obj *scope) Hash() hash.Hash {
	return obj.hash
}

// Prefix returns the prefix
func (obj *scope) Prefix() []string {
	return obj.prefix
}

// Contains returns true if the path is contained in the scope, false otherwise
func (obj *scope) Contains(path []string) bool {
	scopeStr := strings.Join(obj.prefix, "")
	pathStr := strings.Join(path, "")
	return strings.HasPrefix(pathStr, scopeStr)
}
