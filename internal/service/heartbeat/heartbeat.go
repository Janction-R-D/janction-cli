package heartbeat

import (
	"context"
	"encoding/json"
	"fmt"
	"jct/common/config"
	"jct/types"
	"jct/utils"
	"log"
	"time"
)

func SendHeartbeat(info types.SystemInfo, execInfo types.ExecInfo) (*types.HeartbeatResp, error) {
	url := config.TestnetUrl + "/api/v1/node/heartbeat"
	reqData := types.HeartbeatReq{
		GPUInfoData:    types.GPUInfo{},
		SystemInfoData: info,
		ExecInfoData:   execInfo,
	}
	body, err := json.Marshal(reqData)
	token, _ := config.MemCache.GetString(context.Background(), "token")
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	resp, err := utils.PostWithTimeout(url, body, header, 30*time.Second)
	log.Println(string(resp))
	var heartbeatResp types.HeartbeatResp
	err = json.Unmarshal(resp, &heartbeatResp)
	if err != nil {
		return nil, err
	}
	return &heartbeatResp, nil
}
