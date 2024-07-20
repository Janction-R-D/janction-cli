package types

type SystemInfo struct {
	OSType            string `json:"os_type"`      // macos linux windows
	Architecture      string `json:"architecture"` // amd64 arm
	BoardSerialNumber string `json:"board_serial_number"`
	PlatformUUID      string `json:"platform_uuid"`
	MACAddress        string `json:"mac_address"`
}
