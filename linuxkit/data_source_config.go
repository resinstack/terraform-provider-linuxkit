package linuxkit

import (
	"errors"

	"gopkg.in/yaml.v2"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/moby/tool/src/moby"
)

func configDataSource() *schema.Resource {
	return &schema.Resource{
		Read: configRead,

		Schema: map[string]*schema.Schema{
			"kernel": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ID of the kernel resource",
				Required:    true,
				ForceNew:    true,
			},
			"init": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The IDs of init containers",
				Required:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"onboot": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The IDs of the onboot containers",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"onshutdown": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The IDs of the onshutdown containers",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"services": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The IDs of the service containers",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"files": &schema.Schema{
				Type:        schema.TypeList,
				Description: "The IDs of the file config",
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"trust": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The ID of the trust config",
				Optional:    true,
				ForceNew:    true,
			},
			"yaml": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The rendered yaml of the config",
				Computed:    true,
			},
		},
	}
}

func configRead(d *schema.ResourceData, meta interface{}) error {
	config := &moby.Moby{}

	if v, ok := d.GetOk("kernel"); ok {
		if kernel, ok := globalCache.kernels[v.(string)]; ok {
			config.Kernel = *kernel
		} else {
			return errors.New("Kernel not found")
		}
	}

	if v, ok := d.GetOk("init"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if init, ok := globalCache.inits[id]; ok {
				config.Init = append(config.Init, init...)
			} else {
				return errors.New("Init image not found")
			}
		}
	}

	if v, ok := d.GetOk("onboot"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if image, ok := globalCache.images[id]; ok {
				config.Onboot = append(config.Onboot, image)
			} else {
				return errors.New("Onboot image not found")
			}
		}
	}

	if v, ok := d.GetOk("onshutdown"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if image, ok := globalCache.images[id]; ok {
				config.Onshutdown = append(config.Onshutdown, image)
			} else {
				return errors.New("Onshutdown image not found")
			}
		}
	}

	if v, ok := d.GetOk("services"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if image, ok := globalCache.images[id]; ok {
				config.Services = append(config.Services, image)
			} else {
				return errors.New("Services image not found")
			}
		}
	}

	if v, ok := d.GetOk("files"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if file, ok := globalCache.files[id]; ok {
				config.Files = append(config.Files, *file)
			} else {
				return errors.New("File config not found")
			}
		}
	}

	if v, ok := d.GetOk("trust"); ok {
		if trust, ok := globalCache.trust[v.(string)]; ok {
			config.Trust = *trust
		} else {
			return errors.New("Trust config not found")
		}
	}

	byts, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	d.SetId(globalCache.addConfig(config))
	d.Set("yaml", string(byts))

	return nil
}
