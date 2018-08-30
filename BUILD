# Load go_* rules
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# Load gazel rule
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/thatsmrtalbot/terraform-provider-linuxkit

# Gazelle generates the BUILD files for golang
gazelle(
    name = "gazelle",
    external = "vendored",
    command = "fix",
)

# Dep binary
alias(
    name = "dep",
    actual = select({
        "@bazel_tools//src/conditions:darwin": "@io_dep_darwin_amd64//file:downloaded",
        "//conditions:default": "@io_dep_linux_amd64//file:downloaded",
    }),
)

# Go library for main
go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/thatsmrtalbot/terraform-provider-linuxkit",
    visibility = ["//visibility:private"],
    deps = [
        "//linuxkit:go_default_library",
        "//vendor/github.com/hashicorp/terraform/plugin:go_default_library",
    ],
)

# Go binary build
go_binary(
    name = "go_default_binary",
    embed = [":go_default_library"],
    pure = "on",
    visibility = ["//visibility:public"],
    out = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": "linux_amd64/terraform-provider-linuxkit",
        "@io_bazel_rules_go//go/platform:darwin_amd64": "darwin_amd64/terraform-provider-linuxkit",
        "@io_bazel_rules_go//go/platform:windows_amd64": "windows_amd64/terraform-provider-linuxkit",
    }),
)
