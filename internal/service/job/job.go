package job

import (
	"fmt"
	"jct/common/config"
)

func RunAIJob() {
	taskFileName := config.Task + ".py"
	path := config.Path
	runJob(fmt.Sprintf("%s/%s", path, taskFileName))

	// DEBUG
	//runJob(fmt.Sprintf("%s/%s", "/Users/dick/Documents/jac/jct/cmd/scripts", taskFileName))

}

//func RunAIJob(osType, compute, jobType string, jobId int64) {
//	if allowPlatform(osType, jobType) {
//		logrus.Println(fmt.Sprintf("[%s][JobID]:%d, [jobType]:%s Started...", compute, jobId, jobType))
//		path := config.Path
//		fmt.Println("file path:", path)
//		runJob(fmt.Sprintf("python3 task.py", path), jobId)
//	} else {
//		if jobId == 0 {
//			logrus.Error("Node Busy...")
//		} else {
//			logrus.Error("Unsupported job-type")
//			SubmitJobStatus(jobId, types.JobFailed)
//		}
//
//	}
//}

//func RunAIJob(osType, compute, jobType string, jobId int64) {
//	if allowPlatform(osType, jobType) {
//		logrus.Println(fmt.Sprintf("[%s][JobID]:%d, [jobType]:%s Started...", compute, jobId, jobType))
//		path := config.Path
//		runJob(fmt.Sprintf("cd %s && ./darknet detect cfg/yolov3.cfg data/yolov3.weights data/person.jpg", path), jobId)
//	} else {
//		if jobId == 0 {
//			logrus.Error("Node Busy...")
//		} else {
//			logrus.Error("Unsupported job-type")
//			SubmitJobStatus(jobId, types.JobFailed)
//		}
//
//	}
//}

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
