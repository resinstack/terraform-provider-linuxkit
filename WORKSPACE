# Load all workspace dependencies with http_archive
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

# Download golang rules
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.15.0/rules_go-0.15.0.tar.gz"],
    sha256 = "56d946edecb9879aed8dff411eb7a901f687e242da4fa95c81ca08938dd23bb4",
)

# Load golang workspace dependencies and toolchains
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

# Download gazelle rules
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)

# Load gazelle workspace dependencies
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

# Download dep binary for darwin
http_file(
    name = "io_dep_darwin_amd64",
    executable = True,
    urls = ["https://github.com/golang/dep/releases/download/v0.5.0/dep-darwin-amd64"],
    sha256 = "1a7bdb0d6c31ecba8b3fd213a1170adf707657123e89dff234871af9e0498be2",
)

# Download dep binary for linux
http_file(
    name = "io_dep_linux_amd64",
    executable = True,
    urls = ["https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64"],
    sha256 = "287b08291e14f1fae8ba44374b26a2b12eb941af3497ed0ca649253e21ba2f83",
)
