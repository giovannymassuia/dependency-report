package dependencies

type ReportModel struct {
	Project            Project
	ParentProject      *Project // may not exist
	Dependencies       []Dependency
	DependencyManagers []Dependency
	Properties         map[string]string
}

type Project struct {
	Name       string
	GroupId    string
	ArtifactId string
	Version    string
}

type Dependency struct {
	Name       string
	GroupID    string
	ArtifactID string
	Version    string
	Scope      string
}

func NewReport(project Project) *ReportModel {
	return &ReportModel{
		Project:            project,
		Dependencies:       make([]Dependency, 0),
		DependencyManagers: make([]Dependency, 0),
		Properties:         make(map[string]string),
	}
}

func (r *ReportModel) SetParentProject(project Project) {
	r.ParentProject = &project
}

func (r *ReportModel) AddDependency(dependency Dependency) {
	r.Dependencies = append(r.Dependencies, dependency)
}

func (r *ReportModel) AddDependencyManager(dependency Dependency) {
	r.DependencyManagers = append(r.DependencyManagers, dependency)
}

func (r *ReportModel) AddProperty(key, value string) {
	r.Properties[key] = value
}
