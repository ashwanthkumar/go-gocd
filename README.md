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
