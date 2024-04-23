package votes

// Vote represents a vote
type Vote struct {
	Message string `json:"message"`
	Ring    string `json:"ring"`
	Account string `json:"account"`
}
