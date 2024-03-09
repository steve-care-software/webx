package actions

// Action represents an action
type Action interface {
	IsAmount() bool
	IsRetrieve() bool
	Retrieve() string
	IsRollback() bool
	Rollback() string
}
