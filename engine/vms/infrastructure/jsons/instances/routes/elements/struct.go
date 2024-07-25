package elements

// Element represents an element
type Element struct {
	Layer  string `json:"layer"`
	Bytes  []byte `json:"bytes"`
	String string `json:"string"`
}
