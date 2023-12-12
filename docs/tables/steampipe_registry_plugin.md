---
title: "Steampipe Table: steampipe_registry_plugin - Query Steampipe Registry Plugins using SQL"
description: "Allows users to query Steampipe Registry Plugins, specifically the details of each plugin, providing insights into plugin behavior, configuration, and status."
---

# Table: steampipe_registry_plugin - Query Steampipe Registry Plugins using SQL

Steampipe Registry Plugin is a component within Steampipe that allows the integration of different cloud resources and services into SQL based queries and operations. It provides a flexible way to interact with various cloud resources, including AWS, GCP, Azure, and more, using SQL. Steampipe Registry Plugin helps you gain insights into the behavior, configuration, and status of these plugins.

## Table Usage Guide

The `steampipe_registry_plugin` table provides insights into Steampipe Registry Plugins. As a DevOps engineer, explore plugin-specific details through this table, including name, description, and associated metadata. Utilize it to uncover information about plugins, such as their configuration, status, and behavior.

## Examples

### List available plugins
Explore the variety of plugins that are available within the Steampipe registry. This can help users understand the extent of their customization options and identify potential tools for enhancing their data management capabilities.

```sql+postgres
select 
  * 
from 
  steampipe_registry_plugin;
```

```sql+sqlite
select 
    * 
from 
    steampipe_registry_plugin;
```