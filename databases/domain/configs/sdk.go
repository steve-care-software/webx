package configs

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the config builder
type Builder interface {
	Create() Builder
	WithRatios(ratios []float64) Builder
	WithDefault(defValue uint) Builder
	Now() (Config, error)
}

// Config represents a push configuration
type Config interface {
	Amount() uint
	Ratios() []float64
	Default() uint
}
