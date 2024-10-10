package types

type EnvInfo struct {
	Task    string `json:"task"`
	CPU     string `json:"cpu"`
	GPU     string `json:"gpu"`
	GPUUuid string `json:"gpu_uuid"`
}
