package dependencies

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReport(t *testing.T) {
	report := NewReport(Project{
		Name:       "name",
		GroupId:    "groupId",
		ArtifactId: "artifactId",
		Version:    "version",
	})

	assert.Equal(t, "name", report.Project.Name)
	assert.Equal(t, "groupId", report.Project.GroupId)
	assert.Equal(t, "artifactId", report.Project.ArtifactId)
	assert.Equal(t, "version", report.Project.Version)
	assert.Equal(t, 0, len(report.Dependencies))
}

func TestSetParentProject(t *testing.T) {
	report := NewReport(Project{
		Name:       "name",
		GroupId:    "groupId",
		ArtifactId: "artifactId",
		Version:    "version",
	})

	report.SetParentProject(Project{
		Name:       "ParentName",
		GroupId:    "ParentGroupId",
		ArtifactId: "ParentArtifactId",
		Version:    "ParentVersion",
	})

	assert.Equal(t, "ParentName", report.ParentProject.Name)
	assert.Equal(t, "ParentGroupId", report.ParentProject.GroupId)
	assert.Equal(t, "ParentArtifactId", report.ParentProject.ArtifactId)
	assert.Equal(t, "ParentVersion", report.ParentProject.Version)
}

func TestAddDependency(t *testing.T) {
	report := NewReport(Project{
		Name:       "name",
		GroupId:    "groupId",
		ArtifactId: "artifactId",
		Version:    "version",
	})

	report.AddDependency(Dependency{
		Name:       "name",
		GroupID:    "group",
		ArtifactID: "artifact",
		Version:    "version",
		Scope:      "scope",
	})

	assert.Equal(t, 1, len(report.Dependencies))
	assert.Equal(t, "name", report.Dependencies[0].Name)
	assert.Equal(t, "group", report.Dependencies[0].GroupID)
	assert.Equal(t, "artifact", report.Dependencies[0].ArtifactID)
	assert.Equal(t, "version", report.Dependencies[0].Version)
	assert.Equal(t, "scope", report.Dependencies[0].Scope)
}
