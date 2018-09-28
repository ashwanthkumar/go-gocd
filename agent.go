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
	var errors *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve("/go/api/agents")).
		Set("Accept", "application/vnd.go.cd.v4+json").
		End()
	if errs != nil {
		errors = multierror.Append(errors, errs...)
		return []*Agent{}, errors.ErrorOrNil()
	}

	type EmbeddedObj struct {
		Agents []*Agent `json:"agents"`
	}
	type AllAgentsResponse struct {
		Embedded EmbeddedObj `json:"_embedded"`
	}
	var responseFormat *AllAgentsResponse

	jsonErr := json.Unmarshal([]byte(body), &responseFormat)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
		return []*Agent{}, errors.ErrorOrNil()
	}
	return responseFormat.Embedded.Agents, errors.ErrorOrNil()
}

// GetAgent - Gets an agent by its unique identifier (uuid)
func (c *DefaultClient) GetAgent(uuid string) (*Agent, error) {
	var errors *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/agents/%s", uuid))).
		Set("Accept", "application/vnd.go.cd.v2+json").
		End()
	errors = multierror.Append(errors, errs...)
	if errs != nil {
		return nil, errors.ErrorOrNil()
	}

	var agent *Agent

	jsonErr := json.Unmarshal([]byte(body), &agent)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
	}
	return agent, errors.ErrorOrNil()
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
	multierror.Append(errors, errs...)
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
	var errors *multierror.Error
	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/agents/%s/job_run_history/%d", uuid, offset))).
		Set("Accept", "application/vnd.go.cd.v2+json").
		End()
	if errs != nil {
		errors = multierror.Append(errors, errs...)
		return []*JobHistory{}, errors.ErrorOrNil()
	}

	type JobHistoryResponse struct {
		Jobs []*JobHistory `json:"jobs"`
	}
	var jobs *JobHistoryResponse
	jsonErr := json.Unmarshal([]byte(body), &jobs)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
		return []*JobHistory{}, errors.ErrorOrNil()
	}
	return jobs.Jobs, errors.ErrorOrNil()
}
