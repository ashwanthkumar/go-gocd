package gocd

import (
	"encoding/json"
	"errors"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
)

// Agent Object
type Agent struct {
	UUID             string    `json:"uuid,omitempty"`
	Hostname         string    `json:"hostname,omitempty"`
	IPAddress        string    `json:"ip_address,omitempty"`
	Sandbox          string    `json:"sandbox,omitempty"`
	OperatingSystem  string    `json:"operating_system,omitempty"`
	FreeSpace        FreeSpace `json:"free_space,omitempty"`
	AgentConfigState string    `json:"agent_config_state,omitempty"`
	AgentState       string    `json:"agent_state,omitempty"`
	BuildState       string    `json:"build_state,omitempty"`
	BuildDetails     struct {
		PipelineName string `json:"pipeline_name,omitempty"`
		StageName    string `json:"stage_name,omitempty"`
		JobName      string `json:"job_name,omitempty"`
	} `json:"build_details,omitempty"`
	Resources []string `json:"resources,omitempty"`
	Env       []string `json:"environments,omitempty"`
}

// FreeSpace is required for GoCD API inconsistencies in agent free space scrape.
type FreeSpace int

// UnmarshalJSON expects an int or string ("unknown").
func (i *FreeSpace) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	var js interface{}
	if err := json.Unmarshal(data, &js); err != nil {
		return err
	}
	switch v := js.(type) {
	case string:
		// such as "unknown"
		*i = -1
	case float64:
		*i = FreeSpace(v)
	default:
		return errors.New("FreeSpace: unexpected type")
	}
	return nil
}

// GetAllAgents - Lists all available agents, these are agents that are present in the <agents/> tag inside cruise-config.xml and also agents that are in Pending state awaiting registration.
func (c *DefaultClient) GetAllAgents() ([]*Agent, error) {
	type EmbeddedObj struct {
		Agents []*Agent `json:"agents"`
	}
	type AllAgentsResponse struct {
		Embedded EmbeddedObj `json:"_embedded"`
	}
	res := new(AllAgentsResponse)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v6+json"}
	err := c.getJSON("/go/api/agents", headers, res)
	if err != nil {
		return []*Agent{}, err
	}
	return res.Embedded.Agents, nil
}

// GetAgent - Gets an agent by its unique identifier (uuid)
func (c *DefaultClient) GetAgent(uuid string) (*Agent, error) {
	res := new(Agent)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v6+json"}
	err := c.getJSON(fmt.Sprintf("/go/api/agents/%s", uuid), headers, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateAgent - Update some attributes of an agent (uuid).
// Returns the updated agent properties
func (c *DefaultClient) UpdateAgent(uuid string, agent *Agent) (*Agent, error) {
	res := new(Agent)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v6+json"}
	err := c.patchJSON(fmt.Sprintf("/go/api/agents/%s", uuid), headers, agent, res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

// DisableAgent - Disables an agent using it's UUID
func (c *DefaultClient) DisableAgent(uuid string) error {
	var agent = &Agent{
		AgentConfigState: "Disabled",
	}
	_, err := c.UpdateAgent(uuid, agent)
	return err
}

// EnableAgent - Enables an agent using it's UUID
func (c *DefaultClient) EnableAgent(uuid string) error {
	var agent = &Agent{
		AgentConfigState: "Enabled",
	}
	_, err := c.UpdateAgent(uuid, agent)
	return err
}

// DeleteAgent - Deletes an agent.
// PS: You must first disable an agent and ensure that its status is not Building,
// before attempting to deleting it.
func (c *DefaultClient) DeleteAgent(uuid string) error {
	var errors *multierror.Error

	_, _, errs := c.Request.
		Delete(c.resolve(fmt.Sprintf("/go/api/agents/%s", uuid))).
		Set("Accept", "application/vnd.go.cd.v6+json").
		End()
	if len(errs) > 0 {
		errors = multierror.Append(errors, errs...)
	}
	return errors.ErrorOrNil()
}

// AgentRunJobHistory - Lists the jobs that have executed on an agent.
func (c *DefaultClient) AgentRunJobHistory(uuid string, offset int) (*JobRunHistory, error) {
	res := new(JobRunHistory)
	headers := map[string]string{"Accept": "application/json"}
	err := c.getJSON(fmt.Sprintf("/go/api/agents/%s/job_run_history/%d", uuid, offset), headers, res)
	return res, err
}
