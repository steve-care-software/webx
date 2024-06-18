package failures

// Failure represents a failure
type Failure struct {
	Index           uint   `json:"index"`
	Code            uint   `json:"code"`
	IsRaisedInLayer bool   `json:"is_raised_in_layer"`
	Message         string `json:"message"`
}
