## v0.9.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#41](https://github.com/turbot/steampipe-plugin-steampipe/pull/41))
- Recompiled plugin with Go version `1.21`. ([#41](https://github.com/turbot/steampipe-plugin-steampipe/pull/41))

## v0.8.0 [2023-06-20]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.0](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.5.0/CHANGELOG.md#v550-2023-06-16) which significantly reduces API calls and boosts query performance, resulting in faster data retrieval. ([#37](https://github.com/turbot/steampipe-plugin-steampipe/pull/37))

## v0.7.0 [2023-04-06]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#35](https://github.com/turbot/steampipe-plugin-steampipe/pull/35))

## v0.6.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#33](https://github.com/turbot/steampipe-plugin-steampipe/pull/33))
- Recompiled plugin with Go version `1.19`. ([#33](https://github.com/turbot/steampipe-plugin-steampipe/pull/33))

## v0.5.0 [2022-07-22]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v332--2022-07-11) which includes several caching fixes. ([#31](https://github.com/turbot/steampipe-plugin-steampipe/pull/31))

## v0.4.0 [2022-06-24]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v330--2022-6-22). ([#28](https://github.com/turbot/steampipe-plugin-steampipe/pull/28))

## v0.3.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#27](https://github.com/turbot/steampipe-plugin-steampipe/pull/27))

## v0.3.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#25](https://github.com/turbot/steampipe-plugin-steampipe/pull/25))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#24](https://github.com/turbot/steampipe-plugin-steampipe/pull/24))

## v0.2.0 [2021-12-16]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#20](https://github.com/turbot/steampipe-plugin-steampipe/pull/20))
- Recompiled plugin with Go version 1.17 ([#20](https://github.com/turbot/steampipe-plugin-steampipe/pull/20))

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
