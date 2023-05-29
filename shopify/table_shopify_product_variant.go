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
			ParentHydrate: listProducts,
			Hydrate:       listProductVariants,
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
				Description: "The ID of the product variant.",
				Transform:   transform.FromField("ProductID"),
			},
			{
				Name:        "product_title",
				Type:        proto.ColumnType_STRING,
				Description: "The title of the product.",
				Transform:   transform.FromField("Title"),
			},
			{
				Name:        "sku",
				Type:        proto.ColumnType_STRING,
				Description: "The Stock Keeping Unit (SKU) for the product variant.",
			},
			{
				Name:        "position",
				Type:        proto.ColumnType_INT,
				Description: "The position of the variant in the list of variants for the product.",
				Transform:   transform.FromField("Position"),
			},
			{
				Name:        "grams",
				Type:        proto.ColumnType_INT,
				Description: "The weight of the product variant in grams.",
			},
			{
				Name:        "inventory_policy",
				Type:        proto.ColumnType_STRING,
				Description: "The inventory policy of the product.",
			},
			{
				Name:        "price",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The price of the product variant.",
				Transform:   transform.FromField("Price").Transform(convertPrice),
			},
			// TODO: Transform not working for this field
			{
				Name:        "compare_at_price",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The compare at price of the product variant.",
				Transform:   transform.FromField("CompareAtPrice").Transform(convertPrice),
			},
			{
				Name:        "fulfillment_service",
				Type:        proto.ColumnType_STRING,
				Description: "The fulfillment service of the product variant.",
			},
			{
				Name:        "inventory_management",
				Type:        proto.ColumnType_STRING,
				Description: "The inventory management of the product variant.",
			},
			{
				Name:        "inventory_item_id",
				Type:        proto.ColumnType_INT,
				Description: "The inventory policy for the product variant.",
			},
			{
				Name:        "option1",
				Type:        proto.ColumnType_STRING,
				Description: "The first option for the variant.",
			},
			{
				Name:        "option2",
				Type:        proto.ColumnType_STRING,
				Description: "The second option for the variant.",
			},
			{
				Name:        "option3",
				Type:        proto.ColumnType_STRING,
				Description: "The third option for the variant.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time when the product variant was created.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time when the product variant was last updated.",
			},
			{
				Name:        "taxable",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether or not the product variant is taxable.",
			},
			{
				Name:        "tax_code",
				Type:        proto.ColumnType_STRING,
				Description: "The tax code of the product variant.",
			},
			{
				Name:        "barcode",
				Type:        proto.ColumnType_STRING,
				Description: "The barcode of the product variant.",
			},
			{
				Name:        "image_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID for the image associated with the variant.",
				Transform:   transform.FromField("ImageID"),
			},
			{
				Name:        "inventory_quantity",
				Type:        proto.ColumnType_INT,
				Description: "The number of inventory items available for the product variant.",
			},
			{
				Name:        "weight",
				Type:        proto.ColumnType_STRING,
				Description: "The weight of the product variant.",
			},
			{
				Name:        "weight_unit",
				Type:        proto.ColumnType_STRING,
				Description: "The unit of measurement for the product variant weight.",
			},
			{
				Name:        "old_inventory_quantity",
				Type:        proto.ColumnType_INT,
				Description: "The old inventory quantity of the product variant.",
			},
			{
				Name:        "requires_shipping",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the product variant requires shipping.",
			},
			{
				Name:        "admin_graphql_api_id",
				Type:        proto.ColumnType_STRING,
				Description: "The admin graphql API ID of the product.",
				Transform:   transform.FromField("AdminGraphqlAPIID"),
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: "Title of the resource.",
				Transform:   transform.FromField("Title"),
			},
		},
	}
}

func listProductVariants(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_product_variant.listProductVariants", "connection_error", err)
		return nil, err
	}
	id := h.Item.(goshopify.Product).ID

	variants, err := conn.Variant.List(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_product_variant.listProductVariants", "api_error", err)
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

	id := d.EqualsQuals["id"].GetInt64Value()

	// check if the id is 0
	if id == 0 {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_product_variant.getProductVariant", "connection_error", err)
		return nil, err
	}

	result, err := conn.Variant.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_product_variant.getProductVariant", "api_error", err)
		return nil, err
	}

	return result, nil
}
