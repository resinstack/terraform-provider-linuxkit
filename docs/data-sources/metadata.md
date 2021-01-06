---
page_title: "linuxkit_metadata Data Source - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_metadata data source can be used to convert a directory structure rooted at base_path and return it as a JSON formatted string.  Be aware that while linuxkit will parse an arbitrary size of metadata, most cloud providers limit this data.
---

# Data Source `linuxkit_metadata`

`linuxkit_metadata` data source can be used to convert a directory structure rooted at `base_path` and return it as a JSON formatted string.  Be aware that while linuxkit will parse an arbitrary size of metadata, most cloud providers limit this data.



## Schema

### Required

- **base_path** (String) The base path to use when building an archive.

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **json** (String) The packed version of the input file tree


