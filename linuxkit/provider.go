package linuxkit

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
	"github.com/pkg/errors"
)

// Provider linuxkit
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"linuxkit_kernel":   kernelDataSource(),
			"linuxkit_init":     initDataSource(),
			"linuxkit_image":    imageDataSource(),
			"linuxkit_trust":    trustDataSource(),
			"linuxkit_config":   configDataSource(),
			"linuxkit_file":     fileDataSource(),
			"linuxkit_metadata": metadataDataSource(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"linuxkit_build":                   buildResource(),
			"linuxkit_image_kernel_initrd":     imageKernelInitrdResource(),
			"linuxkit_image_aws":               imageAwsResource(),
			"linuxkit_image_dynamic_vhd":       imageDynamicVhdResource(),
			"linuxkit_image_gcp":               imageGcpResource(),
			"linuxkit_image_iso_bios":          imageIsoBiosResource(),
			"linuxkit_image_iso_efi":           imageIsoEfiResource(),
			"linuxkit_image_kernel_squashfs":   imageKernelSquashfsResource(),
			"linuxkit_image_raw_bios":          imageRawBiosResource(),
			"linuxkit_image_raw_efi":           imageRawEfiResource(),
			"linuxkit_image_tar_kernel_initrd": imageTarKernalInitrdResource(),
			"linuxkit_image_vhd":               imageVhdResource(),
			"linuxkit_image_vmdk":              imageVmdkResource(),
			"linuxkit_image_rpi3":              imageRpi3Resource(),
			"linuxkit_image_qcow2_bios":        imageQcow2BiosResource(),
			"linuxkit_image_qcow2_efi":         imageQcow2EfiResource(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(*schema.ResourceData) (meta interface{}, err error) {
	moby.MobyDir, err = defaultMobyConfigDir()
	if err != nil {
		return
	}

	err = errors.Wrap(os.MkdirAll(moby.MobyDir, 0755), "could not create config directory")
	if err != nil {
		return
	}

	err = errors.Wrap(os.MkdirAll(filepath.Join(moby.MobyDir, "tmp"), 0755), "could not create config tmp directory")
	if err != nil {
		return
	}

	return
}
