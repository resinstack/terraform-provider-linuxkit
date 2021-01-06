---
page_title: "linuxkit_image_aws Resource - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_image_aws produces an image suitable for importing into AWS EC2.
---

# Resource `linuxkit_image_aws`

`linuxkit_image_aws` produces an image suitable for importing into AWS EC2.



## Schema

### Required

- **build** (String) The build tarball
- **destination** (String) The destination of the raw generated OS image
- **size** (Number) The size in megabytes of the new image

### Optional

- **id** (String) The ID of this resource.


