package job

import (
	"fmt"
	"jct/common/config"
	"jct/types"

	"github.com/sirupsen/logrus"
)

func RunAIJob(osType, compute, jobType string, jobId int64) {
	if allowPlatform(osType, jobType) {
		logrus.Println(fmt.Sprintf("[%s][JobID]:%d, [jobType]:%s Started...", compute, jobId, jobType))

		// cmd exec
		//macos.RunForMacOS("./yolov3_macos/darknet detect yolov3_macos/cfg/yolov3.cfg yolov3_macos/data/yolov3.weights yolov3_macos/data/person.jpg")
		path := config.Path
		runJob(fmt.Sprintf("cd %s && ./darknet detect cfg/yolov3.cfg data/yolov3.weights data/person.jpg", path), jobId)
		//macos.RunForMacOS("../../../darknet/darknet detect darknet/cfg/yolov3.cfg ../../../darknet/data/yolov3.weights ../../../darknet/data/person.jpg")
	} else {
		if jobId == 0 {
			logrus.Error("Node Busy...")
		} else {
			logrus.Error("Unsupported job-type")
			SubmitJobStatus(jobId, types.JobFailed)
		}

	}
}

func allowPlatform(osType, jobType string) bool {
	// TODO jobType from config
	if osType == "macos" {
		if jobType == "yolov3_arm_cpu" || jobType == "yolov3_amd64_cpu" {
			return true
		}
	} else if osType == "linux" {
		if jobType == "yolov3_arm_cpu" || jobType == "yolov3_arm_gpu" || jobType == "yolov3_amd64_cpu" || jobType == "yolov3_amd64_gpu" {
			return true
		}
	} else if osType == "windows" {
		if jobType == "yolov3_arm_cpu" || jobType == "yolov3_arm_gpu" || jobType == "yolov3_amd64_cpu" || jobType == "yolov3_amd64_gpu" {
			return true
		}
	}
	return false
}
