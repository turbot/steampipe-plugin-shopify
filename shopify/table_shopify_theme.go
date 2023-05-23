package shopify

import (
	"context"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShopifyTheme(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shopify_theme",
		Description: "Shopify themes are pre-designed website templates that allow you to easily customize the look and feel of your online store without requiring advanced web development skills.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getTheme,
		},
		List: &plugin.ListConfig{
			Hydrate: listThemes,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the theme.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the theme.",
			},
			{
				Name:        "previewable",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the website is previewable.",
			},
			{
				Name:        "processing",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the website is processing.",
			},
			{
				Name:        "role",
				Type:        proto.ColumnType_STRING,
				Description: "The role of the theme.",
			},
			{
				Name:        "theme_store_id",
				Type:        proto.ColumnType_INT,
				Description: "The store ID of the theme.",
				Transform:   transform.FromField("ThemeStoreID"),
			},
			{
				Name:        "admin_graphql_api_id",
				Type:        proto.ColumnType_STRING,
				Description: "The admin graphql API ID.",
				Transform:   transform.FromField("AdminGraphQLApiID"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The create time of the theme.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The update time of the theme.",
			},			
			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: "Title of the resource.",
				Transform:   transform.FromField("name"),
			},
		},
	}
}

func listThemes(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listThemes", "connection_error", err)
		return nil, err
	}

	themes, err := conn.Theme.List(nil)
	if err != nil {
		plugin.Logger(ctx).Error("listThemes", "list_api_error", err)
		return nil, err
	}
	for _, theme := range themes {
		d.StreamListItem(ctx, theme)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

func getTheme(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := h.Item.(goshopify.Theme).ID

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getTheme", "connection_error", err)
		return nil, err
	}

	theme, err := conn.Theme.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("getTheme", "get_api_error", err)
		return nil, err
	}

	return theme, nil
}
