package gocd

// Pagination is a structure used in several places when the gocd api paginates
// the results. In the history of jobs and pipelines for example
type Pagination struct {
	Offset   int `json:"offset"`
	Total    int `json:"total"`
	PageSize int `json:"page_size"`
}
