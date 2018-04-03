package gocd_test

import (
	"fmt"

	"github.com/ashwanthkumar/go-gocd"
)

// ExampleDefaultClient_GetPipelineInstance displays an instance of a pipeline
// run using the GetPipelineInstance method
func ExampleDefaultClient_GetPipelineInstance() {
	client := gocd.New("http://localhost:8153", "admin", "badger")
	p, err := client.GetPipelineInstance("my-pipeline-name", 911)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Run #%d pipeline %s was triggered by %s and ran the following stages:\n", p.Counter, p.Name, p.BuildCause.TriggerMessage)
	for _, stg := range p.Stages {
		fmt.Printf(" * %s with the jobs:\n", stg.Name)
		for _, job := range stg.Jobs {
			fmt.Printf("   * %s is currently %s (%s)\n", job.Name, job.State, job.Result)
		}
	}
}

// ExampleDefaultClient_GetPipelineHistoryPage displays gets the pipeline runs
// from 2nd to the last to 15th to the last and displays informations about it.
// Uses the GetPipelineHistoryPage method.
func ExampleDefaultClient_GetPipelineHistoryPage() {
	client := gocd.New("http://localhost:8153", "admin", "badger")

	offset := 2      // we ignore the 2 last pipeline runs
	iterations := 15 // We want to stop iterating after we displayed 15 pipelines
	for {
		h, err := client.GetPipelineHistoryPage("my-pipeline-name", offset)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, p := range h.Pipelines {
			fmt.Printf("Run #%d pipeline %s was triggered by %s and ran the following stages:\n", p.Counter, p.Name, p.BuildCause.TriggerMessage)
			for _, stg := range p.Stages {
				fmt.Printf(" * %s with the jobs:\n", stg.Name)
				for _, job := range stg.Jobs {
					fmt.Printf("   * %s is currently %s (%s)\n", job.Name, job.State, job.Result)
				}
			}
			iterations--
			if iterations <= 0 {
				return
			}
		}
		offset = h.Pagination.Offset + h.Pagination.PageSize
		if h.Pagination.Total-offset <= 0 {
			break
		}
	}
}
