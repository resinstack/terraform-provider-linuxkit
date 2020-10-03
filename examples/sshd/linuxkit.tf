data "linuxkit_kernel" "kernel" {
  image   = "linuxkit/kernel:4.14.62"
  cmdline = "console=tty0 console=ttyS0 console=ttyAMA0"
}

data "linuxkit_init" "init" {
  containers = [
    "linuxkit/init:v0.6",
    "linuxkit/runc:v0.6",
    "linuxkit/containerd:v0.6",
    "linuxkit/ca-certificates:v0.6",
  ]
}

data "linuxkit_image" "sysctl" {
  name  = "sysctl"
  image = "linuxkit/sysctl:v0.6"
}

data "linuxkit_image" "rngd1" {
  name    = "rngd1"
  image   = "linuxkit/rngd:v0.6"
  command = ["/sbin/rngd", "-1"]
}

data "linuxkit_image" "getty" {
  name  = "getty"
  image = "linuxkit/getty:v0.6"
  env   = ["INSECURE=true"]
}

data "linuxkit_image" "rngd" {
  name  = "rngd"
  image = "linuxkit/rngd:v0.6"
}

data "linuxkit_image" "dhcpcd" {
  name  = "dhcpcd"
  image = "linuxkit/dhcpcd:v0.6"
}

data "linuxkit_image" "sshd" {
  name  = "sshd"
  image = "linuxkit/sshd:v0.6"
}

data "linuxkit_trust" "default" {
  org = ["linuxkit"]
}

data "linuxkit_config" "sshd" {
  kernel = "${data.linuxkit_kernel.kernel.id}"
  init   = ["${data.linuxkit_init.init.id}"]

  onboot = [
    "${data.linuxkit_image.sysctl.id}",
    "${data.linuxkit_image.rngd1.id}",
  ]

  services = [
    "${data.linuxkit_image.getty.id}",
    "${data.linuxkit_image.rngd.id}",
    "${data.linuxkit_image.dhcpcd.id}",
    "${data.linuxkit_image.sshd.id}",
  ]

  trust = "${data.linuxkit_trust.default.id}"
}

resource "linuxkit_build" "sshd" {
  config      = "${data.linuxkit_config.sshd.id}"
  destination = "${path.module}/sshd.tar"
}

resource "linuxkit_image_kernel_initrd" "sshd" {
  build = "${linuxkit_build.sshd.destination}"

  destination {
    kernel  = "${path.module}/sshd-kernel"
    initrd  = "${path.module}/sshd-initrd.img"
    cmdline = "${path.module}/sshd-cmdline"
  }
}

resource "linuxkit_image_aws" "sshd" {
  build = "${linuxkit_build.sshd.destination}"

  size        = 1024
  destination = "${path.module}/sshd-aws.raw"
}

resource "linuxkit_image_dynamic_vhd" "sshd" {
  build = "${linuxkit_build.sshd.destination}"

  destination = "${path.module}/sshd-dynamic.vhd"
}

resource "linuxkit_image_gcp" "sshd" {
  build = "${linuxkit_build.sshd.destination}"

  destination = "${path.module}/sshd-gcp.tar.gz"
}
