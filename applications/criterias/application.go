package criterias

import (
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/trees"
	"github.com/steve-care-software/webx/domain/trees/selections"
)

type application struct {
	builder          selections.Builder
	selectionBuilder selections.SelectionBuilder
	childrenBuilder  selections.ChildrenBuilder
	childBuilder     selections.ChildBuilder
}

func createApplication(
	builder selections.Builder,
	selectionBuilder selections.SelectionBuilder,
	childrenBuilder selections.ChildrenBuilder,
	childBuilder selections.ChildBuilder,
) Application {
	out := application{
		builder:          builder,
		selectionBuilder: selectionBuilder,
		childrenBuilder:  childrenBuilder,
		childBuilder:     childBuilder,
	}

	return &out
}

// Retrieve retrieve selections of a tree based on a given criteria
func (app *application) Retrieve(criteria criterias.Criteria, tree trees.Tree) (selections.Selections, error) {
	return nil, nil
}

func (app *application) retrieveByName(name string, tree trees.Tree) {
	block := tree.Block()
	if !block.HasSuccessful() {
		//error
	}

	fetchedElementsList := []trees.Element{}
	successful := block.Successful()
	elementsList := successful.Elements().List()
	for _, oneElement := range elementsList {
		if !oneElement.HasGrammar() {
			continue
		}

		elementName := oneElement.Grammar().Name()
		if elementName == name {
			fetchedElementsList = append(fetchedElementsList, oneElement)
			continue
		}
	}
}

// Execute extracts data from a tree using the provided criteria
func (app *application) Execute(criteria criterias.Criteria, tree trees.Tree) ([]byte, error) {
	return app.extractWithPath([]string{}, criteria, tree)
}

func (app *application) extractWithPath(path []string, criteria criterias.Criteria, tree trees.Tree) ([]byte, error) {
	name := criteria.Name()
	index := criteria.Index()
	subTree, element, err := tree.Fetch(name, *index)
	if err != nil {
		return nil, err
	}

	includeChannels := criteria.IncludeChannels()
	if criteria.HasChild() {
		child := criteria.Child()
		if subTree != nil {
			return app.extractWithPath(append(path, name), child, subTree)
		}

		output := []byte{}
		contents := element.Contents().List()
		for _, oneContent := range contents {
			if !oneContent.IsTree() {
				continue
			}

			subTree := oneContent.Tree()
			data, err := app.extractWithPath(append(path, name), child, subTree)
			if err != nil {
				continue
			}

			output = data
			break
		}

		if len(output) > 0 {
			return output, nil
		}

		str := fmt.Sprintf("the extraction did NOT succeed because it found an element (path: %s) but the criteria had a child", strings.Join(append(path, name), "/"))
		return nil, errors.New(str)
	}

	return element.Bytes(includeChannels), nil
}
