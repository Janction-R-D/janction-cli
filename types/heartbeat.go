package types

type HeartbeatReq struct {
	GPUInfoData    GPUInfo    `json:"gpu_info"`
	SystemInfoData SystemInfo `json:"system_info"`
	ExecInfoData   ExecInfo   `json:"exec_info"`
}

type HeartbeatResp struct {
	Success bool  `json:"success"`
	Jobs    []Job `json:"jobs"`
}
