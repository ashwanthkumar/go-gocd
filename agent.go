package gocd

import (
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
)

// Agent Object
type Agent struct {
	UUID             string   `json:"uuid"`
	Hostname         string   `json:"hostname"`
	IPAddress        string   `json:"ip_address"`
	Sandbox          string   `json:"sandbox"`
	OperatingSystem  string   `json:"operating_system"`
	FreeSpace        int      `json:"free_space"`
	AgentConfigState string   `json:"agent_config_state"`
	AgentState       string   `json:"agent_state"`
	BuildState       string   `json:"build_state"`
	Resources        []string `json:"resources"`
	Env              []string `json:"environments"`
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
