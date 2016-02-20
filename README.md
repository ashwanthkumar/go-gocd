[![Build Status](https://snap-ci.com/ashwanthkumar/go-gocd/branch/master/build_image)](https://snap-ci.com/ashwanthkumar/go-gocd/branch/master) [![GoDoc](https://godoc.org/github.com/ashwanthkumar/go-gocd?status.svg)](https://godoc.org/github.com/ashwanthkumar/go-gocd)

# go-gocd

Go Lang library to access [GoCD API](https://api.go.cd/current/).

## Usage
```go
package main

import (
  gocd "github.com/ashwanthkumar/go-gocd"
)

func main() {
  client := gocd.New("http://localhost:8153", "admin", "badger")
  agents, err := client.GetAllAgents()
  // ... do whatever you want with the agents
}

```

## API Endpoints Pending
- [ ] Agents
-- [x] Get all Agents
-- [ ] Get one Agent
-- [ ] Update an Agent
-- [ ] Delete an Agent
-- [ ] Agent job run history
