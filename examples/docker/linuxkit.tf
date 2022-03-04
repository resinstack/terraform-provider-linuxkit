terraform {
  required_providers {
    linuxkit = {
      source  = "resinstack/linuxkit"
    }
  }
}

data "linuxkit_kernel" "kernel" {
  image = "linuxkit/kernel:5.10.76"
  cmdline = "console=tty0 console=ttyS0 console=ttyAMA0"
}

data "linuxkit_init" "init" {
  containers = [
    "linuxkit/init:7e3d51e6ab5896ecb36a4829450f7430f2878927",
    "linuxkit/runc:9f7aad4eb5e4360cc9ed8778a5c501cce6e21601",
    "linuxkit/containerd:2f0907913dd54ab5186006034eb224a0da12443e",
    "linuxkit/ca-certificates:c1c73ef590dffb6a0138cf758fe4a4305c9864f4",
    "linuxkit/memlogd:fe4a123b619a7dfffc2ba1297dd03b4ac90e3dd7",
  ]
}

data "linuxkit_image" "sysctl" {
  name = "sysctl"
  image = "linuxkit/sysctl:0dc8f792fc3a58afcebcb0fbe6b48de587265c17"
}

data "linuxkit_image" "sysfs" {
  name = "sysfs"
  image = "linuxkit/sysfs:0148c62dbf57948849e8da829d36363b94a76c97"
}

data "linuxkit_image" "dhcp_boot" {
  name = "dhcp_boot"
  image = "linuxkit/dhcpcd:52d2c4df0311b182e99241cdc382ff726755c450"

  command = ["/sbin/dhcpcd", "--nobackground", "-d", "-f", "/dhcpcd.conf", "-1", "-4"]
}

data "linuxkit_image" "dhcp_svc" {
  name = "dhcp"
  image = "linuxkit/dhcpcd:52d2c4df0311b182e99241cdc382ff726755c450"
}

data "linuxkit_image" "logwrite" {
  name = "logwrite"
  image = "linuxkit/logwrite:568325cf294338b37446943c2b86a8cd8dc703db"
}

data "linuxkit_image" "getty" {
  name = "getty"
  image = "linuxkit/getty:76951a596aa5e0867a38e28f0b94d620e948e3e8"
  env   = ["INSECURE=true"]
}

data "linuxkit_image" "ntpd" {
  name = "ntpd"
  image = "linuxkit/openntpd:d6c36ac367ed26a6eeffd8db78334d9f8041b038"
}

data "linuxkit_image" "docker" {
  name  = "dockerd"
  image = "docker:20.10.12-dind"

  command = ["/usr/local/bin/docker-init", "/usr/local/bin/dockerd"]

  capabilities = ["all"]
  net          = "host"
  pid          = "host"

  mounts {
    type    = "cgroup"
    options = ["rw", "nosuid", "noexec", "nodev", "relatime"]
  }

  devices {
    path     = "/dev/console"
    type     = "c"
    major    = 5
    minor    = 1
    filemode = "0666"
  }

  devices {
    path = "all"
    type = "b"
  }

  binds = [
    "/dev:/dev",
    "/etc/docker/daemon.json:/etc/docker/daemon.json",
    "/etc/resolv.conf:/etc/resolv.conf",
    "/lib/modules:/lib/modules",
    "/run:/run:rshared",
    "/var/persist:/var/persist:rshared",
  ]

  rootfs_propagation = "shared"

  runtime {
    mkdir = [
      "/var/persist/docker",
    ]
  }
}

data "linuxkit_file" "docker_config" {
  path = "etc/docker/daemon.json"

  contents = jsonencode({
    data-root = "/var/persist/docker"
    iptables  = false
    bridge    = "none"
  })

  mode     = "0644"
  optional = false
}

data "linuxkit_file" "containerd_toml" {
  path     = "etc/containerd/runtime-config.toml"
  contents = <<EOF
cliopts="--log-level INFO"
stderr="/var/log/containerd.out.log"
stdout="stdout"
EOF
  mode     = "0644"
  optional = false
}

data "linuxkit_file" "ntpd_conf" {
  contents = "servers pool.ntp.org\n"
  path     = "etc/ntpd.conf"
  mode     = "0644"
  optional = false
}

data "linuxkit_config" "build" {
  kernel = data.linuxkit_kernel.kernel.id
  init = [
    data.linuxkit_init.init.id,
  ]

  onboot = flatten([
    data.linuxkit_image.sysctl.id,
    data.linuxkit_image.sysfs.id,
    data.linuxkit_image.dhcp_boot.id,
  ])

  services = flatten([
    data.linuxkit_image.dhcp_svc.id,
    data.linuxkit_image.logwrite.id,
    data.linuxkit_image.getty.id,
    data.linuxkit_image.ntpd.id,
    data.linuxkit_image.docker.id,
  ])

  files = flatten([
    data.linuxkit_file.containerd_toml.id,
    data.linuxkit_file.docker_config.id,
    data.linuxkit_file.ntpd_conf.id,
  ])
}

resource "linuxkit_build" "build" {
  config_yaml         = data.linuxkit_config.build.yaml
  docker_cache_enable = true
  destination         = "dockerd.tar"
}

resource "linuxkit_image_kernel_initrd" "dockerd" {
  build = linuxkit_build.build.destination
  destination = {
    kernel = "dockerd-kernel"
    initrd = "dockerd-initrd.img"
    cmdline = "dockerd-cmdline"
  }
}

resource "local_file" "config_yaml" {
  content = data.linuxkit_config.build.yaml
  filename = "${path.module}/config.yaml"
  file_permission = "0644"
  directory_permission = "0755"
}
