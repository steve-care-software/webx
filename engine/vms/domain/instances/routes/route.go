package routes

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
)

type route struct {
	hash   hash.Hash
	layer  hash.Hash
	tokens tokens.Tokens
	global omissions.Omission
	token  omissions.Omission
}

func createRoute(
	hash hash.Hash,
	layer hash.Hash,
	tokens tokens.Tokens,
) Route {
	return createRouteInternally(hash, layer, tokens, nil, nil)
}

func createRouteWithGlobal(
	hash hash.Hash,
	layer hash.Hash,
	tokens tokens.Tokens,
	global omissions.Omission,
) Route {
	return createRouteInternally(hash, layer, tokens, global, nil)
}

func createRouteWithToken(
	hash hash.Hash,
	layer hash.Hash,
	tokens tokens.Tokens,
	local omissions.Omission,
) Route {
	return createRouteInternally(hash, layer, tokens, nil, local)
}

func createRouteWithGlobalAndToken(
	hash hash.Hash,
	layer hash.Hash,
	tokens tokens.Tokens,
	global omissions.Omission,
	local omissions.Omission,
) Route {
	return createRouteInternally(hash, layer, tokens, global, local)
}

func createRouteInternally(
	hash hash.Hash,
	layer hash.Hash,
	tokens tokens.Tokens,
	global omissions.Omission,
	token omissions.Omission,
) Route {
	out := route{
		hash:   hash,
		layer:  layer,
		tokens: tokens,
		global: global,
		token:  token,
	}

	return &out
}

// Hash returns the hash
func (obj *route) Hash() hash.Hash {
	return obj.hash
}

// Layer returns the layer
func (obj *route) Layer() hash.Hash {
	return obj.layer
}

// Tokens returns the tokens
func (obj *route) Tokens() tokens.Tokens {
	return obj.tokens
}

// HasGlobal returns true if there is global, false otherwise
func (obj *route) HasGlobal() bool {
	return obj.global != nil
}

// Global returns the global, if any
func (obj *route) Global() omissions.Omission {
	return obj.global
}

// HasToken returns true if there is token, false otherwise
func (obj *route) HasToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *route) Token() omissions.Omission {
	return obj.token
}
