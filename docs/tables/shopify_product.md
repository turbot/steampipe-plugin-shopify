# Table: shopify_product

Shopify Products are the goods, digital downloads, services, and gift cards that are sold in Shopify.

## Examples

### Basic info

```sql
select
  id,
  title,
  product_type,
  created_at,
  vendor
from
  shopify_product;
```

###  List products with particular product type

```sql
select
  id,
  title,
  product_type,
  created_at
from
  shopify_product
where
  product_type = 'snowboard';
```

### List products with particular tags

```sql
select
  id,
  title,
  created_at
from
  shopify_product
where
  tags like '%Premium%';
```

### List the product created within 30 days

```sql
select
  id,
  title,
  created_at
from
  shopify_product
where
  created_at >= (created_at - interval '30' day)
order by
  created_at;
```