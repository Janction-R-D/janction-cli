package types

type Job struct {
	JobID   int64  `json:"job_id"`
	JobType string `json:"job_type"`
	Compute string `json:"compute"`
}

// type JobResult struct {
// 	FinishedJobList []JobStatus `json:"finished_job_list"`
// 	FailedJobList   []JobStatus `json:"failed_job_list"`
// 	RunningJobList  []JobStatus `json:"running_job_list"`
// }

// type JobStatus struct {
// 	JobID     string `json:"job_id"`
// 	JobType   string `json:"job_type"`
// 	BeginTime string `json:"begin_time"`
// 	EndTime   string `json:"end_time"`
// }

type SubmitJobReq struct {
	JobID       int64     `json:"job_id"`
	Status      JobStatus `json:"status"`
	EnvInfoData EnvInfo   `json:"env_info"`
	NodeID      string    `json:"node_id" binding:"required"`
	TaskName    string    `json:"task_name" binding:"required"`
	TaskType    string    `json:"task_type" binding:"required"`
}

type SubmitJobFinishReq struct {
	JobID       int64   `json:"job_id"`
	EnvInfoData EnvInfo `json:"env_info_data"`
}

type JobStatus int

const (
	JobRunning = iota + 1
	JobFinished
	JobFailed
)
