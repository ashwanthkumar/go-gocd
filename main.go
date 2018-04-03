package gocd

import (
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/parnurzeal/gorequest"
)

// DefaultClient entrypoint for GoCD
type DefaultClient struct {
	Host    string `json:"host"`
	Request *gorequest.SuperAgent
}

// New GoCD Client
func New(host, username, password string) Client {
	client := DefaultClient{
		Host:    host,
		Request: gorequest.New().Timeout(60*time.Second).SetBasicAuth(username, password),
	}
	return &client
}

func (c *DefaultClient) resolve(resource string) string {
	// TODO: Use a proper URL resolve to parse the string and append the resource
	return c.Host + resource
}

// getJSON executes a query against the given url with the given headers and
// modify the res object given as reference
// usage:
// err := c.getJSON(fmt.Sprintf("/go/api/pipelines/%s/history/%d", name, offset), nil, res)
// err := c.getJSON(fmt.Sprintf("/go/api/pipelines/%s/history/%d", name, offset), map[string]string{"Accept": "application/vnd.go.cd.v2+json"}, res)
func (c *DefaultClient) getJSON(url string, headers map[string]string, res interface{}) error {
	var errors *multierror.Error
	req := c.Request.Get(c.resolve(url))
	for k, v := range headers {
		req.Set(k, v)
	}
	if _, _, errs := req.EndStruct(res); errs != nil {
		errors = multierror.Append(errors, errs...)
	}
	return errors.ErrorOrNil()
}
