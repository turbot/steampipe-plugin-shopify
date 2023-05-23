package shopify

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShopifyDraftOrder(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shopify_draft_order",
		Description: "A draft order is an order a merchant makes in the Shopify admin on behalf of a customer.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDraftOrder,
		},
		List: &plugin.ListConfig{
			Hydrate: listDraftOrders,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the Shopify product.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "order_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the draft order.",
				Transform:   transform.FromField("OrderID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the draft order.",
			},
			{
				Name:        "customer",
				Type:        proto.ColumnType_JSON,
				Description: "The customer associated with the draft order.",
			},
			{
				Name:        "shipping_address",
				Type:        proto.ColumnType_JSON,
				Description: "The shipping address for the draft order.",
			},
			{
				Name:        "billing_address",
				Type:        proto.ColumnType_JSON,
				Description: "The billing address for the draft order.",
			},
			{
				Name:        "note",
				Type:        proto.ColumnType_STRING,
				Description: "An optional note attached to the draft order.",
			},
			{
				Name:        "note_attributes",
				Type:        proto.ColumnType_JSON,
				Description: "Additional metadata about the draft order.",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "The email address associated with the draft order.",
			},
			{
				Name:        "currency",
				Type:        proto.ColumnType_STRING,
				Description: "The currency used for the draft order.",
			},
			{
				Name:        "invoice_sent_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time the invoice was sent, if applicable.",
			},
			{
				Name:        "invoice_url",
				Type:        proto.ColumnType_STRING,
				Description: "The URL of the invoice, if applicable.",
				Transform:   transform.FromField("InvoiceURL"),
			},
			{
				Name:        "line_items",
				Type:        proto.ColumnType_JSON,
				Description: "The line items in the draft order.",
			},
			{
				Name:        "shipping_line",
				Type:        proto.ColumnType_JSON,
				Description: "The shipping details for the draft order.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_STRING,
				Description: "Tags associated with the draft order.",
			},
			{
				Name:        "tax_lines",
				Type:        proto.ColumnType_JSON,
				Description: "The tax lines for the draft order.",
			},
			{
				Name:        "applied_discount",
				Type:        proto.ColumnType_JSON,
				Description: "Discounts applied to the draft order.",
			},
			{
				Name:        "taxes_included",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether taxes are included in the product price.",
			},
			{
				Name:        "total_tax",
				Type:        proto.ColumnType_STRING,
				Description: "The total amount of tax charged for the draft order.",
			},
			{
				Name:        "tax_exempt",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the draft order is tax exempt.",
			},
			{
				Name:        "total_price",
				Type:        proto.ColumnType_STRING,
				Description: "The total price of the draft order.",
			},
			{
				Name:        "subtotal_price",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("SubtotalPrice").Transform(convertPrice),
				Description: "The total price of all the draft order's line items, before taxes and discounts.",
			},
			{
				Name:        "completed_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the draft order was completed.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the draft order was created.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the draft order was last updated.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the draft order.",
			},
			{
				Name:        "use_customer_default_address",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether to use the customer's default address for shipping and/or billing.",
			},
		},
	}
}

func listDraftOrders(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listDraftOrder", "connection_error", err)
		return nil, err
	}

	draftOrders, err := conn.DraftOrder.List(nil)
	if err != nil {
		plugin.Logger(ctx).Error("listDraftOrder", "list_api_error", err)
		return nil, err
	}

	for _, draftOrder := range draftOrders {
		d.StreamListItem(ctx, draftOrder)

		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, nil
}

func getDraftOrder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getDraftOrder", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetInt64Value()

	// check if the id is 0
	if id == 0 {
		return nil, nil
	}
	result, err := conn.DraftOrder.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("getDraftOrder", "api_error", err)
		return nil, err
	}

	return result, nil
}
