package gocd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllEnvironments(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/admin/environments", serveFileAsJSON(t, "GET", "test-fixtures/get_all_environment_configs.json", DummyRequestBodyValidator))
	defer server.Close()
	envs, err := client.GetAllEnvironmentConfigs()
	assert.NoError(t, err)
	assert.NotNil(t, envs)
	assert.Equal(t, 1, len(envs))
	env1 := envs[0]
	assert.NotNil(t, env1)
	assert.Equal(t, "foobar", env1.Name)
	assert.Equal(t, 2, len(env1.EnvironmentVariables))
	var1 := env1.EnvironmentVariables[0]
	assert.False(t, var1.Secure)
	assert.Equal(t, "username", var1.Name)
	assert.Equal(t, "admin", var1.Value)
	assert.Empty(t, var1.EncryptedValue)
	var2 := env1.EnvironmentVariables[1]
	assert.True(t, var2.Secure)
	assert.Equal(t, "password", var2.Name)
	assert.Empty(t, var2.Value)
	assert.Equal(t, "LSd1TI0eLa+DjytHjj0qjA==", var2.EncryptedValue)
	assert.Equal(t, 1, len(env1.Pipelines))
	assert.Equal(t, "up42", env1.Pipelines[0])
	assert.Equal(t, 1, len(env1.Agents))
	assert.Equal(t, "12345678-e2f6-4c78-123456789012", env1.Agents[0])
}

func TestGetEnvironment(t *testing.T) {
	t.Parallel()
	client, server := newTestAPIClient("/go/api/admin/environments/my_environment", serveFileAsJSON(t, "GET", "test-fixtures/get_environment_config.json", DummyRequestBodyValidator))
	defer server.Close()
	env, err := client.GetEnvironmentConfig("my_environment")
	assert.NoError(t, err)
	assert.NotNil(t, env)
	assert.Equal(t, "my_environment", env.Name)
	assert.Equal(t, 2, len(env.EnvironmentVariables))
	var1 := env.EnvironmentVariables[0]
	assert.False(t, var1.Secure)
	assert.Equal(t, "username", var1.Name)
	assert.Equal(t, "admin", var1.Value)
	assert.Empty(t, var1.EncryptedValue)
	var2 := env.EnvironmentVariables[1]
	assert.True(t, var2.Secure)
	assert.Equal(t, "password", var2.Name)
	assert.Empty(t, var2.Value)
	assert.Equal(t, "LSd1TI0eLa+DjytHjj0qjA==", var2.EncryptedValue)
	assert.Equal(t, 1, len(env.Pipelines))
	assert.Equal(t, "up42", env.Pipelines[0])
	assert.Equal(t, 1, len(env.Agents))
	assert.Equal(t, "12345678-e2f6-4c78-123456789012", env.Agents[0])
}
