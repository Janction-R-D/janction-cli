package machine

import "jct/utils/machine/types"

type OsMachineInterface interface {
	GetMachine() types.Information
	GetBoardSerialNumber() (string, error)
	GetPlatformUUID() (string, error)
	GetCpuSerialNumber() (string, error)
}
