package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

	manifest_url := "https://" + defaultRepositoryLocation + "-docker.pkg.dev/v2/" + defaultRepositoryName + "/" +
		defaultPluginPackageName + "/" + p.ImageName + "/manifests/" + p.Digest

	client := &http.Client{}
	req, _ := http.NewRequest("GET", manifest_url, nil)
	req.Header.Set("Accept", "*/*")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Get Manifest failed with error %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Get Manifest failed with response code %s", strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read manifest - %v", err)
	}

	var anyJson map[string]interface{}
	if err := json.Unmarshal(body, &anyJson); err != nil {
		return nil, err
	}

	return anyJson, nil
}
