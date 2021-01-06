---
page_title: "linuxkit_image_tar_kernel_initrd Resource - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_image_tar_kernel_initrd produces a tarball containing a kernel and initrd suitable for PXE booting.
---

# Resource `linuxkit_image_tar_kernel_initrd`

`linuxkit_image_tar_kernel_initrd` produces a tarball containing a kernel and initrd suitable for PXE booting.



## Schema

### Required

- **build** (String) The build tarball
- **destination** (String) The destination of the raw generated OS image

### Optional

- **id** (String) The ID of this resource.


