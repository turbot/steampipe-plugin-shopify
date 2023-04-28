package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-shopify/shopify"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: shopify.Plugin})
}
