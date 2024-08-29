package stackframes

import "github.com/steve-care-software/webx/engine/domain/stacks"

// Application represents a stackframe application
type Application interface {
	Root() error                  // selects the root stack
	Navigate(path []string) error // navigate to the stack where the path points
	Fetch() (stacks.Stack, error) // fetches the stack where our cursor points to
	Push() error                  // push the stack where our cursor points to
	Pop() error                   // pop the stack where our cursor points to
}
