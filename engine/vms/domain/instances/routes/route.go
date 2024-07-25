package routes

import "github.com/steve-care-software/webx/engine/states/domain/hash"

type route struct {
	hash   hash.Hash
	layer  hash.Hash
	tokens Tokens
	global Omission
	token  Omission
}

func createRoute(
	hash hash.Hash,
	layer hash.Hash,
	tokens Tokens,
) Route {
	return createRouteInternally(hash, layer, tokens, nil, nil)
}

func createRouteWithGlobal(
	hash hash.Hash,
	layer hash.Hash,
	tokens Tokens,
	global Omission,
) Route {
	return createRouteInternally(hash, layer, tokens, global, nil)
}

func createRouteWithToken(
	hash hash.Hash,
	layer hash.Hash,
	tokens Tokens,
	local Omission,
) Route {
	return createRouteInternally(hash, layer, tokens, nil, local)
}

func createRouteWithGlobalAndToken(
	hash hash.Hash,
	layer hash.Hash,
	tokens Tokens,
	global Omission,
	local Omission,
) Route {
	return createRouteInternally(hash, layer, tokens, global, local)
}

func createRouteInternally(
	hash hash.Hash,
	layer hash.Hash,
	tokens Tokens,
	global Omission,
	token Omission,
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

// Remaining returns the remaining bytes of the input after trying to matching it
func (obj *route) Remaining(input []byte) []byte {
	remaining := input
	if obj.HasGlobal() {
		remaining = obj.Global().Remaining(input)
	}

	remaining = obj.Tokens().Remaining(remaining)
	return remaining
}

// Layer returns the layer
func (obj *route) Layer() hash.Hash {
	return obj.layer
}

// Tokens returns the tokens
func (obj *route) Tokens() Tokens {
	return obj.tokens
}

// HasGlobal returns true if there is global, false otherwise
func (obj *route) HasGlobal() bool {
	return obj.global != nil
}

// Global returns the global, if any
func (obj *route) Global() Omission {
	return obj.global
}

// HasToken returns true if there is token, false otherwise
func (obj *route) HasToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *route) Token() Omission {
	return obj.token
}
