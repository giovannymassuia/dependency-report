package internal

import "encoding/xml"

type PomFile struct {
	XMLName              xml.Name             `xml:"project"`
	ModelVersion         string               `xml:"modelVersion"`
	Parent               Parent               `xml:"parent"`
	Name                 string               `xml:"name"`
	GroupId              string               `xml:"groupId"`
	ArtifactId           string               `xml:"artifactId"`
	Version              string               `xml:"version"`
	Dependencies         Dependencies         `xml:"dependencies"`
	Properties           Properties           `xml:"properties"`
	DependencyManagement DependencyManagement `xml:"dependencyManagement"`
}

type Properties struct {
	XMLName    xml.Name   `xml:"properties"`
	Properties []Property `xml:",any"` //
}

type Property struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type Parent struct {
	Name         string `xml:"name"`
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
