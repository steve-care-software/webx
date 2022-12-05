package references

import "errors"

type contentBlockchainKeysBuilder struct {
	mp   map[string]BlockchainKey
	list []BlockchainKey
}

func createBlockchainKeysBuilder() BlockchainKeysBuilder {
	out := contentBlockchainKeysBuilder{
		mp:   nil,
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBlockchainKeysBuilder) Create() BlockchainKeysBuilder {
	return createBlockchainKeysBuilder()
}

// WithList add contentBlockchainKeys to the builder
func (app *contentBlockchainKeysBuilder) WithList(list []BlockchainKey) BlockchainKeysBuilder {
	app.list = list
	return app
}

// Now builds BlockchainKeys instance
func (app *contentBlockchainKeysBuilder) Now() (BlockchainKeys, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 BlockchainKey in order to build a BlockchainKeys instance")
	}

	mp := map[string]BlockchainKey{}
	for _, oneBlockchainKey := range app.list {
		contentBlockchainKeyname := oneBlockchainKey.Hash().String()
		mp[contentBlockchainKeyname] = oneBlockchainKey
	}

	return createBlockchainKeys(mp, app.list), nil
}
