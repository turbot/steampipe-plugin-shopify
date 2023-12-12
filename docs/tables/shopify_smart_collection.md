---
title: "Steampipe Table: shopify_smart_collection - Query Shopify Smart Collections using SQL"
description: "Allows users to query Shopify Smart Collections, specifically to retrieve the rules and metadata associated with each collection, providing insights into product categorization and organization."
---

# Table: shopify_smart_collection - Query Shopify Smart Collections using SQL

A Shopify Smart Collection is a feature within Shopify that allows you to create product groups based on predefined conditions. These conditions can include product title, type, price, tag, weight, and more. Smart Collections provide a dynamic way to organize and display products in your Shopify store, improving product discovery and shopping experience for your customers.

## Table Usage Guide

The `shopify_smart_collection` table provides insights into Smart Collections within Shopify. As an e-commerce manager or developer, explore collection-specific details through this table, including the rules that define each collection and associated metadata. Utilize it to uncover information about your product organization, such as collections that include specific product types or tags, and to verify the conditions that define each collection.

## Examples

### Basic info
Explore which smart collections in your Shopify store have been updated recently. This can help you track product trends and manage your inventory more effectively.

```sql+postgres
select
  id,
  title,
  updated_at,
  handle
from
  shopify_smart_collection;
```

```sql+sqlite
select
  id,
  title,
  updated_at,
  handle
from
  shopify_smart_collection;
```

### Get the number of rules for each smart collection
Explore the number of rules applied to each smart collection in your Shopify account, allowing you to assess the complexity of your collection filtering system. This can help streamline your product management by identifying collections with an excessive or minimal number of rules.

```sql+postgres
select
  id,
  title,
  handle,
  updated_at,
  jsonb_array_length(rules) as num_rules
from
  shopify_smart_collection;
```

```sql+sqlite
select
  id,
  title,
  handle,
  updated_at,
  json_array_length(rules) as num_rules
from
  shopify_smart_collection;
```

### Get the total number of published smart collections
Determine the total count of smart collections that have been made publicly available. This information can be useful for assessing the volume of collections that are accessible to customers.

```sql+postgres
select
  count(*) as published_smart_collection
from
  shopify_smart_collection
where
  published;
```

```sql+sqlite
select
  count(*) as published_smart_collection
from
  shopify_smart_collection
where
  published = 1;
```

### Get the smart collections published with in last 30 days
Explore which smart collections were published within the last 30 days. This can help keep track of recent updates and manage your collections more effectively.

```sql+postgres
select
  id,
  title,
  handle,
  updated_at,
  published_at
from
  shopify_smart_collection
where
  published_at >= now() - interval '30' day
order by
  published_at;
```

```sql+sqlite
select
  id,
  title,
  handle,
  updated_at,
  published_at
from
  shopify_smart_collection
where
  published_at >= datetime('now', '-30 day')
order by
  published_at;
```

### Get the IDs, titles, and image URLs of all custom collections that have an image
Discover the custom collections that include images, which can be useful for auditing visual content or identifying collections for promotional campaigns.

```sql+postgres
select
  id,
  title,
  updated_at,
  handle,
  image ->> 'src' as image_url
from
  shopify_smart_collection
where
  image ->> 'src' is not null;
```

```sql+sqlite
select
  id,
  title,
  updated_at,
  handle,
  json_extract(image, '$.src') as image_url
from
  shopify_smart_collection
where
  json_extract(image, '$.src') is not null;
```

### Retrieve all smart collections that contain a specific metafield value
Discover the segments that contain a certain metafield value within all smart collections. This is particularly useful when you want to identify and analyze the collections that are associated with a specific attribute or characteristic.

```sql+postgres
select
  id,
  title,
  handle,
  published,
  updated_at,
  metafields
from
  shopify_smart_collection
where
  metafields @> '[{"value": "hello test 123"}]';
```

```sql+sqlite
select
  id,
  title,
  handle,
  published,
  updated_at,
  metafields
from
  shopify_smart_collection
where
  json_extract(metafields, '$[*].value') = 'hello test 123';
```

### List all disjunctive smart collections
Explore the smart collections on your Shopify store that use the disjunctive condition, allowing for a broader product inclusion in each collection. This can be useful in understanding the diversity of your product range and how it is categorized.

```sql+postgres
select
  id,
  title,
  updated_at,
  handle
from
  shopify_smart_collection
where
  disjunctive;
```

```sql+sqlite
select
  id,
  title,
  updated_at,
  handle
from
  shopify_smart_collection
where
  disjunctive = 1;
```