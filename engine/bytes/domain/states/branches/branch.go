package branches

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/states/branches/layers"
)

type branch struct {
	name      string
	isDeleted bool
	layers    layers.Layers
	metaData  delimiters.Delimiter
	children  Branches
}

func createBranch(
	name string,
	isDeleted bool,
) Branch {
	return createBranchInternally(
		name,
		isDeleted,
		nil,
		nil,
		nil,
	)
}

func createBranchWithLayers(
	name string,
	isDeleted bool,
	layers layers.Layers,
) Branch {
	return createBranchInternally(
		name,
		isDeleted,
		layers,
		nil,
		nil,
	)
}

func createBranchWithMetaData(
	name string,
	isDeleted bool,
	metaData delimiters.Delimiter,
) Branch {
	return createBranchInternally(
		name,
		isDeleted,
		nil,
		metaData,
		nil,
	)
}

func createBranchWithChildren(
	name string,
	isDeleted bool,
	children Branches,
) Branch {
	return createBranchInternally(
		name,
		isDeleted,
		nil,
		nil,
		children,
	)
}

func createBranchWithLayersAndMetaData(
	name string,
	isDeleted bool,
	layers layers.Layers,
	metaData delimiters.Delimiter,
) Branch {
	return createBranchInternally(
		name,
		isDeleted,
		layers,
		metaData,
		nil,
	)
}

func createBranchWithLayersAndChildren(
	name string,
	isDeleted bool,
	layers layers.Layers,
	children Branches,
) Branch {
	return createBranchInternally(
		name,
		isDeleted,
		layers,
		nil,
		children,
	)
}

func createBranchWithMetaDataAndChildren(
	name string,
	isDeleted bool,
	metaData delimiters.Delimiter,
	children Branches,
) Branch {
	return createBranchInternally(
		name,
		isDeleted,
		nil,
		metaData,
		children,
	)
}

func createBranchWithLayersAndMetaDataAndChildren(
	name string,
	isDeleted bool,
	layers layers.Layers,
	metaData delimiters.Delimiter,
	children Branches,
) Branch {
	return createBranchInternally(
		name,
		isDeleted,
		layers,
		metaData,
		children,
	)
}

func createBranchInternally(
	name string,
	isDeleted bool,
	layers layers.Layers,
	metaData delimiters.Delimiter,
	children Branches,
) Branch {
	out := branch{
		name:      name,
		isDeleted: isDeleted,
		layers:    layers,
		metaData:  metaData,
		children:  children,
	}

	return &out
}

// Name returns the name
func (obj *branch) Name() string {
	return obj.name
}

// IsDeleted returns true if deleted, false otherwise
func (obj *branch) IsDeleted() bool {
	return obj.isDeleted
}

// HasLayers returns true if there is layers, false otherwise
func (obj *branch) HasLayers() bool {
	return obj.layers != nil
}

// Layers returns the layers, if any
func (obj *branch) Layers() layers.Layers {
	return obj.layers
}

// HasMetaData returns true if there is metaData, false otherwise
func (obj *branch) HasMetaData() bool {
	return obj.metaData != nil
}

// MetaData returns the metaData, if any
func (obj *branch) MetaData() delimiters.Delimiter {
	return obj.metaData
}

// HasChildren returns true if there is children, false otherwise
func (obj *branch) HasChildren() bool {
	return obj.children != nil
}

// Children returns the children, if any
func (obj *branch) Children() Branches {
	return obj.children
}
