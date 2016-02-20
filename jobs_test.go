package gocd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetScheduledJobs(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/jobs/scheduled.xml", serveFileAsXML(t, "GET", "test-fixtures/get_scheduled_jobs.xml"))
	defer server.Close()
	jobs, err := client.GetScheduledJobs()
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	assert.NoError(t, err)
	assert.NotNil(t, jobs)
	assert.Equal(t, 2, len(jobs))
	job1 := jobs[0]
	assert.NotNil(t, job1)
	assert.Equal(t, "job1", job1.Name)
	assert.Equal(t, "6", job1.JobID)
	assert.Equal(t, "mypipeline/5/defaultStage/1/job1", job1.BuildLocator)
	assert.Equal(t, "https://ci.example.com/go/tab/build/detail/mypipeline/5/defaultStage/1/job1", job1.JobURL())
}

func TestGetJobHistory(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/jobs/pipeline/stage/job/history/0", serveFileAsJSON(t, "GET", "test-fixtures/get_job_history.json", DummyRequestBodyValidator))
	defer server.Close()
	jobs, err := client.GetJobHistory("pipeline", "stage", "job", 0)
	assert.NoError(t, err)
	assert.NotNil(t, jobs)
	assert.Equal(t, 2, len(jobs))
	job1 := jobs[1]
	assert.NotNil(t, job1)
	assert.Equal(t, "278fb0b6-d3b8-47e1-9443-67f26bfb5c15", job1.AgentUUID)
	assert.Equal(t, "job1", job1.Name)
	assert.Equal(t, []JobStateTransition{}, job1.JobStateTransitions)
	assert.Equal(t, 1436519733253, job1.ScheduledDate)
	assert.Equal(t, "", job1.OriginalJobID)
	assert.Equal(t, 4, job1.PipelineCounter)
	assert.Equal(t, false, job1.ReRun)
	assert.Equal(t, "mypipeline", job1.PipelineName)
	assert.Equal(t, "Passed", job1.Result)
	assert.Equal(t, "Completed", job1.State)
	assert.Equal(t, 4, job1.ID)
	assert.Equal(t, "1", job1.StageCounter)
	assert.Equal(t, "defaultStage", job1.StageName)
}
