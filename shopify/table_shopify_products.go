package shopify

import (
	"context"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShopifyProduct(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shopify_product",
		Description: "Shopify Products are the goods, digital downloads, services, and gift cards that are sold in Shopify.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProducts,
		},
		List: &plugin.ListConfig{
			Hydrate: listProducts,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the shopify product.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: "The title of the product.",
			},
			{
				Name:        "body_html",
				Type:        proto.ColumnType_STRING,
				Description: "the BodyHTML of the product.",
				Transform:   transform.FromField("BodyHTML"),
			},
			{
				Name:        "vendor",
				Type:        proto.ColumnType_STRING,
				Description: "The vendor of the shopify products.",
			},
			{
				Name:        "product_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the products sold.",
			},
			{
				Name:        "handle",
				Type:        proto.ColumnType_STRING,
				Description: "The product handle.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The product creation date.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time product was updated.",
			},
			{
				Name:        "published_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The product publish time.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_STRING,
				Description: "The product tags.",
			},
			{
				Name:        "options",
				Type:        proto.ColumnType_JSON,
				Description: "The product options.",
			},
			{
				Name:        "variants",
				Type:        proto.ColumnType_JSON,
				Description: "The product variants.",
			},
			{
				Name:        "images",
				Type:        proto.ColumnType_JSON,
				Description: "The product images.",
			},
		},
	}
}

func listProducts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listProduct", "connection_error", err)
		return nil, err
	}

	options := goshopify.ListOptions{
		Limit: 100000,
	}

	for {
		products, paginator, err := conn.Product.ListWithPagination(options)
		if err != nil {
			plugin.Logger(ctx).Error("listProductsError", "list_api_error", err)
			return nil, err
		}

		for _, product := range products {
			d.StreamListItem(ctx, product)

			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if paginator.NextPageOptions == nil {
			return nil, nil
		}
		options.PageInfo = paginator.NextPageOptions.PageInfo
	}
}

func getProducts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getProduct", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetInt64Value()

	// check if the id is 0
	if id == 0 {
		return nil, nil
	}
	result, err := conn.Product.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("getProduct", "api_error", err)
		return nil, err
	}

	return result, nil
}
