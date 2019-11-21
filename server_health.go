package gocd

type ServerHealthMessage struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Level   string `json:"level"`
	Time    string `json:"time"`
}

func (s *ServerHealthMessage) IsWarning() bool {
	return s.Level == "WARNING"
}

func (s *ServerHealthMessage) IsError() bool {
	return s.Level == "ERROR"
}

func (c *DefaultClient) GetServerHealthMessages() ([]*ServerHealthMessage, error) {
	res := []*ServerHealthMessage{}
	headers := map[string]string{"Accept": "application/vnd.go.cd.v1+json"}
	err := c.getJSON("/go/api/server_health_messages", headers, &res)
	return res, err
}
