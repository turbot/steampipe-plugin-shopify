package shopify

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Plugin creates this (shopify) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-shopify",
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"Not Found"}),
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"shopify_smart_collection":   tableShopifySmartCollection(ctx),
			"shopify_order":              tableShopifyOrder(ctx),
			"shopify_custom_collection":  tableShopifyCustomCollection(ctx),
			"shopify_customer":           tableShopifyCustomer(ctx),
			"shopify_product":            tableShopifyProduct(ctx),
			"shopify_product_variant":    tableShopifyProductVariant(ctx),
			"shopify_collection_product": tableShopifyCollectionProduct(ctx),
			"shopify_theme":              tableShopifyTheme(ctx),
			"shopify_draft_order":        tableShopifyDraftOrder(ctx),
		},
	}
	return p
}
