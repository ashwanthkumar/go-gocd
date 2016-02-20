package gocd

import (
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
)

// Agent Object
type Agent struct {
	UUID             string   `json:"uuid,omitempty"`
	Hostname         string   `json:"hostname,omitempty"`
	IPAddress        string   `json:"ip_address,omitempty"`
	Sandbox          string   `json:"sandbox,omitempty"`
	OperatingSystem  string   `json:"operating_system,omitempty"`
	FreeSpace        int      `json:"free_space,omitempty"`
	AgentConfigState string   `json:"agent_config_state,omitempty"`
	AgentState       string   `json:"agent_state,omitempty"`
	BuildState       string   `json:"build_state,omitempty"`
	Resources        []string `json:"resources,omitempty"`
	Env              []string `json:"environments,omitempty"`
}

// GetAllAgents - Lists all available agents, these are agents that are present in the <agents/> tag inside cruise-config.xml and also agents that are in Pending state awaiting registration.
func (c *Client) GetAllAgents() ([]*Agent, error) {
	var errors *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve("/go/api/agents")).
		Set("Accept", "application/vnd.go.cd.v2+json").
		End()
	multierror.Append(errors, errs...)

	type EmbeddedObj struct {
		Agents []*Agent `json:"agents"`
	}
	type AllAgentsResponse struct {
		Embedded EmbeddedObj `json:"_embedded"`
	}
	var responseFormat *AllAgentsResponse

	jsonErr := json.Unmarshal([]byte(body), &responseFormat)
	multierror.Append(errors, jsonErr)
	return responseFormat.Embedded.Agents, errors.ErrorOrNil()
}

// GetAgent - Gets an agent by its unique identifier (uuid)
func (c *Client) GetAgent(uuid string) (*Agent, error) {
	var errors *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/agents/%s", uuid))).
		Set("Accept", "application/vnd.go.cd.v2+json").
		End()
	multierror.Append(errors, errs...)

	var agent *Agent

	jsonErr := json.Unmarshal([]byte(body), &agent)
	multierror.Append(errors, jsonErr)
	return agent, errors.ErrorOrNil()
}

// UpdateAgent - Update some attributes of an agent (uuid).
// Returns the updated agent properties
func (c *Client) UpdateAgent(uuid string, agent *Agent) (*Agent, error) {
	var errors *multierror.Error

	_, body, errs := c.Request.
		Patch(c.resolve(fmt.Sprintf("/go/api/agents/%s", uuid))).
		Set("Accept", "application/vnd.go.cd.v2+json").
		SendStruct(agent).
		End()
	multierror.Append(errors, errs...)

	var updatedAgent *Agent

	jsonErr := json.Unmarshal([]byte(body), &updatedAgent)
	multierror.Append(errors, jsonErr)
	return updatedAgent, errors.ErrorOrNil()
}

// DeleteAgent - Deletes an agent.
// PS: You must first disable an agent and ensure that its status is not Building,
// before attempting to deleting it.
func (c *Client) DeleteAgent(uuid string) error {
	var errors *multierror.Error

	_, _, errs := c.Request.
		Delete(c.resolve(fmt.Sprintf("/go/api/agents/%s", uuid))).
		Set("Accept", "application/vnd.go.cd.v2+json").
		End()
	multierror.Append(errors, errs...)
	return errors.ErrorOrNil()
}
