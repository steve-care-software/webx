package moves

// Move represents a move
type Move interface {
	Name() string
	DevName() string
	DeleteOriginal() bool
}

/*

	Delete(name string) error
	Recover(name string) error
	Purge(name string) error
	PurgeAll() error
	Move(name string, devName string, deleteOriginal bool) error // moves a development iteration to a production iteration inside the current iteration
	Merge(deleteOriginal bool) error

	// data:
	InsertData(data []byte) (delimiters.Delimiter, error)
	UpdateData(original delimiters.Delimiter, updated []byte) error
	DeleteData(delete delimiters.Delimiter) error

*/
