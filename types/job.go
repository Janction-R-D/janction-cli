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
	JobID  int64     `json:"job_id"`
	Status JobStatus `json:"status"`
}

type JobStatus int

const (
	JobRunning = 1 + iota
	JobFinished
	JobFailed
)
