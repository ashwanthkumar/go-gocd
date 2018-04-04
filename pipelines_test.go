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

func TestGetPipelineHistoryPage(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/pipelines/pipeline1/history/0", serveFileAsJSON(t, "GET", "test-fixtures/get_pipeline_history_page.json", 0, DummyRequestBodyValidator))
	defer server.Close()

	history, err := client.GetPipelineHistoryPage("pipeline1", 0)
	assert.Equal(t, 2, len(history.Pipelines))
	assert.NoError(t, err)
	assert.Equal(t, 0, history.Pagination.Offset)
	assert.Equal(t, 2, history.Pagination.Total)
	assert.Equal(t, 10, history.Pagination.PageSize)
	pipeline := history.Pipelines[1]
	assert.NotNil(t, pipeline)
	assert.Equal(t, "pipeline1", pipeline.Name)

	assert.Equal(t, 1, len(pipeline.Stages))
	stg := pipeline.Stages[0]
	assert.Equal(t, "stage1", stg.Name)
	assert.Equal(t, "Passed", stg.Result)

	assert.Equal(t, 1, len(stg.Jobs))
	assert.Equal(t, "job1", stg.Jobs[0].Name)

	assert.Equal(t, 1, len(pipeline.BuildCause.MaterialRevisions))
	rev := pipeline.BuildCause.MaterialRevisions[0]
	assert.Equal(t, "f6e7a3899c55e1682ffb00383bdf8f882bcee2141e79a8728254190a1fddcf4f", rev.Material.Fingerprint)
	assert.Equal(t, "Git", rev.Material.Type)
	assert.Equal(t, 1, len(rev.Modifications))
	assert.Equal(t, "my hola mundo changes", rev.Modifications[0].Comment)
}

func TestGetPipelineStatus(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/pipelines/pipeline1/status", serveFileAsJSON(t, "GET", "test-fixtures/get_pipeline_status.json", 0, DummyRequestBodyValidator))
	defer server.Close()

	s, err := client.GetPipelineStatus("pipeline1")
	assert.NoError(t, err)
	assert.Equal(t, "Reason for pausing this pipeline", s.PausedCause)
	assert.Equal(t, "admin", s.PausedBy)
	assert.Equal(t, true, s.Paused)
	assert.Equal(t, false, s.Schedulable)
	assert.Equal(t, false, s.Locked)
}
