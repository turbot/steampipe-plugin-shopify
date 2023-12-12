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

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

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

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Shopify Plugin](https://github.com/turbot/steampipe-plugin-shopify/labels/help%20wanted)