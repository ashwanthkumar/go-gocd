package gocd

import (
	"strings"
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

// getJSON executes a Get query against the given url with the given headers and
// modify the out object given as reference. Also returns the value of the ETag
// header if any was returned by the server.
func (c *DefaultClient) getJSON(url string, headers map[string]string, out interface{}) (string, error) {
	var errors *multierror.Error

	req := c.Request.Get(c.resolve(url))
	for k, v := range headers {
		req.Set(k, v)
	}

	resp, _, errs := req.EndStruct(out)
	if errs != nil {
		errors = multierror.Append(errors, errs...)
	}

	etag := ""
	if t, ok := resp.Header["Etag"]; ok {
		etag = strings.Replace(t[0], `"`, "", -1)
	}

	return etag, errors.ErrorOrNil()
}

// postJSON executes a Post query against the given url with the given headers and
// the using the given "in" struct as data, then modify the out object given as reference
// Warning: if the body returned by gocd is empty, you will need
// https://github.com/parnurzeal/gorequest/pull/185 to avoid json unmarshall
// error messages
func (c *DefaultClient) postJSON(url string, headers map[string]string, in, out interface{}) error {
	var errors *multierror.Error

	req := c.Request.Post(c.resolve(url))
	for k, v := range headers {
		req.Set(k, v)
	}

	if _, _, errs := req.SendStruct(in).End(); errs != nil {
		errors = multierror.Append(errors, errs...)
	}
	return errors.ErrorOrNil()
}
