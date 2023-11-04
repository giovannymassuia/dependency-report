package dependencies

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Project struct {
	XMLName              xml.Name             `xml:"project"`
	ModelVersion         string               `xml:"modelVersion"`
	Parent               Parent               `xml:"parent"`
	GroupId              string               `xml:"groupId"`
	ArtifactId           string               `xml:"artifactId"`
	Version              string               `xml:"version"`
	Dependencies         Dependencies         `xml:"dependencies"`
	Properties           Properties           `xml:"properties"`
	DependencyManagement DependencyManagement `xml:"dependencyManagement"`
}

type Properties struct {
	XMLName    xml.Name   `xml:"properties"`
	Properties []Property `xml:",any"`
}

type Property struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type Parent struct {
	GroupId      string `xml:"groupId"`
	ArtifactId   string `xml:"artifactId"`
	Version      string `xml:"version"`
	RelativePath string `xml:"relativePath"`
}

type Dependencies struct {
	Dependency []Dependency `xml:"dependency"`
}

type DependencyManagement struct {
	Dependencies Dependencies `xml:"dependencies"`
}

type Dependency struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Scope      string `xml:"scope"`
}

func FindPomFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == "pom.xml" {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

func ReadDependencyTree(filePath string) ([]Dependency, error) {
	fmt.Println(filePath)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(fileContent), "\n")
	dependencies := []Dependency{}

	fmt.Printf("Lines: %d\n", len(lines))

	for _, line := range lines {
		// Parse each line that represents a dependency
		// Line that starts with: "+- groupId:artifactId:version:scope"
		if strings.HasPrefix(line, "+- ") {
			fmt.Println(line)
		}
	}

	return dependencies, nil
}

func ReadPomFile(pomFilePath string) (*Project, error) {
	xmlFile, err := os.Open(pomFilePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	bytes, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var project Project
	err = xml.Unmarshal(bytes, &project)
	if err != nil {
		return nil, err
	}

	propsMap := make(map[string]string)
	for _, p := range project.Properties.Properties {
		propsMap[p.XMLName.Local] = p.Value
	}

	// Resolve properties within dependencies
	for i, dep := range project.Dependencies.Dependency {
		project.Dependencies.Dependency[i].Version = resolveProperties(dep.Version, propsMap)
	}

	// Resolve dependency management dependencies if they are not already in the main dependencies
	for _, managementDep := range project.DependencyManagement.Dependencies.Dependency {
		found := false
		for _, dep := range project.Dependencies.Dependency {
			if dep.ArtifactId == managementDep.ArtifactId && dep.GroupId == managementDep.GroupId {
				found = true
				break
			}
		}
		if !found {
			managementDep.Version = resolveProperties(managementDep.Version, propsMap)
			project.Dependencies.Dependency = append(project.Dependencies.Dependency, managementDep)
		}
	}

	return &project, nil
}

func resolveProperties(value string, properties map[string]string) string {
	re := regexp.MustCompile(`\$\{(.+?)\}`)
	return re.ReplaceAllStringFunc(value, func(match string) string {
		propName := re.FindStringSubmatch(match)[1]
		if propValue, ok := properties[propName]; ok {
			return propValue
		}
		return match
	})
}
