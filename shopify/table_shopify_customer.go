package shopify

import (
	"context"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShopifyCustomer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shopify_customer",
		Description: "Shopify customer stores information about a shop's customers, such as their contact details, their order history, and whether they've agreed to receive email marketing.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getCustomers,
		},
		List: &plugin.ListConfig{
			Hydrate: listCustomers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "The customer ID.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The Email ID of the customer.",
			},
			{
				Name:        "first_name",
				Type:        proto.ColumnType_STRING,
				Description: "First name of the customer.",
			},
			{
				Name:        "last_name",
				Type:        proto.ColumnType_STRING,
				Description: "Last name of the customer.",
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "Customer state.",
			},
			{
				Name:        "note",
				Type:        proto.ColumnType_STRING,
				Description: "A specific note related to the customer.",
			},
			{
				Name:        "verified_email",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the customer email is verified.",
			},
			{
				Name:        "multipass_identifier",
				Type:        proto.ColumnType_STRING,
				Description: "Customer multi-pass identifier.",
			},
			{
				Name:        "orders_count",
				Type:        proto.ColumnType_INT,
				Description: "The customer order count.",
			},
			{
				Name:        "tax_exempt",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the customer is tax exempted.",
			},
			{
				Name:        "total_spent",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("TotalSpent").Transform(transform.ToString).Transform(transform.ToDouble),
				Description: "Total amount spent by the customer.",
			},
			{
				Name:        "phone",
				Type:        proto.ColumnType_STRING,
				Description: "Customer phone.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_STRING,
				Description: "Tags attached to the customer.",
			},
			{
				Name:        "last_order_id",
				Type:        proto.ColumnType_INT,
				Description: "The last order ID of the customer.",
			},
			{
				Name:        "last_order_name",
				Type:        proto.ColumnType_STRING,
				Description: "The last order name of the customer.",
			},
			{
				Name:        "accepts_marketing",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the customer subscribed to email marketing campaign.",
			},
			{
				Name:        "default_address",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
				Description: "Customer default address.",
			},
			{
				Name:        "addresses",
				Type:        proto.ColumnType_JSON,
				Description: "Other addresses of the customer.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the customer created their account.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the customer last updated their account.",
			},
			{
				Name:        "metafields",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromJSONTag(),
				Description: "Customer metafields.",
			},
		},
	}
}

func listCustomers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listCustomers", "connection_error", err)
		return nil, err
	}
	// the max limit defined by the API is 250
	options := goshopify.ListOptions{}

	// set the limit if a lower limit is passed in query context
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < 250 {
			options.Limit = int(*limit)
		}
	}

	for {
		customers, paginator, err := conn.Customer.ListWithPagination(options)
		if err != nil {
			plugin.Logger(ctx).Error("listCustomersError", "list_api_error", err)
			return nil, err
		}

		for _, customer := range customers {
			d.StreamListItem(ctx, customer)

			// Context can be cancelled due to manual cancellation or the limit has been hit
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

func getCustomers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getCustomer", "connection_error", err)
		return nil, err
	}

	id := d.EqualsQuals["id"].GetInt64Value()
	// check if the id is empty
	if id == 0 {
		return nil, nil
	}

	result, err := conn.Customer.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("getCustomer", "get_api_error", err)
		return nil, err
	}

	return result, nil
}
