package types

type Information struct {
	PlatformUUID      string `json:"platformUUID"`
	BoardSerialNumber string `json:"boardSerialNumber"`
	CpuSerialNumber   string `json:"cpuSerialNumber"`
	LocalMacInfo      string `json:"localMacInfo"`
}
