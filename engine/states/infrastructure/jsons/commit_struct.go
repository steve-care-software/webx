package jsons

type commit struct {
	Executions []execution `json:"executions"`
	Parent     string      `json:"parent"`
}

type execution struct {
	Data  string `json:"data"`
	Chunk *chunk `json:"chunk"`
}

type chunk struct {
	Path        []string `json:"path"`
	Fingerprint string   `json:"fingerprint"`
}
