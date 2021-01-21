package registry

import (
	"time"
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
