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

// Stage represents a stage configuration object
// See https://api.gocd.org/current/#the-stage-object
type Stage struct {
	Name                  string                `json:"name"`
	FetchMaterials        bool                  `json:"fetch_materials"`
	CleanWorkingDirectory bool                  `json:"clean_working_directory"`
	NeverCleanupArtifacts bool                  `json:"never_cleanup_artifacts"`
	Approval              StageApproval         `json:"approval"`
	EnvironmentVariables  []EnvironmentVariable `json:"environment_variables"`
	Jobs                  []JobConfig           `json:"jobs"`
}

// StageApproval defines the approval needed to trigger the stage
// See: https://api.gocd.org/current/#pipeline-config-approval-object
type StageApproval struct {
	Type          string             `json:"type"`
	Authorization StageAuthorization `json:"authorization"`
}

// StageAuthorization provide the list of users and roles authorized
// to operate (run) on this stage.
// See: https://api.gocd.org/current/#the-authorization-object
type StageAuthorization struct {
	Roles []string `json:"roles"`
	Users []string `json:"users"`
}
