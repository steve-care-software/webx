package databases

type sizeInBytesBuilder struct {
	maxAmount uint
	amount    uint
}

func createSizeInBytesBuilder() SizeInBytesBuilder {
	out := sizeInBytesBuilder{
		maxAmount: 0,
		amount:    0,
	}

	return &out
}

// Create initializes the builder
func (app *sizeInBytesBuilder) Create() SizeInBytesBuilder {
	return createSizeInBytesBuilder()
}

// WithMaxAmount adds a max amount to the builder
func (app *sizeInBytesBuilder) WithMaxAmount(maxAmount uint) SizeInBytesBuilder {
	app.maxAmount = maxAmount
	return app
}

// WithAmount adds an amount to the builder
func (app *sizeInBytesBuilder) WithAmount(amount uint) SizeInBytesBuilder {
	app.amount = amount
	return app
}

// Now builds a new SizeInBytes instance
func (app *sizeInBytesBuilder) Now() (SizeInBytes, error) {
	return createSizeInBytes(app.maxAmount, app.amount), nil
}
