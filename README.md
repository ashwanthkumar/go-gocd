[![Build Status](https://snap-ci.com/ashwanthkumar/go-gocd/branch/master/build_image)](https://snap-ci.com/ashwanthkumar/go-gocd/branch/master) [![GoDoc](https://godoc.org/github.com/ashwanthkumar/go-gocd?status.svg)](https://godoc.org/github.com/ashwanthkumar/go-gocd)

# go-gocd

Go Lang library to access [GoCD API](https://api.go.cd/current/).

## Usage
```go
package main

import (
  "github.com/ashwanthkumar/go-gocd"
)

func main() {
  client := gocd.New("http://localhost:8153", "admin", "badger")
  agents, err := client.GetAllAgents()
  // ... do whatever you want with the agents
}

```

## API Endpoints Pending
- [x] Agents
  - [x] Get all Agents
  - [x] Get one Agent
  - [x] Update an Agent
  - [x] Disable Agent
  - [x] Delete an Agent
  - [x] Agent job run history
- [ ] Users
  - [ ] Get all Users
  - [ ] Get one user
  - [ ] Create a user
  - [ ] Update a user
  - [ ] Delete a user
- [ ] Materials
  - [ ] Get all Materials
  - [ ] Get material modifications
  - [ ] Notify SVN materials
  - [ ] Notify git materials
- [ ] Backups
  - [ ] Create a backup
- [ ] Pipeline Group
  - [ ] Config listing
- [ ] Artifacts
  - [ ] Get all Artifacts
  - [ ] Get artifact file
  - [ ] Get artifact directory
  - [ ] Create artifact
  - [ ] Append to artifact
- [ ] Pipelines
  - [X] Get pipeline instance
  - [X] Get pipeline status
  - [X] Pause a pipeline
  - [X] Unpause a pipeline
  - [X] Releasing a pipeline lock
  - [ ] Scheduling pipelines
  - [ ] Create a pipeline
  - [ ] Delete a pipeline
- [ ] Stages
  - [ ] Cancel Stage
  - [ ] Get Stage instance
  - [ ] Get stage history
- [x] Jobs
  - [x] Get Scheduled Jobs
  - [x] Get Job history
- [ ] Properties
  - [ ] Get all job Properties
  - [ ] Get one property
  - [ ] Get historical properties
  - [ ] Create property
- [ ] Configurations
  - [ ] List all modifications
  - [ ] Get repository modification diff
  - [ ] Get Configuration
  - Create pipeline (Deprecated API)
- Feeds (will not support)
- [ ] Dashboard
  - [ ] Get Dashboard
- [ ] Pipeline Config
  - [ ] Get pipeline Configuration
  - [ ] Edit Pipeline configuration
  - [ ] Create Pipeline
- [ ] Environment Config
  - [x] Get all environments
  - [x] Get environment config
  - [ ] Create an environment
  - [ ] Update an environment
  - [ ] Patch an environment
  - [ ] Delete an environment
