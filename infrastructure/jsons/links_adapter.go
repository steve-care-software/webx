package jsons

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
	structs "github.com/steve-care-software/datastencil/infrastructure/jsons/structs/libraries/layers/links"
)

type linksAdapter struct {
	builder      links.Builder
	pLinkAdapter *linkAdapter
}

func createLinksAdapter(
	builder links.Builder,
	pLinkAdapter *linkAdapter,
) links.Adapter {
	out := linksAdapter{
		builder:      builder,
		pLinkAdapter: pLinkAdapter,
	}

	return &out
}

// ToData converts links to data
func (app *linksAdapter) ToData(ins links.Links) ([]byte, error) {
	str := app.toStructLinks(ins)
	data, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ToInstance converts bytes to links
func (app *linksAdapter) ToInstance(data []byte) (links.Links, error) {
	ins := []structs.Link{}
	err := json.Unmarshal(data, &ins)
	if err != nil {
		return nil, err
	}

	return app.toInstanceLinks(ins)
}

func (app *linksAdapter) toInstanceLinks(list []structs.Link) (links.Links, error) {
	output := []links.Link{}
	for _, oneStr := range list {
		ins, err := app.pLinkAdapter.toInstanceLink(oneStr)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return app.builder.Create().WithList(output).Now()
}

func (app *linksAdapter) toStructLinks(ins links.Links) []structs.Link {
	list := ins.List()
	output := []structs.Link{}
	for _, oneLink := range list {
		output = append(output, app.pLinkAdapter.toStructLink(oneLink))
	}

	return output
}
