package instructions

type failure struct {
	index         uint
	code          uint
	raisedInLayer bool
	message       string
}
