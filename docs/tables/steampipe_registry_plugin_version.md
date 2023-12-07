---
title: "Steampipe Table: steampipe_registry_plugin_version - Query Steampipe Plugin Versions using SQL"
description: "Allows users to query Steampipe Plugin Versions, specifically version details of the Steampipe plugins, providing insights into plugin versions and their metadata."
---

# Table: steampipe_registry_plugin_version - Query Steampipe Plugin Versions using SQL

Steampipe is a open source database for cloud infrastructure. It is designed to make it easy to query cloud resources as if they were tables in a SQL database. The steampipe_registry_plugin_version resource provides information about the versions of the Steampipe plugins.

## Table Usage Guide

The `steampipe_registry_plugin_version` table provides insights into the versions of Steampipe plugins. As a developer or cloud engineer, explore version-specific details through this table, including version number, release date, and associated metadata. Utilize it to keep track of the plugin versions you are using, and to decide when to update to a newer version.

## Examples

### List available plugin versions
Explore the different versions of available plugins in a chronological order to better manage updates and ensure compatibility with your system.

```sql+postgres
select 
  * 
from 
  steampipe_registry_plugin_version
order by
  name asc,
  create_time desc;
```

```sql+sqlite
select 
  * 
from 
  steampipe_registry_plugin_version
order by
  name asc,
  create_time desc;
```



### List latest available plugin versions
Discover the most recent versions of plugins available in the Steampipe registry. This information can be useful when planning updates or investigating compatibility issues.

```sql+postgres
select 
    * 
from 
    steampipe_registry_plugin_version
where
    tags ? 'latest'
order by
    name asc;
```

```sql+sqlite
Error: The corresponding SQLite query is unavailable.
```