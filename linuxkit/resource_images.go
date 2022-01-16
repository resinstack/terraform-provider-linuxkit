package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	resourceImage map[string]*schema.Resource

	descriptions = map[string]string{
		"aws":               "`linuxkit_image_aws` produces an image suitable for importing into AWS EC2.",
		"dynamic-vhd":       "`linuxkit_dynamic_vhd` produces a sparse VHD from a build tarball.",
		"gcp":               "`linuxkit_image_gcp` produces an image tarball suitable for uploading to GCP.",
		"iso-bios":          "`linuxkit_image_iso_bios` produces an ISO file suitable for booting on systems that support BIOS booting.",
		"iso-efi":           "`linuxkit_image_iso_efi` produces an ISO suitable for booting a machine with UEFI booting.",
		"kernel+initrd":     "`linuxkit_image_kernel_initrd` produces a kernel and initrd suitable for PXE booting.  If you want an archive of the output, use `linuxkit_image_tar_kernel_initrd` instead.",
		"kernel+squashfs":   "`linuxkit_image_kernel_squashfs` process a build tarball and places the root filesystem on a squashfs.",
		"qcow2-efi":         "`linuxkit_image_qcow_efi` produces a qcow2 filesystem bootable on EFI systems.",
		"raw-bios":          "`linuxkit_image_raw_bios` produces an MBR image suitable for booting BIOS systems.",
		"raw-efi":           "`linuxkit_image_raw_efi` produces a GPT image file suitable for booting an UEFI enabled system.",
		"rpi3":              "`linuxkit_image_rpi3` produces a filesystem image suitable for booting a Raspberry Pi model 3, and by extension compatible boards.",
		"tar-kernel-initrd": "`linuxkit_image_tar_kernel_initrd` produces a tarball containing a kernel and initrd suitable for PXE booting.",
		"vhd":               "`linuxkit_image_vhd` produces a VHD file with provisioned storage.",
		"vmdk":              "`linuxkit_image_vmdk` produces a VMDK file from the build tarball.",
	}
)

func createResources() {
	resourceImage = make(map[string]*schema.Resource, len(descriptions))
	for out, desc := range descriptions {
		impl := newOutput(out)
		resourceImage[out] = &schema.Resource{
			Description: desc,

			Create: impl.create,
			Read:   impl.read,
			Delete: impl.delete,

			Schema: map[string]*schema.Schema{
				"build": &schema.Schema{
					Type:        schema.TypeString,
					Description: "The build tarball",
					Required:    true,
					ForceNew:    true,
				},

				"destination": &schema.Schema{
					Type:        schema.TypeString,
					Description: "The destination of linuxkit artifact(s).",
					Required:    true,
					ForceNew:    true,
				},
			},
		}
	}

	// Some providers take a little tweaking, these are applied
	// below:

	// kernel+initrd outputs multiple files and as a result has a
	// different destination type
	resourceImage["kernel+initrd"].Schema["destination"] = &schema.Schema{
		Type:     schema.TypeMap,
		Required: true,
		ForceNew: true,

		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	// For some reason that is not adequately documented, AWS
	// images have to have a size set.
	resourceImage["aws"].Schema["size"] = &schema.Schema{
		Type:        schema.TypeInt,
		Description: "The size in megabytes of the new image",
		Required:    true,
		ForceNew:    true,
	}
}
