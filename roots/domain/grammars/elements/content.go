package elements

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

type content struct {
	pValue      *uint8
	pExternal   *hash.Hash
	pToken      *hash.Hash
	pEverything *hash.Hash
	pRecursive  *hash.Hash
}

func createContentWithValue(
	pValue *uint8,
) Content {
	return createContentInternally(pValue, nil, nil, nil, nil)
}

func createContentWithExternal(
	pExternal *hash.Hash,
) Content {
	return createContentInternally(nil, pExternal, nil, nil, nil)
}

func createContentWithToken(
	pToken *hash.Hash,
) Content {
	return createContentInternally(nil, nil, pToken, nil, nil)
}

func createContentWithEverything(
	pEverything *hash.Hash,
) Content {
	return createContentInternally(nil, nil, nil, pEverything, nil)
}

func createContentWithRecursive(
	pRecursive *hash.Hash,
) Content {
	return createContentInternally(nil, nil, nil, nil, pRecursive)
}

func createContentInternally(
	pValue *uint8,
	pExternal *hash.Hash,
	pToken *hash.Hash,
	pEverything *hash.Hash,
	pRecursive *hash.Hash,
) Content {
	out := content{
		pValue:      pValue,
		pExternal:   pExternal,
		pToken:      pToken,
		pEverything: pEverything,
		pRecursive:  pRecursive,
	}

	return &out
}

// IsValue returns true if there is a value, false otherwise
func (obj *content) IsValue() bool {
	return obj.pValue != nil
}

// Value returns the value, if any
func (obj *content) Value() *uint8 {
	return obj.pValue
}

// IsExternal returns true if there is an pExternal, false otherwise
func (obj *content) IsExternal() bool {
	return obj.pExternal != nil
}

// External returns the pExternal, if any
func (obj *content) External() *hash.Hash {
	return obj.pExternal
}

// IsToken returns true if there is a pToken, false otherwise
func (obj *content) IsToken() bool {
	return obj.pToken != nil
}

// Token returns the pToken, if any
func (obj *content) Token() *hash.Hash {
	return obj.pToken
}

// IsEverything returns true if there is an pEverything, false otherwise
func (obj *content) IsEverything() bool {
	return obj.pEverything != nil
}

// Everything returns the pEverything, if any
func (obj *content) Everything() *hash.Hash {
	return obj.pEverything
}

// IsRecursive returns true if there is a pRecursive, false otherwise
func (obj *content) IsRecursive() bool {
	return obj.pRecursive != nil
}

// Recursive returns the pRecursive, if any
func (obj *content) Recursive() *hash.Hash {
	return obj.pRecursive
}
