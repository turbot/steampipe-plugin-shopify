---
title: "Steampipe Table: shopify_collection_product - Query Shopify Collection Products using SQL"
description: "Allows users to query Collection Products in Shopify, specifically the associated product details and collection information, providing insights into product categorization and organization."
---

# Table: shopify_collection_product - Query Shopify Collection Products using SQL

Shopify Collection Products represent the relationship between products and their associated collections within the Shopify platform. Collections are a way to group and organize products, making them easier to manage and find. They can be created manually, or they can be automatically generated based on conditions that you specify.

## Table Usage Guide

The `shopify_collection_product` table provides insights into the relationship between products and collections within Shopify. As a store manager or developer, explore product-specific details and their collection associations through this table, including collection IDs, product IDs, and position in the collection. Utilize it to uncover information about product organization, such as which products belong to which collections, and the order of products within collections.

## Examples

### Basic info
Explore which products are included in your Shopify collections, along with their vendors and status. This can help you manage and organize your products more effectively.

```sql+postgres
select
  collection_id,
  collection_title,
  product_id,
  product_title,
  vendor,
  status
from
  shopify_collection_product;
```

```sql+sqlite
select
  collection_id,
  collection_title,
  product_id,
  product_title,
  vendor,
  status
from
  shopify_collection_product;
```

### List all products in a specific collection by collection ID
Explore which products belong to a certain collection using a unique identifier, allowing you to assess the range and variety within that collection. This could be useful for inventory management or identifying gaps in your product offerings.

```sql+postgres
select
  collection_id,
  product_id,
  product_title,
  handle,
  vendor,
  status
from
  shopify_collection_product
where
  collection_id = 444300460327;
```

```sql+sqlite
select
  collection_id,
  product_id,
  product_title,
  handle,
  vendor,
  status
from
  shopify_collection_product
where
  collection_id = 444300460327;
```

### List all products in a specific collection by collection title
Explore which products fall under a specific collection in your Shopify store. This is particularly useful to assess the range of items within a given collection, aiding in inventory management and marketing efforts.

```sql+postgres
select
  collection_id,
  product_id,
  product_title,
  collection_title,
  handle,
  vendor,
  status
from
  shopify_collection_product
where
  collection_title = 'Jelly';
```

```sql+sqlite
select
  collection_id,
  product_id,
  product_title,
  collection_title,
  handle,
  vendor,
  status
from
  shopify_collection_product
where
  collection_title = 'Jelly';
```

### Get the total count of products in each collection
Explore which collections have the most products to better manage inventory and sales strategies. This allows for a comprehensive understanding of product distribution across different collections.

```sql+postgres
select
  collection_title,
  count(*) as total_count
from
  shopify_collection_product
group by
  collection_title;
```

```sql+sqlite
select
  collection_title,
  count(*) as total_count
from
  shopify_collection_product
group by
  collection_title;
```

### Get the number of products in each collection that are currently in stock
Explore which collections have products currently in stock. This query is useful for inventory management, allowing you to see the distribution of available products across different collections.

```sql+postgres
select
  collection_title,
  count(*) as total_count 
from
  shopify_collection_product
where
  status = 'in_stock' 
group by
  collection_title;
```

```sql+sqlite
select
  collection_title,
  count(*) as total_count 
from
  shopify_collection_product
where
  status = 'in_stock' 
group by
  collection_title;
```

### List all the products in a collection created in the last one month
Identify all the products added to any collection within the past month. This can assist in understanding recent inventory changes and tracking product performance.

```sql+postgres
select
  collection_id,
  collection_title,
  product_id,
  product_title,
  vendor,
  status,
  created_at
from
  shopify_collection_product
where
  created_at >= now() - interval '30' day;
```

```sql+sqlite
select
  collection_id,
  collection_title,
  product_id,
  product_title,
  vendor,
  status,
  created_at
from
  shopify_collection_product
where
  created_at >= datetime('now', '-30 day');
```