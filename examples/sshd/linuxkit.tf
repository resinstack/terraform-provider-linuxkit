data "linuxkit_kernel" "kernel" {
  image   = "linuxkit/kernel:5.6.11"
  cmdline = "console=tty0 console=ttyS0 console=ttyAMA0"
}

data "linuxkit_init" "init" {
  containers = [
    "linuxkit/init:v0.8",
    "linuxkit/runc:v0.8",
    "linuxkit/containerd:v0.8",
    "linuxkit/ca-certificates:v0.8",
  ]
}

data "linuxkit_image" "sysctl" {
  name  = "sysctl"
  image = "linuxkit/sysctl:v0.8"
}

data "linuxkit_image" "rngd1" {
  name    = "rngd1"
  image   = "linuxkit/rngd:v0.8"
  command = ["/sbin/rngd", "-1"]
}

data "linuxkit_image" "getty" {
  name  = "getty"
  image = "linuxkit/getty:v0.8"
  env   = ["INSECURE=true"]
}

data "linuxkit_image" "rngd" {
  name  = "rngd"
  image = "linuxkit/rngd:v0.8"
}

data "linuxkit_image" "dhcpcd" {
  name  = "dhcpcd"
  image = "linuxkit/dhcpcd:v0.8"
}

data "linuxkit_image" "sshd" {
  name  = "sshd"
  image = "linuxkit/sshd:v0.8"
}

data "linuxkit_config" "sshd" {
  kernel = data.linuxkit_kernel.kernel.id
  init   = [data.linuxkit_init.init.id]

  onboot = [
    data.linuxkit_image.sysctl.id,
    data.linuxkit_image.rngd1.id,
  ]

  services = [
    data.linuxkit_image.getty.id,
    data.linuxkit_image.rngd.id,
    data.linuxkit_image.dhcpcd.id,
    data.linuxkit_image.sshd.id,
  ]
}

resource "linuxkit_build" "sshd" {
  config_yaml = data.linuxkit_config.sshd.yaml
  destination = "${path.module}/sshd.tar"
}

resource "linuxkit_image_raw_bios" "sshd" {
  build = linuxkit_build.sshd.destination
  destination = "${path.module}/sshd.raw"
}
