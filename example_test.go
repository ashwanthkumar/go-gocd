package gocd_test

import (
	"fmt"

	"github.com/pagero/go-gocd-1"
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

// ExampleDefaultClient_GetPipelineStatus shows an example on how to use
// GetPipelineStatus
func ExampleDefaultClient_GetPipelineStatus() {
	client := gocd.New("http://localhost:8153", "admin", "badger")
	name := "my-pipeline-name"
	p, err := client.GetPipelineStatus(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Pipeline %s status: %#v\n", name, p)
}

// ExampleDefaultClient_UnpausePipeline shows an example on how to use
// UnpausePipeline and double-checking the status with GetPipelineStatus
func ExampleDefaultClient_UnpausePipeline() {
	client := gocd.New("http://localhost:8153", "admin", "badger")
	name := "my-pipeline-name"
	p, err := client.GetPipelineStatus(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !p.Paused {
		_, err = client.UnpausePipeline(name)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	c, err := client.GetPipelineStatus(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if c.Paused {
		fmt.Printf("Pipeline %s is now paused\n", name)
	} else {
		fmt.Printf("Pipeline %s does seem to still be unpaused\n", name)
	}

}
