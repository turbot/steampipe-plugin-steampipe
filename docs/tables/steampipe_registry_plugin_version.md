# Table: steampipe_registry_plugin_version

The steampipe_registry_plugin_version table provides information about the versions of plugins that are available in the public Steampipe hub registry.

## Examples

### List available plugin versions

```sql
select 
    * 
from 
    steampipe_registry_plugin_version
order by
    name asc,
    create_time desc;
```



### List latest available plugin versions

```sql
select 
    * 
from 
    steampipe_registry_plugin_version
where
    tags ? 'latest'
order by
    name asc;
```
