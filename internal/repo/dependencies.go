package repo

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"os/exec"
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

func ScanRepository(repoName string, cloneRepo func(name string) error) ([]Project, error) {

	// if repo is not present, try to clone it
	if !GitRepositoryExists(fmt.Sprintf("%s/%s", TempDir, repoName)) {
		err := cloneRepo(repoName)
		if err != nil {
			return nil, err
		}
	}

	pomFiles, err := findPomFiles(fmt.Sprintf("%s/%s", TempDir, repoName))
	if err != nil {
		return nil, err
	}

	var projects []Project
	for _, pomFile := range pomFiles {
		var project *Project
		//if pomFile.HasMvnw {
		//	project, err = processPomFileWithMvnw(pomFile)
		//} else {
		project, err = readPomFile(pomFile)
		//}
		if err != nil {
			return nil, err
		}
		projects = append(projects, *project)
	}

	// remove temp folder
	// TODO: uncomment this
	//err = os.RemoveAll(fmt.Sprintf("%s/%s", repo.TempDir, repoName))
	//if err != nil {
	//	return err
	//}

	return projects, nil
}

func processPomFileWithMvnw(pomFile file) (*Project, error) {
	// execute mvnw dependency:tre
	cmd := exec.Command("./mvnw", "dependency:tree", "-DoutputFile=dependency-tree.txt")
	cmd.Dir = pomFile.ParentFolderPath
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	// read the generated file
	return readDependencyTree(fmt.Sprintf("%s/dependency-tree.txt", pomFile.ParentFolderPath))
}

type file struct {
	Path             string
	ParentFolderPath string
	ParentFolderName string
	HasMvnw          bool
}

func findPomFiles(root string) ([]file, error) {
	var files []file
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == "pom.xml" {

			hasMvnw := false
			_, err := os.Stat(fmt.Sprintf("%s/mvnw", filepath.Dir(path)))
			if err == nil {
				hasMvnw = true
			}

			files = append(files, file{
				Path:             path,
				ParentFolderPath: strings.Replace(path, "/pom.xml", "", 1),
				ParentFolderName: filepath.Base(strings.Replace(path, "/pom.xml", "", 1)),
				HasMvnw:          hasMvnw,
			})
		}

		return nil
	})

	return files, err
}

func readDependencyTree(filePath string) (*Project, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(fileContent), "\n")

	project := &Project{}
	var dependencies []Dependency
	for idx, line := range lines {
		// First line is the project itself
		if idx == 0 {
			project.GroupId = strings.Split(line, ":")[0]
			project.ArtifactId = strings.Split(line, ":")[1]
			project.Version = strings.Split(line, ":")[3]
			continue
		}

		// Parse each line that represents a dependency
		// Line that starts with: "+- groupId:artifactId:version:scope"
		if strings.HasPrefix(line, "+- ") {
			// remove "+- " prefix
			line = strings.Replace(line, "+- ", "", 1)
			parts := strings.Split(line, ":")
			dependency := Dependency{
				GroupId:    parts[0],
				ArtifactId: parts[1],
				Version:    parts[3],
				Scope:      parts[4],
			}
			dependencies = append(dependencies, dependency)
		}
	}

	project.Dependencies.Dependency = dependencies

	return project, nil
}

func readPomFile(pomFilePath file) (*Project, error) {
	xmlFile, err := os.Open(pomFilePath.Path)
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
	for i, managementDep := range project.DependencyManagement.Dependencies.Dependency {
		//found := false
		//for _, dep := range project.Dependencies.Dependency {
		//	if dep.ArtifactId == managementDep.ArtifactId && dep.GroupId == managementDep.GroupId {
		//		found = true
		//		break
		//	}
		//}
		//if !found {
		project.DependencyManagement.Dependencies.Dependency[i].Version = resolveProperties(managementDep.Version, propsMap)
		//	managementDep.Version = resolveProperties(managementDep.Version, propsMap)
		//	project.Dependencies.Dependency = append(project.Dependencies.Dependency, managementDep)
		//}
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
