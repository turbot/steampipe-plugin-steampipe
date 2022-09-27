package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-steampipe/steampipe"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: steampipe.Plugin})
}
