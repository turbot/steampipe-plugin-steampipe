---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/steampipe.svg"
brand_color: "#a42a2d"
display_name: Steampipe
short_name: steampipe
description: Steampipe plugin for querying Steampipe components, such as the available plugins in the steampipe hub.
og_description: "Query Steampipe with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/steampipe-social-graphic.png"
---

# Steampipe

The Steampipe plugin is used to query Steampipe components, such as the available plugins in the Steampipe hub.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

This plugin currently requires no authentication as it only queries public endpoints.

For example:

```sql
select
  name,
  update_time
from
  steampipe_registry_plugin
```

```
+--------------------+---------------------+
| name               | update_time         |
+--------------------+---------------------+
| turbot/alicloud    | 2021-09-13 16:18:49 |
| theapsgroup/gitlab | 2021-09-16 12:22:14 |
| turbot/aws         | 2021-10-12 12:45:35 |
| turbot/jira        | 2021-09-22 18:59:46 |
| turbot/steampipe   | 2021-07-22 20:36:14 |
+--------------------+---------------------+
```

## Documentation

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/turbot/steampipe/tables)**

## Get started

### Install

Download and install the latest Steampipe plugin:

```bash
steampipe plugin install steampipe
```

### Configuration

No configuration needed.

Installing the latest steampipe plugin will create a config file (`~/.steampipe/config/steampipe.spc`) with a single connection named `steampipe`:

```hcl
connection "steampipe" {
  plugin = "steampipe"

  # token - No token/creds required.
}
```

## Get Involved

* Open source: https://github.com/turbot/steampipe-plugin-steampipe
* Community: [Slack Channel](https://steampipe.io/community/join)
