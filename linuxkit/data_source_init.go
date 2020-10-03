package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func initDataSource() *schema.Resource {
	return &schema.Resource{
		Read: initRead,

		Schema: map[string]*schema.Schema{
			"containers": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Containers to mount on the root filesystem",
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func initRead(d *schema.ResourceData, meta interface{}) error {
	init := []string{}

	if containers, ok := d.GetOk("containers"); ok {
		for _, container := range containers.([]interface{}) {
			init = append(init, container.(string))
		}
	}

	d.SetId(globalCache.addInit(init))

	return nil
}
