package gocd

type Material struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Fingerprint string `json:"fingerprint"`
}
