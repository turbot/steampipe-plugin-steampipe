## v0.1.3 [2021-07-22]

_Enhancements_

- Updated: Add columns `image_annotations` and `image_layers` to `steampipe_registry_plugin_version` table ([#540](https://github.com/turbot/steampipe-plugin-aws/pull/540))
- Recompiled plugin with [steampipe-plugin-sdk v1.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v141--2021-07-20) ([#540](https://github.com/turbot/steampipe-plugin-aws/pull/540))

## v0.1.2 [2021-03-19]

_Bug fixes_

- Use API paging to get full results for `steampipe_registry_plugin` and `steampipe_registry_plugin_version`.


## v0.1.1 [2021-02-25]

_Bug fixes_

- Update to display the version of the plugin.
- Recompiled plugin with latest [steampipe-plugin-sdk](https://github.com/turbot/steampipe-plugin-sdk) to resolve SDK issues:
  - Fix error for missing required quals [#40](https://github.com/turbot/steampipe-plugin-sdk/issues/42).
  - Queries fail with error socket: too many open files [#190](https://github.com/turbot/steampipe/issues/190)
