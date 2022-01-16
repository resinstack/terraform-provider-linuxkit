package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func imageAwsResource() *schema.Resource {
	out := newOutput("aws")

	return &schema.Resource{
		Description: "`linuxkit_image_aws` produces an image suitable for importing into AWS EC2.",

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

			"size": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The size in megabytes of the new image",
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
