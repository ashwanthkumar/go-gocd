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
	AgentRunJobHistory(uuid string, offset int) ([]*JobHistory, error)

	// Pipeline Groups API
	GetPipelineGroups() ([]*PipelineGroup, error)

	// Pipelines API
	PipelineGetInstance(string, int) (*PipelineInstance, error)
	PipelineGetHistoryPage(string, int) (*PipelineHistoryPage, error)
	PipelineGetStatus(string) (*PipelineStatus, error)
	PipelinePause(string, string) (*SimpleMessage, error)
	PipelineUnpause(string) (*SimpleMessage, error)
	PipelineUnlock(string) (*SimpleMessage, error)
	PipelineGetConfig(string) (*PipelineConfig, string, error)

	// Jobs API
	GetScheduledJobs() ([]*ScheduledJob, error)
	GetJobHistory(pipeline, stage, job string, offset int) ([]*JobHistory, error)

	// Environment Config API
	GetAllEnvironmentConfigs() ([]*EnvironmentConfig, error)
	GetEnvironmentConfig(name string) (*EnvironmentConfig, error)
}
