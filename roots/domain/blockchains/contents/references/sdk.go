package references

import (
	"time"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

const pointerSize = 8 * 2
const blockchainKeySize = hash.Size + pointerSize + 8
const contentKeySize = blockchainKeySize + 8 + hash.Size

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	contentAdapter := NewContentAdapter()
	blockchainAdapter := NewBlockchainAdapter()
	builder := NewBuilder()
	return createAdapter(contentAdapter, blockchainAdapter, builder)
}

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	contentFactory := NewContentFactory()
	return createBuilder(contentFactory)
}

// NewContentAdapter creates a new content adapter
func NewContentAdapter() ContentAdapter {
	contentKeysAdapter := NewContentKeysAdapter()
	builder := NewContentBuilder()
	return createContentAdapter(contentKeysAdapter, builder)
}

// NewContentFactory creates a new content factory
func NewContentFactory() ContentFactory {
	builder := NewContentBuilder()
	return createContentFactory(builder)
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// NewBlockchainAdapter creates a new blockchain adapter
func NewBlockchainAdapter() BlockchainAdapter {
	blockchainKeyAdapter := NewBlockchainKeyAdapter()
	blockchainKeysAdapter := NewBlockchainKeysAdapter()
	builder := NewBlockchainBuilder()
	return createBlockchainAdapter(blockchainKeyAdapter, blockchainKeysAdapter, builder)
}

// NewBlockchainBuilder creates a new blockchain builder
func NewBlockchainBuilder() BlockchainBuilder {
	return createBlockchainBuilder()
}

// NewContentKeysAdapter creates a new content keys adapter
func NewContentKeysAdapter() ContentKeysAdapter {
	adapter := NewContentKeyAdapter()
	builder := NewContentKeysBuilder()
	return createContentKeysAdapter(adapter, builder)
}

// NewContentKeysBuilder creates a new content keys builder
func NewContentKeysBuilder() ContentKeysBuilder {
	return createContentKeysBuilder()
}

// NewContentKeyAdapter creates a new content key adapter
func NewContentKeyAdapter() ContentKeyAdapter {
	hashAdapter := hash.NewAdapter()
	pointerAdapter := NewPointerAdapter()
	builder := NewContentKeyBuilder()
	return createContentKeyAdapter(hashAdapter, pointerAdapter, builder)
}

// NewContentKeyBuilder createsa new content key builder
func NewContentKeyBuilder() ContentKeyBuilder {
	return createContentKeyBuilder()
}

// NewBlockchainKeysAdapter creates a new blockchain keys adapter
func NewBlockchainKeysAdapter() BlockchainKeysAdapter {
	adapter := NewBlockchainKeyAdapter()
	builder := NewBlockchainKeysBuilder()
	return createBlockchainKeysAdapter(adapter, builder)
}

// NewBlockchainKeysBuilder creates a new blockchain keys builder
func NewBlockchainKeysBuilder() BlockchainKeysBuilder {
	return createBlockchainKeysBuilder()
}

// NewBlockchainKeyAdapter creates a new blockchain key adapter
func NewBlockchainKeyAdapter() BlockchainKeyAdapter {
	hashAdapter := hash.NewAdapter()
	pointerAdapter := NewPointerAdapter()
	builder := NewBlockchainKeyBuilder()
	return createBlockchainKeyAdapter(hashAdapter, pointerAdapter, builder)
}

// NewBlockchainKeyBuilder creates a new blockchain key builder
func NewBlockchainKeyBuilder() BlockchainKeyBuilder {
	return createBlockchainKeyBuilder()
}

// NewPointerAdapter creates a new pointer adapter
func NewPointerAdapter() PointerAdapter {
	builder := NewPointerBuilder()
	return createPointerAdapter(builder)
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	return createPointerBuilder()
}

// Adapter represents a reference adapter
type Adapter interface {
	ToContent(ins Reference) ([]byte, error)
	ToReference(content []byte) (Reference, error)
}

// Factory represents a reference factory
type Factory interface {
	Create() (Reference, error)
}

// Builder represents a reference builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithBlockchain(blockchain Blockchain) Builder
	Now() (Reference, error)
}

// Reference represents the reference
type Reference interface {
	Content() Content
	HasBlockchain() bool
	Blockchain() Blockchain
}

// ContentFactory represents a content factory
type ContentFactory interface {
	Create() (Content, error)
}

