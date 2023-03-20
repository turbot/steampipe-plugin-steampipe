package steampipe

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-steampipe/registry"
)

func tableSteampipeRegistryPlugin(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "steampipe_registry_plugin",
		Description: "Steampipe Registry Plugins",
		List: &plugin.ListConfig{
			Hydrate: listRegistryPlugins,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the plugin.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_time",
				Description: "The time the plugin was created in the repository.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "update_time",
				Description: "The time the plugin was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
		},
	}
}

func listRegistryPlugins(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugins, err := registry.ListPlugins()
	if err != nil {
		plugin.Logger(ctx).Warn("steampipe_registry_plugin.listRegistryPlugins", "query_error", err)
		return nil, err
	}
	for _, p := range plugins {
		d.StreamListItem(ctx, p)
	}
	return nil, nil
}
