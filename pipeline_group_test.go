package gocd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPipelineGroups(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/config/pipeline_groups", serveFileAsJSON(t, "GET", "test-fixtures/get_pipeline_groups.json", 0, DummyRequestBodyValidator))
	defer server.Close()
	pipelineGroups, err := client.GetPipelineGroups()
	assert.NoError(t, err)
	assert.NotNil(t, pipelineGroups)
	assert.Equal(t, 1, len(pipelineGroups))
	pg := pipelineGroups[0]
	assert.NotNil(t, pg)
	assert.Equal(t, "first", pg.Name)
	assert.Equal(t, 1, len(pg.Pipelines))
	p := pg.Pipelines[0]
	assert.Equal(t, "up42", p.Name)
	assert.Equal(t, "${COUNT}", p.Label)
	assert.Equal(t, 1, len(p.Materials))
	material := p.Materials[0]
	assert.Equal(t, "2d05446cd52a998fe3afd840fc2c46b7c7e421051f0209c7f619c95bedc28b88", material.Fingerprint)
	assert.Equal(t, "Git", material.Type)
	assert.Equal(t, "URL: https://github.com/gocd/gocd, Branch: master", material.Description)
	assert.Equal(t, 1, len(p.Stages))
	assert.Equal(t, "up42_stage", p.Stages[0])
}
