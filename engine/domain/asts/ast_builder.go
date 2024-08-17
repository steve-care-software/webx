package asts

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/hash"
)

type astAstBuilder struct {
	hashAdapter hash.Adapter
	library     NFTs
	entry       hash.Hash
}

func createAstBuilder(
	hashAdapter hash.Adapter,
) AstBuilder {
	out := astAstBuilder{
		hashAdapter: hashAdapter,
		library:     nil,
		entry:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *astAstBuilder) Create() AstBuilder {
	return createAstBuilder(
		app.hashAdapter,
	)
}

// WithLibrary adds a library to the builder
func (app *astAstBuilder) WithLibrary(library NFTs) AstBuilder {
	app.library = library
	return app
}

// WithEntry adds an entry to the builder
func (app *astAstBuilder) WithEntry(entry hash.Hash) AstBuilder {
	app.entry = entry
	return app
}

func (app *astAstBuilder) fetchComplexity(hashStr string) (map[string]uint, error) {
	// create the hash from the string:
	pHash, err := app.hashAdapter.FromString(hashStr)
	if err != nil {
		return nil, err
	}

	// fetch the nft from the library:
	nft, err := app.library.Fetch(*pHash)
	if err != nil {
		return nil, err
	}

	// fetch the complexity of that nft:
	complexity := map[string]uint{}

	// for each sub-nft, fetch their sub-complexity and add the current complexity:
	directComplexity := nft.Complexity()
	for subHashStr, score := range directComplexity {
		// fetch the sub complexity:
		subComplexity, err := app.fetchComplexity(subHashStr)
		if err != nil {
			return nil, err
		}

		// merge the sub complexity to the output:
		for subHashStr, subScore := range subComplexity {
			if currentScore, ok := complexity[subHashStr]; ok {
				complexity[subHashStr] = currentScore + subScore
				continue
			}

			complexity[subHashStr] = subScore
		}

		// add the current sub score to the score if it already exists:
		if currentScore, ok := complexity[subHashStr]; ok {
			complexity[subHashStr] = currentScore + score
			continue
		}

		// it only exists directly, so add the score:
		complexity[subHashStr] = score
	}

	return complexity, nil
}

// Now builds a new AST instance
func (app *astAstBuilder) Now() (AST, error) {
	if app.library == nil {
		return nil, errors.New("the library is mandatory in order to build an AST instance")
	}

	if app.entry == nil {
		return nil, errors.New("the entry hash is mandatory in order to build an AST instance")
	}

	complexity, err := app.fetchComplexity(app.entry.String())
	if err != nil {
		return nil, err
	}

	return createAST(
		app.library,
		app.entry,
		complexity,
	), nil
}