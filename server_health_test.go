package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetServerHealthMessages(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/server_health_messages", serveFileAsJSON(t, "GET", "test-fixtures/get_server_health_messages.json", 1, DummyRequestBodyValidator))
	defer server.Close()
	messages, err := client.GetServerHealthMessages()
	assert.NoError(t, err)
	assert.NotNil(t, messages)
	assert.Equal(t, 2, len(messages))
	m := messages[0]
	assert.False(t, m.IsError())
	assert.True(t, m.IsWarning())
	assert.Equal(t, m.Message, "Job 'foo/bar/job' is not responding")
	assert.Equal(t, m.Detail, "This job may be hung.")
	assert.Equal(t, m.Time, "2018-02-27T07:36:30Z")
	m = messages[1]
	assert.True(t, m.IsError())
}
