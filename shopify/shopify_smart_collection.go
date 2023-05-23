package shopify

import (
	"context"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShopifySmartCollection(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shopify_smart_collection",
		Description: "Shopify smart collection is a grouping of products defined by rules that are set by the merchant. Shopify automatically changes the contents of a smart collection based on the rules.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSmartCollection,
		},
		List: &plugin.ListConfig{
			Hydrate: listSmartCollections,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the smart collection.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: "Title of the smart collection.",
			},
			{
				Name:        "handle",
				Type:        proto.ColumnType_STRING,
				Description: "The handle of the smart collection",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the smart collection was last updated.",
			},
			{
				Name:        "body_html",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("BodyHTML"),
				Description: "Body HTML of the smart collection.",
			},
			{
				Name:        "sort_order",
				Type:        proto.ColumnType_STRING,
				Description: "A specific sort order for the smart collection.",
			},
			{
				Name:        "template_suffix",
				Type:        proto.ColumnType_STRING,
				Description: "The template suffix of the smart collection.",
			},
			{
				Name:        "image",
				Type:        proto.ColumnType_JSON,
				Description: "The image of the smart collection.",
			},
			{
				Name:        "published",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the smart collection is published or not.",
			},
			{
				Name:        "published_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the smart collection was published.",
			},
			{
				Name:        "published_scope",
				Type:        proto.ColumnType_STRING,
				Description: "Scope of the published smart collection.",
			},
			{
				Name:        "rules",
				Type:        proto.ColumnType_JSON,
				Description: "Smart collection rules.",
			},
			{
				Name:        "disjunctive",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the smart collection is disjunctive or not.",
			},
			{
				Name:        "metafields",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listSmartCollectionMetafields,
				Transform:   transform.FromValue(),
				Description: "Smart collection metafields.",
			},
		},
	}
}

func listSmartCollections(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSmartCollection", "connection_error", err)
		return nil, err
	}

	smartCols, err := conn.SmartCollection.List(nil)
	if err != nil {
		plugin.Logger(ctx).Error("listSmartCollection", "list_api_error", err)
		return nil, err
	}

	for _, smart := range smartCols {
		d.StreamListItem(ctx, smart)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

// HYDRATE FUNCTIONS

func getSmartCollection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getSmartCollection", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetInt64Value()

	// check if the id is empty
	if id == 0 {
		return nil, nil
	}
	result, err := conn.SmartCollection.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("getSmartCollection", "api_error", err)
		return nil, err
	}

	return *result, nil
}

func listSmartCollectionMetafields(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(goshopify.SmartCollection).ID

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSmartCollectionMetafields", "connection_error", err)
		return nil, err
	}

	meta, err := conn.SmartCollection.ListMetafields(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("listSmartCollectionMetafields", "list_api_error", err)
		return nil, err
	}

	return meta, nil
}
