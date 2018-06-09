package gocd

import (
	"encoding/json"

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
	var errors *multierror.Error

	// Somehow GoCD will return "The resource you requested was not found!" if you specify an Accept header
	_, body, errs := c.Request.
		Get(c.resolve("/go/api/config/pipeline_groups")).
		//Set("Accept", "application/vnd.go.cd.v2+json").
		End()

	if errs != nil {
		errors = multierror.Append(errors, errs...)
		return []*PipelineGroup{}, errors.ErrorOrNil()
	}

	// first parse the json into temporary structure, so we parse stages object
	// with a single name string attribute as simple string
	type tmpStage struct {
		Name string `json:"name,omitempty"`
	}
	type tmpPipeline struct {
		Name      string     `json:"name,omitempty"`
		Label     string     `json:"label,omitempty"`
		Materials []Material `json:"materials,omitempty"`
		Stages    []tmpStage `json:"stages,omitempty"`
	}

	type tmpPipelineGroup struct {
		Name      string        `json:"name,omitempty"`
		Pipelines []tmpPipeline `json:"pipelines,omitempty"`
	}
	var tmpPipelineGroups []tmpPipelineGroup

	jsonErr := json.Unmarshal([]byte(body), &tmpPipelineGroups)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
		return []*PipelineGroup{}, errors.ErrorOrNil()
	}
	// create the good pipeline groups structures and copy data from the temporary structs
	pipelineGroups := make([]*PipelineGroup, len(tmpPipelineGroups))
	for i, pg := range tmpPipelineGroups {
		pipelineGroups[i] = &PipelineGroup{Name: pg.Name}
		pipelineGroups[i].Pipelines = make([]Pipeline, len(pg.Pipelines))
		for j, p := range pg.Pipelines {
			pipelineGroups[i].Pipelines[j] = Pipeline{Name: p.Name, Label: p.Label, Materials: p.Materials, Stages: make([]string, len(p.Stages))}
			for k, s := range p.Stages {
				pipelineGroups[i].Pipelines[j].Stages[k] = s.Name
			}
		}

	}
	return pipelineGroups, errors.ErrorOrNil()
}
