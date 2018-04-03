package gocd

// Material represents a material (Can be Git, Mercurial, Perforce, Subversion, Tfs, Pipeline, SCM)
type Material struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Fingerprint string `json:"fingerprint"`
}

// MaterialModification represents a modification done on a material configuration
type MaterialModification struct {
	ID           int    `json:"id"`
	ModifiedTime int64  `json:"modified_time"`
	UserName     string `json:"user_name"`
	EmailAddress string `json:"email_address"`
	Comment      string `json:"comment"`
	Revision     string `json:"revision"`
}

// MaterialRevision is a given revision of a material
type MaterialRevision struct {
	Material      Material               `json:"material"`
	Modifications []MaterialModification `json:"modifications"`
	Changed       bool                   `json:"changed"`
}
