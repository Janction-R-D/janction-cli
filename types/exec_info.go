package types

type ExecInfo struct {
	UseGPU int `json:"use_gpu"` // 0 , 1 可接受使用 gpu
	UseCPU int `json:"use_cpu"` // 0 , 1 可接受使用 cpu
}
