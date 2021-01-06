---
page_title: "linuxkit_image_kernel_initrd Resource - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_image_kernel_initrd produces a kernel and initrd suitable for PXE booting.  If you want an archive of the output, use linuxkit_image_tar_kernel_initrd instead.
---

# Resource `linuxkit_image_kernel_initrd`

`linuxkit_image_kernel_initrd` produces a kernel and initrd suitable for PXE booting.  If you want an archive of the output, use `linuxkit_image_tar_kernel_initrd` instead.



## Schema

### Required

- **build** (String) The build tarball
- **destination** (Block Set, Min: 1) (see [below for nested schema](#nestedblock--destination))

### Optional

- **id** (String) The ID of this resource.

<a id="nestedblock--destination"></a>
### Nested Schema for `destination`

Required:

- **cmdline** (String) The destination of the generated cmdline
- **initrd** (String) The destination of the generated initrd
- **kernel** (String) The destination of the generated kernel


