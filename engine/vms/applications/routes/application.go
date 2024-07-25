package routes

import (
	"bytes"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/elements"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/omissions"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
)

type application struct {
	routeRepository routes.Repository
}

func createApplication(
	routeRepository routes.Repository,
) Application {
	out := application{
		routeRepository: routeRepository,
	}

	return &out
}

// Execute executes the route
func (app *application) Execute(input []byte, route routes.Route) (bool, []byte, error) {
	remaining, err := app.route(input, route)
	if err != nil {
		return false, nil, err
	}

	// if the remaining is the same as the input, it does not match:
	if bytes.Equal(remaining, input) {
		return false, input, nil
	}

	return true, remaining, nil
}

func (app *application) route(input []byte, route routes.Route) ([]byte, error) {
	remaining := input
	if route.HasGlobal() {
		retRemaining, err := app.omission(input, route.Global())
		if err != nil {
			return nil, err
		}

		remaining = retRemaining
	}

	tokenOmission := route.Token()
	return app.tokens(
		remaining,
		route.Tokens(),
		tokenOmission,
	)
}

func (app *application) tokens(input []byte, tokens tokens.Tokens, tokenOmission omissions.Omission) ([]byte, error) {
	remaining := input
	list := tokens.List()
	for _, oneToken := range list {
		if len(remaining) <= 0 {
			return input, nil // return the whole input as the list of tokens did not all match
		}

		retRemaining, err := app.token(remaining, oneToken, tokenOmission)
		if err != nil {
			return nil, err
		}
		remaining = retRemaining
	}

	return remaining, nil
}

func (app *application) token(input []byte, token tokens.Token, tokenOmission omissions.Omission) ([]byte, error) {
	remaining := input
	if tokenOmission != nil {
		retRemaining, err := app.omission(remaining, tokenOmission)
		if err != nil {
			return nil, err
		}

		remaining = retRemaining
	}

	if token.HasOmission() {
		retRemaining, err := app.omission(remaining, token.Omission())
		if err != nil {
			return nil, err
		}

		remaining = retRemaining
	}

	cardinality := token.Cardinality()
	hasMax := cardinality.HasMax()
	pMax := cardinality.Max()
	amount := 0
	for {
		if hasMax {
			max := *pMax
			if amount >= int(max) {
				break
			}
		}

		retRemaining, err := app.elements(remaining, token.Elements())
		if err != nil {
			return nil, err
		}

		if len(retRemaining) < len(remaining) {
			remaining = retRemaining
			amount++
			continue
		}
	}

	min := cardinality.Min()
	if int(min) < amount {
		return input, nil // the minimum has not been reached so return the input
	}

	return remaining, nil
}

func (app *application) omission(input []byte, omission omissions.Omission) ([]byte, error) {
	remaining := input
	if omission.HasPrefix() {
		retRemaining, err := app.element(input, omission.Prefix())
		if err != nil {
			return nil, err
		}

		remaining = retRemaining
	}

	if omission.HasSuffix() {
		toRemove := []byte{}
		data := input
		length := len(data) - 1
		suffix := omission.Suffix()
		for i := 0; i < length; i++ {
			data = data[i:]
			retRemaining, err := app.element(data, suffix)
			if err != nil {
				return nil, err
			}

			// if they are NOT equal, means there is no match:
			if !bytes.Equal(data, retRemaining) {
				continue
			}

			toRemove = retRemaining
		}

		// remove the discovered suffix, if there is any:
		if len(toRemove) >= 0 {
			remaining = bytes.TrimSuffix(remaining, toRemove)
		}
	}

	return remaining, nil
}

func (app *application) elements(input []byte, elements elements.Elements) ([]byte, error) {
	remaining := input
	list := elements.List()
	for _, oneElement := range list {
		// if the remaining is empty, there is no match:
		if len(remaining) <= 0 {
			return input, nil
		}

		retRemaining, err := app.element(remaining, oneElement)
		if err != nil {
			return nil, err
		}

		// if the returned remaining is in the suffix of the input, there is a match:
		if bytes.HasSuffix(remaining, retRemaining) {
			remaining = retRemaining
			continue
		}

		// if the returned remaining is the same as the input, there is no match:
		if bytes.Equal(input, retRemaining) {
			return input, nil
		}
	}

	return remaining, nil
}

func (app *application) element(input []byte, element elements.Element) ([]byte, error) {
	if element.IsLayer() {
		layerHash := element.Layer()
		route, err := app.routeRepository.RetrieveFromLayer(layerHash)
		if err != nil {
			return nil, err
		}

		return app.route(input, route)
	}

	if element.IsBytes() {
		prefix := element.Bytes()
		if bytes.HasPrefix(input, prefix) {
			return bytes.TrimPrefix(input, prefix), nil
		}

		return input, nil
	}

	str := element.String()
	prefix := []byte(str)
	if bytes.HasPrefix(input, prefix) {
		return bytes.TrimPrefix(input, prefix), nil
	}

	return input, nil
}
