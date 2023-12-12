---
title: "Steampipe Table: shopify_product - Query Shopify Products using SQL"
description: "Allows users to query Products in Shopify, specifically the product details, providing insights into product inventory, price, and other related information."
---

# Table: shopify_product - Query Shopify Products using SQL

Shopify is an e-commerce platform that allows businesses to create online stores and sell their products. It provides a variety of tools for managing products, inventory, payments, and shipping. Shopify's Product resource represents individual items that a store's customers can purchase.

## Table Usage Guide

The `shopify_product` table provides insights into products within Shopify. As a store manager or business analyst, explore product-specific details through this table, including pricing, inventory, and associated metadata. Utilize it to uncover information about products, such as their availability status, variant details, and the verification of product details.

## Examples

### Basic info
Explore the basic attributes of your Shopify products such as their ID, title, type, creation date, and vendor. This query is beneficial in providing a quick overview of your product catalog.

```sql+postgres
select
  id,
  title,
  product_type,
  created_at,
  vendor
from
  shopify_product;
```

```sql+sqlite
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
Explore the variety of items in your Shopify store by assessing the total count of each product type. This can aid in inventory management and understanding your product diversity.

```sql+postgres
select
  product_type,
  count(*) as product_count
from
  shopify_product
group by 
  product_type;
```

```sql+sqlite
select
  product_type,
  count(*) as product_count
from
  shopify_product
group by 
  product_type;
```

### List products with a specific tag
Explore which products in your Shopify store are labeled as 'Premium'. This can help you identify your high-end offerings and analyze their performance.

```sql+postgres
select
  id,
  title,
  created_at
from
  shopify_product
where
  tags like '%Premium%';
```

```sql+sqlite
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
Explore which products have been added to your Shopify store in the past month. This is useful for tracking inventory updates and identifying recent additions to your product catalog.

```sql+postgres
select
  id,
  title,
  created_at
from
  shopify_product
where
  created_at >= now() - interval '30' day
order by
  created_at;
```

```sql+sqlite
select
  id,
  title,
  created_at
from
  shopify_product
where
  created_at >= datetime('now', '-30 day')
order by
  created_at;
```

### List archived products
Discover the segments that contain archived products in your Shopify store. This is beneficial for assessing inventory management and identifying products that are no longer active.

```sql+postgres
select
  id,
  title,
  created_at
from
  shopify_product
where
  status = 'archived';
```

```sql+sqlite
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
Discover the segments that have product variants with low stock levels. This can help businesses to plan for restocking and prevent potential sales losses due to unavailability of popular products.

```sql+postgres
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

```sql+sqlite
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
Explore which product variants require shipping, helping you to plan logistics and shipping costs more effectively. Similarly, determine the product variants that are taxable, providing crucial information for accurate financial planning and tax compliance.

```sql+postgres
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

```sql+sqlite
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

```sql+postgres
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

```sql+sqlite
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

### Get the best selling product of the last month
Determine the top-selling product over the past month for strategic business insights. This aids in understanding consumer behavior, product performance, and can guide inventory management and marketing efforts.

```sql+postgres
select
  p.id,
  p.title,
  p.product_type,
  p.created_at,
  p.vendor,
  q.c as sales_count
from
  shopify_product as p
  join (
    select
      item ->> 'product_id' as id,
      count(*) as c,
      min(created_at) AS order_date
    from
      shopify_order,
      jsonb_array_elements(line_items) as item
    where 
      created_at >= (CURRENT_DATE - interval '30' day)
    group by
      item ->> 'product_id'
    order by
      c desc,
      order_date
    limit 1
  ) as q on p.id = q.id::bigint;
```

```sql+sqlite
select
  p.id,
  p.title,
  p.product_type,
  p.created_at,
  p.vendor,
  q.c as sales_count
from
  shopify_product as p
  join (
    select
      json_extract(item.value, '$.product_id') as id,
      count(*) as c,
      min(created_at) AS order_date
    from
      shopify_order,
      json_each(line_items) as item
    where 
      date(created_at) >= date(julianday('now') - 30)
    group by
      json_extract(item.value, '$.product_id')
    order by
      c desc,
      order_date
    limit 1
  ) as q on p.id = q.id;
```