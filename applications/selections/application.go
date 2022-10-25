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
	elementBuilder   selections.ElementBuilder
	childrenBuilder  selections.ChildrenBuilder
	childBuilder     selections.ChildBuilder
}

func createApplication(
	builder selections.Builder,
	selectionBuilder selections.SelectionBuilder,
	elementBuilder selections.ElementBuilder,
	childrenBuilder selections.ChildrenBuilder,
	childBuilder selections.ChildBuilder,
) Application {
	out := application{
		builder:          builder,
		selectionBuilder: selectionBuilder,
		elementBuilder:   elementBuilder,
		childrenBuilder:  childrenBuilder,
		childBuilder:     childBuilder,
	}

	return &out
}

// Convert converts a tree to a selections instance
func (app *application) Convert(tree trees.Tree, includeChannelBytes bool) (selections.Selections, error) {
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
				child, err := app.childBuilder.Create().WithBytes(content).Now()
				if err != nil {
					return nil, err
				}

				childList = append(childList, child)
				content = []byte{}
			}

			subTree := oneContent.Tree()
			subSelections, err := app.Convert(subTree, includeChannelBytes)
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
			child, err := app.childBuilder.Create().WithBytes(content).Now()
			if err != nil {
				return nil, err
			}

			childList = append(childList, child)
		}

		childrenBuilder := app.childrenBuilder.Create().WithElementName("reverse").WithList(childList)
		if oneElement.HasGrammar() {
			elementName := oneElement.Grammar().Name()
			childrenBuilder.WithElementName(elementName)
		}

		children, err := childrenBuilder.Now()
		if err != nil {
			return nil, err
		}

		elementBuilder := app.elementBuilder.Create().WithValue(oneElement)
		if includeChannelBytes {
			elementBuilder.IncludeChannelBytes()
		}

		element, err := elementBuilder.Now()
		if err != nil {
			return nil, err
		}

		selection, err := app.selectionBuilder.Create().
			WithElement(element).
			WithChildren(children).
			Now()

		if err != nil {
			return nil, err
		}

		selectionsList = append(selectionsList, selection)
	}

	return app.builder.Create().
		WithTreeName(treeName).
		WithList(selectionsList).
		Now()
}

// Search search in the selections using a criteria instance
func (app *application) Search(selections selections.Selections, criteria criterias.Criteria) (selections.Selections, error) {
	name := criteria.Name()
	selected, err := app.refine(selections, name)
	if err != nil {
		return nil, err
	}

	if criteria.HasChild() {
		childCriteria := criteria.Child()
		return app.Search(selected, childCriteria)
	}

	return selected, nil
}

func (app *application) refine(ins selections.Selections, name string) (selections.Selections, error) {
	list := ins.List()
	selectedList := []selections.Selection{}
	for _, oneSelection := range list {
		if oneSelection.ElementName() == name {
			selectedList = append(selectedList, oneSelection)
			continue
		}

		content := oneSelection.Content()
		if content.HasChildren() {
			selectedChildList := []selections.Child{}
			children := content.Children()
			childList := children.List()
			for _, oneChild := range childList {
				if oneChild.IsBytes() {
					fmt.Printf("\n%s\n", oneChild.Bytes())
					continue
				}

				next := oneChild.Selections()
				subSelections, err := app.refine(next, name)
				if err != nil {
					return nil, err
				}

				selectedChild, err := app.childBuilder.Create().WithSelections(subSelections).Now()
				if err != nil {
					return nil, err
				}

				selectedChildList = append(selectedChildList, selectedChild)
			}

			if len(selectedChildList) <= 0 {
				continue
			}

			elementName := children.ElementName()
			selectedChildren, err := app.childrenBuilder.Create().
				WithElementName(elementName).
				WithList(selectedChildList).
				Now()

			if err != nil {
				return nil, err
			}

			selectedSelection, err := app.selectionBuilder.Create().
				WithChildren(selectedChildren).
				Now()

			if err != nil {
				return nil, err
			}

			selectedList = append(selectedList, selectedSelection)
		}
	}

	return app.builder.Create().
		WithList(selectedList).
		Now()
}
