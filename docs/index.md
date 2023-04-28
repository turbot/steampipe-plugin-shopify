---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/shopify.svg"
brand_color: "#6BBF4E"
display_name: "Shopify"
short_name: "shopify"
description: "Steampipe plugin to query products, order, customers and more from Shopify."
og_description: "Query Shopify with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/shopify-social-graphic.png"
---

# Shopify + Steampipe

[Shopify](https://shopify.com/) is an e-commerce platform that allows businesses to create and manage online stores.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Get Shopify order details:

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

## Get started

### Install

Download and install the latest Shopify plugin:

```bash
steampipe plugin install shopify
```

### Configuration

Installing the latest shopify plugin will create a config file (`~/.steampipe/config/shopify.spc`) with a single connection named `shopify`:

```hcl
connection "shopify" {
  plugin = "shopify"

  # API access token to request data from the Admin API. e.g., `shpat_ab0a4zaa19c3faketoken924176b387d`.
  # Please see https://www.shopify.com/partners/blog/17056443-how-to-generate-a-shopify-api-token for more information.
  # Can also be set with the SHOPIFY_API_TOKEN environment variable.
  # token = "shpat_ab0a4zaa19c3faketoken924176b387d"

  # The shop_name parameter is the shop's myshopify domain, e.g. "theshop.myshopify.com", or simply "theshop".
  # Can also be set with the SHOPIFY_SHOP_NAME environment variable.
  # shop_name = "theshop"
}
```

- `shop_name` - Shopify shop_name refers to the unique name that a merchant chooses for their online store built on the Shopify platform.

- `token` - Shopify api_token is a secure and unique identifier that is used to authenticate and authorize a third-party app or service to access a Shopify store's data and perform actions on behalf of the store owner.

Environment variables are also available as an alternate configuration method:

- `SHOPIFY_SHOP_NAME`
- `SHOPIFY_API_TOKEN`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-shopify
- Community: [Slack Channel](https://steampipe.io/community/join)