BAZEL   = bazel --host_jvm_args=-Xmx500m --host_jvm_args=-Xms500m
VERSION = $(shell git describe --tags 2> /dev/null)

.PHONY: vendor
vendor:
	$(BAZEL) run //:dep -- ensure 

.PHONY: gazelle
gazelle:
	$(BAZEL) run //:gazelle

.PHONY: test
test:
	$(BAZEL) test //linuxkit:go_default_test

.PHONY: clean
clean:
	$(BAZEL) clean
	rm -rf bin

.PHONY: build
build: build-darwin_amd64 build-linux_amd64

.PHONY: build-%
build-windows_%: EXTENSION = .exe
build-%:
	$(BAZEL) build //:go_default_binary --platforms=@io_bazel_rules_go//go/toolchain:$* 
	mkdir -p bin/$*
	cp -fL bazel-bin/$*/terraform-provider-linuxkit bin/$*/terraform-provider-linuxkit$(if $(VERSION),_$(VERSION))$(EXTENSION)

.PHONY: build
bundle: build
	tar -zcvf release.tar.gz -C bin .
