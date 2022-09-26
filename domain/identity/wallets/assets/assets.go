package assets

import "github.com/steve-care-software/syntax/domain/identity/cryptography/hash"

type assets struct {
	list []Asset
}

func createAssets(
	list []Asset,
) Assets {
	out := assets{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *assets) List() []Asset {
	return obj.list
}

// FetchByUnits retrieves an asset list by unit hashes
func (obj *assets) FetchByUnits(unitHashes []hash.Hash) ([]Asset, error) {
	return nil, nil
}
