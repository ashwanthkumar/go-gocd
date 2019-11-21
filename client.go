package gocd

// Client interface that exposes all the API methods supported by the underlying Client
type Client interface {
	// Agents API
	GetAllAgents() ([]*Agent, error)
	GetAgent(uuid string) (*Agent, error)
	UpdateAgent(uuid string, agent *Agent) (*Agent, error)
	DisableAgent(uuid string) error
	EnableAgent(uuid string) error
	DeleteAgent(uuid string) error
	AgentRunJobHistory(uuid string, offset int) (*JobRunHistory, error)

	// Pipeline Groups API
	GetPipelineGroups() ([]*PipelineGroup, error)

	// Pipelines API
	GetPipelineInstance(string, int) (*PipelineInstance, error)
	GetPipelineHistoryPage(string, int) (*PipelineHistoryPage, error)
	GetPipelineStatus(string) (*PipelineStatus, error)
	PausePipeline(string, string) (*SimpleMessage, error)
	UnpausePipeline(string) (*SimpleMessage, error)
	UnlockPipeline(string) (*SimpleMessage, error)

	// Jobs API
	GetScheduledJobs() ([]*ScheduledJob, error)
	GetJobHistory(pipeline, stage, job string, offset int) ([]*JobHistory, error)

	// Environment Config API
	GetAllEnvironmentConfigs() ([]*EnvironmentConfig, error)
	GetEnvironmentConfig(name string) (*EnvironmentConfig, error)

	// Server health
	GetServerHealthMessages() ([]*ServerHealthMessage, error)
}
