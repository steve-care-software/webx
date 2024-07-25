package tokens

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/tokens"
	json_elements "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/elements"
	json_omissions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/omissions"
	json_cardinalities "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/routes/tokens/cardinalities"
)

// Adapter represents the token adapter
type Adapter struct {
	elementAdapter     *json_elements.Adapter
	omissionAdapter    *json_omissions.Adapter
	cardinalityAdapter *json_cardinalities.Adapter
	builder            tokens.Builder
	tokenBuilder       tokens.TokenBuilder
}

func createAdapter(
	elementAdapter *json_elements.Adapter,
	omissionAdapter *json_omissions.Adapter,
	cardinalityAdapter *json_cardinalities.Adapter,
	builder tokens.Builder,
	tokenBuilder tokens.TokenBuilder,
) tokens.Adapter {
	out := Adapter{
		elementAdapter:     elementAdapter,
		omissionAdapter:    omissionAdapter,
		cardinalityAdapter: cardinalityAdapter,
		builder:            builder,
		tokenBuilder:       tokenBuilder,
	}

	return &out
}

// InstancesToBytes converts instances to bytes
func (app *Adapter) InstancesToBytes(ins tokens.Tokens) ([]byte, error) {
	ptr, err := app.TokensToStructs(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// BytesToInstances converts bytes to instances
func (app *Adapter) BytesToInstances(data []byte) (tokens.Tokens, error) {
	ins := new([]Token)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructsToTokens(*ins)
}

// InstanceToBytes converts instance to bytes
func (app *Adapter) InstanceToBytes(ins tokens.Token) ([]byte, error) {
	ptr, err := app.TokenToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// BytesToInstance converts bytes to instance
func (app *Adapter) BytesToInstance(data []byte) (tokens.Token, error) {
	ins := new(Token)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToToken(*ins)
}

// TokensToStructs converts instances to structs
func (app *Adapter) TokensToStructs(ins tokens.Tokens) ([]Token, error) {
	list := ins.List()
	out := []Token{}
	for _, oneIns := range list {
		ptr, err := app.TokenToStruct(oneIns)
		if err != nil {
			return nil, err
		}

		out = append(out, *ptr)
	}

	return out, nil
}

// StructsToTokens converts a structs to tokens
func (app *Adapter) StructsToTokens(str []Token) (tokens.Tokens, error) {
	list := []tokens.Token{}
	builder := app.builder.Create()
	for _, oneStr := range str {
		ins, err := app.StructToToken(oneStr)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return builder.WithList(list).
		Now()
}

// TokenToStruct converts an instance to struct
func (app *Adapter) TokenToStruct(ins tokens.Token) (*Token, error) {
	cardinality := app.cardinalityAdapter.CardinalityToStruct(ins.Cardinality())
	elements, err := app.elementAdapter.ElementsToStructs(ins.Elements())
	if err != nil {
		return nil, err
	}

	out := Token{
		Elements:    elements,
		Cardinality: cardinality,
	}

	if ins.HasOmission() {
		pOmission, err := app.omissionAdapter.OmissionToStruct(ins.Omission())
		if err != nil {
			return nil, err
		}

		out.Omission = pOmission
	}

	return &out, nil
}

// StructToToken converts a struct to token
func (app *Adapter) StructToToken(str Token) (tokens.Token, error) {
	elements, err := app.elementAdapter.StructsToElements(str.Elements)
	if err != nil {
		return nil, err
	}

	cardinality, err := app.cardinalityAdapter.StructToCardinality(str.Cardinality)
	if err != nil {
		return nil, err
	}

	builder := app.tokenBuilder.Create().
		WithElements(elements).
		WithCardinality(cardinality)

	if str.Omission != nil {
		omission, err := app.omissionAdapter.StructToOmission(*str.Omission)
		if err != nil {
			return nil, err
		}

		builder.WithOmission(omission)
	}

	return builder.Now()
}
