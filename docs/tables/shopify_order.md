# Table: shopify_product

Shopify Orders record all the orders that customers have made in the Shopify store.

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
  shopify_order;
```

### List all fulfilled orders shipped to a specific postal code

```sql
select
  id,
  name,
  email,
  shipping_address ->> 'zip' as zip_code,
  fulfillment_status
from
  shopify_order
where
  fulfillment_status = 'fulfilled'
  and
  shipping_address ->> 'zip' = '712136';
```

### List the total price of each order where at least one item has a price greater than $100

```sql
select
  id,
  name,
  email,
  total_price
from
  shopify_order,
  jsonb_array_elements(line_items) as item
where
  (item->>'price')::numeric > 100;
```

### List number of items in each order

```sql
select
  id,
  name,
  email,
  jsonb_array_length(line_items) as number_of_items
from
  shopify_order;
```

### List all orders where the customer's email belong to the domain "gmail.com"

```sql
select
  id,
  name,
  email,
  customer->>'name' as customer_name
from
  shopify_order
where
  customer->>'email' like '%@gmail.com';
```

### List the orders cancelled within last 30 days

```sql
select
  id,
  name,
  email,
  cancel_reason,
  cancelled_at
from
  shopify_order
where
  cancelled_at >= (cancelled_at - interval '30' day)
order by
  cancelled_at;
```

### List pending or partially paid orders

```sql
select
  id,
  name,
  email,
  financial_status
from
  shopify_order
where
  financial_status in ('pending','partially_paid');
```

### Get the order details of refunded orders

```sql
select
  id,
  name,
  email,
  jsonb_array_elements(refund -> 'transactions') ->> 'status' as refund_status,
  jsonb_array_elements(refund -> 'transactions') ->> 'amount' as refund_amount,
  coalesce(jsonb_array_length(refund -> 'refund_line_items'),0) as number_of_products_refunded
from
  shopify_order,
  jsonb_array_elements(refunds) as refund;
```

### Get the orders processed on a particular date

```sql
select
  id,
  name,
  email,
  cancel_reason,
  processed_at
from
  shopify_order
where
  processed_at::date = '2023-04-25';
```

### List the orders which are fulfilled but receipts are not send

```sql
select
  id,
  name,
  email,
  fulfillment_status,
  send_fulfillment_receipt
from
  shopify_order
where
  fulfillment_status = 'fulfilled'
  and
  not send_fulfillment_receipt;
```

### Count number of orders paid manually

```sql
select
  count(*)
from
  shopify_order
where
  gateway = 'manual';
```

### Get the tax details of the products ordered

```sql
select
  id as order_id,
  name as order_name,
  email,
  item ->> 'product_id' as product_id,
  item ->> 'price' as product_price,
  jsonb_array_elements(item -> 'tax_lines') ->> 'rate' as tax_rate,
  jsonb_array_elements(item -> 'tax_lines') ->> 'title' as tax_type,
  jsonb_array_elements(item -> 'tax_lines') ->> 'price' as tax_price
from
  shopify_order,
  jsonb_array_elements(line_items) as item;
```

### List the orders with discounts

```sql
select
  id,
  name,
  email,
  total_discounts
from
  shopify_order
where
  total_discounts > 0;
```
