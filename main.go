package gocd

import "github.com/parnurzeal/gorequest"

// Client entrypoint for GoCD
type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Request  *gorequest.SuperAgent
}

// New GoCD Client
func New(host, username, password string) *Client {
	client := Client{
		Username: username,
		Password: password,
		Host:     host,
		Request:  gorequest.New(),
	}
	return &client
}

func (c *Client) resolve(resource string) string {
	return c.Host + resource
}
