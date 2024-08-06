package commands

// Commands represents a commands
type Commands interface {
	List() []Command
}

// Command represents a command
type Command interface {
}
