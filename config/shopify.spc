connection "shopify" {
  plugin = "shopify"

  # `api_token`: API access token to request data from the Admin API. e.g., `shpat_ab0a4zaa19c3faketoken924176b387d`.
  # Please see https://www.shopify.com/partners/blog/17056443-how-to-generate-a-shopify-api-token for more information.
  # Can also be set with the SHOPIFY_API_TOKEN environment variable.
  # api_token = "shpat_ab0a4zaa19c3faketoken924176b387d"

  # `shop_name`: The shop_name parameter is the shop's myshopify domain, e.g. "theshop.myshopify.com", or simply "theshop".
  # Can also be set with the SHOPIFY_SHOP_NAME environment variable.
  # shop_name = "theshop"
}