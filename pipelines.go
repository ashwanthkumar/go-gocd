package gocd

import (
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
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

// GetPipelineInstance returns the pipeline instance corresponding to the given
// pipeline name and counter
func (c *DefaultClient) GetPipelineInstance(name string, counter int) (*PipelineInstance, error) {
	var errors *multierror.Error
	res := new(PipelineInstance)

	_, body, errs := c.Request.Get(c.resolve(fmt.Sprintf("/go/api/pipelines/%s/instance/%d", name, counter))).End()
	if errs != nil {
		errors = multierror.Append(errors, errs...)
		return res, errors.ErrorOrNil()
	}

	err := json.Unmarshal([]byte(body), res)
	if err != nil {
		errors = multierror.Append(errors, err)
	}
	return res, errors.ErrorOrNil()
}
