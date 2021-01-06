---
page_title: "linuxkit_build Resource - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_build assembles an image based on a linuxkit_config.  It will produce a tarball that contains the complete image in an intermediate format.
---

# Resource `linuxkit_build`

`linuxkit_build` assembles an image based on a `linuxkit_config`.  It will produce a tarball that contains the complete image in an intermediate format.



## Schema

### Required

- **destination** (String) The destination of the generated OS image

### Optional

- **config** (String) The linuxkit config id
- **config_yaml** (String) The linuxkit config yaml
- **files** (List of String) The IDs of the file config
- **id** (String) The ID of this resource.
- **init** (List of String) The IDs of init containers
- **kernel** (String) The ID of the kernel resource
- **onboot** (List of String) The IDs of the onboot containers
- **onshutdown** (List of String) The IDs of the onshutdown containers
- **services** (List of String) The IDs of the service containers
- **trust** (String) The ID of the trust config
- **type** (String) Type of build, can be tar or docker


