---
page_title: "linuxkit_config Data Source - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_config represents a complete configuration for a full machine.  It unifies the kernel, init, onboot, and service components into a single configuration to be built.
---

# Data Source `linuxkit_config`

`linuxkit_config` represents a complete configuration for a full machine.  It unifies the kernel, init, onboot, and service components into a single configuration to be built.



## Schema

### Required

- **init** (List of String) The IDs of init containers
- **kernel** (String) The ID of the kernel resource

### Optional

- **files** (List of String) The IDs of the file config
- **id** (String) The ID of this resource.
- **onboot** (List of String) The IDs of the onboot containers
- **onshutdown** (List of String) The IDs of the onshutdown containers
- **services** (List of String) The IDs of the service containers

### Read-only

- **yaml** (String) The rendered yaml of the config


