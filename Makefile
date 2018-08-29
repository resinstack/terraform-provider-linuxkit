.PHONY: vendor
vendor:
	bazel run //:dep -- ensure 

.PHONY: gazelle
gazelle:
	bazel run //:gazelle

.PHONY: build-%
build-%: 
	bazel build //:terraform-provider-linuxkit --platforms=@io_bazel_rules_go//go/toolchain:$*

.PHONY: build
build: build-darwin_amd64 build-linux_amd64

.PHONY: test
test:
	bazel test //linuxkit:go_default_test