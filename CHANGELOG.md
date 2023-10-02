## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#32](https://github.com/turbot/steampipe-plugin-shopify/pull/32))
- Recompiled plugin with Go version `1.21`. ([#32](https://github.com/turbot/steampipe-plugin-shopify/pull/32))

## v0.0.2 [2023-08-25]

_Bug fixes_

- Fixed pagination in the following tables: ([#26](https://github.com/turbot/steampipe-plugin-shopify/pull/26))
  - `shopify_collection_product`
  - `shopify_customer`
  - `shopify_order`
  - `shopify_product`

## v0.0.1 [2023-05-29]

_What's new?_

- New tables added
  - [shopify_collection_product](https://hub.steampipe.io/plugins/turbot/shopify/tables/shopify_collection_product)
  - [shopify_custom_collection](https://hub.steampipe.io/plugins/turbot/shopify/tables/shopify_custom_collection)
  - [shopify_customer](https://hub.steampipe.io/plugins/turbot/shopify/tables/shopify_customer)
  - [shopify_draft_order](https://hub.steampipe.io/plugins/turbot/shopify/tables/shopify_draft_order)
  - [shopify_order](https://hub.steampipe.io/plugins/turbot/shopify/tables/shopify_order)
  - [shopify_product](https://hub.steampipe.io/plugins/turbot/shopify/tables/shopify_product)
  - [shopify_product_variant](https://hub.steampipe.io/plugins/turbot/shopify/tables/shopify_product_variant)
  - [shopify_smart_collection](https://hub.steampipe.io/plugins/turbot/shopify/tables/shopify_smart_collection)
  - [shopify_theme](https://hub.steampipe.io/plugins/turbot/shopify/tables/shopify_theme)
