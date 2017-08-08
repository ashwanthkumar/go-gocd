package gocd

// EnvironmentVariable Object
type EnvironmentVariable struct {
	Secure         bool   `json:"secure"`
	Name           string `json:"name"`
	Value          string `json:"value,omitempty"`
	EncryptedValue string `json:"encrypted_value,omitempty"`
}
