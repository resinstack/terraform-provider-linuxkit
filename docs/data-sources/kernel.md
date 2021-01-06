---
page_title: "linuxkit_kernel Data Source - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_kernel is a single kernel to be included in a linuxkit_config.
---

# Data Source `linuxkit_kernel`

`linuxkit_kernel` is a single kernel to be included in a `linuxkit_config`.



## Schema

### Required

- **image** (String) The Docker image which should contain a kernel file that will be booted

### Optional

- **binary** (String) Name of the kernel file that will be booted
- **cmdline** (String) Kernel command line options if required.
- **id** (String) The ID of this resource.
- **tar** (String) Name of tarball that unpacked into the root.
- **ucode** (String) Name of cpio archive containing CPU microcode which needs prepending to the initrd.


