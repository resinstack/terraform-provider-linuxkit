# terraform-provider-linuxkit

This is, as the name suggests, a terraform provider for working with
LinuxKit.  Right now all it does is provide a data source which
formats metadata files.

To install it you can either compile it locally with `go build` or you
can obtain a precompiled release.

Once you have obtained a copy, you can install it in
`~/.terraform.d/plugins/`.

This terraform provider builds
[LinuxKit](https://github.com/linuxkit/linuxkit) images in all
supported formats.

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
