package gocd

import (
	"encoding/json"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"
)

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
	var errors *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve("/go/api/admin/environments")).
		Set("Accept", "application/vnd.go.cd.v2+json").
		End()
	if errs != nil {
		errors = multierror.Append(errors, errs...)
		return []*EnvironmentConfig{}, errors.ErrorOrNil()
	}

	type EmbeddedObj struct {
		Environments []*EnvironmentConfig `json:"environments"`
	}
	type AllAgentsResponse struct {
		Embedded EmbeddedObj `json:"_embedded"`
	}
	var responseFormat *AllAgentsResponse

	jsonErr := json.Unmarshal([]byte(body), &responseFormat)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
		return []*EnvironmentConfig{}, errors.ErrorOrNil()
	}
	return responseFormat.Embedded.Environments, errors.ErrorOrNil()
}

// GetEnvironmentConfig - Gets environment config for specified environment name.
func (c *DefaultClient) GetEnvironmentConfig(name string) (*EnvironmentConfig, error) {
	var errors *multierror.Error

	_, body, errs := c.Request.
		Get(c.resolve(fmt.Sprintf("/go/api/admin/environments/%s", name))).
		Set("Accept", "application/vnd.go.cd.v2+json").
		End()
	errors = multierror.Append(errors, errs...)
	if errs != nil {
		return nil, errors.ErrorOrNil()
	}

	var environment *EnvironmentConfig

	jsonErr := json.Unmarshal([]byte(body), &environment)
	if jsonErr != nil {
		errors = multierror.Append(errors, jsonErr)
	}
	return environment, errors.ErrorOrNil()
}