// ContentAdapter represents a content adapter
type ContentAdapter interface {
	ToContent(ins Content) ([]byte, error)
	ToInstance(content []byte) (Content, error)
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithActive(active ContentKeys) ContentBuilder
	WithPendings(pendings ContentKeys) ContentBuilder
	WithDeleted(deleted ContentKeys) ContentBuilder
	Now() (Content, error)
}

// Content represents the content reference
type Content interface {
	HasActive() bool
	Active() ContentKeys
	HasPendings() bool
	Pendings() ContentKeys
	HasDeleted() bool
	Deleted() ContentKeys
}

// BlockchainAdapter represents a blockchain adapter
type BlockchainAdapter interface {
	ToContent(ins Blockchain) ([]byte, error)
	ToBlockchain(content []byte) (Blockchain, error)
}

// BlockchainBuilder represents a blockchain builder
type BlockchainBuilder interface {
	Create() BlockchainBuilder
	WithChain(chain BlockchainKey) BlockchainBuilder
	WithBlocks(blocks BlockchainKeys) BlockchainBuilder
	WithTransactions(trx BlockchainKeys) BlockchainBuilder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain reference
type Blockchain interface {
	Chain() BlockchainKey
	Blocks() BlockchainKeys
	Transactions() BlockchainKeys
}

// ContentKeysAdapter represents the content keys adapter
type ContentKeysAdapter interface {
	ToContent(ins ContentKeys) ([]byte, error)
	ToContentKeys(content []byte) (ContentKeys, error)
}

// ContentKeysBuilder represents a content keys builder
type ContentKeysBuilder interface {
	Create() ContentKeysBuilder
	WithList(list []ContentKey) ContentKeysBuilder
	Now() (ContentKeys, error)
}

// ContentKeys represents content keys
type ContentKeys interface {
	List() []ContentKey
	Fetch(hash hash.Hash) (ContentKey, error)
}

// ContentKeyAdapter represents the content key adapter
type ContentKeyAdapter interface {
	ToContent(ins ContentKey) ([]byte, error)
	ToContentKey(content []byte) (ContentKey, error)
}

// ContentKeyBuilder represents a content key builder
type ContentKeyBuilder interface {
	Create() ContentKeyBuilder
	WithHash(hash hash.Hash) ContentKeyBuilder
	WithKind(kind uint) ContentKeyBuilder
	WithContent(content Pointer) ContentKeyBuilder
	WithTransaction(trx hash.Hash) ContentKeyBuilder
	CreatedOn(createdOn time.Time) ContentKeyBuilder
	Now() (ContentKey, error)
}

// ContentKey represents a content key
type ContentKey interface {
	BlockchainKey
	Kind() uint
	Transaction() hash.Hash
}

// BlockchainKeysAdapter represents the blockchain keys adapter
type BlockchainKeysAdapter interface {
	ToContent(ins BlockchainKeys) ([]byte, error)
	ToBlockchainKeys(content []byte) (BlockchainKeys, error)
}

// BlockchainKeysBuilder represents a blockchain keys builder
type BlockchainKeysBuilder interface {
	Create() BlockchainKeysBuilder
	WithList(list []BlockchainKey) BlockchainKeysBuilder
	Now() (BlockchainKeys, error)
}

// BlockchainKeys represents blockchain keys
type BlockchainKeys interface {
	List() []BlockchainKey
	Fetch(hash hash.Hash) (BlockchainKey, error)
}

// BlockchainKeyAdapter represents the blockchain key adapter
type BlockchainKeyAdapter interface {
	ToContent(ins BlockchainKey) ([]byte, error)
	ToBlockchainKey(content []byte) (BlockchainKey, error)
}

// BlockchainKeyBuilder represents a blockchain key builder
type BlockchainKeyBuilder interface {
	Create() BlockchainKeyBuilder
	WithHash(hash hash.Hash) BlockchainKeyBuilder
	WithContent(content Pointer) BlockchainKeyBuilder
	CreatedOn(createdOn time.Time) BlockchainKeyBuilder
	Now() (BlockchainKey, error)
}

// BlockchainKey represents a blockchain key
type BlockchainKey interface {
	Hash() hash.Hash
	Content() Pointer
	CreatedOn() time.Time
}

// PointerAdapter represents the pointer adapter
type PointerAdapter interface {
	ToContent(ins Pointer) ([]byte, error)
	ToPointer(content []byte) (Pointer, error)
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithLength(length uint) PointerBuilder
	From(from uint) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	From() uint
	Length() uint
}
