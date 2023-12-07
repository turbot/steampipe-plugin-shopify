---
title: "Steampipe Table: shopify_theme - Query Shopify Themes using SQL"
description: "Allows users to query Shopify Themes, specifically the active and inactive themes in a Shopify store, providing insights into the current theme setup and potential customization options."
---

# Table: shopify_theme - Query Shopify Themes using SQL

Shopify Themes are the templates that determine the look and feel of a Shopify store. They provide the framework for all the content seen by customers when they visit the online store. Themes can be customized, and Shopify stores can have multiple themes installed, but only one can be published and active at a time.

## Table Usage Guide

The `shopify_theme` table provides insights into the themes used within a Shopify store. As a store owner or a web developer, explore theme-specific details through this table, including theme roles, theme names, and associated metadata. Utilize it to uncover information about themes, such as the currently active theme, the previewable themes, and the customization options available for each theme.

## Examples

### Basic info
Discover the segments that include the unique identifiers, names, and roles of your Shopify themes to assess their availability for preview. This can be useful in managing and organizing your store's visual presentation.

```sql+postgres
select
  id,
  name,
  previewable,
  role
from
  shopify_theme;
```

```sql+sqlite
select
  id,
  name,
  previewable,
  role
from
  shopify_theme;
```

### Get the names and IDs of all themes that are previewable
Discover the themes that are previewable to understand which themes are available for preview, helping you to make informed decisions about theme selection.

```sql+postgres
select
  id,
  name
from
  shopify_theme
where
  previewable = true;
```

```sql+sqlite
select
  id,
  name
from
  shopify_theme
where
  previewable = 1;
```

### Get the name and creation date of the oldest processing theme
Explore which theme has been processing for the longest time on your Shopify store. This can help identify potential performance issues or bottlenecks in your store's theme management.

```sql+postgres
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

```sql+sqlite
select 
  name,
  created_at
from 
  shopify_theme
where
  processing = 1
order by
  created_at asc
limit 1;
```

### Get the themes that were last updated within the past 30 days
Gain insights into the themes that have been updated recently. This query is particularly useful for keeping track of theme modifications and ensuring your website stays current and functional.

```sql+postgres
select
  name,
  id,
  updated_at
from
  shopify_theme
where
  updated_at >= now() - interval '30 days';
```

```sql+sqlite
select
  name,
  id,
  updated_at
from
  shopify_theme
where
  updated_at >= datetime('now', '-30 days');
```