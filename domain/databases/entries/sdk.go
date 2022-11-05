package entries

import "github.com/steve-care-software/webx/domain/databases/entities"

// Entries represents entries
type Entries interface {
	List() []Entry
}

// Entry represents an entry
type Entry interface {
	Entity() entities.Entity
	Kind() Kind
	Content() []byte
}

// Kind represents a content kind
type Kind interface {
	IsGrammar() bool
	Grammar() Grammar
	IsProgram() bool
	Program() Program
	IsSelector() bool
	Selector() Selector
}

// Grammar represents a grammar kind
type Grammar interface {
	IsGrammar() bool
	IsChannel() bool
	IsCondition() bool
	IsToken() bool
	IsSuite() bool
	IsLine() bool
	IsElement() bool
	IsCardinality() bool
	IsInstance() bool
	IsEverything() bool
}

// Program represents a program kind
type Program interface {
	IsProgram() bool
	IsInstruction() bool
	IsApplication() bool
	IsAssignment() bool
	IsAttachment() bool
	IsModule() bool
	IsValue() bool
}

// Selector represents a selector kind
type Selector interface {
	IsSelector() bool
	IsToken() bool
	IsInside() bool
	IsElement() bool
	IsFetcher() bool
	IsFn() bool
}

// Tree represents a tree
type Tree interface {
	IsTrees() bool
	IsTree() bool
	IsLine() bool
	IsElements() bool
	IsElement() bool
	IsContents() bool
	IsContent() bool
}
