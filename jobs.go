package gocd

import (
	"encoding/xml"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
)

// Job definition used also by other elements like the pipeline and stages
type Job struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Result        string `json:"result"`
	State         string `json:"state"`
	ScheduledDate int64  `json:"scheduled_date"`
}

// JobConfig is used to configure a job inside a stage of a pipeline
// See: https://api.gocd.org/current/#the-job-object
type JobConfig struct {
	Name                 string                `json:"name"`
	RunInstanceCount     interface{}           `json:"run_instance_count"`
	Timeout              int                   `json:"timeout"`
	EnvironmentVariables []EnvironmentVariable `json:"environment_variables"`
	Resources            []string              `json:"resources,omitempty"`
	Tasks                []Task                `json:"tasks,omitempty"`
	Tabs                 []Tab                 `json:"tabs,omitempty"`
	Artifacts            []ArtifactConfig      `json:"artifacts,omnitempty"`
	Properties           []JobProperty         `json:"properties,omitempty"`
	ElasticProfileID     string                `json:"elastic_profile_id,omitempty"`
}

// Task is used to configure tasks inside a JobConfig
// See: https://api.gocd.org/current/#the-task-object
type Task struct {
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}

// ExecTaskAttributes define the attributes of an exec task
// See: https://api.gocd.org/current/#the-exec-task-attributes
type ExecTaskAttributes struct {
	RunIf            []string `json:"run_if"`
	Command          string   `json:"command,omitempty"`
	Arguments        []string `json:"arguments,omitempty"`
	WorkingDirectory string   `json:"working_directory,omitempty"`
	OnCancel         Task     `json:"on_cancel,omitempty"`
}

// AntTaskAttributes define the attributes of an ant task
// See: https://api.gocd.org/current/#the-ant-task-attributes
type AntTaskAttributes struct {
	RunIf            []string `json:"run_if"`
	BuildFile        string   `json:"build_file,omitempty"`
	Target           string   `json:"target,omitempty"`
	WorkingDirectory string   `json:"working_directory,omitempty"`
	OnCancel         Task     `json:"on_cancel,omitempty"`
}

// NantTaskAttributes define the attributes of a nant task
// See: https://api.gocd.org/current/#the-nant-task-attributes
type NantTaskAttributes struct {
	RunIf            []string `json:"run_if"`
	BuildFile        string   `json:"build_file,omitempty"`
	Target           string   `json:"target,omitempty"`
	NantPath         string   `json:"nant_path,omitempty"`
	WorkingDirectory string   `json:"working_directory,omitempty"`
	OnCancel         Task     `json:"on_cancel,omitempty"`
}

// RakeTaskAttributes define the attributes of a rake task
// See: https://api.gocd.org/current/#the-rake-task-attributes
type RakeTaskAttributes struct {
	RunIf            []string `json:"run_if"`
	BuildFile        string   `json:"build_file,omitempty"`
	Target           string   `json:"target,omitempty"`
	WorkingDirectory string   `json:"working_directory,omitempty"`
	OnCancel         Task     `json:"on_cancel,omitempty"`
}

// FetchTaskAttributes define the attributes of a fetch task
// See: https://api.gocd.org/current/#the-fetch-task-attributes
type FetchTaskAttributes struct {
	RunIf         []string `json:"run_if"`
	Pipeline      string   `json:"pipeline,omitempty"`
	Stage         string   `json:"stage,omitempty"`
	Job           string   `json:"job,omitempty"`
	Source        string   `json:"source,omitempty"`
	IsASourceFile bool     `json:"is_source_a_file,omitempty"`
	Destination   string   `json:"destination,omitempty"`
	OnCancel      Task     `json:"on_cancel,omitempty"`
}

// PluggableTaskAttributes define the attributes of a pluggable task
// See: https://api.gocd.org/current/#the-pluggable-task-attributes
type PluggableTaskAttributes struct {
	RunIf               []string                               `json:"run_if"`
	PluginConfiguration TaskPluginConfiguration                `json:"plugin_configuration,omitempty"`
	Configuration       []PluggableTaskAttributesConfiguration `json:"configuration,omitempty"`
	OnCancel            Task                                   `json:"on_cancel,omitempty"`
}

