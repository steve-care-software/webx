package jsons

// Pointer represents a database pointer
type Pointer struct {
	Head     string   `json:"head"`
	MetaData MetaData `json:"meta_data"`
}
