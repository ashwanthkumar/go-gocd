package mocks

import "github.com/ashwanthkumar/go-gocd"
import "github.com/stretchr/testify/mock"

type Client struct {
	mock.Mock
}

// GetAllAgents provides a mock function with given fields:
func (_m *Client) GetAllAgents() ([]*gocd.Agent, error) {
	ret := _m.Called()

	var r0 []*gocd.Agent
	if rf, ok := ret.Get(0).(func() []*gocd.Agent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gocd.Agent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAgent provides a mock function with given fields: uuid
func (_m *Client) GetAgent(uuid string) (*gocd.Agent, error) {
	ret := _m.Called(uuid)

	var r0 *gocd.Agent
	if rf, ok := ret.Get(0).(func(string) *gocd.Agent); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gocd.Agent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAgent provides a mock function with given fields: uuid, agent
func (_m *Client) UpdateAgent(uuid string, agent *gocd.Agent) (*gocd.Agent, error) {
	ret := _m.Called(uuid, agent)

	var r0 *gocd.Agent
	if rf, ok := ret.Get(0).(func(string, *gocd.Agent) *gocd.Agent); ok {
		r0 = rf(uuid, agent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gocd.Agent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *gocd.Agent) error); ok {
		r1 = rf(uuid, agent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableAgent provides a mock function with given fields: uuid
func (_m *Client) DisableAgent(uuid string) error {
	ret := _m.Called(uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableAgent provides a mock function with given fields: uuid
func (_m *Client) EnableAgent(uuid string) error {
	ret := _m.Called(uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAgent provides a mock function with given fields: uuid
func (_m *Client) DeleteAgent(uuid string) error {
	ret := _m.Called(uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AgentRunJobHistory provides a mock function with given fields: uuid, offset
func (_m *Client) AgentRunJobHistory(uuid string, offset int) ([]*gocd.JobHistory, error) {
	ret := _m.Called(uuid, offset)

	var r0 []*gocd.JobHistory
	if rf, ok := ret.Get(0).(func(string, int) []*gocd.JobHistory); ok {
		r0 = rf(uuid, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gocd.JobHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(uuid, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScheduledJobs provides a mock function with given fields:
func (_m *Client) GetScheduledJobs() ([]*gocd.ScheduledJob, error) {
	ret := _m.Called()

	var r0 []*gocd.ScheduledJob
	if rf, ok := ret.Get(0).(func() []*gocd.ScheduledJob); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gocd.ScheduledJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobHistory provides a mock function with given fields: pipeline, stage, job, offset
func (_m *Client) GetJobHistory(pipeline string, stage string, job string, offset int) ([]*gocd.JobHistory, error) {
	ret := _m.Called(pipeline, stage, job, offset)

	var r0 []*gocd.JobHistory
	if rf, ok := ret.Get(0).(func(string, string, string, int) []*gocd.JobHistory); ok {
		r0 = rf(pipeline, stage, job, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gocd.JobHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, int) error); ok {
		r1 = rf(pipeline, stage, job, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
