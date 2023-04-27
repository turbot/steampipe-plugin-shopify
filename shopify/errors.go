package shopify

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// isNotFoundError:: function which returns an ErrorPredicate for Shopify API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		// Added to support regex in not found errors
		for _, pattern := range notFoundErrors {
			if strings.Contains(err.Error(), pattern) {
				return true
			}
		}
		return false
	}
}

func shouldRetryError(retryErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {

		if strings.Contains(err.Error(), "Exceeded 2 calls per second for api client. Reduce request rates to resume uninterrupted service.") {
			plugin.Logger(ctx).Debug("shopify_product_variant.shouldRetryError", "rate_limit_error", err)
			return true
		}
		return false
	}
}
