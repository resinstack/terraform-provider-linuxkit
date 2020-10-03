VERSION = $(shell git describe --tags 2> /dev/null)

.PHONY: vendor
vendor:
	bazel run //:dep -- ensure 

.PHONY: gazelle
gazelle:
	bazel run //:gazelle

.PHONY: test
test:
	bazel test //linuxkit:go_default_test

.PHONY: clean
clean:
	bazel clean
	rm -rf bin

.PHONY: build
build: build-darwin_amd64 build-linux_amd64

.PHONY: build-%
build-windows_%: EXTENSION = .exe
build-%:
	bazel build //:go_default_binary --platforms=@io_bazel_rules_go//go/toolchain:$* 
	mkdir -p bin/$*
	cp -fL bazel-bin/$*/terraform-provider-linuxkit bin/$*/terraform-provider-linuxkit$(if $(VERSION),_$(VERSION))$(EXTENSION)

.PHONY: build
bundle: build
	tar -zcvf release.tar.gz -C bin .
