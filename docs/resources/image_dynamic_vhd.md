---
page_title: "linuxkit_image_dynamic_vhd Resource - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_dynamic_vhd produces a sparse VHD from a build tarball.
---

# Resource `linuxkit_image_dynamic_vhd`

`linuxkit_dynamic_vhd` produces a sparse VHD from a build tarball.



## Schema

### Required

- **build** (String) The build tarball
- **destination** (String) The destination of the raw generated OS image

### Optional

- **id** (String) The ID of this resource.


