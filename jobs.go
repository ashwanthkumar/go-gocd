package gocd

import (
	"encoding/xml"

	multierror "github.com/hashicorp/go-multierror"
)

// ScheduledJob instance
type ScheduledJob struct {
	Name         string    `xml:"name,attr"`
	JobID        string    `xml:"id,attr"`
	BuildLocator string    `xml:"buildLocator"`
	Link         LinkInXML `xml:"link"`
}

// JobURL - Full URL location of the scheduled job
func (job *ScheduledJob) JobURL() string {
	return job.Link.Href
}

// LinkInXML - <link rel="..." href="..."> tag
type LinkInXML struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

// GetScheduledJobs - Lists all the current job instances which are scheduled but not yet assigned to any agent.
func (c *Client) GetScheduledJobs() ([]*ScheduledJob, error) {
	var errors *multierror.Error

	type ScheduledJobsResponse struct {
		XMLName xml.Name        `xml:"scheduledJobs"`
		Jobs    []*ScheduledJob `xml:"job"`
	}

	var jobs ScheduledJobsResponse
	_, body, errs := c.Request.
		Get(c.resolve("/go/api/jobs/scheduled.xml")).
		End()
	multierror.Append(errors, errs...)
	xmlErr := xml.Unmarshal([]byte(body), &jobs)
	multierror.Append(errors, xmlErr)

	return jobs.Jobs, errors.ErrorOrNil()
}
