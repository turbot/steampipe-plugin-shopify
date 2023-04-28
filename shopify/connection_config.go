package shopify

import (
	"context"
	"errors"
	"os"

	goshopify "github.com/bold-commerce/go-shopify/v3"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type shopifyConfig struct {
	Token    *string `cty:"token"`
	ShopName *string `cty:"shop_name"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"shop_name": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &shopifyConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) shopifyConfig {
	if connection == nil || connection.Config == nil {
		return shopifyConfig{}
	}
	config, _ := connection.Config.(shopifyConfig)
	return config
}

func connect(_ context.Context, d *plugin.QueryData) (*goshopify.Client, error) {
	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "shopify"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*goshopify.Client), nil
	}

	// Default to env var settings
	token := os.Getenv("SHOPIFY_API_TOKEN")
	shopName := os.Getenv("SHOPIFY_SHOP_NAME")

	// Prefer config settings
	shopifyConfig := GetConfig(d.Connection)
	if shopifyConfig.Token != nil {
		token = *shopifyConfig.Token
	}
	if shopifyConfig.ShopName != nil {
		shopName = *shopifyConfig.ShopName
	}

	// Error if the minimum config is not set
	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}
	if shopName == "" {
		return nil, errors.New("'shop_name' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// Currently we don't need to specify the API Key/API Secret Key, to create the
	// conn, just the API token is enough to fetch the data(for the initial tables).
	// TODO: Look into if we need to use keys/secret keys in the future.
	conn := goshopify.NewClient(goshopify.App{}, shopName, token)

	return conn, nil
}
