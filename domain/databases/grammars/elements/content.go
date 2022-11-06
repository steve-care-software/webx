package elements

import "github.com/steve-care-software/webx/domain/databases/entities"

type content struct {
	pValue     *uint8
	external   entities.Identifier
	token      entities.Identifier
	everything entities.Identifier
	recursive  entities.Identifier
}

func createContentWithValue(
	pValue *uint8,
) Content {
	return createContentInternally(pValue, nil, nil, nil, nil)
}

func createContentWithExternal(
	external entities.Identifier,
) Content {
	return createContentInternally(nil, external, nil, nil, nil)
}

func createContentWithToken(
	token entities.Identifier,
) Content {
	return createContentInternally(nil, nil, token, nil, nil)
}

func createContentWithEverything(
	everything entities.Identifier,
) Content {
	return createContentInternally(nil, nil, nil, everything, nil)
}

func createContentWithRecursive(
	recursive entities.Identifier,
) Content {
	return createContentInternally(nil, nil, nil, nil, recursive)
}

func createContentInternally(
	pValue *uint8,
	external entities.Identifier,
	token entities.Identifier,
	everything entities.Identifier,
	recursive entities.Identifier,
) Content {
	out := content{
		pValue:     pValue,
		external:   external,
		token:      token,
		everything: everything,
		recursive:  recursive,
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

// IsExternal returns true if there is an external, false otherwise
func (obj *content) IsExternal() bool {
	return obj.external != nil
}

// External returns the external, if any
func (obj *content) External() entities.Identifier {
	return obj.external
}

// IsToken returns true if there is a token, false otherwise
func (obj *content) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *content) Token() entities.Identifier {
	return obj.token
}

// IsEverything returns true if there is an everything, false otherwise
func (obj *content) IsEverything() bool {
	return obj.everything != nil
}

// Everything returns the everything, if any
func (obj *content) Everything() entities.Identifier {
	return obj.everything
}

// IsRecursive returns true if there is a recursive, false otherwise
func (obj *content) IsRecursive() bool {
	return obj.recursive != nil
}

// Recursive returns the recursive, if any
func (obj *content) Recursive() entities.Identifier {
	return obj.recursive
}
