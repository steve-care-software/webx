package branches

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
)

type branch struct {
	name        string
	description string
	isDeleted   bool
	states      delimiters.Delimiter
	metaData    delimiters.Delimiter
	children    Branches
}

func createBranch(
	name string,
	description string,
	isDeleted bool,
) Branch {
	return createBranchInternally(
		name,
		description,
		isDeleted,
		nil,
		nil,
		nil,
	)
}

func createBranchWithStates(
	name string,
	description string,
	isDeleted bool,
	states delimiters.Delimiter,
) Branch {
	return createBranchInternally(
		name,
		description,
		isDeleted,
		states,
		nil,
		nil,
	)
}

func createBranchWithMetaData(
	name string,
	description string,
	isDeleted bool,
	metaData delimiters.Delimiter,
) Branch {
	return createBranchInternally(
		name,
		description,
		isDeleted,
		nil,
		metaData,
		nil,
	)
}

func createBranchWithChildren(
	name string,
	description string,
	isDeleted bool,
	children Branches,
) Branch {
	return createBranchInternally(
		name,
		description,
		isDeleted,
		nil,
		nil,
		children,
	)
}

func createBranchWithStatesAndMetaData(
	name string,
	description string,
	isDeleted bool,
	states delimiters.Delimiter,
	metaData delimiters.Delimiter,
) Branch {
	return createBranchInternally(
		name,
		description,
		isDeleted,
		states,
		metaData,
		nil,
	)
}

func createBranchWithStatesAndChildren(
	name string,
	description string,
	isDeleted bool,
	states delimiters.Delimiter,
	children Branches,
) Branch {
	return createBranchInternally(
		name,
		description,
		isDeleted,
		states,
		nil,
		children,
	)
}

func createBranchWithMetaDataAndChildren(
	name string,
	description string,
	isDeleted bool,
	metaData delimiters.Delimiter,
	children Branches,
) Branch {
	return createBranchInternally(
		name,
		description,
		isDeleted,
		nil,
		metaData,
		children,
	)
}

func createBranchWithStatesAndMetaDataAndChildren(
	name string,
	description string,
	isDeleted bool,
	states delimiters.Delimiter,
	metaData delimiters.Delimiter,
	children Branches,
) Branch {
	return createBranchInternally(
		name,
		description,
		isDeleted,
		states,
		metaData,
		children,
	)
}

func createBranchInternally(
	name string,
	description string,
	isDeleted bool,
	states delimiters.Delimiter,
	metaData delimiters.Delimiter,
	children Branches,
) Branch {
	out := branch{
		name:        name,
		description: description,
		isDeleted:   isDeleted,
		states:      states,
		metaData:    metaData,
		children:    children,
	}

	return &out
}

// Name returns the name
func (obj *branch) Name() string {
	return obj.name
}

// Description returns the description
func (obj *branch) Description() string {
	return obj.description
}

// IsDeleted returns true if deleted, false otherwise
func (obj *branch) IsDeleted() bool {
	return obj.isDeleted
}

// HasStates returns true if there is states, false otherwise
func (obj *branch) HasStates() bool {
	return obj.states != nil
}

// States returns the states, if any
func (obj *branch) States() delimiters.Delimiter {
	return obj.states
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
