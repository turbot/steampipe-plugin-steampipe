package steampipe

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-steampipe"

// Plugin creates this (steamipee) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"steampipe_registry_plugin":         tableSteampipeRegistryPlugin(ctx),
			"steampipe_registry_plugin_version": tableSteampipeRegistryPluginVersion(ctx),
		},
	}

	return p
}
