package trees

import "github.com/steve-care-software/webx/domain/databases/entities"

// Builder represents a trees builder
type Builder interface {
	Create() Trees
	WithEntity(entity entities.Entity) Trees
	WithList(list []Tree) Trees
	Now() (Trees, error)
}

// Trees represents trees
type Trees interface {
	Entity() entities.Entity
	List() []Tree
}

// TreeBuilder represents a tree builder
type TreeBuilder interface {
	Create() TreeBuilder
	WithEntity(entity entities.Entity) TreeBuilder
	WithGrammar(grammar entities.Identifier) TreeBuilder
	WithLine(line entities.Identifier) TreeBuilder
	WithSuffix(suffix entities.Identifier) TreeBuilder
	Now() (Tree, error)
}

// Tree represents a tree
type Tree interface {
	Entity() entities.Entity
	Grammar() entities.Identifier
	Line() entities.Identifier
	HasSuffix() bool
	Suffix() entities.Identifier
}
