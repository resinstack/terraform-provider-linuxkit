package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func imageRpi3Resource() *schema.Resource {
	out := newOutput("rpi3")

	return &schema.Resource{
		Description: "`linuxkit_image_rpi3` produces a filesystem image suitable for booting a Raspberry Pi model 3, and by extension compatible boards.",

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
