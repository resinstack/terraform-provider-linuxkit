package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func imageKernelInitrdResource() *schema.Resource {
	out := newOutput("kernel+initrd")

	return &schema.Resource{
		Description: "`linuxkit_image_kernel_initrd` produces a kernel and initrd suitable for PXE booting.  If you want an archive of the output, use `linuxkit_image_tar_kernel_initrd` instead.",

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
				Type:     schema.TypeMap,
				Required: true,
				ForceNew: true,

				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
