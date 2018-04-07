package gocd

// ArtifactConfig is used to specify how to configure your artifacts in a job
// when creating a pipeline.
// See: https://api.gocd.org/current/#the-pipeline-config-artifact-object
type ArtifactConfig struct {
	Type        string `json:"type"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}
