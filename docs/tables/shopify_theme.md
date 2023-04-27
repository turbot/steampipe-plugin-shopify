# Table: shopify_theme

Shopify themes are pre-designed website templates that allow you to easily customize the look and feel of your online store without requiring advanced web development skills.

## Examples

### Basic info

```sql
select
  id,
  name,
  previewable,
  role
from
  shopify_theme;
```

### Get the names and IDs of all themes that are previewable

```sql
select
  id,
  name
from
  shopify_theme
where
  previewable = true;
```

### Get the name and creation date of the oldest processing theme

```sql
select 
  name,
  created_at
from 
  shopify_theme
where
  processing = true
order by
  created_at asc
limit 1;
```

### Get the themes that were last updated within the past 30 days

```sql
select
  name,
  id,
  updated_at
from
  shopify_theme
where
  updated_at >= now() - interval '30 days';
```
