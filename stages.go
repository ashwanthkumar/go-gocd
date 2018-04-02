package gocd

// StageRun represent a stage run history event
type StageRun struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ApprovedBy        string `json:"approved_by"`
	Jobs              []Job  `json:"jobs"`
	CanRun            bool   `json:"can_run"`
	Result            string `json:"result"`
	ApprovalType      string `json:"approval_type"`
	Counter           string `json:"counter"`
	OperatePermission bool   `json:"operate_permission"`
	RerunOfCounter    bool   `json:"rerun_of_counter"`
	Scheduled         bool   `json:"scheduled"`
}
