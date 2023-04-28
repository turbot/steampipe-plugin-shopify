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
			Hydrate:    getProduct,
		},
		List: &plugin.ListConfig{
			Hydrate: listProduct,
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
				Description: "The BodyHTML of the product.",
				Transform:   transform.FromField("BodyHTML"),
			},
			{
				Name:        "vendor",
				Type:        proto.ColumnType_STRING,
				Description: "The vendor of the shopify product.",
			},
			{
				Name:        "product_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the product.",
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
				Description: "Time when the product was last updated.",
			},
			{
				Name:        "published_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the product was published.",
			},
			{
				Name:        "published_scope",
				Type:        proto.ColumnType_STRING,
				Description: "The product publish scope.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_STRING,
				Description: "The product tags.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The product status.",
			},
			{
				Name:        "options",
				Type:        proto.ColumnType_JSON,
				Description: "The product options.",
			},
			{
				Name:        "image",
				Type:        proto.ColumnType_JSON,
				Description: "The product image.",
			},
			{
				Name:        "images",
				Type:        proto.ColumnType_JSON,
				Description: "List of images associated with the product.",
			},
			{
				Name:        "template_suffix",
				Type:        proto.ColumnType_STRING,
				Description: "The product template suffix.",
			},
			{
				Name:        "metafields_global_title_tag",
				Type:        proto.ColumnType_STRING,
				Description: "The product metafields global title tag.",
			},
			{
				Name:        "metafields_global_description_tag",
				Type:        proto.ColumnType_STRING,
				Description: "The product metafields global description tag.",
			},
			{
				Name:        "metafields",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listProductMetafields,
				Transform:   transform.FromValue(),
				Description: "The product metafields.",
			},
			{
				Name:        "admin_graphql_api_id",
				Type:        proto.ColumnType_STRING,
				Description: "The admin graphql api id.",
				Transform:   transform.FromField("AdminGraphqlAPIID"),
			},
		},
	}
}

func listProduct(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listProduct", "connection_error", err)
		return nil, err
	}

	// max limit defined by the api is 250
	options := goshopify.ListOptions{}

	// set the limit if a lower limit is passed in query context
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < 250 {
			options.Limit = int(*limit)
		}
	}

	for {
		products, paginator, err := conn.Product.ListWithPagination(options)
		if err != nil {
			plugin.Logger(ctx).Error("listProductError", "list_api_error", err)
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

func getProduct(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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

func listProductMetafields(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(goshopify.Product).ID

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listProductMetafields", "connection_error", err)
		return nil, err
	}

	meta, err := conn.Product.ListMetafields(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("listProductMetafields", "list_api_error", err)
		return nil, err
	}

	return meta, nil
}
