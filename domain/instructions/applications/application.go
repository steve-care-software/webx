package applications

type application struct {
	module string
	name   string
}

func createApplication(
	module string,
	name string,
) Application {
	out := application{
		module: module,
		name:   name,
	}

	return &out
}

// Module returns the module
func (obj *application) Module() string {
	return obj.module
}

// Name returns the name
func (obj *application) Name() string {
	return obj.name
}
