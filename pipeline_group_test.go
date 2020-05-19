package gocd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPipelineGroups(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/admin/pipeline_groups", serveFileAsJSON(t, "GET", "test-fixtures/get_pipeline_groups.json", 0, DummyRequestBodyValidator))
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
	assert.Empty(t, p.Label)
	assert.Nil(t, p.Materials)
}
