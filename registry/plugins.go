package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const (
	defaultRepositoryName     = "steampipe"
	defaultRepositoryLocation = "us"
	defaultPluginPackageName  = "plugins"
)

// Plugin respresents Plugin image from the repository
type Plugin struct {
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
	Package    Package
}

// PluginVersion respresents Plugin image image version
type PluginVersion struct {
	ImageName  string
	Digest     string
	CreateTime time.Time
	UpdateTime time.Time
	Tags       []string
}

// Versions gets a list of versions (and related tags) of a plugin
func (p *Plugin) Versions() ([]PluginVersion, error) {
	var versions []PluginVersion

	versionsRaw, err := p.Package.Versions()
	if err != nil {
		return nil, err
	}
	tagsRaw, err := p.Package.Tags()
	if err != nil {
		return nil, err
	}

	for _, v := range versionsRaw {
		version := PluginVersion{
			ImageName:  p.Name,
			Digest:     v.Digest,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			Tags:       tagsForVersion(tagsRaw, v.Name),
		}
		versions = append(versions, version)
	}
	return versions, nil
}

// ListPlugins gets a list of plugins from the repository
func ListPlugins() ([]Plugin, error) {
	var plugins []Plugin
	packages, err := ListPackages(defaultRepositoryName, defaultRepositoryLocation, defaultPluginPackageName)
	if err != nil {
		return nil, err
	}

	for _, pkg := range packages {
		plugin := Plugin{
			Package:    pkg,
			Name:       pkg.ImageName,
			CreateTime: pkg.CreateTime,
			UpdateTime: pkg.UpdateTime,
		}
		plugins = append(plugins, plugin)
	}

	return plugins, err
}

func tagsForVersion(tagsRaw []Tag, version string) []string {
	var matches []string
	for _, t := range tagsRaw {
		if t.Version == version {
			matches = append(matches, t.Tag)
		}
	}

	return matches
}

func (p *PluginVersion) GetManifest(ctx context.Context) (interface{}, error) {
	log := plugin.Logger(ctx)
	log.Trace("GetManifest")

	// Construct the manifest URL
	manifestURL := "https://ghcr.io/v2/turbot/" + p.ImageName + "/manifests/" + p.Digest
	log.Trace("GetManifestURL", manifestURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", manifestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add the Authorization header with the GitHub token
	token := os.Getenv("GITHUB_TOKEN") // Make sure GITHUB_TOKEN is set in your environment
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable not set")
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// Set the Accept header to request the Docker manifest
	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	// Make the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Get Manifest failed with error: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get Manifest failed with response code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read manifest: %v", err)
	}

	// Parse the response body as JSON
	var manifest map[string]interface{}
	if err := json.Unmarshal(body, &manifest); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return manifest, nil
}
