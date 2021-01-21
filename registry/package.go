package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://artifactregistry.googleapis.com/v1beta2/"
	userAgent      = "go-steampipe"
)

// Package represents a google artifact repository package
type Package struct {
	ImageName  string
	Name       string    `json:"name,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
}

// Packages respresents the object returned from the packages endpoint
type Packages struct {
	Packages []Package `json:"packages,omitempty"`
}

//ListPackages lists all the packages in the repository
func ListPackages(projectName, location, repository string) ([]Package, error) {
	var packages Packages

	url := packagesURL(projectName, location, repository)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("List packages failed with response code %s", strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &packages); err != nil {
		return nil, err
	}

	p := insertName(packages.Packages)

	return p, nil
}

func packagesURL(projectName, location, repository string) string {
	url := defaultBaseURL +
		"projects/" + projectName +
		"/locations/" + location +
		"/repositories/" + repository +
		"/packages"

	return url
}

// Versions returns all the versions of this package
func (p *Package) Versions() ([]Version, error) {
	var versions Versions
	url := defaultBaseURL + p.Name + "/versions"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("List Versions failed with response code %s", strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &versions); err != nil {
		return nil, err
	}

	v := insertVersionDigest(versions.Versions)

	return v, nil
}

// Tags returns all the tags on this package
func (p *Package) Tags() ([]Tag, error) {
	var tags Tags

	url := defaultBaseURL + p.Name + "/tags"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("List Tags failed with response code %s", strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &tags); err != nil {
		return nil, err
	}

	t := insertTagsTagName(tags.Tags)

	return t, nil
}

func insertVersionDigest(p []Version) []Version {
	var versions []Version
	for _, pkg := range p {
		parts := strings.Split(pkg.Name, "/")

		pkg.Digest = parts[len(parts)-1]
		versions = append(versions, pkg)
	}
	return versions
}

func insertTagsTagName(p []Tag) []Tag {
	var tags []Tag
	for _, pkg := range p {
		parts := strings.Split(pkg.Name, "/")

		pkg.Tag = parts[len(parts)-1]
		tags = append(tags, pkg)
	}
	return tags
}

func insertName(p []Package) []Package {
	var packages []Package
	for _, pkg := range p {
		pkg.ImageName, _ = stripPackageName(pkg.Name)
		packages = append(packages, pkg)
	}
	return packages
}

func stripPackageName(s string) (string, error) {
	dirs := strings.Split(s, "/")
	decodedValue, err := url.QueryUnescape(dirs[len(dirs)-1])
	return decodedValue, err
}
