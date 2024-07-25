package routes

import "github.com/steve-care-software/webx/engine/states/domain/hash"

type route struct {
	hash   hash.Hash
	tokens Tokens
	global Omission
	token  Omission
}

func createRoute(
	hash hash.Hash,
	tokens Tokens,
) Route {
	return createRouteInternally(hash, tokens, nil, nil)
}

func createRouteWithGlobal(
	hash hash.Hash,
	tokens Tokens,
	global Omission,
) Route {
	return createRouteInternally(hash, tokens, global, nil)
}

func createRouteWithToken(
	hash hash.Hash,
	tokens Tokens,
	local Omission,
) Route {
	return createRouteInternally(hash, tokens, nil, local)
}

func createRouteWithGlobalAndToken(
	hash hash.Hash,
	tokens Tokens,
	global Omission,
	local Omission,
) Route {
	return createRouteInternally(hash, tokens, global, local)
}

func createRouteInternally(
	hash hash.Hash,
	tokens Tokens,
	global Omission,
	token Omission,
) Route {
	out := route{
		hash:   hash,
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
