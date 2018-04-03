package gocd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPipelineInstance(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/pipelines/PipelineName/instance/1", serveFileAsJSON(t, "GET", "test-fixtures/get_pipeline_instance.json", 0, DummyRequestBodyValidator))
	defer server.Close()

	pipeline, err := client.GetPipelineInstance("PipelineName", 1)
	assert.NoError(t, err)
	assert.NotNil(t, pipeline)
	assert.Equal(t, "PipelineName", pipeline.Name)

	assert.Equal(t, 1, len(pipeline.Stages))
	stg := pipeline.Stages[0]
	assert.Equal(t, "stage1", stg.Name)
	assert.Equal(t, "Passed", stg.Result)

	assert.Equal(t, 1, len(stg.Jobs))
	assert.Equal(t, "jsunit", stg.Jobs[0].Name)

	assert.Equal(t, 1, len(pipeline.BuildCause.MaterialRevisions))
	rev := pipeline.BuildCause.MaterialRevisions[0]
	assert.Equal(t, "61e2da369d0207a7ef61f326eed837f964471b35072340a03f8f55d993afe01d", rev.Material.Fingerprint)
	assert.Equal(t, "Git", rev.Material.Type)
	assert.Equal(t, 1, len(rev.Modifications))
	assert.Equal(t, "my hola mundo changes", rev.Modifications[0].Comment)
}
