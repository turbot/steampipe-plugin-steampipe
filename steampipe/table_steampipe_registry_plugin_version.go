package steampipe

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"github.com/turbot/steampipe-plugin-steampipe/registry"
)

func tableSteampipeRegistryPluginVersion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "steampipe_registry_plugin_version",
		Description: "Steampipe Registry Plugin Version",

		List: &plugin.ListConfig{
			ParentHydrate: listRegistryPlugins,
			Hydrate:       listRegistryPluginVersions,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the plugin.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ImageName"),
			},
			{
				Name:        "digest",
				Description: "The digest (shasum) of the plugin version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "tags",
				Description: "The tags on the plugin version.",
				Type:        proto.ColumnType_JSON,
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
			{
				Name:        "image_annotations",
				Description: "The annotations from the OCI image manifest.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getImageManifest,
				Transform:   transform.FromField("annotations"),
			},
			{
				Name:        "image_layers",
				Description: "The layers from the OCI image manifest.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getImageManifest,
				Transform:   transform.FromField("layers"),
			},
		},
	}
}

func listRegistryPluginVersions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	// read the parent item from the hydrate data
	plugin := h.Item.(registry.Plugin)

	versions, err := plugin.Versions()
	if err != nil {
		return nil, err
	}

	for _, p := range versions {
		d.StreamLeafListItem(ctx, p)
	}
	return nil, nil

}

//// HYDRATE FUNCTIONS
func getImageManifest(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	p := h.Item.(registry.PluginVersion)
	manifest, err := p.GetManifest(ctx)
	if err != nil {
		return nil, err
	}

	return manifest, err
}
