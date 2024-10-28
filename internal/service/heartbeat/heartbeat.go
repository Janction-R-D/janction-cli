package heartbeat

import (
	"context"
	"encoding/json"
	"fmt"
	"jct/common/config"
	"jct/types"
	"jct/utils"
	"time"
)

//func SendHeartbeat(info types.SystemInfo, execInfo types.ExecInfo) (*types.HeartbeatResp, error) {
//	url := config.TestnetUrl + "/api/v1/node/heartbeat"
//	reqData := types.HeartbeatReq{
//		GPUInfoData:    types.GPUInfo{},
//		SystemInfoData: info,
//		ExecInfoData:   execInfo,
//	}
//	nodeId := reqData.SystemInfoData.BoardSerialNumber
//	fmt.Println(nodeId)
//	body, err := json.Marshal(reqData)
//	token, _ := config.MemCache.GetString(context.Background(), "token")
//	header := map[string]string{
//		"Authorization": fmt.Sprintf("Bearer %s", token),
//	}
//	resp, err := utils.PostWithTimeout(url, body, header, 30*time.Second)
//	log.Println(string(resp))
//	var heartbeatResp types.HeartbeatResp
//	err = json.Unmarshal(resp, &heartbeatResp)
//	if err != nil {
//		return nil, err
//	}
//	return &heartbeatResp, nil
//}

func SendHeartbeat(info types.SystemInfo, execInfo types.ExecInfo) (bool, error) {
	url := config.TestnetUrl + "/api/v1/node/heartbeat"
	reqData := types.HeartbeatReq{
		GPUInfoData:    types.GPUInfo{},
		SystemInfoData: info,
		ExecInfoData:   execInfo,
		EnvInfoData: types.EnvInfo{
			Task:    config.Task,
			CPU:     config.JCT_CPU,
			GPU:     config.JCT_GPU,
			GPUUuid: config.JCT_GPU_ID,
		},
	}
	//nodeId := reqData.SystemInfoData.BoardSerialNumber
	body, err := json.Marshal(reqData)
	token, _ := config.MemCache.GetString(context.Background(), "token")
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	_, err = utils.PostWithTimeout(url, body, header, 30*time.Second)
	if err != nil {
		return false, err
	}
	return true, nil
}
