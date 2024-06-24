package shopify

import (
	"context"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "shop_name",
			Description: "The name of the shop.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getShopName,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getShopNameMemoized = plugin.HydrateFunc(getShopNameUncached).Memoize(memoize.WithCacheKeyFunction(getShopNameCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getShopName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getShopNameMemoized(ctx, d, h)
}

// Build a cache key for the call to getShopNameCacheKey.
func getShopNameCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getShopName"
	return key, nil
}

func getShopNameUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	shopName := os.Getenv("SHOPIFY_SHOP_NAME")
	shopifyConfig := GetConfig(d.Connection)
	if shopifyConfig.ShopName != nil {
		shopName = *shopifyConfig.ShopName
	}

	return shopName, nil
}
