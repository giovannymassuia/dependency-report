package managers

import (
	"github.com/giovannymassuia/dependency-report/internal/dependencies"
	"github.com/giovannymassuia/dependency-report/test/testutils"
	"github.com/stretchr/testify/assert"
	"path"
	"strings"
	"testing"
)

func TestMaven_Scan(t *testing.T) {
	maven := NewMaven()
	modulePath := testutils.GetModulePath()
	testPath := path.Join(modulePath, "test", "data", "maven_repo")

	result, err := maven.Scan(testPath)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))

	testProject := result[0]

	// project
	assert.Equal(t, "api", testProject.Project.Name)
	assert.Equal(t, "io.giovannymassuia.dependency-report.test", testProject.Project.GroupId)
	assert.Equal(t, "test-api", testProject.Project.ArtifactId)
	assert.Equal(t, "1.0", testProject.Project.Version)

	// parent project
	assert.Equal(t, "", testProject.ParentProject.Name)
	assert.Equal(t, "org.springframework.boot", testProject.ParentProject.GroupId)
	assert.Equal(t, "spring-boot-starter-parent", testProject.ParentProject.ArtifactId)
	assert.Equal(t, "2.3.3.RELEASE", testProject.ParentProject.Version)

	// properties
	assert.Equal(t, 3, len(testProject.Properties))
	assert.Equal(t, "14", testProject.Properties["java.version"])
	assert.Equal(t, "1.11.608", testProject.Properties["aws.sdk.version"])
	assert.Equal(t, "6.4.4", testProject.Properties["flyway.version"])

	// dependencies managers
	assert.Equal(t, 1, len(testProject.DependencyManagers))
	expectedDependencyManager := buildDependency("com.amazonaws:aws-java-sdk-bom:1.11.608:import")
	assert.Equal(t, expectedDependencyManager, testProject.DependencyManagers[0])

	// dependencies
	assert.Equal(t, 7, len(testProject.Dependencies))

	expectedDependencies := make(map[string]dependencies.Dependency)
	addDependency(expectedDependencies, buildDependency("org.springframework.boot:spring-boot-starter-web:2.3.3.RELEASE:compile"))
	addDependency(expectedDependencies, buildDependency("org.postgresql:postgresql:42.2.14:runtime"))
	addDependency(expectedDependencies, buildDependency("org.flywaydb:flyway-core:6.4.4:compile"))
	addDependency(expectedDependencies, buildDependency("com.amazonaws:aws-java-sdk-ses:1.11.608:compile"))
	addDependency(expectedDependencies, buildDependency("io.jsonwebtoken:jjwt:0.9.1:compile"))
	addDependency(expectedDependencies, buildDependency("org.apache.commons:commons-collections4:4.4:compile"))
	addDependency(expectedDependencies, buildDependency("org.springframework.security:spring-security-test:5.3.4.RELEASE:test"))

	for _, dependency := range testProject.Dependencies {
		key := dependency.GroupID + ":" + dependency.ArtifactID
		assert.Equal(t, expectedDependencies[key], dependency)
	}

}

func addDependency(values map[string]dependencies.Dependency, dependency dependencies.Dependency) {
	values[dependency.GroupID+":"+dependency.ArtifactID] = dependency
}

func buildDependency(dependencyString string) dependencies.Dependency {
	parts := strings.Split(dependencyString, ":")
	return dependencies.Dependency{
		Name:       parts[0] + ":" + parts[1],
		GroupID:    parts[0],
		ArtifactID: parts[1],
		Version:    parts[2],
		Scope:      parts[3],
	}
}
