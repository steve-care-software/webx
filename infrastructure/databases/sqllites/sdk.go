package sqllites

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
)

const contextDoesNotExistsErrorStr = "the context (%d) does NOT exists"

// listInstanceToElementHashesListFn takes a list instance and return its element's hashes
type listInstanceToElementHashesListFn func(ins instances.Instance) ([]hash.Hash, error)

// elementsToListInstanceFn takes a list of elements and returns a list instance
type elementsToListInstanceFn func(input []interface{}) (instances.Instance, error)

// buildInstanceFn takes values and build an Instance instance
type buildInstanceFn func(values map[string]interface{}) (instances.Instance, error)

// callMethodOnInstanceFn calls a method related to a field on instance, then returns its value
type callMethodOnInstanceFn func(ins instances.Instance, fieldName string) (bool, interface{}, error)

const resourceNameDelimiter = "_"
const endOfLine = "\n"
const connectionNameDelimiter = "$"
