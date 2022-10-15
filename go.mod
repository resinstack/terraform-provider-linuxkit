module github.com/resinstack/terraform-provider-linuxkit

go 1.15

require (
	github.com/hashicorp/go-hclog v1.1.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1
	github.com/linuxkit/linuxkit/src/cmd/linuxkit v0.0.0-20221012113451-61a07e26cf97
	github.com/mitchellh/go-homedir v1.1.0
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	// these are for the delicate dance of docker/docker, moby/moby, moby/buildkit, estesp/manifest-tool, oras.land/oras-go, linuxkit/linuxkit
	github.com/docker/docker => github.com/moby/moby v20.10.3-0.20220728162118-71cb54cec41e+incompatible
	oras.land/oras-go => oras.land/oras-go v1.1.0
)
