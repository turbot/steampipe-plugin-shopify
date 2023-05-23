package shopify

import (
	"context"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableShopifyOrder(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "shopify_order",
		Description: "Shopify orders are records of purchases made by customers through a Shopify-powered online store, containing information such as products ordered, customer details, and transaction details.",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getOrder,
		},
		List: &plugin.ListConfig{
			Hydrate: listOrders,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "The order ID.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The order name.",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "Email id of the customer who placed the order.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the order was placed.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the order was updated.",
			},
			{
				Name:        "cancelled_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the order was cancelled.",
			},
			{
				Name:        "closed_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the order was closed.",
			},
			{
				Name:        "processed_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the order was processed.",
			},
			{
				Name:        "customer",
				Type:        proto.ColumnType_JSON,
				Description: "The details of the customer who placed the order.",
			},
			{
				Name:        "billing_address",
				Type:        proto.ColumnType_JSON,
				Description: "The billing address of the customer who placed the order.",
			},
			{
				Name:        "shipping_address",
				Type:        proto.ColumnType_JSON,
				Description: "The shipping address of the customer who placed the order.",
			},
			{
				Name:        "currency",
				Type:        proto.ColumnType_STRING,
				Description: "The currency used for a particular order.",
			},
			{
				Name:        "total_price",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The total price of the order.",
				Transform:   transform.FromField("TotalPrice").Transform(transform.ToString).Transform(transform.ToDouble),
			},
			{
				Name:        "current_total_price",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The current total price of the order.",
				Transform:   transform.FromField("CurrentTotalPrice").Transform(transform.ToString).Transform(transform.ToDouble),
			},
			{
				Name:        "subtotal_price",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The subtotal price of the order.",
				Transform:   transform.FromField("SubtotalPrice").Transform(transform.ToString).Transform(transform.ToDouble),
			},
			{
				Name:        "total_discounts",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The total discounts on the order.",
				Transform:   transform.FromField("TotalDiscounts").Transform(transform.ToString).Transform(transform.ToDouble),
			},
			{
				Name:        "total_line_items_price",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The total line items price on the order.",
				Transform:   transform.FromField("TotalLineItemsPrice").Transform(transform.ToString).Transform(transform.ToDouble),
			},
			{
				Name:        "taxes_included",
				Type:        proto.ColumnType_BOOL,
				Description: "The taxes are included in the order or not.",
			},
			{
				Name:        "total_tax",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The total tax in the order.",
				Transform:   transform.FromField("TotalTax").Transform(transform.ToString).Transform(transform.ToDouble),
			},
			{
				Name:        "tax_lines",
				Type:        proto.ColumnType_JSON,
				Description: "The tax lines in the order.",
			},
			{
				Name:        "total_weight",
				Type:        proto.ColumnType_INT,
				Description: "The total weight of the order placed.",
			},
			{
				Name:        "financial_status",
				Type:        proto.ColumnType_STRING,
				Description: "The financial status of the order.",
			},
			{
				Name:        "fulfillments",
				Type:        proto.ColumnType_JSON,
				Description: "The fulfillments of the order.",
			},
			{
				Name:        "fulfillment_status",
				Type:        proto.ColumnType_STRING,
				Description: "The fulfillment status of the order.",
			},
			{
				Name:        "token",
				Type:        proto.ColumnType_STRING,
				Description: "The order token.",
			},
			{
				Name:        "cart_token",
				Type:        proto.ColumnType_STRING,
				Description: "The cart token of the order.",
			},
			{
				Name:        "number",
				Type:        proto.ColumnType_INT,
				Description: "The number of orders.",
			},
			{
				Name:        "order_number",
				Type:        proto.ColumnType_INT,
				Description: "The order number.",
			},
			{
				Name:        "note",
				Type:        proto.ColumnType_STRING,
				Description: "The order note.",
			},
			{
				Name:        "test",
				Type:        proto.ColumnType_STRING,
				Description: "The test of the order.",
			},
			{
				Name:        "browser_ip",
				Type:        proto.ColumnType_STRING,
				Description: "The browser ip of the order.",
			},
			{
				Name:        "buyer_accepts_marketing",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the buyer accepts email marketing or not.",
			},
			{
				Name:        "cancel_reason",
				Type:        proto.ColumnType_STRING,
				Description: "The reason for the order cancellation.",
			},
			{
				Name:        "note_attributes",
				Type:        proto.ColumnType_JSON,
				Description: "Any notes associated with the order.",
			},
			{
				Name:        "discount_codes",
				Type:        proto.ColumnType_JSON,
				Description: "Any discount codes applied to the order.",
			},
			{
				Name:        "line_items",
				Type:        proto.ColumnType_JSON,
				Description: "The line items of the order.",
			},
			{
				Name:        "shipping_lines",
				Type:        proto.ColumnType_JSON,
				Description: "The shipping lines of the order.",
			},
			{
				Name:        "transactions",
				Type:        proto.ColumnType_JSON,
				Description: "The transactions in the order.",
			},
			{
				Name:        "app_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the app that created the order.",
				Transform:   transform.FromField("AppID"),
			},
			{
				Name:        "customer_locale",
				Type:        proto.ColumnType_STRING,
				Description: "The locale associated with the customer.",
			},
			{
				Name:        "landing_site",
				Type:        proto.ColumnType_STRING,
				Description: "The landing site of the order.",
			},
			{
				Name:        "referring_site",
				Type:        proto.ColumnType_STRING,
				Description: "The referring site of the order.",
			},
			{
				Name:        "source_name",
				Type:        proto.ColumnType_STRING,
				Description: "The source name of the order.",
			},
			{
				Name:        "client_details",
				Type:        proto.ColumnType_JSON,
				Description: "The details of the client who placed the order.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_STRING,
				Description: "The order tags.",
			},
			{
				Name:        "location_id",
				Type:        proto.ColumnType_INT,
				Description: "The location id of the order.",
			},
			{
				Name:        "payment_gateway_names",
				Type:        proto.ColumnType_JSON,
				Description: "The payment gateway names associated with the order.",
			},
			{
				Name:        "processing_method",
				Type:        proto.ColumnType_STRING,
				Description: "The processing method used for the order.",
			},
			{
				Name:        "refunds",
				Type:        proto.ColumnType_JSON,
				Description: "The refunds associated with the order.",
			},
			{
				Name:        "user_id",
				Type:        proto.ColumnType_INT,
				Description: "The user id of the users who placed the order.",
			},
			{
				Name:        "order_status_url",
				Type:        proto.ColumnType_STRING,
				Description: "The URL to check the status of the order.",
			},
			{
				Name:        "gateway",
				Type:        proto.ColumnType_STRING,
				Description: "The payment gateway used for the orders.",
			},
			{
				Name:        "confirmed",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether the order has been confirmed.",
			},
			{
				Name:        "total_price_usd",
				Type:        proto.ColumnType_JSON,
				Description: "The total price of the order in USD.",
			},
			{
				Name:        "checkout_token",
				Type:        proto.ColumnType_STRING,
				Description: "The token for the checkout process.",
			},
			{
				Name:        "reference",
				Type:        proto.ColumnType_STRING,
				Description: "The reference number for the order.",
			},
			{
				Name:        "source_identifier",
				Type:        proto.ColumnType_STRING,
				Description: "The source identifier of the order.",
			},
			{
				Name:        "source_url",
				Type:        proto.ColumnType_STRING,
				Description: "The source URL of the order.",
			},
			{
				Name:        "device_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the device used to place the order.",
			},
			{
				Name:        "phone",
				Type:        proto.ColumnType_STRING,
				Description: "The phone number associated with the order.",
			},
			{
				Name:        "landing_site_ref",
				Type:        proto.ColumnType_STRING,
				Description: "The reference number for the landing site of the order.",
			},
			{
				Name:        "checkout_id",
				Type:        proto.ColumnType_INT,
				Description: "The ID of the checkout associated with the order.",
			},
			{
				Name:        "contact_email",
				Type:        proto.ColumnType_STRING,
				Description: "The email address associated with the order.",
			},
			{
				Name:        "send_receipt",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether or not a receipt was sent for the order.",
			},
			{
				Name:        "send_fulfillment_receipt",
				Type:        proto.ColumnType_BOOL,
				Description: "Whether or not a fulfillment receipt was sent for the order.",
			},
		},
	}
}

func listOrders(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listOrder", "connection_error", err)
		return nil, err
	}

	// max limit defined by the api is 250
	// We are setting status to 'any' to get all the orders(open, closed, cancelled)
	options := goshopify.OrderListOptions{
		ListOptions: goshopify.ListOptions{},
		Status:      "any",
	}

	// set the limit if a lower limit is passed in query context
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < 250 {
			options.ListOptions.Limit = int(*limit)
		}
	}

	for {
		orders, paginator, err := conn.Order.ListWithPagination(options)
		if err != nil {
			plugin.Logger(ctx).Error("listOrder", "list_api_error", err)
			return nil, err
		}

		for _, order := range orders {
			d.StreamListItem(ctx, order)

			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if paginator.NextPageOptions == nil {
			return nil, nil
		}
		options.ListOptions.PageInfo = paginator.NextPageOptions.PageInfo
	}
}

func getOrder(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getOrder", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["id"].GetInt64Value()

	// check if the id is 0
	if id == 0 {
		return nil, nil
	}
	result, err := conn.Order.Get(id, nil)
	if err != nil {
		plugin.Logger(ctx).Error("getOrder", "api_error", err)
		return nil, err
	}

	return result, nil
}
