package gocd

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllAgents(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/agents", serveFileAsJSON(t, "GET", "test-fixtures/get_all_agents.json", 6, DummyRequestBodyValidator))
	defer server.Close()
	agents, err := client.GetAllAgents()
	assert.NoError(t, err)
	assert.NotNil(t, agents)
	assert.Equal(t, 1, len(agents))
	agent1 := agents[0]
	assert.NotNil(t, agent1)
	assert.Equal(t, "adb9540a-b954-4571-9d9b-2f330739d4da", agent1.UUID)
	assert.Equal(t, "agent01.example.com", agent1.Hostname)
	assert.Equal(t, "10.12.20.47", agent1.IPAddress)
	assert.Equal(t, "/Users/ketanpadegaonkar/projects/gocd/gocd/agent", agent1.Sandbox)
	assert.Equal(t, "Mac OS X", agent1.OperatingSystem)
	assert.Equal(t, "Enabled", agent1.AgentConfigState)
	assert.Equal(t, "Idle", agent1.AgentState)
	assert.Equal(t, "Idle", agent1.BuildState)
	assert.Equal(t, []string{"java", "linux", "firefox"}, agent1.Resources)
	assert.Equal(t, []string{"perf", "UAT"}, agent1.Env)
}

func TestGetAgent(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/agents/uuid", serveFileAsJSON(t, "GET", "test-fixtures/get_agent.json", 6, DummyRequestBodyValidator))
	defer server.Close()
	agent, err := client.GetAgent("uuid")
	assert.NoError(t, err)
	assert.NotNil(t, agent)
	assert.Equal(t, "adb9540a-b954-4571-9d9b-2f330739d4da", agent.UUID)
	assert.Equal(t, "ketanpkr.corporate.thoughtworks.com", agent.Hostname)
	assert.Equal(t, "10.12.20.47", agent.IPAddress)
	assert.Equal(t, "/Users/ketanpadegaonkar/projects/gocd/gocd/agent", agent.Sandbox)
	assert.Equal(t, "Mac OS X", agent.OperatingSystem)
	assert.Equal(t, "Enabled", agent.AgentConfigState)
	assert.Equal(t, "Idle", agent.AgentState)
	assert.Equal(t, "Idle", agent.BuildState)
	assert.Equal(t, []string{"java", "linux", "firefox"}, agent.Resources)
	assert.Equal(t, []string{"perf", "UAT"}, agent.Env)
}

func TestUpdateAgent(t *testing.T) {
	t.Parallel()
	requestBodyValidator := func(body string) error {
		expectedBody := "{\"build_details\":{},\"hostname\":\"agent02.example.com\"}"
		if body != expectedBody {
			return fmt.Errorf("Request body (%s) didn't match the expected body (%s)", body, expectedBody)
		}
		return nil
	}

	client, server := newTestAPIClient("/go/api/agents/uuid", serveFileAsJSON(t, "PATCH", "test-fixtures/patch_agent.json", 6, requestBodyValidator))
	defer server.Close()
	var agent = Agent{
		Hostname: "agent02.example.com",
	}
	updatedAgent, err := client.UpdateAgent("uuid", &agent)
	assert.NoError(t, err)
	assert.NotNil(t, updatedAgent)
	assert.Equal(t, "agent02.example.com", updatedAgent.Hostname)
}

func TestDeleteAgent(t *testing.T) {
	t.Parallel()

	client, server := newTestAPIClient("/go/api/agents/uuid", serveFileAsJSON(t, "DELETE", "test-fixtures/delete_agent.json", 6, DummyRequestBodyValidator))
	defer server.Close()
	err := client.DeleteAgent("uuid")
	assert.NoError(t, err)
}

func TestAgentRunHistory(t *testing.T) {
	t.Parallel()

	client, server := newTestAPIClient("/go/api/agents/uuid/job_run_history/0", serveFileAsJSON(t, "GET", "test-fixtures/get_agent_run_history.json", 0, DummyRequestBodyValidator))
	defer server.Close()
	history, err := client.AgentRunJobHistory("uuid", 0)
	assert.NoError(t, err)
	jobs := history.Jobs
	assert.NotNil(t, jobs)
	assert.Equal(t, 1, len(jobs))
	job1 := jobs[0]
	assert.NotNil(t, job1)
	assert.Equal(t, "5c5c318f-e6d3-4299-9120-7faff6e6030b", job1.AgentUUID)
	assert.Equal(t, "upload", job1.Name)
	assert.Equal(t, []JobStateTransition{{
		StateChangeTime: 1435631497131,
		ID:              539906,
		State:           "Scheduled",
	}}, job1.JobStateTransitions)
	assert.Equal(t, 1435631497131, job1.ScheduledDate)
	assert.Equal(t, "", job1.OriginalJobID)
	assert.Equal(t, 251, job1.PipelineCounter)
	assert.Equal(t, false, job1.ReRun)
	assert.Equal(t, "distributions-all", job1.PipelineName)
	assert.Equal(t, "Passed", job1.Result)
	assert.Equal(t, "Completed", job1.State)
	assert.Equal(t, 100129, job1.ID)
	assert.Equal(t, "1", job1.StageCounter)
	assert.Equal(t, "upload-installers", job1.StageName)

	pagination := history.Pagination
	assert.Equal(t, pagination.Total, 1292)
	assert.Equal(t, pagination.Offset, 0)
	assert.Equal(t, pagination.PageSize, 10)
}

func TestFreeSpace(t *testing.T) {
	tc := map[string]map[string]FreeSpace{
		"numbers": {
			"1337": 1337,
			"0":    0,
		},
		"strings": {
			"\"unknown\"": -1,
			"\"\"":        -1,
		},
	}
	for desc, c := range tc {
		t.Run(desc, func(t *testing.T) {
			for js, expected := range c {
				var f FreeSpace
				if err := json.Unmarshal([]byte(js), &f); err != nil {
					t.Fatal(err)
				}
				if f != expected {
					t.Error("invalid value", f)
				}
			}
		})
	}

}
