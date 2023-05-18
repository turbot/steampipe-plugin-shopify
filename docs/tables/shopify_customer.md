# Table: shopify_customer

Shopify Customer stores information about a shop's customers, such as their contact details, their order history, and whether they've agreed to receive email marketing.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

```sql
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
  total_spent = (
  select
    max(total_spent)
  from
    shopify_customer
  );
```

### Get the default address for a customer

```sql
select
  id,
  email,
  jsonb_pretty(default_address) as default_address
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

### Get all the addresses for a customer

```sql
select
  id,
  email,
  jsonb_pretty(addresses) as addresses
from
  shopify_customer
where
  email = 'russel.winfield@example.com';
```

### List customers with unverified emails

```sql
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

### List customers who opted-in for email marketing.

```sql
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

```sql
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

### List disabled customers

```sql
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

```sql
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
  created_at >= (created_at - interval '30' day)
order by
  created_at;
```
