package rules

import "github.com/steve-care-software/webx/engine/domain/nfts"

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	rulesBuilder := NewBuilder()
	ruleBuilder := NewRuleBuilder()
	nftsBuilder := nfts.NewBuilder()
	nftBuilder := nfts.NewNFTBuilder()
	return createAdapter(
		rulesBuilder,
		ruleBuilder,
		nftsBuilder,
		nftBuilder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewRuleBuilder creates a new rule builder
func NewRuleBuilder() RuleBuilder {
	return createRuleBuilder()
}

// Adapter represents the rule adapter
type Adapter interface {
	RuleToNFT(ins Rule) (nfts.NFT, error)
	NFTToInstance(root nfts.NFT) (Rule, error)
}

// Builder represents a rule list
type Builder interface {
	Create() Builder
	WithList(list []Rule) Builder
	Now() (Rules, error)
}

// Rules represents rules
type Rules interface {
	List() []Rule
	Fetch(name string) (Rule, error)
}

// RuleBuilder represents a rule builder
type RuleBuilder interface {
	Create() RuleBuilder
	WithName(name string) RuleBuilder
	WithBytes(bytes []byte) RuleBuilder
	Now() (Rule, error)
}

// Rule represents a rule
type Rule interface {
	Name() string
	Bytes() []byte
}
