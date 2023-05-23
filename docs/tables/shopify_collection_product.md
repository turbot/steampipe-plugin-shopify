# Table: shopify_collection_product

Shopify collection products are a group of related products that are curated and organized together by a store owner on the Shopify platform.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

```sql
select
  collection_title,
  count(*) as total_count
from
  shopify_collection_product
group by
  collection_title;
```

### Get the number of products in each collection that are currently in stock

```sql
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

```sql
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
  created_at >= now()- interval '30' day;
```