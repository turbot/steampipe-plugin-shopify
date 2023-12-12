---
title: "Steampipe Table: shopify_order - Query Shopify Orders using SQL"
description: "Allows users to query Orders from Shopify, specifically the order details, providing insights into sales performance, customer purchasing habits, and inventory management."
---

# Table: shopify_order - Query Shopify Orders using SQL

Shopify is a widely-used, e-commerce platform that allows businesses to set up an online store and sell their products. It provides an array of features such as product inventory management, order management, and customer relationship management. An Order in Shopify is a record of a sale made through your store, it contains information about the products, customer details, and shipping.

## Table Usage Guide

The `shopify_order` table provides insights into orders made within a Shopify store. As a store manager or a business analyst, explore order-specific details through this table, including customer information, product details, and shipping details. Utilize it to analyze sales performance, understand customer purchasing habits, and manage inventory effectively.

## Examples

### Basic info
Discover the segments that include customer details and their associated shipping and billing addresses. This can help to gain insights into customer location data for potential marketing strategies or logistical planning.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have successfully completed orders delivered to a specific area. This can be useful in understanding customer distribution and analyzing sales performance in targeted locations.

```sql+postgres
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
  and shipping_address ->> 'zip' = '712136';
```

```sql+sqlite
select
  id,
  name,
  email,
  json_extract(shipping_address, '$.zip') as zip_code,
  fulfillment_status
from
  shopify_order
where
  fulfillment_status = 'fulfilled'
  and json_extract(shipping_address, '$.zip') = '712136';
```

### List the total price of each order where at least one item has a price greater than $100
Assess the total costs of orders containing items priced over $100 to gain insights into high-value transactions. This can aid in identifying potential revenue streams and understanding customer purchasing habits.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  email,
  total_price
from
  shopify_order,
  json_each(line_items) as item
where
  CAST(json_extract(item.value, '$.price') AS REAL) > 100;
```

### List number of items in each order
Analyze the settings to understand the quantity of products in each customer order. This is useful for inventory management and understanding customer purchasing behavior.

```sql+postgres
select
  id,
  name,
  email,
  jsonb_array_length(line_items) as number_of_items
from
  shopify_order;
```

```sql+sqlite
select
  id,
  name,
  email,
  json_array_length(line_items) as number_of_items
from
  shopify_order;
```

### List all orders where the customer's email belong to the domain "gmail.com"
Explore which orders were made by customers using a Gmail account. This is useful for understanding customer demographics and tailoring marketing efforts towards specific email service users.

```sql+postgres
select
  id,
  name,
  email,
  ((customer->>'first_name') || ' ' ||  (customer->>'last_name')) as customer_name
from
  shopify_order
where
  customer->>'email' like '%@gmail.com';
```

```sql+sqlite
select
  id,
  name,
  email,
  (json_extract(customer, '$.first_name') || ' ' ||  json_extract(customer, '$.last_name')) as customer_name
from
  shopify_order
where
  json_extract(customer, '$.email') like '%@gmail.com';
```

### List the orders cancelled within last 30 days
Discover the instances of order cancellations in the past month. This helps in analyzing the reasons for cancellation and aids in making informed decisions to reduce such instances in the future.

```sql+postgres
select
  id,
  name,
  email,
  cancel_reason,
  cancelled_at
from
  shopify_order
where
  cancelled_at >= now() - interval '30' day
order by
  cancelled_at;
```

```sql+sqlite
select
  id,
  name,
  email,
  cancel_reason,
  cancelled_at
from
  shopify_order
where
  cancelled_at >= datetime('now', '-30 day')
order by
  cancelled_at;
```

### List pending or partially paid orders
Determine the status of your online store's orders to identify those that are pending or partially paid. This could be useful for tracking incomplete transactions and improving your revenue collection process.

```sql+postgres
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

```sql+sqlite
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
Explore which orders have been refunded, including the status of each refund and the total amount refunded. This can be useful for gaining insights into refund patterns and tracking customer satisfaction.

```sql+postgres
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

```sql+sqlite
select
  shop.id,
  shop.name,
  shop.email,
  json_extract(ref.value, '$.status') as refund_status,
  json_extract(ref.value, '$.amount') as refund_amount,
  coalesce(json_array_length(json_extract(shop.refunds, '$.refund_line_items')),0) as number_of_products_refunded
from
  shopify_order shop,
  json_each(shop.refunds) as ref;
```

### Get the orders processed on a particular date
Discover the segments that had their orders processed on a specific date. This is useful for tracking business performance and customer activity on a day-to-day basis.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  email,
  cancel_reason,
  processed_at
from
  shopify_order
where
  date(processed_at) = '2023-04-25';
```

### List the orders which are fulfilled but receipts are not send
Identify instances where orders have been fulfilled but the corresponding receipts have not been sent. This is useful for ensuring all customers receive their receipts and for maintaining accurate records.

```sql+postgres
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
  and not send_fulfillment_receipt;
```

```sql+sqlite
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
  and not send_fulfillment_receipt;
```

### Count number of orders paid manually
Explore the volume of orders that have been manually paid for. This can help in identifying the extent of manual interventions in the payment process and potentially streamline operations.

```sql+postgres
select
  count(*) as orders_paid_manually
from
  shopify_order
where
  gateway = 'manual';
```

```sql+sqlite
select
  count(*) as orders_paid_manually
from
  shopify_order
where
  gateway = 'manual';
```

### Get the tax details of the products ordered
Determine the tax details associated with each product order to understand the different tax types and rates applied. This can help in analyzing the tax structure for various products, facilitating better financial planning and management.

```sql+postgres
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

```sql+sqlite
select
  id as order_id,
  name as order_name,
  email,
  json_extract(item.value, '$.product_id') as product_id,
  json_extract(item.value, '$.price') as product_price,
  json_extract(tax.value, '$.rate') as tax_rate,
  json_extract(tax.value, '$.title') as tax_type,
  json_extract(tax.value, '$.price') as tax_price
from
  shopify_order,
  json_each(line_items) as item,
  json_each(json_extract(item.value, '$.tax_lines')) as tax;
```

### List the orders with discounts
Discover the segments that have been granted discounts on their orders. This is useful for analyzing the effectiveness of discount strategies and identifying popular discount trends among customers.

```sql+postgres
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

```sql+sqlite
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

### Get the most ordered product of the last month
Determine the most popular product sold in the last month to identify customer preferences and guide inventory decisions. This query is useful for businesses seeking to optimize their product offering based on recent sales trends.

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
      created_at >= date('now', '-30 day')
    group by
      json_extract(item.value, '$.product_id')
    order by
      c desc,
      order_date
    limit 1
  ) as q on p.id = q.id;
```