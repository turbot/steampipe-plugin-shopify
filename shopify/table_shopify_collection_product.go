package shopify

import (
	"context"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type Collection struct {
	ID    int64
	Title string
}

type CollectionProduct struct {
	Collection Collection
	Product    goshopify.Product
}

func tableShopifyCollectionProduct(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shopify_collection_product",
		Description: "Shopify collection product is a group of related items that are organized together on an online store for easy browsing and purchasing by customers.",
		List: &plugin.ListConfig{
			ParentHydrate: listCollections,
			Hydrate:       listCollectionProducts,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "collection_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the custom collection.",
				Transform:   transform.FromField("Collection.ID"),
			},
			{
				Name:        "collection_title",
				Type:        proto.ColumnType_STRING,
				Description: "Title of the custom collection.",
				Transform:   transform.FromField("Collection.Title"),
			},
			{
				Name:        "product_id",
				Type:        proto.ColumnType_INT,
				Description: "Product ID.",
				Transform:   transform.FromField("Product.ID"),
			},
			{
				Name:        "product_title",
				Type:        proto.ColumnType_STRING,
				Description: "Product title.",
				Transform:   transform.FromField("Product.Title"),
			},
			{
				Name:        "handle",
				Type:        proto.ColumnType_STRING,
				Description: "Product handle.",
				Transform:   transform.FromField("Product.Handle"),
			},
			{
				Name:        "vendor",
				Type:        proto.ColumnType_STRING,
				Description: "The vendor of the shopify products.",
				Transform:   transform.FromField("Product.Vendor"),
			},
			{
				Name:        "product_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the products sold.",
				Transform:   transform.FromField("Product.ProductType"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The product creation date.",
				Transform:   transform.FromField("Product.CreatedAt"),
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time product was updated.",
				Transform:   transform.FromField("Product.UpdatedAt"),
			},
			{
				Name:        "published_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The product publish time.",
				Transform:   transform.FromField("Product.PublishedAt"),
			},
			{
				Name:        "published_scope",
				Type:        proto.ColumnType_STRING,
				Description: "The product publish scope.",
				Transform:   transform.FromField("Product.PublishedScope"),
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_STRING,
				Description: "The product tags.",
				Transform:   transform.FromField("Product.Tags"),
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The product status.",
				Transform:   transform.FromField("Product.Status"),
			},
			{
				Name:        "options",
				Type:        proto.ColumnType_JSON,
				Description: "The product options.",
				Transform:   transform.FromField("Product.Options"),
			},
			{
				Name:        "variants",
				Type:        proto.ColumnType_JSON,
				Description: "The product variants.",
				Transform:   transform.FromField("Product.Variants"),
			},
			{
				Name:        "image",
				Type:        proto.ColumnType_JSON,
				Description: "The product image.",
				Transform:   transform.FromField("Product.Image"),
			},
			{
				Name:        "images",
				Type:        proto.ColumnType_JSON,
				Description: "List of images associated with the product.",
				Transform:   transform.FromField("Product.Images"),
			},
			{
				Name:        "template_suffix",
				Type:        proto.ColumnType_STRING,
				Description: "The product template suffix.",
				Transform:   transform.FromField("Product.TemplateSuffix"),
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: "Title of the resource.",
				Transform:   transform.FromField("Product.Title"),
			},
		}),
	}
}

func listCollections(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_collection_product.listCollections", "connection_error", err)
		return nil, err
	}

	// TODO: try and use existing listCustomCollections and listSmartCollections funcs
	// from existing customCollections and smartCollection tables
	// customCols, err := listCustomCollections(ctx, d, h)
	// if err != nil {
	// 	return nil, err
	// }
	customCollections, err := conn.CustomCollection.List(nil)
	if err != nil {
		return nil, err
	}
	smartCollections, err := conn.SmartCollection.List(nil)
	if err != nil {
		return nil, err
	}

	// Add custom collections to the slice
	for _, c := range customCollections {
		d.StreamListItem(ctx, Collection{
			ID:    c.ID,
			Title: c.Title,
		})
	}
	// Add smart collections to the slice
	for _, c := range smartCollections {
		d.StreamListItem(ctx, Collection{
			ID:    c.ID,
			Title: c.Title,
		})
	}
	return nil, nil
}

func listCollectionProducts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(Collection)
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("shopify_collection_product.listCollectionProducts", "connection_error", err)
		return nil, err
	}
	// the max limit defined by the API is 250
	options := goshopify.ListOptions{}

	// set the limit if a lower limit is passed in query context
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < 250 {
			options.Limit = int(*limit)
		} else {
			options.Limit = 250
		}
	}

	for {
		products, paginator, err := conn.Collection.ListProductsWithPagination(data.ID, nil)
		if err != nil {
			plugin.Logger(ctx).Error("shopify_collection_product.listProducts", "api_error", err)
			return nil, err
		}

		for _, product := range products {
			d.StreamListItem(ctx, CollectionProduct{
				Collection: data,
				Product:    product,
			})
		}

		if paginator.NextPageOptions == nil {
			return nil, nil
		}
		options.PageInfo = paginator.NextPageOptions.PageInfo
	}

}
