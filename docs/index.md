---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/shopify.svg"
brand_color: "#95BF47"
display_name: "Shopify"
short_name: "shopify"
description: "Steampipe plugin to query products, order, customers and more from Shopify."
og_description: "Query Shopify with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/shopify-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Shopify + Steampipe

[Shopify](https://shopify.com/) is an e-commerce platform that allows businesses to create and manage online stores.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List Shopify order details:

```sql
select
  id,
  name,
  email,
  billing_address ->> 'address1' as address,
  billing_address ->> 'city' as city,
  billing_address ->> 'country'as country,
  billing_address ->> 'zip'as zip_code
from
  shopify_order;
```

```
+---------------+-------+-----------------------------+-----------------+---------+---------+----------+
| id            | name  | email                       | address         | city    | country | zip_code |
+---------------+-------+-----------------------------+-----------------+---------+---------+----------+
| 5367225188647 | #1007 | russel.winfield@example.com | 105 Victoria St | Toronto | Canada  | M5C1N7   |
| 5367225057575 | #1003 | russel.winfield@example.com | 105 Victoria St | Toronto | Canada  | M5C1N7   |
| 5367225254183 | #1009 | russel.winfield@example.com | 105 Victoria St | Toronto | Canada  | M5C1N7   |
| 5367225221415 | #1008 | russel.winfield@example.com | 105 Victoria St | Toronto | Canada  | M5C1N7   |
+---------------+-------+-----------------------------+-----------------+---------+---------+----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/shopify/tables)**

## Quick start

### Install

Download and install the latest Shopify plugin:

```bash
steampipe plugin install shopify
```

### Credentials

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Shopify requires a `Shop name` and an [API token](https://shopify.dev/docs/apps/auth/admin-app-access-tokens) for all requests.                                                                |
| Permissions | API tokens have the same permissions as the user who creates them, and if the user permissions change, the token permissions also change.                                                         |
| Radius      | Each connection represents a single Shopify Installation.                                                                                                                                           |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/shopify.spc`)<br />2. Credentials specified in environment variables, e.g., `SHOPIFY_API_TOKEN`, `SHOPIFY_SHOP_NAME`. |

### Configuration

Installing the latest shopify plugin will create a config file (`~/.steampipe/config/shopify.spc`) with a single connection named `shopify`:

Configure your account details in `~/.steampipe/config/shopify.spc`:

```hcl
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
```

Alternatively, you can also use the standard Shopify environment variables to obtain credentials **only if other arguments (`api_token` and `shop_name`) are not specified** in the connection:

```sh
export SHOPIFY_API_TOKEN=shpat_ab0a4zaa19c3faketoken924176b387d
export SHOPIFY_SHOP_NAME=theshop
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-shopify
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)