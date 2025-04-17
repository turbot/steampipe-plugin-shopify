## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#48](https://github.com/turbot/steampipe-plugin-shopify/pull/48))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#48](https://github.com/turbot/steampipe-plugin-shopify/pull/48))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#46](https://github.com/turbot/steampipe-plugin-shopify/pull/46))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#46](https://github.com/turbot/steampipe-plugin-shopify/pull/46))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#42](https://github.com/turbot/steampipe-plugin-shopify/pull/42))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#42](https://github.com/turbot/steampipe-plugin-shopify/pull/42))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-shopify/blob/main/docs/LICENSE). ([#42](https://github.com/turbot/steampipe-plugin-shopify/pull/42))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#41](https://github.com/turbot/steampipe-plugin-shopify/pull/41))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#34](https://github.com/turbot/steampipe-plugin-shopify/pull/34))

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
