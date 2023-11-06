package dependencies

type ReportModel struct {
	Project       Project
	ParentProject *Project // may not exist
	Dependencies  []Dependency
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
		Project:      project,
		Dependencies: make([]Dependency, 0),
	}
}

func (r *ReportModel) SetParentProject(project Project) {
	r.ParentProject = &project
}

func (r *ReportModel) AddDependency(dependency Dependency) {
	r.Dependencies = append(r.Dependencies, dependency)
}
