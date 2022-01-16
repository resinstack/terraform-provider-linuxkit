package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func imageTarKernalInitrdResource() *schema.Resource {
	out := newOutput("tar-kernel-initrd")

	return &schema.Resource{
		Description: "`linuxkit_image_tar_kernel_initrd` produces a tarball containing a kernel and initrd suitable for PXE booting.",

		Create: out.create,
		Read:   out.read,
		Delete: out.delete,

		Schema: map[string]*schema.Schema{
			"build": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The build tarball",
				Required:    true,
				ForceNew:    true,
			},

			"destination": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The destination of the raw generated OS image",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}
