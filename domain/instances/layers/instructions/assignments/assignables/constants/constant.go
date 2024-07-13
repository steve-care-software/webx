package constants

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

type constant struct {
	hash    hash.Hash
	pBool   *bool
	pString *string
	pInt    *int
	pUint   *uint
	pFloat  *float64
	list    Constants
}

func createConstantWithBool(
	hash hash.Hash,
	pBool *bool,
) Constant {
	return createConstantInternally(hash, pBool, nil, nil, nil, nil, nil)
}

func createConstantWithString(
	hash hash.Hash,
	pString *string,
) Constant {
	return createConstantInternally(hash, nil, pString, nil, nil, nil, nil)
}

func createConstantWithInt(
	hash hash.Hash,
	pInt *int,
) Constant {
	return createConstantInternally(hash, nil, nil, pInt, nil, nil, nil)
}

func createConstantWithUint(
	hash hash.Hash,
	pUint *uint,
) Constant {
	return createConstantInternally(hash, nil, nil, nil, pUint, nil, nil)
}

func createConstantWithFloat(
	hash hash.Hash,
	pFloat *float64,
) Constant {
	return createConstantInternally(hash, nil, nil, nil, nil, pFloat, nil)
}

func createConstantWithList(
	hash hash.Hash,
	list Constants,
) Constant {
	return createConstantInternally(hash, nil, nil, nil, nil, nil, list)
}

func createConstantInternally(
	hash hash.Hash,
	pBool *bool,
	pString *string,
	pInt *int,
	pUint *uint,
	pFloat *float64,
	list Constants,
) Constant {
	out := constant{
		hash:    hash,
		pBool:   pBool,
		pString: pString,
		pInt:    pInt,
		pUint:   pUint,
		pFloat:  pFloat,
		list:    list,
	}

	return &out
}

// Hash returns the hash
func (obj *constant) Hash() hash.Hash {
	return obj.hash
}

// IsBool returns true if bool, false otherwise
func (obj *constant) IsBool() bool {
	return obj.pBool != nil
}

// Bool returns the bool, if any
func (obj *constant) Bool() *bool {
	return obj.pBool
}

// IsString returns true if string, false otherwise
func (obj *constant) IsString() bool {
	return obj.pString != nil
}

// String returns the string, if any
func (obj *constant) String() *string {
	return obj.pString
}

// IsInt returns true if int, false otherwise
func (obj *constant) IsInt() bool {
	return obj.pInt != nil
}

// Int returns the int, if any
func (obj *constant) Int() *int {
	return obj.pInt
}

// IsUint returns true if uint, false otherwise
func (obj *constant) IsUint() bool {
	return obj.pUint != nil
}

// Uint returns the uint, if any
func (obj *constant) Uint() *uint {
	return obj.pUint
}

// IsFloat returns true if float, false otherwise
func (obj *constant) IsFloat() bool {
	return obj.pFloat != nil
}

// Float returns the float, if any
func (obj *constant) Float() *float64 {
	return obj.pFloat
}

// IsList returns true if list, false otherwise
func (obj *constant) IsList() bool {
	return obj.list != nil
}

// List returns the list, if any
func (obj *constant) List() Constants {
	return obj.list
}
