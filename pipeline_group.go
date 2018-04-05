package gocd

type pipeline struct {
	Name      string     `json:"name,omitempty"`
	Label     string     `json:"label,omitempty"`
	Materials []Material `json:"materials,omitempty"`
	Stages    []string   `json:"stages,omitempty"`
}

// PipelineGroup Object
type PipelineGroup struct {
	Name      string     `json:"name,omitempty"`
	Pipelines []pipeline `json:"pipelines,omitempty"`
}

// GetPipelineGroups List pipeline groups along with the pipelines, stages and materials for each pipeline.
func (c *DefaultClient) GetPipelineGroups() ([]*PipelineGroup, error) {
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

	tmpPipelineGroups := new([]*tmpPipelineGroup)
	_, err := c.getJSON("/go/api/config/pipeline_groups", nil, tmpPipelineGroups)
	if err != nil {
		return []*PipelineGroup{}, err
	}
	// create the good pipeline groups structures and copy data from the temporary structs
	pipelineGroups := make([]*PipelineGroup, len(*tmpPipelineGroups))
	for i, pg := range *tmpPipelineGroups {
		pipelineGroups[i] = &PipelineGroup{Name: pg.Name}
		pipelineGroups[i].Pipelines = make([]pipeline, len(pg.Pipelines))
		for j, p := range pg.Pipelines {
			pipelineGroups[i].Pipelines[j] = pipeline{Name: p.Name, Label: p.Label, Materials: p.Materials, Stages: make([]string, len(p.Stages))}
			for k, s := range p.Stages {
				pipelineGroups[i].Pipelines[j].Stages[k] = s.Name
			}
		}

	}
	return pipelineGroups, err
}
