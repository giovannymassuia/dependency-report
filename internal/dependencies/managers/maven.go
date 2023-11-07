package managers

import (
	"encoding/xml"
	"github.com/giovannymassuia/dependency-report/internal/dependencies"
	"github.com/giovannymassuia/dependency-report/internal/files"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Maven struct{}

func NewMaven() Maven {
	return Maven{}
}

// TODO: consider using mvn help:effective-pom -Doutput=effective-pom.xml

// Scan scans a maven repository and returns a list of dependencies.ReportModel
func (m Maven) Scan(path string) ([]dependencies.ReportModel, error) {
	var result []dependencies.ReportModel

	// find all pom.xml files in path
	pomFilePaths, err := files.FindFiles(path, "pom.xml")
	if err != nil {
		return nil, err
	}

	for _, pomFile := range pomFilePaths {
		parsedPom, err := readPomFile(pomFile.FilePath)
		if err != nil {
			return nil, err
		}

		versions, err := parseVersions(pomFile.FolderPath)
		if err != nil {
			return nil, err
		}

		project := dependencies.NewReport(
			dependencies.Project{
				Name:       parsedPom.Name,
				GroupId:    parsedPom.GroupId,
				ArtifactId: parsedPom.ArtifactId,
				Version:    parsedPom.Version,
			})

		if parsedPom.Parent.GroupId != "" {
			project.SetParentProject(dependencies.Project{
				Name:       parsedPom.Parent.Name,
				GroupId:    parsedPom.Parent.GroupId,
				ArtifactId: parsedPom.Parent.ArtifactId,
				Version:    parsedPom.Parent.Version,
			})
		}

		for _, property := range parsedPom.Properties.Properties {
			project.AddProperty(property.XMLName.Local, property.Value)
		}

		if parsedPom.DependencyManagement.Dependencies.Dependency != nil {
			for _, dependency := range parsedPom.DependencyManagement.Dependencies.Dependency {
				version := dependency.Version
				if strings.HasPrefix(version, "${") {
					key := strings.ReplaceAll(version, "${", "")[:len(version)-3]
					version = project.Properties[key]
				}
				project.AddDependencyManager(dependencies.Dependency{
					Name:       dependency.GroupId + ":" + dependency.ArtifactId,
					GroupID:    dependency.GroupId,
					ArtifactID: dependency.ArtifactId,
					Version:    version,
					Scope:      dependency.Scope,
				})
			}
		}

		if parsedPom.Dependencies.Dependency != nil {
			for _, dependency := range parsedPom.Dependencies.Dependency {
				key := dependency.GroupId + ":" + dependency.ArtifactId
				version := versions[key]
				if version == "" {
					version = dependency.Version
				}
				scope := dependency.Scope
				if scope == "" {
					scope = "compile"
				}
				project.AddDependency(dependencies.Dependency{
					Name:       dependency.GroupId + ":" + dependency.ArtifactId,
					GroupID:    dependency.GroupId,
					ArtifactID: dependency.ArtifactId,
					Version:    version,
					Scope:      scope,
				})
			}
		}

		result = append(result, *project)
	}

	return result, nil
}

func parseVersions(rootPath string) (map[string]string, error) {

	// check if mvnw exists
	mvnwExists := files.IsExists(rootPath, "mvnw")
	if !mvnwExists {
		return nil, nil
	}

	// generate dependency-tree.txt
	cmd := exec.Command("./mvnw", "dependency:tree", "-DoutputFile=dependency-tree.txt")
	cmd.Dir = rootPath
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// read dependency-tree.txt
	dependencyTreeFilePath := path.Join(rootPath, "dependency-tree.txt")
	dependencyTreeFile, err := os.ReadFile(dependencyTreeFilePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(dependencyTreeFile), "\n")

	versions := make(map[string]string)

	for _, line := range lines {
		if strings.HasPrefix(line, "+- ") || strings.HasPrefix(line, "\\- ") {
			// remove "+- " or "\-" prefix
			line = strings.Replace(line, "+- ", "", 1)
			line = strings.Replace(line, "\\- ", "", 1)
			parts := strings.Split(line, ":")

			// +- org.springframework.boot:spring-boot-starter-web:jar:2.3.3:compile
			versions[parts[0]+":"+parts[1]] = parts[3]
		}
	}

	return versions, nil
}

func readPomFile(pomFilePath string) (*PomFile, error) {
	xmlFile, err := os.Open(pomFilePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	bytes, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var pomFile PomFile
	err = xml.Unmarshal(bytes, &pomFile)
	if err != nil {
		return nil, err
	}

	return &pomFile, nil
}
