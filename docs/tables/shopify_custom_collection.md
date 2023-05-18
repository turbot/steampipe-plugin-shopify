# Table: shopify_custom_collection

Shopify custom collection is a group of products that a merchant can create to make their store easier to browse.

## Examples

### Basic info

```sql
select
  id,
  title,
  handle,
  published_scope
from
  shopify_custom_collection;
```

### Get the total number of custom collections

```sql
select
  count(*)
from
  shopify_custom_collection;
```

### Retrieve all published custom collections

```sql
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

### Retrieve the custom collection details with a specific ID

```sql
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

```sql
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

```sql
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

```sql
select
  count(*)
from
  shopify_custom_collection
where
  published=true
and
  date_trunc('month', published_at) = '2023-04-01';
```

### Retrieve all custom collections that contain a specific metafield key

```sql
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
