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

### Count how many items of a specific type are there

```sql
select
  product_type,
  count(*) as product_count
from
  shopify_product
group by 
  product_type;
```

### List products with a specific tag

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

### List products created within the last 30 days

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
  p.id as product_id,
  p.title as product_title,
  v.inventory_item_id as variant_inventory_item_id,
  v.inventory_quantity as variant_inventory_quantity
from
  shopify_product as p,
  shopify_product_variant as v
where
  v.inventory_quantity < 20;
```

### List the product variants which require shipping

```sql
select
  p.id as product_id,
  p.title as product_title,
  v.inventory_item_id as variant_inventory_item_id,
  v.requires_shipping as requires_shipping
from
  shopify_product as p,
  shopify_product_variant as v
where
  v.requires_shipping;
```

## List the product variants which are taxable

```sql
select
  p.id as product_id,
  p.title as product_title,
  v.inventory_item_id as variant_inventory_item_id,
  v.taxable as taxable,
  v.tax_code as tax_code
from
  shopify_product as p,
  shopify_product_variant as v
where
  v.taxable;
```

### Get the best selling product

```sql
select
  p.id,
  p.title,
  p.product_type,
  p.created_at,
  p.vendor
  q.c as sales_count
from
  shopify_product as p
  join (
    select
      item ->> 'product_id' as id,
      count(*) as c
    from
      shopify_order,
      jsonb_array_elements(line_items) as item
    group by
      item ->> 'product_id'
    order by
      c desc
    limit 1
  ) as q on p.id = q.id::bigint;
```