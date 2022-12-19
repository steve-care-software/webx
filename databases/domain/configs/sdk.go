package configs

// Builder represents the config builder
type Builder interface {
	Create() Builder
	WithAmount(amount uint) Builder
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
