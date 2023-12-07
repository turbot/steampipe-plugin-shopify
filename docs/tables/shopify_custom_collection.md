---
title: "Steampipe Table: shopify_custom_collection - Query Shopify Custom Collections using SQL"
description: "Allows users to query Shopify Custom Collections, specifically the details of each custom collection, providing insights into product organization and categorization."
---

# Table: shopify_custom_collection - Query Shopify Custom Collections using SQL

Shopify Custom Collections are a feature within Shopify that allows you to group your products into categories. These collections can be manually curated or automatically generated based on conditions you set. Custom collections are useful for organizing products by type, season, sale, and other variables that suit your store's needs.

## Table Usage Guide

The `shopify_custom_collection` table provides insights into custom collections within Shopify. As a store manager or developer, explore collection-specific details through this table, including the collection title, product count, and related metadata. Utilize it to uncover information about collections, such as their organization, the products they contain, and their visibility status.

## Examples

### Basic info
Explore the basic details of your custom collections on Shopify, such as their unique identifiers, titles, handles, and the scope in which they are published. This can help you better manage and organize your collections.

```sql+postgres
select
  id,
  title,
  handle,
  published_scope
from
  shopify_custom_collection;
```

```sql+sqlite
select
  id,
  title,
  handle,
  published_scope
from
  shopify_custom_collection;
```

### Get the total number of custom collections
Discover the total number of custom collections in your Shopify store. This can be useful for managing inventory and understanding the diversity of your product offerings.

```sql+postgres
select
  count(*) as total_custom_collections
from
  shopify_custom_collection;
```

```sql+sqlite
select
  count(*) as total_custom_collections
from
  shopify_custom_collection;
```

### Retrieve all published custom collections
Discover the segments that are part of all the published custom collections. This is useful in understanding which collections are available for public viewing, aiding in inventory management and marketing strategies.

```sql+postgres
select
  id,
  title,
  handle,
  published
from
  shopify_custom_collection
where
  published = true;
```

```sql+sqlite
select
  id,
  title,
  handle,
  published
from
  shopify_custom_collection
where
  published = 1;
```

### Retrieve the custom collection details with a specific ID
Explore the specific details of a custom collection in your Shopify store using its unique ID. This is particularly useful to quickly assess the status and properties of a collection without having to manually search for it.

```sql+postgres
select
  id,
  title,
  handle,
  published
from
  shopify_custom_collection
where
  id = 444300460327;
```

```sql+sqlite
select
  id,
  title,
  handle,
  published
from
  shopify_custom_collection
where
  id = 444300460327;
```

### Retrieve the custom collection with a specific handle
Explore which custom collections are associated with a specific handle in Shopify to assess their publication status. This can be particularly useful for managing and organizing your collections efficiently.

```sql+postgres
select
  id,
  title,
  handle,
  published
from
  shopify_custom_collection
where
  handle = 'jelly';
```

```sql+sqlite
select
  id,
  title,
  handle,
  published
from
  shopify_custom_collection
where
  handle = 'jelly';
```

### Retrieve all custom collections updated after a specific date
Explore which custom collections have been updated after a certain date. This is particularly useful for keeping track of recent changes to your collections on Shopify.

```sql+postgres
select
  id,
  title,
  handle,
  published,
  updated_at
from
  shopify_custom_collection
where
  updated_at > '2023-01-01';
```

```sql+sqlite
select
  id,
  title,
  handle,
  published,
  updated_at
from
  shopify_custom_collection
where
  updated_at > '2023-01-01';
```

### Retrieve the number of custom collections published in a specific month
Discover the volume of custom collections made public within a specific month. This could be useful for understanding publishing trends or assessing the effectiveness of content strategies during that period.

```sql+postgres
select
  count(*) as custom_collection_count
from
  shopify_custom_collection
where
  published = true
  and date_trunc('month', published_at) = '2023-04-01';
```

```sql+sqlite
select
  count(*) as custom_collection_count
from
  shopify_custom_collection
where
  published = 1
  and strftime('%Y-%m', published_at) = '2023-04';
```

### Retrieve all custom collections that contain a specific metafield key
Discover the segments that have a particular type of metafield key in your custom collections. This can help you understand and manage the collections that share common characteristics or settings.

```sql+postgres
select
  id,
  title,
  handle,
  published,
  updated_at,
  jsonb_pretty(metafields)
from
  shopify_custom_collection
where
  metafields @> '[{"key": "description_tag"}]';
```

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```