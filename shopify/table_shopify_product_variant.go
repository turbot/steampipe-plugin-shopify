package shopify

import (
	"context"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShopifyProductVariant(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shopify_product_variant",
		Description: "Shopify Products variant can be added to a Product resource to represent one version of a product with several options. The Product resource will have a variant for every possible combination of its options.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getProductVariant,
		},
		List: &plugin.ListConfig{
			ParentHydrate: listProduct,
			Hydrate:       listProductVariant,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the shopify product.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "product_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the product.",
				Transform:   transform.FromField("ProductID"),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: "The title of the product.",
			},
			{
				Name:        "sku",
				Type:        proto.ColumnType_STRING,
				Description: "The SKU of the product.",
			},
			{
				Name:        "position",
				Type:        proto.ColumnType_INT,
				Description: "The position of the product.",
				Transform:   transform.FromField("Position"),
			},
			{
				Name:        "grams",
				Type:        proto.ColumnType_INT,
				Description: "The weight of the product in grams.",
			},
			{
				Name:        "inventory_policy",
				Type:        proto.ColumnType_STRING,
				Description: "The inventory policy of the product.",
			},
			{
				Name:        "price",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The price of the product.",
				Transform:   transform.FromField("Price").Transform(transform.ToString).Transform(transform.ToDouble),
			},
			// TODO: Transform not working for this field
			// {
			// 	Name:        "compare_at_price",
			// 	Type:        proto.ColumnType_DOUBLE,
			// 	Description: "The compare at price of the product.",
			// 	// Transform:   transform.FromField("CompareAtPrice."),
			// 	Transform: transform.FromField("CompareAtPrice").Transform(transform.ToString).Transform(transform.ToDouble),
			// },
			{
				Name:        "fulfillment_service",
				Type:        proto.ColumnType_STRING,
				Description: "The fulfillment service of the product.",
			},
			{
				Name:        "inventory_management",
				Type:        proto.ColumnType_STRING,
				Description: "The inventory management of the product.",
			},
			{
				Name:        "inventory_item_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the inventory item.",
			},
			{
				Name:        "option1",
				Type:        proto.ColumnType_STRING,
				Description: "The first product option.",
			},
			{
				Name:        "option2",
				Type:        proto.ColumnType_STRING,
				Description: "The second product option.",
			},
			{
				Name:        "option3",
				Type:        proto.ColumnType_STRING,
				Description: "The third product option.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time when the product was created.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time when the product was last updated.",
			},
			{
				Name:        "taxable",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether or not the product is taxable.",
			},
			{
				Name:        "tax_code",
				Type:        proto.ColumnType_STRING,
				Description: "The tax code of the product.",
			},
			{
				Name:        "barcode",
				Type:        proto.ColumnType_STRING,
				Description: "The barcode of the product.",
			},
			{
				Name:        "image_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the product image.",
				Transform:   transform.FromField("ImageID"),
			},
			{
				Name:        "inventory_quantity",
				Type:        proto.ColumnType_INT,
				Description: "The number of inventory items available for the product.",
			},
			{
				Name:        "weight",
				Type:        proto.ColumnType_STRING,
				Description: "The weight of the product.",
			},
			{
				Name:        "weight_unit",
				Type:        proto.ColumnType_STRING,
				Description: "The unit of measurement for the product weight.",
			},
			{
				Name:        "old_inventory_quantity",
				Type:        proto.ColumnType_INT,
				Description: "The old inventory quantity of the product.",
			},
			{
				Name:        "requires_shipping",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the product requires shipping.",
			},
			{
				Name:        "admin_graphql_api_id",
				Type:        proto.ColumnType_STRING,
				Description: "The admin graphql API ID of the product.",
				Transform:   transform.FromField("AdminGraphqlAPIID"),
			},
		},
	}
}

func listProductVariant(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listProductVariant", "connection_error", err)
		return nil, err
	}
	id := h.Item.(goshopify.Product).ID

	variants, err := conn.Variant.List(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("listProductVariantError", "list_api_error", err)
		return nil, err
	}

	for _, variant := range variants {
		d.StreamListItem(ctx, variant)

		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

func getProductVariant(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getProductVariant", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetInt64Value()

	// check if the id is 0
	if id == 0 {
		return nil, nil
	}
	result, err := conn.Variant.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("getProductVariant", "api_error", err)
		return nil, err
	}

	return result, nil
}
