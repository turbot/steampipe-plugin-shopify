---
title: "Steampipe Table: shopify_product_variant - Query Shopify Product Variants using SQL"
description: "Allows users to query Shopify Product Variants, providing detailed information on different versions of products available in a Shopify store."
---

# Table: shopify_product_variant - Query Shopify Product Variants using SQL

A Shopify Product Variant represents different versions of a product in a Shopify store. These versions can differ in many aspects, such as size, color, material, or other attributes. They are crucial in providing customers with a variety of options and increasing the diversity of product offerings in a Shopify store.

## Table Usage Guide

The `shopify_product_variant` table provides insights into the product variants within a Shopify store. As a store manager or data analyst, you can explore variant-specific details through this table, including price, SKU, and inventory quantities. Utilize it to uncover information about product diversity, pricing strategies, and inventory management in your Shopify store.

## Examples

### Basic info
Explore which product variants are available in your Shopify store, helping you to assess inventory and manage product listings effectively.

```sql+postgres
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant;
```

```sql+sqlite
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant;
```

### Retrieve all products where the weight is greater than 5.5 pounds
Explore which product variants in your Shopify store weigh more than 5.5 pounds. This can be useful for determining shipping costs or identifying heavy items in your inventory.

```sql+postgres
select
  id,
  product_id,
  title,
  weight,
  weight_unit,
  inventory_item_id
from
  shopify_product_variant
where
  weight::decimal > 5.5
  and weight_unit = 'lb';
```

```sql+sqlite
select
  id,
  product_id,
  title,
  weight,
  weight_unit,
  inventory_item_id
from
  shopify_product_variant
where
  cast(weight as decimal) > 5.5
  and weight_unit = 'lb';
```

### Retrieve all products that have a price less than $50 and are not taxable
Explore which products are priced under $50 and are exempt from taxation. This can be beneficial for customers seeking affordable and tax-free options.

```sql+postgres
select
  id,
  product_id,
  title,
  weight,
  weight_unit,
  inventory_item_id
from
  shopify_product_variant
where
  price::numeric < 50
  and not taxable;
```

```sql+sqlite
select
  id,
  product_id,
  title,
  weight,
  weight_unit,
  inventory_item_id
from
  shopify_product_variant
where
  CAST(price AS REAL) < 50
  and not taxable;
```

### Retrieve all products that have inventory management enabled and have less than 10 items in stock
Explore which products are running low on stock and have inventory management enabled. This query can help you proactively manage inventory and prevent product shortages.

```sql+postgres
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant
where
  inventory_management is not null
  and inventory_quantity < 10;
```

```sql+sqlite
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant
where
  inventory_management is not null
  and inventory_quantity < 10;
```

### Get the variant with the lowest price for a specific product
Discover the variant of a specific product that offers the lowest price. This can be useful in identifying the most cost-effective option for purchasing or selling a particular product.

```sql+postgres
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant
where
  product_id = '8264171749671'
order by
  price
limit 1;
```

```sql+sqlite
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant
where
  product_id = '8264171749671'
order by
  price
limit 1;
```

### Get the variants that are out of stock
Discover the variants that are currently out of stock. This is beneficial in managing inventory and understanding which products need to be restocked.

```sql+postgres
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant
where
  inventory_quantity = 0;
```

```sql+sqlite
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant
where
  inventory_quantity = 0;
```

### Retrieve all products that have a barcode and are fulfilled by a specific service
Explore which products are fulfilled by a specific service and also have a barcode. This can be useful for tracking inventory or managing product distribution.

```sql+postgres
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant
where
  barcode is not null
  and fulfillment_service = 'my_fulfillment_service';
```

```sql+sqlite
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant
where
  barcode is not null
  and fulfillment_service = 'my_fulfillment_service';
```