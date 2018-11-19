package gocd

import (
	"encoding/json"
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
	var errors *multierror.Error
	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/jobs/%s/%s/%s/history/%d", pipeline, stage, job, offset))).
		// 18.6.0: providing API version here results in "resource not found"
		Set("Accept", "application/json").
		End()
	if errs != nil {
		errors = multierror.Append(errors, errs...)
		return []*JobHistory{}, errors.ErrorOrNil()
	}

	type JobHistoryResponse struct {
		Jobs []*JobHistory `json:"jobs"`
	}
	jobs := JobHistoryResponse{}
	jsonErr := json.Unmarshal([]byte(body), &jobs)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
		return []*JobHistory{}, errors.ErrorOrNil()
	}
	return jobs.Jobs, nil
}
