package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"jct/common/config"
	"jct/internal/service/heartbeat"
	aiJob "jct/internal/service/job"
	"jct/internal/service/login"
	"jct/internal/service/node_controller"
	"jct/internal/service/system_info"
	"jct/types"
	"jct/utils/snowflake"
	"os"
	"strings"
	"time"
)

type JanctionService struct{}

func (j *JanctionService) InitLogin() error {
	nonce, err := login.FetchNonce()
	if nonce == nil {
		return errors.New("failed to fetch nonce")
	}
	nonceStr := nonce.Data.Nonce
	if err != nil {
		return err
	}
	fmt.Println("[nonce] ", nonceStr)
	osType := config.OsType
	architecture := config.Architecture
	sysInfo := j.getSystemInfo(osType, architecture)
	loginResp, err := login.Login(nonceStr, sysInfo.BoardSerialNumber)
	if err != nil {
		return err
	}
	tokenStr := loginResp.Data.Token
	fmt.Println("[token] ", tokenStr)
	err = config.MemCache.SetString(context.Background(), "token", tokenStr, 8760*time.Hour)
	if err != nil {
		return err
	}
	return nil
}

func (j *JanctionService) InitController() error {
	osType := config.OsType
	architecture := config.Architecture
	useCPU := config.UseCPU
	useGPU := config.UseGPU
	sysInfo := j.getSystemInfo(osType, architecture)
	return node_controller.JoinController(sysInfo.BoardSerialNumber, architecture, useCPU, useGPU)
}

func (j *JanctionService) Run() error {
	osType := config.OsType
	architecture := config.Architecture
	useGPU := config.UseGPU
	useCPU := config.UseCPU
	sysInfo := j.getSystemInfo(osType, architecture)
	err := j.sendHeartbeat(sysInfo, types.ExecInfo{
		UseCPU: useCPU,
		UseGPU: useGPU,
	})
	return err
}

func (j *JanctionService) ExecTask() {
	aiJob.RunAIJob()
}

func (j *JanctionService) getSystemInfo(osType, architecture string) types.SystemInfo {
	var systemInfo types.SystemInfo
	systemInfo = system_info.GetLinuxInfo()
	systemInfo.OSType = osType
	systemInfo.Architecture = architecture
	_, err := os.Stat(".id") //os.Stat获取文件信息
	if os.IsNotExist(err) {
		f, err := os.Create(".id")
		defer f.Close()
		if err != nil {
			logrus.Error(err)
		} else {
			_, err = f.Write((snowflake.GenID()))
			if err != nil {
				logrus.Error(err)
			}
		}
	}

	uuid, err := os.ReadFile(".id")
	if err != nil {
		logrus.Error(err)
	}
	nodeId := strings.Replace(string(uuid), "\n", "", -1)
	systemInfo.BoardSerialNumber = nodeId
	return systemInfo
}

func (j *JanctionService) sendHeartbeat(info types.SystemInfo, execInfo types.ExecInfo) error {
	heartbeatResp, err := heartbeat.SendHeartbeat(info, execInfo)
	if err != nil {
		logrus.Info(err)
		return err
	}
	if !heartbeatResp {
		return errors.New("Connect Failed!")
	}

	//
	//j.execAIJobs
	return nil
}

//
//func (j *JanctionService) postHeartbeat(info types.SystemInfo, execInfo types.ExecInfo) error {
//	// Platform : macos linux windows
//	heartbeatResp, err := heartbeat.SendHeartbeat(info, execInfo)
//	fmt.Println(heartbeatResp)
//	if err != nil {
//		logrus.Info(err)
//		return err
//	}
//
//	///// test
//
//	// var jobs []types.Job
//	// jobs = append(jobs, types.Job{
//	// 	JobID:   time.Now().Unix(),
//	// 	JobType: "yolov3_arm_cpu",
//	// 	Compute: "CPU",
//	// })
//
//	// var heartbeatResp *types.HeartbeatResp
//	// heartbeatResp = &types.HeartbeatResp{
//	// 	Jobs: jobs,
//	// }
//	////
//
//	if heartbeatResp == nil {
//		return errors.New("Connect Failed!")
//	}
//
//	// 之前 从后端拿到 jobs 现在后端不返回 jobs
//
//	logrus.Println("[heartbeat] ", heartbeatResp)
//	if len(heartbeatResp.Jobs) > 0 {
//		go j.execAIJobs(info.OSType, heartbeatResp.Jobs)
//	}
//	// 发一个请求个 controller 获取 jobs， 请求参数可以一模一样
//	//if len(jobs > 0) {
//	//	go j.execAIJobs(info.OSType, heartbeatResp.Jobs)
//	//}
//
//	return nil
//}

//func (j *JanctionService) execAIJobs(osType string, jobs []types.Job) {
//
//	for _, job := range jobs {
//		jobType := job.JobType
//		jobID := job.JobID
//		compute := job.Compute
//		aiJob.RunAIJob(osType, compute, jobType, jobID)
//	}
//}

//
//func (j *JanctionService) Run(platform string, useGPU int) error {
//	macInfo := j.getSystemInfo(platform)
//	err := j.postHeartbeat(platform, macInfo)
//	return err
//}
//
//func (j *JanctionService) getSystemInfo(platform string) types.SystemInfo {
//	var systemInfo types.SystemInfo
//	if platform == "Darwin" {
//		systemInfo = system_info.GetMacInfo()
//	} else if platform == "Linux" {
//		systemInfo = system_info.GetLinuxInfo()
//	} else if platform == "Windows" {
//		systemInfo = system_info.GetWindowsInfo()
//	}
//	systemInfo.OSType = platform
//	return systemInfo
//}
//
//func (j *JanctionService) postHeartbeat(platform string, info types.SystemInfo) error {
//	heartbeatResp, err := heartbeat.SendHeartbeat(info)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	if heartbeatResp == nil {
//		return errors.New("Connect Failed!")
//	}
//	if len(heartbeatResp.Jobs) > 0 {
//		j.execAIJobs(platform, heartbeatResp.Jobs)
//	}
//	return nil
//}
//
//func (j *JanctionService) execAIJobs(platform string, jobs []types.Job) {
//	for _, job := range jobs {
//		jobType := job.JobType
//		jobID := job.JobID
//		aiJob.RunAIJob(platform, jobType, jobID)
//	}
//}
