package system_info

import (
	"jct/types"
	"jct/utils/machine"
)

func GetLinuxInfo() types.SystemInfo {
	serialNumber, _ := machine.GetBoardSerialNumber()
	uuid, _ := machine.GetPlatformUUID()
	macInfo, _ := machine.GetMACAddress()
	systemInfo := types.SystemInfo{
		BoardSerialNumber: serialNumber,
		PlatformUUID:      uuid,
		MACAddress:        macInfo,
	}
	return systemInfo
}
