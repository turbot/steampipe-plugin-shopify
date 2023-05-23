# Table : shopify_draft_order

A draft order is an order a merchant makes in the Shopify admin on behalf of a customer.

## Examples

### Basic info

```sql
select
  id,
  name,
  email,
  customer,
  billing_address,
  shipping_address
from
  shopify_draft_order;
```

### List all draft orders that have a status of "open" and include taxes in the product price

```sql
select
  id,
  name,
  email,
  customer,
  billing_address,
  shipping_address,
  status,
  taxes_included
from
  shopify_draft_order
where
  status = 'open'
  and taxes_included = true;
```

### Find all draft orders with a specific tag

```sql
select
  id,
  name,
  email,
  customer,
  billing_address,
  shipping_address,
  status,
  tags,
  taxes_included
from
  shopify_draft_order
where
  tags like '%Discount%';
```

### Find all draft orders that contain a specific product in their line items

```sql
select
  id,
  name,
  email,
  customer,
  li ->> 'name' as product_name,
  li ->> 'price' as product_price,
  li -> 'product_id' as product_id,
  billing_address
from
  shopify_draft_order,
  jsonb_array_elements(line_items) as li
where
  li -> 'product_id' = '8264171716903';
```

### List all the draft orders from a particular city

```sql
select
  id,
  name,
  email,
  shipping_address
from
  shopify_draft_order
where
  shipping_address ->> 'city' = 'Toronto';
```

### List all draft orders for a particular customer

```sql
select
  id,
  name,
  email,
  customer,
  billing_address,
  shipping_address
from
  shopify_draft_order
where
  customer ->> 'first_name' = 'Karine'
  and customer ->> 'last_name' = 'Ruby';
```

### List the draft orders that have total tax greater than value 100

```sql
select
  id,
  name,
  email,
  customer,
  billing_address,
  shipping_address,
  total_tax
from
  shopify_draft_order
where
  (total_tax)::numeric > 100;
```