package gocd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetScheduledJobs(t *testing.T) {
	client := newTestClient(t)
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
