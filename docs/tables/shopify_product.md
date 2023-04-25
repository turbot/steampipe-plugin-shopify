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

###  Count the number of products of a particular type 

```sql
select
  product_type,
  count(*) as product_count
from
  shopify_product
where
  product_type = 'snowboard'
group by 
  product_type;
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

### List the products created within last 30 days

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

### List archived products

```sql
select
  id,
  title,
  created_at
from
  shopify_product
where
  status = 'archived';
```

### List the product variants with quantities less than 20 

```sql
select
  id as product_id,
  title as product_title,
  v -> 'inventory_item_id' as variant_inventory_item_id,
  v -> 'inventory_quantity' as variant_inventory_quantity
from
  shopify_product,
  jsonb_array_elements(variants) as v
where 
  (v ->> 'inventory_quantity')::integer < 20;
```

### List the product variants which require shipping

```sql
select
  id as product_id,
  title as product_title,
  v -> 'inventory_item_id' as variant_inventory_item_id,
  v ->> 'requires_shipping' as requires_shipping
from
  shopify_product,
  jsonb_array_elements(variants) as v
where 
  (v ->> 'requires_shipping')::boolean;
```

## List the product variants which are taxable

```sql
select
  id as product_id,
  title as product_title,
  v -> 'inventory_item_id' as variant_inventory_item_id,
  v ->> 'taxable' as taxable
from
  shopify_product,
  jsonb_array_elements(variants) as v
where 
  (v ->> 'taxable')::boolean;
```