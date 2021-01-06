---
page_title: "linuxkit_trust Data Source - terraform-provider-linuxkit"
subcategory: ""
description: |-
  linuxkit_trust is a set of images and organizations that should be trusted based on their docker notary signatures.
---

# Data Source `linuxkit_trust`

`linuxkit_trust` is a set of images and organizations that should be trusted based on their docker notary signatures.



## Schema

### Optional

- **id** (String) The ID of this resource.
- **image** (List of String) Images to trust
- **org** (List of String) Orgs to trust