// PluggableTaskAttributesConfiguration is used for the configuration attribute
// of the PluggableTaskAttributes
type PluggableTaskAttributesConfiguration struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// TaskPluginConfiguration used in the PluginConfiguration of a pluggable task
// See: https://api.gocd.org/current/#the-pluggable-task-object
type TaskPluginConfiguration struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

// Tab defines the tabs to configure on a job
// See: https://api.gocd.org/current/#the-tab-object
type Tab struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// JobProperty for a JobConfig struct
// See: https://api.gocd.org/current/#the-property-object
type JobProperty struct {
	Name   string `json:"name"`
	Source string `json:"source"`
	XPath  string `json:"xpath"`
}

// ScheduledJobResource wrapper for resources > resource
type ScheduledJobResource struct {
	Name string `xml:",chardata"`
}

// ScheduledJob instance
type ScheduledJob struct {
	Name         string                 `xml:"name,attr"`
	JobID        string                 `xml:"id,attr"`
	BuildLocator string                 `xml:"buildLocator"`
	Link         LinkInXML              `xml:"link"`
	Environment  string                 `xml:"environment,omitempty"`
	RawResources []ScheduledJobResource `xml:"resources>resource,omitempty"`
}

// Resources - return resources as []string
func (sj *ScheduledJob) Resources() []string {
	stringSlice := make([]string, len(sj.RawResources))
	for index, resource := range sj.RawResources {
		stringSlice[index] = resource.Name
	}

	return stringSlice
}

// JobURL - Full URL location of the scheduled job
func (sj *ScheduledJob) JobURL() string {
	return sj.Link.Href
}

// LinkInXML - <link rel="..." href="..."> tag
type LinkInXML struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

// GetScheduledJobs - Lists all the current job instances which are scheduled but not yet assigned to any agent.
func (c *DefaultClient) GetScheduledJobs() ([]*ScheduledJob, error) {
	var errors *multierror.Error

	type ScheduledJobsResponse struct {
		XMLName xml.Name        `xml:"scheduledJobs"`
		Jobs    []*ScheduledJob `xml:"job"`
	}

	var jobs ScheduledJobsResponse
	_, body, errs := c.Request.
		Get(c.resolve("/go/api/jobs/scheduled.xml")).
		End()
	if errs != nil {
		errors = multierror.Append(errors, errs...)
		return []*ScheduledJob{}, errors.ErrorOrNil()
	}
	xmlErr := xml.Unmarshal([]byte(body), &jobs)
	if xmlErr != nil {
		errors = multierror.Append(errors, xmlErr)
	}

	return jobs.Jobs, errors.ErrorOrNil()
}

// JobHistory - Represents an instance of a job from the past
type JobHistory struct {
	AgentUUID           string               `json:"agent_uuid"`
	Name                string               `json:"name"`
	JobStateTransitions []JobStateTransition `json:"job_state_transitions"`
	ScheduledDate       int                  `json:"scheduled_date"`
	OriginalJobID       string               `json:"original_job_id"`
	PipelineCounter     int                  `json:"pipeline_counter"`
	PipelineName        string               `json:"pipeline_name"`
	Result              string               `json:"result"`
	State               string               `json:"state"`
	ID                  int                  `json:"id"`
	StageCounter        string               `json:"stage_counter"`
	StageName           string               `json:"stage_name"`
	ReRun               bool                 `json:"rerun"`
}

// JobStateTransition - Represents an instance of StateTransition the job went through
type JobStateTransition struct {
	StateChangeTime int    `json:"state_change_time,omitempty"`
	ID              int    `json:"id,omitempty"`
	State           string `json:"state,omitempty"`
}

// GetJobHistory - The job history allows users to list job instances of specified job. Supports pagination using offset which tells the API how many instances to skip.
func (c *DefaultClient) GetJobHistory(pipeline, stage, job string, offset int) ([]*JobHistory, error) {
	type JobHistoryResponse struct {
		Jobs []*JobHistory `json:"jobs"`
	}
	res := new(JobHistoryResponse)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v2+json"}
	_, err := c.getJSON(fmt.Sprintf("/go/api/jobs/%s/%s/%s/history/%d", pipeline, stage, job, offset), headers, res)
	return res.Jobs, err
}
