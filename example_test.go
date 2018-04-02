package gocd_test

import (
	"fmt"

	"github.com/ashwanthkumar/go-gocd"
)

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
