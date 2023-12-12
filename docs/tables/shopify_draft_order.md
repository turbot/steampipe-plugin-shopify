---
title: "Steampipe Table: shopify_draft_order - Query Shopify Draft Orders using SQL"
description: "Allows users to query Draft Orders in Shopify, specifically providing detailed information about the orders that are in draft status, which can aid in understanding order patterns and potential discrepancies."
---

# Table: shopify_draft_order - Query Shopify Draft Orders using SQL

A Draft Order in Shopify is a customizable order that allows merchants to create items and orders for customers that can be modified before purchase. It provides a way to create orders for customers that can be customized and updated before they are finalized. Draft Orders are used for orders that are not placed through the usual online checkout process, such as orders taken over the phone or for custom products.

## Table Usage Guide

The `shopify_draft_order` table provides detailed insights into Draft Orders within Shopify. As an e-commerce manager or data analyst, explore draft order-specific details through this table, including customer information, line items, and associated metadata. Utilize it to uncover information about draft orders, such as those that have been pending for a long time, the common products in draft orders, and the verification of customer details.

## Examples

### Basic info
Explore which draft orders exist in your Shopify store, including customer details and shipping information. This can help you understand your pending transactions and plan your inventory and shipping strategies accordingly.

```sql+postgres
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

```sql+sqlite
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
Determine the open draft orders which have incorporated taxes into the product pricing. This is particularly useful for financial analysis and tax auditing purposes.

```sql+postgres
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

```sql+sqlite
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
  and taxes_included = 1;
```

### Find all draft orders with a specific tag
Discover the segments that contain draft orders tagged with discounts. This is beneficial for identifying potential sales opportunities or for tracking promotional campaigns.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that contain a specific product in their draft orders, allowing you to analyze customer purchase behavior and product popularity. This information can be used to tailor marketing strategies and optimize inventory management.

```sql+postgres
select
  id,
  name,
  email,
  customer,
  li ->> 'name' as product_name,
  li ->> 'price' as product_price,
  li ->> 'product_id' as product_id,
  billing_address
from
  shopify_draft_order,
  jsonb_array_elements(line_items) as li
where
  li ->> 'product_id' = '8264171716903';
```

```sql+sqlite
select
  id,
  name,
  email,
  customer,
  json_extract(li.value, '$.name') as product_name,
  json_extract(li.value, '$.price') as product_price,
  json_extract(li.value, '$.product_id') as product_id,
  billing_address
from
  shopify_draft_order,
  json_each(line_items) as li
where
  json_extract(li.value, '$.product_id') = '8264171716903';
```

### List all the draft orders from a particular city
Discover the segments that contain draft orders from a particular city. This is useful for businesses that want to analyze their order patterns geographically, specifically focusing on areas with pending transactions.

```sql+postgres
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

```sql+sqlite
select
  id,
  name,
  email,
  shipping_address
from
  shopify_draft_order
where
  json_extract(shipping_address, '$.city') = 'Toronto';
```

### List all draft orders for a particular customer
Explore all pending orders associated with a specific customer to keep track of their purchase history and manage their orders effectively. This can be particularly useful for businesses seeking to enhance their customer service and ensure timely order processing.

```sql+postgres
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

```sql+sqlite
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
  json_extract(customer, '$.first_name') = 'Karine'
  and json_extract(customer, '$.last_name') = 'Ruby';
```

### List the draft orders that have total tax greater than 100
Determine the areas in which draft orders have a tax amount exceeding 100. This can be useful for identifying potential high-value transactions or regions with higher tax rates.

```sql+postgres
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

```sql+sqlite
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
  CAST(total_tax AS REAL) > 100;
```