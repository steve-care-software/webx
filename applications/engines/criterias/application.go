package criterias

import (
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/trees"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute extracts data from a tree using the provided criteria
func (app *application) Execute(criteria criterias.Criteria, tree trees.Tree) ([]byte, error) {
	return app.extractWithPath([]string{}, criteria, tree)
}

func (app *application) extractWithPath(path []string, criteria criterias.Criteria, tree trees.Tree) ([]byte, error) {
	name := criteria.Name()
	index := criteria.Index()
	subTree, element, err := tree.Fetch(name, index)
	if err != nil {
		return nil, err
	}

	includeChannels := criteria.IncludeChannels()
	if criteria.HasContent() {
		content := criteria.Content()
		if subTree != nil {
			if content.IsChild() {
				child := content.Child()
				return app.extractWithPath(append(path, name), child, subTree)
			}

			return subTree.Bytes(includeChannels), nil
		}

		if content.IsChild() {
			str := fmt.Sprintf("the extraction did NOT succeed because it found an element (path: %s) but the criteria had a child", strings.Join(append(path, name), "/"))
			return nil, errors.New(str)
		}

		if content.IsMatch() {

		}
	}

	return element.Bytes(includeChannels), nil
}
