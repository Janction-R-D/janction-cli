package node_controller

import (
	"encoding/json"
	"fmt"
	"jct/common/config"
	"jct/utils"
	"log"
	"time"
)

type FormRegisterNode struct {
	NodeID           string `json:"node_id" binding:"required"`
	ArchitectureType string `json:"architecture_type" binding:"required"`
	UseCPU           int    `json:"use_cpu" binding:"required"`
	UseGPU           int    `json:"use_gpu" binding:"required"`
}

func JoinController(nodeId, architectureType string, useCPU, useGPU int) error {
	url := config.ControllerUrl + "/api/controller/v1/register"
	reqData := FormRegisterNode{
		NodeID:           nodeId,
		ArchitectureType: architectureType,
		UseCPU:           useCPU,
		UseGPU:           useGPU,
	}

	fmt.Println(nodeId, architectureType, useCPU, useGPU)
	body, err := json.Marshal(reqData)
	//token, _ := config.MemCache.GetString(context.Background(), "token")
	//header := map[string]string{
	//	"Authorization": fmt.Sprintf("Bearer %s", token),
	//}
	header := map[string]string{}
	resp, err := utils.PostWithTimeout(url, body, header, 30*time.Second)
	log.Println(string(resp))
	return err
}
