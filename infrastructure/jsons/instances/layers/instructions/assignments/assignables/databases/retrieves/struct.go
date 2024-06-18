package retrieves

// Retrieve represents a retrieve
type Retrieve struct {
	IsList   bool   `json:"is_list"`
	Exists   string `json:"exists"`
	Retrieve string `json:"retrieve"`
}
