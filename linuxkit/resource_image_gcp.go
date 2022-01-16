package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func imageGcpResource() *schema.Resource {
	out := newOutput("gcp")

	return &schema.Resource{
		Description: "`linuxkit_image_gcp` produces an image tarball suitable for uploading to GCP.",

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
