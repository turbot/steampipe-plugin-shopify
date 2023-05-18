# Table: shopify_product

Shopify Products Variant can be added to a Product resource to represent one version of a product with several options. The Product resource will have a variant for every possible combination of its options.

## Examples

### Basic info

```sql
select
  id,
  product_id,
  title,
  inventory_item_id
from
  shopify_product_variant;
```

### Retrieve all products where the weight is greater than 5.5 pounds

```sql
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

### Retrieve all products that have a price less than $50 and are not taxable

```sql
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

### Retrieve all products that have inventory management enabled and have less than 10 items in stock

```sql
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

```sql
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

```sql
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

```sql
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