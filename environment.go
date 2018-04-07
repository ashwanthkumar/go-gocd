package gocd

import (
	"encoding/json"
	"fmt"
)

// EnvironmentVariable is used to configure environment variables of stages and
// jobs as well as in the environment config.
// See: https://api.gocd.org/current/#the-environment-variable-object
type EnvironmentVariable struct {
	Secure         bool   `json:"secure"`
	Name           string `json:"name"`
	Value          string `json:"value,omitempty"`
	EncryptedValue string `json:"encrypted_value,omitempty"`
}

// EnvironmentConfig Object
type EnvironmentConfig struct {
	Name                 string                `json:"name,omitempty"`
	Pipelines            []string              `json:"pipelines"`
	Agents               []string              `json:"agents"`
	EnvironmentVariables []EnvironmentVariable `json:"environment_variables"`
}

// UnmarshalJSON overriding it for the dynamic material attributes type
func (ec *EnvironmentConfig) UnmarshalJSON(b []byte) error {
	type pipeline struct {
		Name string `json:"name"`
	}
	type agent struct {
		UUID string `json:"uuid"`
	}
	type tmpConfig struct {
		Name                 string                `json:"name,omitempty"`
		Pipelines            []pipeline            `json:"pipelines"`
		Agents               []agent               `json:"agents"`
		EnvironmentVariables []EnvironmentVariable `json:"environment_variables"`
	}
	var tmpec *tmpConfig
	if err := json.Unmarshal(b, &tmpec); err != nil {
		return err
	}
	ec.Name = tmpec.Name
	for _, v := range tmpec.Pipelines {
		ec.Pipelines = append(ec.Pipelines, v.Name)
	}
	for _, v := range tmpec.Agents {
		ec.Agents = append(ec.Agents, v.UUID)
	}
	ec.EnvironmentVariables = tmpec.EnvironmentVariables
	return nil
}

// GetAllEnvironmentConfigs - Lists all available environments.
func (c *DefaultClient) GetAllEnvironmentConfigs() ([]*EnvironmentConfig, error) {
	type EmbeddedObj struct {
		Environments []*EnvironmentConfig `json:"environments"`
	}
	type AllAgentsResponse struct {
		Embedded EmbeddedObj `json:"_embedded"`
	}

	res := new(AllAgentsResponse)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v2+json"}
	_, err := c.getJSON("/go/api/admin/environments", headers, res)
	return res.Embedded.Environments, err
}

// GetEnvironmentConfig - Gets environment config for specified environment name.
func (c *DefaultClient) GetEnvironmentConfig(name string) (*EnvironmentConfig, error) {
	res := new(EnvironmentConfig)
	headers := map[string]string{"Accept": "application/vnd.go.cd.v2+json"}
	_, err := c.getJSON(fmt.Sprintf("/go/api/admin/environments/%s", name), headers, res)
	return res, err
}
