package selectors

import "github.com/steve-care-software/webx/domain/selectors"

// Application represents the create's selector application
type Application interface {
	Execute() (selectors.Selector, error)
}
