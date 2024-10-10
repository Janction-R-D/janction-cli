package job

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"jct/common/config"
	"jct/types"
	"jct/utils"
	"jct/utils/snowflake"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

//func runJob(cmdStr string, jobId int64) {
//	command := cmdStr
//	SubmitJobStatus(jobId, types.JobRunning)
//	ok, outString, errString := run(jobId, command, 10*60*1000)
//	if ok {
//		if strings.Contains(outString, "Predicted") {
//			log.Printf("[Success][Predicted][JobID:%d]: %s", jobId, outString)
//			SubmitJobStatus(jobId, types.JobFinished)
//		} else {
//			//log.Printf("[ERROR][Predicted]: %s", errString[len(errString)-83:])
//			log.Printf("[ERROR][Predicted][JobID:%d]", jobId)
//			SubmitJobStatus(jobId, types.JobFailed)
//		}
//	} else {
//		log.Printf("[Failed][Error][JobID:%d]: [%s][%s] ", jobId, outString, errString)
//		SubmitJobStatus(jobId, types.JobFailed)
//	}
//}

// TODO
func runJob(cmdStr string) {
	command := cmdStr
	jobId := snowflake.GenIntID()
	SubmitJobStatus(jobId, types.JobRunning)
	ok, outString, errString := run(jobId, command, 10*60*1000)
	// check success or fail
	if ok {
		log.Printf("[Failed][Error][JobID:%d]: [%s][%s] ", jobId, outString, errString)
		SubmitJobStatus(jobId, types.JobFinished)

	}
}

//func runJob(cmdStr string) {
//	command := cmdStr
//	jobId := snowflake.GenIntID()
//	SubmitJobStatus(jobId, types.JobRunning)
//	ok, outString, errString := run(jobId, command, 10*60*1000)
//	// check success or fail
//	if ok {
//		if strings.Contains(outString, "Predicted") {
//			log.Printf("[Success][Predicted][JobID:%d]: %s", jobId, outString)
//			SubmitJobStatus(jobId, types.JobFinished)
//		} else {
//			//log.Printf("[ERROR][Predicted]: %s", errString[len(errString)-83:])
//			log.Printf("[ERROR][Predicted][JobID:%d]", jobId)
//			SubmitJobStatus(jobId, types.JobFailed)
//		}
//	} else {
//		log.Printf("[Failed][Error][JobID:%d]: [%s][%s] ", jobId, outString, errString)
//		SubmitJobStatus(jobId, types.JobFailed)
//	}
//}

func run(jobId int64, command string, killInMilliSeconds time.Duration) (okResult bool, stdout, stderr string) {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", command)
	var buf1 bytes.Buffer
	r1, w1, _ := os.Pipe()
	cmd.Stdout = w1
	var buf2 bytes.Buffer
	r2, w2, _ := os.Pipe()
	cmd.Stderr = w2
	okResult = true
	go buf1.ReadFrom(r1)
	go buf2.ReadFrom(r2)
	err := cmd.Start()
	log.Printf("[Run][JobId:%d]Waiting For Jacntion AI Job To Finish...", jobId)
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(killInMilliSeconds * time.Millisecond):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill: ", err)
			okResult = false
		}
		<-done
		log.Println("process killed")
	case err := <-done:
		if err != nil {
			log.Printf("process done with error = %v", err)
			okResult = false
		}
	}
	if err != nil {
		log.Fatal(err)
		okResult = false
	}
	return true, buf1.String(), buf2.String()
}

func SubmitJobStatus(jobId int64, jobStatus types.JobStatus) {
	url := config.TestnetUrl + "/api/v1/node/job"
	reqSubmitJob := types.SubmitJobReq{
		JobID:    jobId,
		Status:   jobStatus,
		TaskName: config.Task,
		NodeID:   config.NodeID,
		TaskType: config.TaskType,
		EnvInfoData: types.EnvInfo{
			Task:    config.Task,
			CPU:     config.JCT_CPU,
			GPU:     config.JCT_GPU,
			GPUUuid: config.JCT_GPU_ID,
		},
	}

	token, _ := config.MemCache.GetString(context.Background(), "token")
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	body, err := json.Marshal(reqSubmitJob)
	if err != nil {
		logrus.Error(err)
	}

	_, err = utils.PostWithTimeout(url, body, header, 30*time.Second)
	if err != nil {
		logrus.Error(err)
	}
}

func SubmitJobFinishStatus(jobId int64) {
	url := config.TestnetUrl + "/api/v1/node/job_finish"

	reqSubmitJob := types.SubmitJobFinishReq{
		JobID: jobId,
		EnvInfoData: types.EnvInfo{
			Task:    config.Task,
			CPU:     config.JCT_CPU,
			GPU:     config.JCT_GPU,
			GPUUuid: config.JCT_GPU_ID,
		},
	}

	token, _ := config.MemCache.GetString(context.Background(), "token")
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	body, err := json.Marshal(reqSubmitJob)
	if err != nil {
		logrus.Error(err)
	}

	_, err = utils.PostWithTimeout(url, body, header, 30*time.Second)
	if err != nil {
		logrus.Error(err)
	}
}
