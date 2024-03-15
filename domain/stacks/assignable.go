package stacks

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/queries"
	stack_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

type assignable struct {
	pBool        *bool
	bytes        []byte
	hash         hash.Hash
	hashList     []hash.Hash
	stringList   []string
	pUnsignedInt *uint
	account      stack_accounts.Account
	instance     instances.Instance
	query        queries.Query
}

func createAssignableWithBool(
	pBool *bool,
) Assignable {
	return createAssignableInternally(
		pBool,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithBytes(
	bytes []byte,
) Assignable {
	return createAssignableInternally(
		nil,
		bytes,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithHash(
	hash hash.Hash,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		hash,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithHashList(
	hashList []hash.Hash,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		hashList,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithStringList(
	stringList []string,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		stringList,
		nil,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithUnsignedInt(
	pUnsignedInt *uint,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		pUnsignedInt,
		nil,
		nil,
		nil,
	)
}

func createAssignableWithAccount(
	account stack_accounts.Account,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		account,
		nil,
		nil,
	)
}

func createAssignableWithInstance(
	instance instances.Instance,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		instance,
		nil,
	)
}

func createAssignableWithQuery(
	query queries.Query,
) Assignable {
	return createAssignableInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		query,
	)
}

func createAssignableInternally(
	pBool *bool,
	bytes []byte,
	hash hash.Hash,
	hashList []hash.Hash,
	stringList []string,
	pUnsignedInt *uint,
	account stack_accounts.Account,
	instance instances.Instance,
	query queries.Query,
) Assignable {
	out := assignable{
		pBool:        pBool,
		bytes:        bytes,
		hash:         hash,
		hashList:     hashList,
		stringList:   stringList,
		pUnsignedInt: pUnsignedInt,
		account:      account,
		instance:     instance,
		query:        query,
	}

	return &out
}

// IsBool returns true if bool, false otherwise
func (obj *assignable) IsBool() bool {
	return obj.pBool != nil
}

// Bool returns bool, if any
func (obj *assignable) Bool() *bool {
	return obj.pBool
}

// IsBytes returns true if bytes, false otherwise
func (obj *assignable) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns bytes, if any
func (obj *assignable) Bytes() []byte {
	return obj.bytes
}

// IsHash returns true if hash, false otherwise
func (obj *assignable) IsHash() bool {
	return obj.hash != nil
}

// Hash returns hash, if any
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}

// IsHashList returns true if hash list, false otherwise
func (obj *assignable) IsHashList() bool {
	return obj.hashList != nil
}

// HashList returns hashList, if any
func (obj *assignable) HashList() []hash.Hash {
	return obj.hashList
}

// IsStringList returns true if string list, false otherwise
func (obj *assignable) IsStringList() bool {
	return obj.stringList != nil
}

// StringList returns stringList, if any
func (obj *assignable) StringList() []string {
	return obj.stringList
}

// IsUnsignedInt returns true if unsigned int, false otherwise
func (obj *assignable) IsUnsignedInt() bool {
	return obj.pUnsignedInt != nil
}

// UnsignedInt returns unsigned int, if any
func (obj *assignable) UnsignedInt() *uint {
	return obj.pUnsignedInt
}

// IsAccount returns true if account, false otherwise
func (obj *assignable) IsAccount() bool {
	return obj.account != nil
}

// Account returns account, if any
func (obj *assignable) Account() stack_accounts.Account {
	return obj.account
}

// IsInstance returns true if instance, false otherwise
func (obj *assignable) IsInstance() bool {
	return obj.instance != nil
}

// Instance returns instance, if any
func (obj *assignable) Instance() instances.Instance {
	return obj.instance
}

// IsQuery returns true if query, false otherwise
func (obj *assignable) IsQuery() bool {
	return obj.query != nil
}

// Query returns query, if any
func (obj *assignable) Query() queries.Query {
	return obj.query
}
