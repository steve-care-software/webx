package selectors

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	grammar_applications "github.com/steve-care-software/webx/roots/applications/grammars"
	"github.com/steve-care-software/webx/roots/domain/grammars/grammars"
	"github.com/steve-care-software/webx/roots/domain/grammars/trees"
	"github.com/steve-care-software/webx/roots/domain/selectors/selectors"
)

type application struct {
	grammarSoftware grammar_applications.Software
}

func createApplication(
	grammarSoftware grammar_applications.Software,
) Application {
	out := application{
		grammarSoftware: grammarSoftware,
	}

	return &out
}

// Retrieve retrieves a selector
func (app *application) Retrieve(context uint, hash hash.Hash) (selectors.Selector, error) {
	return nil, nil
}

// Scan scans the database for the best selector
func (app *application) Scan(context uint, input trees.Tree, output interface{}) (selectors.Selector, error) {
	return nil, nil
}

// Insert inserts a selector
func (app *application) Insert(context uint, selector selectors.Selector) error {
	return nil
}

// InsertAll inserts a list of selectors
func (app *application) InsertAll(context uint, selectors []selectors.Selector) error {
	return nil
}

// Matches returns true if the selector matches the grammar, false otherwise
func (app *application) Matches(grammar grammars.Grammar, selector selectors.Selector) (bool, error) {
	return true, nil
}

// Execute executes a selector on a data tree
func (app *application) Execute(selector selectors.Selector, script []byte) (interface{}, bool, []byte, error) {
	grammar := selector.Grammar()
	treeIns, err := app.grammarSoftware.Execute(grammar, script)
	if err != nil {
		return nil, false, nil, err
	}

	ins, isValid, err := app.selectorFetch(selector, treeIns, nil)
	if err != nil {
		return nil, false, nil, err
	}

	remaining := []byte{}
	if treeIns.HasRemaining() {
		remaining = treeIns.Remaining()
	}

	return ins, isValid, remaining, nil
}

func (app *application) selectorFetch(selector selectors.Selector, tree trees.Tree, previous map[string]selectors.Selector) (interface{}, bool, error) {
	contentsList, err := app.selectorContents(selector, tree)
	if err != nil {
		return nil, false, err
	}

	if previous == nil {
		previous = map[string]selectors.Selector{}
	}

	previous[selector.Token().Name()] = selector
	return app.selectorInstance(selector, contentsList, previous)
}

func (app *application) selectorContents(selector selectors.Selector, tree trees.Tree) ([]trees.Content, error) {
	selectorToken := selector.Token()
	treeName := tree.Grammar().Name()
	if selectorToken.Name() != treeName {
		str := fmt.Sprintf("the contents cannot be retrieved because the tree (token: %s) do not match the selector (token: %s)", treeName, selectorToken.Name())
		return nil, errors.New(str)
	}

	block := tree.Block()
	if !block.HasSuccessful() {
		str := fmt.Sprintf("the contents cannot be retrieved because the tree (token: %s) contains no successful line", treeName)
		return nil, errors.New(str)
	}

	cpt := uint(0)
	selectorElement := selectorToken.Element()
	elements := tree.Block().Successful().Elements().List()
	for _, oneElement := range elements {
		contents := oneElement.Contents()
		if !oneElement.HasGrammar() {
			if selectorElement.Name() != selectorToken.ReverseName() {
				continue
			}

			if selectorElement.Index() == cpt {
				return app.tokenRefine(selectorToken, contents)
			}

			cpt++
		}

		if oneElement.Grammar().Name() != selectorElement.Name() {
			continue
		}

		if selectorElement.Index() == cpt {
			return app.tokenRefine(selectorToken, contents)
		}

		cpt++
	}

	return []trees.Content{}, nil
}

func (app *application) selectorInstance(selector selectors.Selector, contentList []trees.Content, previous map[string]selectors.Selector) (interface{}, bool, error) {
	fn := selector.Fn()
	inside := selector.Inside()
	insList, err := app.insideInstances(inside, contentList, previous)
	if err != nil {
		return nil, false, err
	}

	if len(insList) <= 0 {
		return nil, false, nil
	}

	if fn.IsSingle() {
		var param interface{}
		if len(insList) > 0 {
			param = insList[0]
		}

		singleFn := fn.Single()
		return singleFn(param)
	}

	multiFn := fn.Multi()
	return multiFn(insList)
}

func (app *application) tokenRefine(token selectors.Token, contents trees.Contents) ([]trees.Content, error) {
	if !token.HasContent() {
		return contents.List(), nil
	}

	pIndex := token.Content()
	list := contents.List()
	listLength := uint(len(list))
	if listLength <= *pIndex {
		str := fmt.Sprintf("the contents cannot be refined because the token selector requires a content (index: %d) but the list (length: %d) is too small", *pIndex, listLength)
		return nil, errors.New(str)
	}

	return []trees.Content{
		list[*pIndex],
	}, nil
}

func (app *application) insideInstances(inside selectors.Inside, contentList []trees.Content, previous map[string]selectors.Selector) ([]interface{}, error) {
	if inside.IsFn() {
		fn := inside.Fn()
		if fn.IsSingle() {
			if len(contentList) < 0 {
				str := fmt.Sprintf("%d content instances were expected in the Content list, %d were provided", 1, len(contentList))
				return nil, errors.New(str)
			}

			singleFn := fn.Single()
			return singleFn(contentList[0])
		}

		multiFn := fn.Multi()
		return multiFn(contentList)
	}

	output := []interface{}{}
	fetchers := inside.Fetchers()
	fetchersList := fetchers.List()
	for _, oneContent := range contentList {
		row := []interface{}{}
		for _, oneFetcher := range fetchersList {
			if oneContent.IsValue() {
				continue
			}

			tree := oneContent.Tree()
			if oneFetcher.IsSelector() {
				selector := oneFetcher.Selector()
				ins, isValid, err := app.selectorFetch(selector, tree, previous)
				if err != nil {
					return nil, err
				}

				if !isValid {
					continue
				}

				row = append(row, ins)
				continue
			}

			recursive := oneFetcher.Recursive()
			if selector, ok := previous[recursive]; ok {
				ins, isValid, err := app.selectorFetch(selector, tree, previous)
				if err != nil {
					return nil, err
				}

				if !isValid {
					continue
				}

				row = append(row, ins)
				continue
			}

			str := fmt.Sprintf("the recursive Selector's Token (name: %s) could not be found in the previous iterations", recursive)
			return nil, errors.New(str)

		}

		output = append(output, row...)
	}

	return output, nil
}
