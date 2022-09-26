package commands

type commands struct {
	list []Command
}

func createCommands(
	list []Command,
) Commands {
	out := commands{
		list: list,
	}

	return &out
}

// List returns the commands
func (obj *commands) List() []Command {
	return obj.list
}
