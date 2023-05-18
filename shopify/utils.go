package shopify

import (
	"context"

	"github.com/shopspring/decimal"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)
func convertPrice(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	price := d.Value.(*decimal.Decimal)
	if price != nil {
		value, _ := price.Float64()
		return value, nil
	}
	return nil, nil
}