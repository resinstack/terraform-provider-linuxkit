---
page_title: "linuxkit_file Data Source - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_file is a file you would like to include in the finished system image.
---

# Data Source `linuxkit_file`

`linuxkit_file` is a file you would like to include in the finished system image.



## Schema

### Required

- **path** (String) The path to create the file or directory

### Optional

- **contents** (String) The contents of the file
- **directory** (Boolean) If true a directory is created
- **gid** (Number) The gid of the file/directory owner
- **gid_name** (String) The gid name of the file/directory owner
- **id** (String) The ID of this resource.
- **metadata** (String) Format to write image metadata, only yaml is currently supported
- **mode** (String) The mode to create the file or directory
- **optional** (Boolean) File is optional, dont fail if source does not exist
- **source** (String) The path to the source of the file
- **symlink** (String) The path to link to
- **uid** (Number) The uid of the file/directory owner
- **uid_name** (String) The uid name of the file/directory owner


