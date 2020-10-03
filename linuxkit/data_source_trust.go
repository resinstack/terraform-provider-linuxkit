package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
)

func trustDataSource() *schema.Resource {
	return &schema.Resource{
		Read: trustRead,

		Schema: map[string]*schema.Schema{
			"image": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Images to trust",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"org": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Orgs to trust",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func trustRead(d *schema.ResourceData, meta interface{}) error {
	trust := &moby.TrustConfig{}

	if v, ok := d.GetOk("image"); ok {
		trust.Image = interfaceSliceToStringSlice(v.([]interface{}))
	}

	if v, ok := d.GetOk("org"); ok {
		trust.Org = interfaceSliceToStringSlice(v.([]interface{}))
	}

	d.SetId(globalCache.addTrust(trust))

	return nil
}
