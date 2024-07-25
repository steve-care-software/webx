package routes

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
)

// NewRouteWithGobalAndTokenForTests creates a new route with global and token omission for tests
func NewRouteWithGobalAndTokenForTests(tokens tokens.Tokens, global omissions.Omission, token omissions.Omission) Route {
	ins, err := NewBuilder().Create().WithTokens(tokens).WithGlobal(global).WithToken(token).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRouteWithGobalForTests creates a new route with global omission for tests
func NewRouteWithGobalForTests(tokens tokens.Tokens, global omissions.Omission) Route {
	ins, err := NewBuilder().Create().WithTokens(tokens).WithGlobal(global).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRouteWithTokenForTests creates a new route with token omission for tests
func NewRouteWithTokenForTests(tokens tokens.Tokens, token omissions.Omission) Route {
	ins, err := NewBuilder().Create().WithTokens(tokens).WithToken(token).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRouteForTests creates a new route for tests
func NewRouteForTests(tokens tokens.Tokens) Route {
	ins, err := NewBuilder().Create().WithTokens(tokens).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
