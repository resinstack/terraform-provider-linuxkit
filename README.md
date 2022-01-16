# terraform-provider-linuxkit

This is, as the name suggests, a terraform provider for working with
LinuxKit.  It functions as a complete frontend to the LinuxKit config
format enabling composable images and easily reused blocks.

More information can be found [in the Terraform
registry](https://registry.terraform.io/providers/resinstack/linuxkit).

## Features

- Build LinuxKit images in all the formats that LinuxKit supports
- Reuse config blocks in multiple images 

## Docs

Documentation can be found in the [docs](/docs/index.md) folder.

## Notes

Though this is built on the same libraries as LinuxKit, those
libraries can call out to the LinuxKit binary. For this reason you
will need LinuxKit and QEMU installed on the machine in order to build
some images (the AWS image for example).

This provider will also not support "pushing" images to cloud
platforms, this is because there are better providers for doing this.


---

This is a continuation of the work originally started by
[SkyBet](https://www.github.com/skybet/terraform-provider-linuxkit/).
It is the result of the merge of the ResinStack provider which deals
with metadata, and the SkyBet provider which deals with building
images.
