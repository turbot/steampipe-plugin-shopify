package shopify

import (
	"context"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShopifyCustomCollection(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shopify_custom_collection",
		Description: "Shopify custom collection is a group of products that a merchant can create to make their store easier to browse.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getCustomCollection,
		},
		List: &plugin.ListConfig{
			Hydrate: listCustomCollections,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the custom collection.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "custom_collection_title",
				Type:        proto.ColumnType_STRING,
				Description: "Title of the custom collection.",
				Transform:   transform.FromField("Title"),
			},
			{
				Name:        "handle",
				Type:        proto.ColumnType_STRING,
				Description: "The handle of the custom collection.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the custom collection was last updated.",
			},
			{
				Name:        "body_html",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("BodyHTML"),
				Description: "Body HTML of the custom collection.",
			},
			{
				Name:        "sort_order",
				Type:        proto.ColumnType_STRING,
				Description: "A specific sort order for the custom collection.",
			},
			{
				Name:        "template_suffix",
				Type:        proto.ColumnType_STRING,
				Description: "The template suffix of the custom collection.",
			},
			{
				Name:        "image",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
				Description: "The image of the custom collection.",
			},
			{
				Name:        "published",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the custom collection is published or not.",
			},
			{
				Name:        "published_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the custom collection was published.",
			},
			{
				Name:        "published_scope",
				Type:        proto.ColumnType_STRING,
				Description: "Scope of the published custom collection.",
			},
			{
				Name:        "metafields",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listCustomCollectionMetafields,
				Transform:   transform.FromValue(),
				Description: "Custom collection metafields.",
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: "Title of the resource.",
				Transform:   transform.FromField("Title"),
			},
		}),
	}
}

func listCustomCollections(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_custom_collection.listCustomCollections", "connection_error", err)
		return nil, err
	}

	customCols, err := conn.CustomCollection.List(nil)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_custom_collection.listCustomCollections", "api_error", err)
		return nil, err
	}

	for _, custom := range customCols {
		d.StreamListItem(ctx, custom)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

// HYDRATE FUNCTIONS

func getCustomCollection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	id := d.EqualsQuals["id"].GetInt64Value()

	// check if the id is empty
	if id == 0 {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_custom_collection.getCustomCollection", "connection_error", err)
		return nil, err
	}

	result, err := conn.CustomCollection.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_custom_collection.getCustomCollection", "api_error", err)
		return nil, err
	}

	return *result, nil
}

func listCustomCollectionMetafields(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	id := h.Item.(goshopify.CustomCollection).ID

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_custom_collection.listCustomCollectionMetafields", "connection_error", err)
		return nil, err
	}

	meta, err := conn.CustomCollection.ListMetafields(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_custom_collection.listCustomCollectionMetafields", "api_error", err)
		return nil, err
	}

	return meta, nil
}
