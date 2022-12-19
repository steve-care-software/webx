package configs

type config struct {
	ratios       []float64
	defaultValue uint
}

func createConfig(
	ratios []float64,
	defaultValue uint,
) Config {
	out := config{
		ratios:       ratios,
		defaultValue: defaultValue,
	}

	return &out
}

// Amount returns the amount
func (obj *config) Amount() uint {
	return uint(len(obj.ratios))
}

// Ratios returns the ratios
func (obj *config) Ratios() []float64 {
	return obj.ratios
}

// Default returns the default amount
func (obj *config) Default() uint {
	return obj.defaultValue
}
