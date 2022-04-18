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
	Packages      []Package `json:"packages,omitempty"`
	NextPageToken string    `json:"nextPageToken"`
}

//ListPackages lists all the packages in the repository
func ListPackages(projectName, location, repository string) ([]Package, error) {
	var allPackages []Package

	packagesUrl := packagesURL(projectName, location, repository)
	url := packagesUrl
	for {
		var packages Packages

		resp, err := http.Get(packagesUrl)
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

		allPackages = append(allPackages, insertName(packages.Packages)...)
		if packages.NextPageToken == "" {
			break
		}
		packagesUrl = url + "?pageToken=" + packages.NextPageToken
	}

	return allPackages, nil
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
	var allVersions []Version

	versionsUrl := defaultBaseURL + p.Name + "/versions"
	url := versionsUrl
	for {
		var versions Versions

		resp, err := http.Get(versionsUrl)
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

		allVersions = append(allVersions, insertVersionDigest(versions.Versions)...)
		if versions.NextPageToken == "" {
			break
		}
		versionsUrl = url + "?pageToken=" + versions.NextPageToken

	}
	return allVersions, nil

}

// Tags returns all the tags on this package
func (p *Package) Tags() ([]Tag, error) {
	var allTags []Tag

	tagsUrl := defaultBaseURL + p.Name + "/tags"
	url := tagsUrl
	for {
		var tags Tags

		resp, err := http.Get(tagsUrl)
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

		allTags = append(allTags, insertTagsTagName(tags.Tags)...)
		if tags.NextPageToken == "" {
			break
		}
		tagsUrl = url + "?pageToken=" + tags.NextPageToken

	}
	return allTags, nil
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
