package assets

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
