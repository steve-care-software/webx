package resources

// Resource represents a resource
type Resource struct {
	Layer       string `json:"layer"`
	IsMandatory bool   `json:"is_mandatory"`
}
