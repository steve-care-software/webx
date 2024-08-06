package routes

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes"
	json_omissions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/omissions"
	json_tokens "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/tokens"
)

// Adapter represents a route adapter
type Adapter struct {
	omissionAdapter *json_omissions.Adapter
	tokenAdapter    *json_tokens.Adapter
	builder         routes.Builder
	hashAdapter     hash.Adapter
}

func createAdapter(
	omissionAdapter *json_omissions.Adapter,
	tokenAdapter *json_tokens.Adapter,
	builder routes.Builder,
	hashAdapter hash.Adapter,
) routes.Adapter {
	out := Adapter{
		omissionAdapter: omissionAdapter,
		tokenAdapter:    tokenAdapter,
		builder:         builder,
		hashAdapter:     hashAdapter,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins routes.Route) ([]byte, error) {
	ptr, err := app.RouteToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(data []byte) (routes.Route, error) {
	ins := new(Route)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToRoute(*ins)
}

// RouteToStruct converts an instance to struct
func (app *Adapter) RouteToStruct(ins routes.Route) (*Route, error) {
	tokens, err := app.tokenAdapter.TokensToStructs(ins.Tokens())
	if err != nil {
		return nil, err
	}

	out := Route{
		Layer:  ins.Layer().String(),
		Tokens: tokens,
	}

	if ins.HasGlobal() {
		ptr, err := app.omissionAdapter.OmissionToStruct(ins.Global())
		if err != nil {
			return nil, err
		}

		out.Global = ptr
	}

	if ins.HasToken() {
		ptr, err := app.omissionAdapter.OmissionToStruct(ins.Token())
		if err != nil {
			return nil, err
		}

		out.TokenOmission = ptr
	}

	return &out, nil
}

// StructToRoute converts a struct to route
func (app *Adapter) StructToRoute(str Route) (routes.Route, error) {
	pHash, err := app.hashAdapter.FromString(str.Layer)
	if err != nil {
		return nil, err
	}

	tokens, err := app.tokenAdapter.StructsToTokens(str.Tokens)
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().
		WithLayer(*pHash).
		WithTokens(tokens)

	if str.Global != nil {
		ins, err := app.omissionAdapter.StructToOmission(*str.Global)
		if err != nil {
			return nil, err
		}

		builder.WithGlobal(ins)
	}

	if str.TokenOmission != nil {
		ins, err := app.omissionAdapter.StructToOmission(*str.TokenOmission)
		if err != nil {
			return nil, err
		}

		builder.WithToken(ins)
	}

	return builder.Now()
}
