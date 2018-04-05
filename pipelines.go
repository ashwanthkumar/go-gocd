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

// PipelineConfig is used to manage a pipeline configuration
// See: https://api.gocd.org/current/#the-pipeline-config-object
type PipelineConfig struct {
	LabelTemplate        string                `json:"label_template"`
	LockBehavior         string                `json:"lock_behavior"`
	Name                 string                `json:"name"`
	Template             string                `json:"template"`
	Origin               RepoOrigin            `json:"origin"`
	Parameters           []PipelineParameter   `json:"parameters"`
	EnvironmentVariables []EnvironmentVariable `json:"environment_variables"`
	Materials            []Material            `json:"materials"`
	Stages               []Stage               `json:"stages"`
	TrackingTool         TrackingTool          `json:"tracking_tool,omitempty"`
	Timer                Timer                 `json:"timer,omitempty"`
}

// PipelineParameter is used to provide parameters to a pipeline configuration
// See: https://api.gocd.org/current/#the-pipeline-parameter-object
type PipelineParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// RepoOrigin is used to configure the repo origin in a pipeline configuration
// See: https://api.gocd.org/current/#the-config-repo-origin-object
type RepoOrigin struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// TrackingTool defines which tracking tool to use on the pipeline
// See: https://api.gocd.org/current/#the-tracking-tool-object
type TrackingTool struct {
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}

// TrackingToolAttributesGeneric is used in the TrackingTool.Attributes field
// for a tracking tool of type generic.
// See: https://api.gocd.org/current/#the-generic-tracking-tool-object
type TrackingToolAttributesGeneric struct {
	URLPattern string `json:"url_pattern"`
	Regex      string `json:"regex"`
}

// TrackingToolAttributesMingle is used in the TrackingTool.Attributes field
// for a tracking tool of type mingle.
// See: https://api.gocd.org/current/#the-mingle-tracking-tool-object
type TrackingToolAttributesMingle struct {
	BaseURL               string `json:"base_url"`
	ProjectIdentifier     string `json:"project_identifier"`
	MqlGroupingConditions string `json:"mql_grouping_conditions"`
}

// Timer is used to define the cron-like schedule used to buid a pipeline
// See: https://api.gocd.org/current/#the-timer-object
type Timer struct {
	Spec          string `json:"spec"`
	OnlyOnChanges bool   `json:"only_on_changes"`
}

// PipelineGetInstance returns the pipeline instance corresponding to the given
// pipeline name and counter
func (c *DefaultClient) PipelineGetInstance(name string, counter int) (*PipelineInstance, error) {
	res := new(PipelineInstance)
	_, err := c.getJSON(fmt.Sprintf("/go/api/pipelines/%s/instance/%d", name, counter), nil, res)
	return res, err
}

// PipelineGetHistoryPage allows users to list pipeline instances. Supports
// pagination using offset which tells the API how many instances to skip.
// Note that te history is listed in reverse chronological order meaning the
// setting an offset to 1 will skip the last run of the pipeline and will give
// you a page of pipeline runs history which is 10 by default.
func (c *DefaultClient) PipelineGetHistoryPage(name string, offset int) (*PipelineHistoryPage, error) {
	res := new(PipelineHistoryPage)
	_, err := c.getJSON(fmt.Sprintf("/go/api/pipelines/%s/history/%d", name, offset), nil, res)
	return res, err
}

// PipelineGetStatus allows users to check if the pipeline is paused, locked and
// schedulable.
func (c *DefaultClient) PipelineGetStatus(name string) (*PipelineStatus, error) {
	res := new(PipelineStatus)
	_, err := c.getJSON(fmt.Sprintf("/go/api/pipelines/%s/status", name), nil, res)
	return res, err
}

// PipelinePause pauses the specified pipeline using the given cause
func (c *DefaultClient) PipelinePause(name, cause string) (*SimpleMessage, error) {
	data := struct{ PauseCause string }{cause}
	res := new(SimpleMessage)
	// The headers bellow only work with gocd 18.2
	// headers := map[string]string{"Accept": "application/vnd.go.cd.v1+json", "X-GoCD-Confirm": "true"}
	headers := map[string]string{"Accept": "application/json", "Confirm": "true"}
	err := c.postJSON(fmt.Sprintf("/go/api/pipelines/%s/pause", name), headers, data, res)
	return res, err
}

// PipelineUnpause unpauses the specified pipeline
func (c *DefaultClient) PipelineUnpause(name string) (*SimpleMessage, error) {
	res := new(SimpleMessage)
	// The headers bellow only work with gocd 18.2 and over
	//headers := map[string]string{"Accept": "application/vnd.go.cd.v1+json", "X-GoCD-Confirm": "true"}
	headers := map[string]string{"Accept": "application/json", "Confirm": "true"}
	err := c.postJSON(fmt.Sprintf("/go/api/pipelines/%s/unpause", name), headers, nil, res)
	return res, err
}

// PipelineUnlock releases a lock on a pipeline so that you can start up a new
// instance without having to wait for the earlier instance to finish.
// Note: A pipeline lock can only be released when a pipeline is locked, AND there is no
// running instance of the pipeline.
// Requires GoCD version 18.2.0+
func (c *DefaultClient) PipelineUnlock(name string) (*SimpleMessage, error) {
	res := new(SimpleMessage)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v1+json", "X-GoCD-Confirm": "true"}
	err := c.postJSON(fmt.Sprintf("/go/api/pipelines/%s/unlock", name), headers, nil, res)
	return res, err
}

// PipelineGetConfig returns the configuration of the given pipeline along with
// the ETag header value
func (c *DefaultClient) PipelineGetConfig(name string) (*PipelineConfig, string, error) {
	res := new(PipelineConfig)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v5+json"}
	etag, err := c.getJSON(fmt.Sprintf("/go/api/admin/pipelines/%s", name), headers, res)
	return res, etag, err
}
