package gocd

import (
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
)

// Agent Object
type Agent struct {
	UUID            string `json:"uuid,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	IPAddress       string `json:"ip_address,omitempty"`
	Sandbox         string `json:"sandbox,omitempty"`
	OperatingSystem string `json:"operating_system,omitempty"`
	// FreeSpace        int      `json:"free_space,string,omitempty"` - There's inconsistency on how this field is being returned in the API
	AgentConfigState string   `json:"agent_config_state,omitempty"`
	AgentState       string   `json:"agent_state,omitempty"`
	BuildState       string   `json:"build_state,omitempty"`
	Resources        []string `json:"resources,omitempty"`
	Env              []string `json:"environments,omitempty"`
}

// GetAllAgents - Lists all available agents, these are agents that are present in the <agents/> tag inside cruise-config.xml and also agents that are in Pending state awaiting registration.
func (c *DefaultClient) GetAllAgents() ([]*Agent, error) {
	type EmbeddedObj struct {
		Agents []*Agent `json:"agents"`
	}
	type AllAgentsResponse struct {
		Embedded EmbeddedObj `json:"_embedded"`
	}

	responseFormat := new(AllAgentsResponse)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v2+json"}
	_, err := c.getJSON("/go/api/agents", headers, responseFormat)
	return responseFormat.Embedded.Agents, err
}

// GetAgent - Gets an agent by its unique identifier (uuid)
func (c *DefaultClient) GetAgent(uuid string) (*Agent, error) {
	res := new(Agent)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v2+json"}
	_, err := c.getJSON(fmt.Sprintf("/go/api/agents/%s", uuid), headers, res)
	return res, err
}

// UpdateAgent - Update some attributes of an agent (uuid).
// Returns the updated agent properties
func (c *DefaultClient) UpdateAgent(uuid string, agent *Agent) (*Agent, error) {
	var errors *multierror.Error

	_, body, errs := c.Request.
		Patch(c.resolve(fmt.Sprintf("/go/api/agents/%s", uuid))).
		Set("Accept", "application/vnd.go.cd.v2+json").
		SendStruct(agent).
		End()
	errors = multierror.Append(errors, errs...)
	if errs != nil {
		return nil, errors.ErrorOrNil()
	}

	var updatedAgent *Agent

	jsonErr := json.Unmarshal([]byte(body), &updatedAgent)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
	}
	return updatedAgent, errors.ErrorOrNil()
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
		Set("Accept", "application/vnd.go.cd.v2+json").
		End()
	if len(errs) > 0 {
		errors = multierror.Append(errors, errs...)
	}
	return errors.ErrorOrNil()
}

// AgentRunJobHistory - Lists the jobs that have executed on an agent.
func (c *DefaultClient) AgentRunJobHistory(uuid string, offset int) ([]*JobHistory, error) {
	type JobHistoryResponse struct {
		Jobs []*JobHistory `json:"jobs"`
	}
	res := new(JobHistoryResponse)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v2+json"}
	_, err := c.getJSON(fmt.Sprintf("/go/api/agents/%s/job_run_history/%d", uuid, offset), headers, res)
	return res.Jobs, err
}
