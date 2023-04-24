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
	token *string `cty:"token"`
	shopName *string `cty:"shop_name"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"shopName": {
		Type: schema.TypeString,
	}
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

func connect(_ context.Context, d *plugin.QueryData) (*shopifyapi.Client, error) {
	shopifyConfig := GetConfig(d.Connection)

	// Default to env var settings
	token := os.Getenv("SHOPIFY_API_TOKEN")
	shopName := os.Getenv("SHOPIFY_SHOP_NAME")

	// Prefer config settings
	if shopifyConfig.token != nil {
		token = *shopifyConfig.token
	}
	if shopifyConfig.shopName != nil {
		token = *shopifyConfig.shopName
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if shopName == "" {
		return nil, errors.New("'shop_name' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	client := goshopify.NewClient(goshopify.App{}, shopifyConfig.shopName, token)

	return client, nil
}
