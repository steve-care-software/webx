package applications

type application struct {
	module []byte
	name   []byte
}

func createApplication(
	module []byte,
	name []byte,
) Application {
	out := application{
		module: module,
		name:   name,
	}

	return &out
}

// Module returns the module
func (obj *application) Module() []byte {
	return obj.module
}

// Name returns the name
func (obj *application) Name() []byte {
	return obj.name
}
