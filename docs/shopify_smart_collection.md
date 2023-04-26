# Table: shopify_smart_collection

Shopify smart collection is a grouping of products defined by rules that are set by the merchant. Shopify automatically changes the contents of a smart collection based on the rules. Smart collections, like other types of collections, are used to break down the catalog of products into categories and make the shop easier to browse.

## Examples

### Basic info

```sql
select 
  id,
  title,
  updated_at,
  handle 
from 
  shopify_smart_collection;
```

### Get the number of rules for each smart collection

```sql
select
  id,
  title,
  handle,
  updated_at,
  jsonb_array_length(rules) as num_rules
from
  shopify_smart_collection;
```

### Get the total number of published smart collections

```sql
select 
  count(*)
from 
  shopify_smart_collection
where 
  published;
```

### Get the smart collections published with in last 30 days

```sql
select
  id,
  title,
  handle,
  updated_at,
  published_at
from
  shopify_smart_collection
where
  published_at >= (published_at - interval '30' day)
order by
  published_at;
```

### Get the IDs, titles, and image URLs of all custom collections that have an image

```sql
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

### Retrieve all smart collections that contain a specific metafield value

```sql
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