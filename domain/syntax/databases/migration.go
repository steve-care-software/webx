package databases

import (
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/syntax/databases/programs"
)

type migration struct {
	hash        hash.Hash
	previous    Content
	height      uint
	description string
	program     programs.Program
}

func createMigration(
	hash hash.Hash,
	previous Content,
	height uint,
	description string,
) Migration {
	return createMigrationInternally(hash, previous, height, description, nil)
}

func createMigrationWithProgram(
	hash hash.Hash,
	previous Content,
	height uint,
	description string,
	program programs.Program,
) Migration {
	return createMigrationInternally(hash, previous, height, description, program)
}

func createMigrationInternally(
	hash hash.Hash,
	previous Content,
	height uint,
	description string,
	program programs.Program,
) Migration {
	out := migration{
		hash:        hash,
		previous:    previous,
		height:      height,
		description: description,
		program:     program,
	}

	return &out
}

// Hash returns the hash
func (obj *migration) Hash() hash.Hash {
	return obj.hash
}

// Previous returns the previous content
func (obj *migration) Previous() Content {
	return obj.previous
}

// Height returns the height
func (obj *migration) Height() uint {
	return obj.height
}

// Description returns the description
func (obj *migration) Description() string {
	return obj.description
}

// HasProgram returns true if there is a program, false otherwise
func (obj *migration) HasProgram() bool {
	return obj.program != nil
}

// Program returns the program, if any
func (obj *migration) Program() programs.Program {
	return obj.program
}
