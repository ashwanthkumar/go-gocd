package gocd

import (
	multierror "github.com/hashicorp/go-multierror"
)

// Pipeline Object
type Pipeline struct {
	Name      string     `json:"name,omitempty"`
	Label     string     `json:"label,omitempty"`
	Materials []Material `json:"materials,omitempty"`
	Stages    []string   `json:"stages,omitempty"`
}

// PipelineGroup Object
type PipelineGroup struct {
	Name      string     `json:"name,omitempty"`
	Pipelines []Pipeline `json:"pipelines,omitempty"`
}

// GetPipelineGroups List pipeline groups along with the pipelines, stages and materials for each pipeline.
func (c *DefaultClient) GetPipelineGroups() ([]*PipelineGroup, error) {

	type EmbeddedObj struct {
		PipelineGroup []*PipelineGroup `json:"groups"`
	}
	type AllPipelineGroupsResponse struct {
		Embedded EmbeddedObj `json:"_embedded"`
	}
	res := new(AllPipelineGroupsResponse)
	headers := map[string]string{"Accept": "application/vnd.go.cd+json"}
	err := c.getJSON("/go/api/admin/pipeline_groups", headers, res)
	if err != nil {
		return []*PipelineGroup{}, err
	}

	var errors *multierror.Error

	return res.Embedded.PipelineGroup, errors.ErrorOrNil()
}
