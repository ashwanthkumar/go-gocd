package gocd

import (
	"fmt"
)

// PipelineInstance represents a pipeline instance (for a given run)
type PipelineInstance struct {
	ID                  int                `json:"id"`
	Name                string             `json:"name"`
	Label               string             `json:"label"`
	NaturalOrder        float32            `json:"natural_order"`
	CanRun              bool               `json:"can_run"`
	Comment             string             `json:"comment"`
	Counter             int                `json:"counter"`
	PreparingToSchedule bool               `json:"preparing_to_schedule"`
	Stages              []StageRun         `json:"stages"`
	BuildCause          PipelineBuildCause `json:"build_cause"`
}

// PipelineBuildCause represent what triggered the build of the pipeline
type PipelineBuildCause struct {
	Approver          string             `json:"approver"`
	MaterialRevisions []MaterialRevision `json:"material_revisions"`
	TriggerForced     bool               `json:"trigger_forced"`
	TriggerMessage    string             `json:"trigger_message"`
}

// PipelineHistoryPage represents a page of the history of run of a pipeline
type PipelineHistoryPage struct {
	Pipelines  []PipelineInstance `json:"pipelines"`
	Pagination Pagination         `json:"pagination"`
}

// PipelineStatus represents the status of a pipeline
type PipelineStatus struct {
	PausedCause string `json:"pausedCause"`
	PausedBy    string `json:"pausedBy"`
	Paused      bool   `json:"paused"`
	Schedulable bool   `json:"schedulable"`
	Locked      bool   `json:"locked"`
}

// GetPipelineInstance returns the pipeline instance corresponding to the given
// pipeline name and counter
func (c *DefaultClient) GetPipelineInstance(name string, counter int) (*PipelineInstance, error) {
	res := new(PipelineInstance)
	err := c.getJSON(fmt.Sprintf("/go/api/pipelines/%s/instance/%d", name, counter), nil, res)
	return res, err
}

// GetPipelineHistoryPage allows users to list pipeline instances. Supports
// pagination using offset which tells the API how many instances to skip.
// Note that te history is listed in reverse chronological order meaning the
// setting an offset to 1 will skip the last run of the pipeline and will give
// you a page of pipeline runs history which is 10 by default.
func (c *DefaultClient) GetPipelineHistoryPage(name string, offset int) (*PipelineHistoryPage, error) {
	res := new(PipelineHistoryPage)
	err := c.getJSON(fmt.Sprintf("/go/api/pipelines/%s/history/%d", name, offset), nil, res)
	return res, err
}

// GetPipelineStatus allows users to check if the pipeline is paused, locked and
// schedulable.
func (c *DefaultClient) GetPipelineStatus(name string) (*PipelineStatus, error) {
	res := new(PipelineStatus)
	err := c.getJSON(fmt.Sprintf("/go/api/pipelines/%s/status", name), nil, res)
	return res, err
}

// PausePipeline pauses the specified pipeline using the given cause
func (c *DefaultClient) PausePipeline(name, cause string) (*SimpleMessage, error) {
	data := struct{ PauseCause string }{cause}
	res := new(SimpleMessage)
	// The headers bellow only work with gocd 18.2
	// headers := map[string]string{"Accept": "application/vnd.go.cd.v1+json", "X-GoCD-Confirm": "true"}
	headers := map[string]string{"Accept": "application/json", "Confirm": "true"}
	err := c.postJSON(fmt.Sprintf("/go/api/pipelines/%s/pause", name), headers, data, res)
	return res, err
}

// UnpausePipeline unpauses the specified pipeline
func (c *DefaultClient) UnpausePipeline(name string) (*SimpleMessage, error) {
	res := new(SimpleMessage)
	// The headers bellow only work with gocd 18.2 and over
	//headers := map[string]string{"Accept": "application/vnd.go.cd.v1+json", "X-GoCD-Confirm": "true"}
	headers := map[string]string{"Accept": "application/json", "Confirm": "true"}
	err := c.postJSON(fmt.Sprintf("/go/api/pipelines/%s/unpause", name), headers, nil, res)
	return res, err
}

// UnlockPipeline releases a lock on a pipeline so that you can start up a new
// instance without having to wait for the earlier instance to finish.
// Note: A pipeline lock can only be released when a pipeline is locked, AND there is no
// running instance of the pipeline.
// Requires GoCD version 18.2.0+
func (c *DefaultClient) UnlockPipeline(name string) (*SimpleMessage, error) {
	res := new(SimpleMessage)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v1+json", "X-GoCD-Confirm": "true"}
	err := c.postJSON(fmt.Sprintf("/go/api/pipelines/%s/unlock", name), headers, nil, res)
	return res, err
}
