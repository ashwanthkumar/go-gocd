package gocd

// Material represents a material (Can be Git, Mercurial, Perforce, Subversion, Tfs, Pipeline, SCM)
// It is also used for other stuffs like pipeline history
// The "Attributes" field is only used in the configuration of a pipeline
// See: https://api.gocd.org/current/#the-pipeline-material-object
type Material struct {
	ID          int         `json:"id"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Fingerprint string      `json:"fingerprint"`
	Attributes  interface{} `json:"attributes,omitempty"`
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

// MaterialAttributesGit represents the attributes of a git material.
// See: https://api.gocd.org/current/#the-git-material-attributes
type MaterialAttributesGit struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Branch          string `json:"branch,omitempty"`
	Destination     string `json:"destination,omitempty"`
	AutoUpdate      bool   `json:"auto_update,omitempty"`
	Filter          Filter `json:"filter,omitempty"`
	InvertFilter    bool   `json:"invert_filter,omitempty"`
	SubmoduleFolder string `json:"submodule_folder,omitempty"`
	ShallowClone    bool   `json:"shallow_clone,omitempty"`
}

// MaterialAttributesSvn represents the attributes of a svn material.
// See: https://api.gocd.org/current/#the-subversion-material-attributes
type MaterialAttributesSvn struct {
	Name              string `json:"name"`
	URL               string `json:"url"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"encrypted_password,omitempty"`
	Destination       string `json:"destination,omitempty"`
	Filter            Filter `json:"filter,omitempty"`
	InvertFilter      bool   `json:"invert_filter,omitempty"`
	AutoUpdate        bool   `json:"auto_update,omitempty"`
	CheckExternals    bool   `json:"check_external,omitempty"`
}

// MaterialAttributesMercurial represents the attributes of a mercurial material.
// See: https://api.gocd.org/current/#the-mercurial-material-attributes
type MaterialAttributesMercurial struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	Destination  string `json:"destination,omitempty"`
	Filter       Filter `json:"filter,omitempty"`
	InvertFilter bool   `json:"invert_filter,omitempty"`
	AutoUpdate   bool   `json:"auto_update,omitempty"`
}

// MaterialAttributesPerforce represents the attributes of a perforce material.
// See: https://api.gocd.org/current/#the-perforce-material-attributes
type MaterialAttributesPerforce struct {
	Name              string `json:"name"`
	Port              string `json:"port,omitempty"`
	UseTickets        bool   `json:"use_tickets,omitempty"`
	View              string `json:"view,omitempty"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"encrypted_password,omitempty"`
	Destination       string `json:"destination,omitempty"`
	Filter            Filter `json:"filter,omitempty"`
	InvertFilter      bool   `json:"invert_filter,omitempty"`
	AutoUpdate        bool   `json:"auto_update,omitempty"`
}

// MaterialAttributesTfs represents the attributes of a tfs material.
// See: https://api.gocd.org/current/#the-tfs-material-attributes
type MaterialAttributesTfs struct {
	Name              string `json:"name"`
	URL               string `json:"url"`
	ProjectPath       string `json:"project_path,omitempty"`
	Domain            string `json:"domain,omitempty"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"encrypted_password,omitempty"`
	Destination       string `json:"destination,omitempty"`
	AutoUpdate        bool   `json:"auto_update,omitempty"`
	Filter            Filter `json:"filter,omitempty"`
	InvertFilter      bool   `json:"invert_filter,omitempty"`
}

// MaterialAttributesDependency represents the attributes of a dependency material.
// See: https://api.gocd.org/current/#the-dependency-material-attributes
type MaterialAttributesDependency struct {
	Name       string `json:"name"`
	Pipeline   string `json:"pipeline,omitempty"`
	Stage      string `json:"stage,omitempty"`
	AutoUpdate bool   `json:"auto_update,omitempty"`
}

// MaterialAttributesPackage represents the attributes of a package material.
// See: https://api.gocd.org/current/#the-package-material-attributes
type MaterialAttributesPackage struct {
	Ref string `json:"ref,omitempty"`
}

// PluginMaterialAttributes represents the attributes of a plugin material.
// See: https://api.gocd.org/current/#the-plugin-material-attributes
type PluginMaterialAttributes struct {
	Ref          string `json:"ref,omitempty"`
	Destination  string `json:"destination,omitempty"`
	Filter       Filter `json:"filter,omitempty"`
	InvertFilter bool   `json:"invert_filter,omitempty"`
}

// Filter is a filter structure. The filter specifies files in changesets that
// should not trigger a pipeline automatically.
// See: https://api.gocd.org/current/#the-filter-object
type Filter struct {
	Ignore string `json:"ignore"`
}
