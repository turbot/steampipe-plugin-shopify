![image](https://hub.steampipe.io/images/plugins/turbot/shopify-social-graphic.png)

# Shopify Plugin for Steampipe

Use SQL to query products, order, customers and more from Shopify.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/shopify)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/shopify/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-shopify/issues)

## Quick start

### Install

Download and install the latest Shopify plugin:

```shell
steampipe plugin install shopify
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/shopify#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/shopify#configuration).

Configure your account details in `~/.steampipe/config/shopify.spc`:

```hcl
connection "shopify" {
  plugin = "shopify"

  # Authentication information
  api_token = "shpat_ab0a4zaa19c3faketoken924176b387d"
  shop_name = "myshop"
}
```

Or through environment variables:

```sh
export SHOPIFY_API_TOKEN=shpat_ab0a4zaa19c3faketoken924176b387d
export SHOPIFY_SHOP_NAME=theshop
```

Run steampipe:

```shell
steampipe query
```

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

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-shopify.git
cd steampipe-plugin-shopify
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/shopify.spc
```

Try it!

```
steampipe query
> .inspect shopify
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-shopify/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Shopify Plugin](https://github.com/turbot/steampipe-plugin-shopify/labels/help%20wanted)