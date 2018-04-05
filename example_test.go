package gocd_test

import (
	"fmt"

	"github.com/ashwanthkumar/go-gocd"
)

// ExampleDefaultClient_PipelineGetInstance displays an instance of a pipeline
// run using the PipelineGetInstance method
func ExampleDefaultClient_PipelineGetInstance() {
	client := gocd.New("http://localhost:8153", "admin", "badger")
	p, err := client.PipelineGetInstance("my-pipeline-name", 911)
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

// ExampleDefaultClient_PipelineGetHistoryPage gets the pipeline runs
// from 2nd to the last to 15th to the last and displays informations about it.
// Uses the PipelineGetHistoryPage method.
func ExampleDefaultClient_PipelineGetHistoryPage() {
	client := gocd.New("http://localhost:8153", "admin", "badger")

	offset := 2      // we ignore the 2 last pipeline runs
	iterations := 15 // We want to stop iterating after we displayed 15 pipelines
	for {
		h, err := client.PipelineGetHistoryPage("my-pipeline-name", offset)
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

// ExampleDefaultClient_PipelineGetStatus shows an example on how to use
// PipelineGetStatus
func ExampleDefaultClient_PipelineGetStatus() {
	client := gocd.New("http://localhost:8153", "admin", "badger")
	name := "my-pipeline-name"
	p, err := client.PipelineGetStatus(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Pipeline %s status: %#v\n", name, p)
}

// ExampleDefaultClient_PipelineUnpause shows an example on how to use
// PipelineUnpause and double-checking the status with PipelineGetStatus
func ExampleDefaultClient_PipelineUnpause() {
	client := gocd.New("http://localhost:8153", "admin", "badger")
	name := "my-pipeline-name"
	p, err := client.PipelineGetStatus(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !p.Paused {
		_, err = client.PipelineUnpause(name)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	c, err := client.PipelineGetStatus(name)
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
