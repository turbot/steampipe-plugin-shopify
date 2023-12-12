---
title: "Steampipe Table: shopify_customer - Query Shopify Customers using SQL"
description: "Allows users to query Customers in Shopify, specifically the customer details and their associated data, providing insights into customer behavior and preferences."
---

# Table: shopify_customer - Query Shopify Customers using SQL

Shopify is a commerce platform that allows anyone to set up an online store and sell their products. It is also used by merchants to sell products in person with Shopify POS. Shopify Customers are the group of people who have made an account or purchased from the Shopify store, their data includes personal information, purchase history, and interaction with the store.

## Table Usage Guide

The `shopify_customer` table provides insights into customers within Shopify. As a store owner or a sales analyst, explore customer-specific details through this table, including personal information, purchase history, and interaction with the store. Utilize it to uncover information about customers, such as their preferences, buying behavior, and interaction with the store.

## Examples

### Basic info
Explore customer details to understand their contact information, which can be useful for customer service or marketing purposes. This query can help identify key customer data, enhancing customer relationships and driving personalized engagement strategies.

```sql+postgres
select
  id,
  email,
  first_name,
  last_name,
  phone
from
  shopify_customer;
```

```sql+sqlite
select
  id,
  email,
  first_name,
  last_name,
  phone
from
  shopify_customer;
```

### Get the total number of orders placed and the total amount spent by a customer
Explore the purchasing behavior of a specific customer by determining the total number of purchases made and the overall amount spent. This can aid in understanding customer loyalty and spending habits, which is crucial for targeted marketing and sales strategies.

```sql+postgres
select
  id,
  email,
  first_name,
  last_name,
  phone,
  orders_count,
  total_spent
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

```sql+sqlite
select
  id,
  email,
  first_name,
  last_name,
  phone,
  orders_count,
  total_spent
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

### Get the latest order details of a customer
Explore the most recent order made by a specific customer to gain insights into their purchase history and preferences. This can aid in personalizing their shopping experience and improving customer retention strategies.

```sql+postgres
select
  id,
  email,
  first_name,
  last_name,
  last_order_id,
  last_order_name
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

```sql+sqlite
select
  id,
  email,
  first_name,
  last_name,
  last_order_id,
  last_order_name
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

### Get the details of the customer who spent the most
Discover the customer who has made the highest total purchases. This can be useful for identifying your most valuable customer for targeted marketing or reward programs.

```sql+postgres
select 
  id,
  email,
  phone,
  first_name,
  last_name,
  total_spent
from
  shopify_customer
where
  total_spent = (select max(total_spent) from shopify_customer);
```

```sql+sqlite
select 
  id,
  email,
  phone,
  first_name,
  last_name,
  total_spent
from
  shopify_customer
where
  total_spent = (select max(total_spent) from shopify_customer);
```

### Get the default address of a customer
Discover the default address linked to a specific customer's email. This is useful for verifying the primary location tied to a customer's account.

```sql+postgres
select
  id,
  email,
  jsonb_pretty(default_address) as default_address
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

```sql+sqlite
select
  id,
  email,
  default_address
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

### Get all the addresses for a customer
Discover the segments that contain all the addresses associated with a specific customer's email. This can be useful for businesses looking to understand customer location distribution or for customer service purposes.

```sql+postgres
select
  id,
  email,
  jsonb_pretty(addresses) as addresses
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

```sql+sqlite
select
  id,
  email,
  addresses
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

### List customers with unverified emails
Pinpoint the specific customers who have yet to verify their email addresses. This is beneficial for sending reminders or follow-ups to increase the rate of verified customers.

```sql+postgres
select
  id,
  email,
  phone,
  verified_email
from
  shopify_customer
where
  not verified_email;
```

```sql+sqlite
select
  id,
  email,
  phone,
  verified_email
from
  shopify_customer
where
  verified_email = 0;
```

### List customers who opted-in for email marketing.
Explore which customers have chosen to receive email marketing, allowing for targeted communication and promotional strategies. This can aid in understanding customer preferences and enhancing marketing efforts.

```sql+postgres
select
  id,
  email,
  phone,
  first_name,
  last_name,
  accepts_marketing
from
  shopify_customer
where
  accepts_marketing;
```

```sql+sqlite
select
  id,
  email,
  phone,
  first_name,
  last_name,
  accepts_marketing
from
  shopify_customer
where
  accepts_marketing;
```

### List tax exempted customers
Discover the segments of customers who have been granted tax exemption. This can be useful for understanding the demographics of customers who may be more likely to make purchases due to their tax exempt status.

```sql+postgres
select
  id,
  email,
  first_name,
  last_name,
  phone,
  tax_exempt
from
  shopify_customer
where
  tax_exempt;
```

```sql+sqlite
select
  id,
  email,
  first_name,
  last_name,
  phone,
  tax_exempt
from
  shopify_customer
where
  tax_exempt = 1;
```

### List disabled customers
Explore which customers have been disabled in your Shopify store. This can be useful for understanding customer engagement, identifying potential issues or errors, and maintaining accurate customer records.

```sql+postgres
select
  id,
  email,
  first_name,
  last_name,
  phone,
  state
from
  shopify_customer
where
  state = 'disabled';
```

```sql+sqlite
select
  id,
  email,
  first_name,
  last_name,
  phone,
  state
from
  shopify_customer
where
  state = 'disabled';
```

### List customers created within the last 30 days
Discover the segments that have newly joined your customer base in the past month. This can help in tailoring new marketing strategies or promotional offers to engage them effectively.

```sql+postgres
select
  id,
  email,
  first_name,
  last_name,
  phone,
  created_at
from
  shopify_customer
where
  created_at >= now() - interval '30' day
order by
  created_at;
```

```sql+sqlite
select
  id,
  email,
  first_name,
  last_name,
  phone,
  created_at
from
  shopify_customer
where
  created_at >= datetime('now', '-30 day')
order by
  created_at;
```