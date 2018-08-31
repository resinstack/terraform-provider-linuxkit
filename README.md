# terraform-provider-linuxkit

This terraform provider builds [LinuxKit](https://github.com/linuxkit/linuxkit) images in all supported formats.

## Features

- Build LinuxKit images in all the formats that LinuxKit supports
- Reuse config blocks in multiple images 

# Docs

Documentation can be found in the [docs](/docs/index.md) folder.

## Notes

Though this is built on the same libraries as LinuxKit, those libraries can call out to the LinuxKit binary. For this reason you will need LinuxKit and QEMU installed on the machine in order to build some images (the AWS image for example).

This provider will also not support "pushing" images to providers, this is because there are better providers for doing this.