package selections

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/trees"
	"github.com/steve-care-software/webx/domain/trees/selections"
)

type application struct {
	builder          selections.Builder
	selectionBuilder selections.SelectionBuilder
	childBuilder     selections.ChildBuilder
}

func createApplication(
	builder selections.Builder,
	selectionBuilder selections.SelectionBuilder,
	childBuilder selections.ChildBuilder,
) Application {
	out := application{
		builder:          builder,
		selectionBuilder: selectionBuilder,
		childBuilder:     childBuilder,
	}

	return &out
}

// Convert converts a tree to a selections instance
func (app *application) Convert(tree trees.Tree, includeChannelBytes bool) (selections.Selection, error) {
	selectionsIns, err := app.convert(tree, includeChannelBytes)
	if err != nil {
		return nil, err
	}

	child, err := app.childBuilder.Create().WithSelections(selectionsIns).Now()
	if err != nil {
		return nil, err
	}

	name := tree.Grammar().Name()
	return app.selectionBuilder.Create().WithElementName(name).WithList([]selections.Child{
		child,
	}).Now()
}

func (app *application) convert(tree trees.Tree, includeChannelBytes bool) (selections.Selections, error) {
	block := tree.Block()
	treeName := tree.Grammar().Name()
	if !block.HasSuccessful() {
		str := fmt.Sprintf("the tree (root: %s) contains no successful line in its root block", treeName)
		return nil, errors.New(str)
	}

	selectionsList := []selections.Selection{}
	elementsList := block.Successful().Elements().List()
	for _, oneElement := range elementsList {
		content := []byte{}
		childList := []selections.Child{}
		contentsList := oneElement.Contents().List()
		for _, oneContent := range contentsList {
			if oneContent.IsValue() {
				val := oneContent.Value()
				if includeChannelBytes && val.HasPrefix() {
					content = append(content, val.Prefix().Bytes(includeChannelBytes)...)
				}

				content = append(content, val.Content())
				continue
			}

			if len(content) > 0 {
				child, err := app.childBuilder.Create().WithContent(content).Now()
				if err != nil {
					return nil, err
				}

				childList = append(childList, child)
				content = []byte{}
			}

			subTree := oneContent.Tree()
			subSelections, err := app.convert(subTree, includeChannelBytes)
			if err != nil {
				return nil, err
			}

			child, err := app.childBuilder.Create().WithSelections(subSelections).Now()
			if err != nil {
				return nil, err
			}

			childList = append(childList, child)
		}

		if len(content) > 0 {
			child, err := app.childBuilder.Create().WithContent(content).Now()
			if err != nil {
				return nil, err
			}

			childList = append(childList, child)
		}

		selectionBuilder := app.selectionBuilder.Create().WithElementName("reverse").WithList(childList)
		if oneElement.HasGrammar() {
			elementName := oneElement.Grammar().Name()
			selectionBuilder.WithElementName(elementName)
		}

		selection, err := selectionBuilder.Now()
		if err != nil {
			return nil, err
		}

		selectionsList = append(selectionsList, selection)
	}

	if includeChannelBytes && tree.HasSuffix() {
		suffixes := tree.Suffix().List()
		for _, oneSuffix := range suffixes {
			suffixSelections, err := app.convert(oneSuffix, includeChannelBytes)
			if err != nil {
				return nil, err
			}

			selectionsList = append(selectionsList, suffixSelections.List()...)
		}
	}

	return app.builder.Create().
		WithTreeName(treeName).
		WithList(selectionsList).
		Now()
}

// Search search in the selections using a criteria instance
func (app *application) Search(selection selections.Selection, criteria criterias.Criteria) (selections.Selection, error) {
	if !criteria.HasNext() {
		current := criteria.Current()
		return app.searchTail(selection, current)
	}

	next := criteria.Next()
	return app.searchNode(selection, next)
}

func (app *application) searchNode(selectionIns selections.Selection, node criterias.Node) (selections.Selection, error) {
	if node.IsNext() {
		next := node.Next()
		return app.Search(selectionIns, next)
	}

	tail := node.Tail()
	return app.searchTail(selectionIns, tail)
}

func (app *application) searchTail(selectionIns selections.Selection, tail criterias.Tail) (selections.Selection, error) {
	name := tail.Name()
	elementName := selectionIns.ElementName()
	if elementName == name {
		childList := selectionIns.List()
		if !tail.HasDelimiter() {
			return app.selectionBuilder.Create().
				WithElementName(elementName).
				WithList(childList).
				Now()
		}

		delimiter := tail.Delimiter()
		index := delimiter.Index()
		pAmount := delimiter.Amount()
		if uint(len(childList)) <= index {
			str := fmt.Sprintf("there is %d children in the element (name: %s), therefore the requested child (index: %d) cannot be selected", len(childList), elementName, index)
			return nil, errors.New(str)
		}

		childList = childList[index:]
		if pAmount != nil {
			if uint(len(childList)) > *pAmount {
				childList = childList[:*pAmount]
			}
		}

		return app.selectionBuilder.Create().
			WithElementName(elementName).
			WithList(childList).
			Now()
	}

	childResults := []selections.Child{}
	childList := selectionIns.List()
	for _, oneChild := range childList {
		if oneChild.IsContent() {
			continue
		}

		childSections := oneChild.Selections()
		childSelections := childSections.List()
		for _, oneSelection := range childSelections {
			res, err := app.searchTail(oneSelection, tail)
			if err != nil {
				continue
			}

			subChild := res.List()
			childResults = append(childResults, subChild...)
		}
	}

	return app.selectionBuilder.Create().
		WithElementName(elementName).
		WithList(childResults).
		Now()

}
