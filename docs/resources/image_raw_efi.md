---
page_title: "linuxkit_image_raw_efi Resource - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_image_raw_efi produces a GPT image file suitable for booting an UEFI enabled system.
---

# Resource `linuxkit_image_raw_efi`

`linuxkit_image_raw_efi` produces a GPT image file suitable for booting an UEFI enabled system.



## Schema

### Required

- **build** (String) The build tarball
- **destination** (String) The destination of the raw generated OS image

### Optional

- **id** (String) The ID of this resource.


