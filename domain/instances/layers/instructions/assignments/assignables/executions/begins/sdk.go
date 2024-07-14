package begins

// Begin represents a begin
type Begin interface {
	DbPath() string
	Context() string
}
