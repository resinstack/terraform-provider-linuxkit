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
	createResources()

	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"linuxkit_kernel":   kernelDataSource(),
			"linuxkit_init":     initDataSource(),
			"linuxkit_image":    imageDataSource(),
			"linuxkit_config":   configDataSource(),
			"linuxkit_file":     fileDataSource(),
			"linuxkit_metadata": metadataDataSource(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"linuxkit_build":                   buildResource(),
			"linuxkit_image_aws":               resourceImage["aws"],
			"linuxkit_image_dynamic_vhd":       resourceImage["dynamic-vhd"],
			"linuxkit_image_gcp":               resourceImage["gcp"],
			"linuxkit_image_iso_bios":          resourceImage["iso-bios"],
			"linuxkit_image_iso_efi":           resourceImage["iso-efi"],
			"linuxkit_image_kernel_initrd":     resourceImage["kernel+initrd"],
			"linuxkit_image_kernel_squashfs":   resourceImage["kernel+squashfs"],
			"linuxkit_image_qcow2_efi":         resourceImage["qcow2-efi"],
			"linuxkit_image_raw_bios":          resourceImage["raw-bios"],
			"linuxkit_image_raw_efi":           resourceImage["raw-efi"],
			"linuxkit_image_rpi3":              resourceImage["rpi3"],
			"linuxkit_image_tar_kernel_initrd": resourceImage["tar-kernel-initrd"],
			"linuxkit_image_vhd":               resourceImage["vhd"],
			"linuxkit_image_vmdk":              resourceImage["vmdk"],
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
