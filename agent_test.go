package gocd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllAgents(t *testing.T) {
	client := newTestClient(t)
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
	assert.Equal(t, 84983328768, agent1.FreeSpace)
	assert.Equal(t, "Enabled", agent1.AgentConfigState)
	assert.Equal(t, "Idle", agent1.AgentState)
	assert.Equal(t, "Idle", agent1.BuildState)
	assert.Equal(t, []string{"java", "linux", "firefox"}, agent1.Resources)
	assert.Equal(t, []string{"perf", "UAT"}, agent1.Env)
}
