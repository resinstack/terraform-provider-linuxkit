package linuxkit

import (
	"errors"

	"gopkg.in/yaml.v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
)

func withConfigSchema(s map[string]*schema.Schema) {
	s["kernel"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The ID of the kernel resource",
		Required:    true,
		ForceNew:    true,
	}

	s["init"] = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The IDs of init containers",
		Required:    true,
		ForceNew:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}

	s["onboot"] = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The IDs of the onboot containers",
		Optional:    true,
		ForceNew:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}

	s["onshutdown"] = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The IDs of the onshutdown containers",
		Optional:    true,
		ForceNew:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}

	s["services"] = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The IDs of the service containers",
		Optional:    true,
		ForceNew:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}

	s["files"] = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The IDs of the file config",
		Optional:    true,
		ForceNew:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}

	s["trust"] = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The ID of the trust config",
		Optional:    true,
		ForceNew:    true,
	}
}

func fromConfigSchema(d *schema.ResourceData) (*moby.Moby, []byte, error) {
	config := &moby.Moby{}

	if v, ok := d.GetOk("kernel"); ok {
		if kernel, ok := globalCache.kernels[v.(string)]; ok {
			config.Kernel = *kernel
		} else {
			return nil, nil, errors.New("kernel not found")
		}
	}

	if v, ok := d.GetOk("init"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if init, ok := globalCache.inits[id]; ok {
				config.Init = append(config.Init, init...)
			} else {
				return nil, nil, errors.New("init image not found")
			}
		}
	}

	if v, ok := d.GetOk("onboot"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if image, ok := globalCache.images[id]; ok {
				config.Onboot = append(config.Onboot, image)
			} else {
				return nil, nil, errors.New("onboot image not found")
			}
		}
	}

	if v, ok := d.GetOk("onshutdown"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if image, ok := globalCache.images[id]; ok {
				config.Onshutdown = append(config.Onshutdown, image)
			} else {
				return nil, nil, errors.New("onshutdown image not found")
			}
		}
	}

	if v, ok := d.GetOk("services"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if image, ok := globalCache.images[id]; ok {
				config.Services = append(config.Services, image)
			} else {
				return nil, nil, errors.New("services image not found")
			}
		}
	}

	if v, ok := d.GetOk("files"); ok {
		for _, id := range interfaceSliceToStringSlice(v.([]interface{})) {
			if file, ok := globalCache.files[id]; ok {
				config.Files = append(config.Files, *file)
			} else {
				return nil, nil, errors.New("file config not found")
			}
		}
	}

	if v, ok := d.GetOk("trust"); ok {
		if trust, ok := globalCache.trust[v.(string)]; ok {
			config.Trust = *trust
		} else {
			return nil, nil, errors.New("trust config not found")
		}
	}

	byts, err := yaml.Marshal(config)
	if err != nil {
		return nil, nil, err
	}

	return config, byts, nil
}

func configDataSource() *schema.Resource {
	resource := &schema.Resource{
		Read: configRead,

		Schema: map[string]*schema.Schema{
			"yaml": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The rendered yaml of the config",
				Computed:    true,
			},
		},
	}

	withConfigSchema(resource.Schema)

	return resource
}

func configRead(d *schema.ResourceData, meta interface{}) error {
	config, byts, err := fromConfigSchema(d)
	if err != nil {
		return err
	}

	d.SetId(globalCache.addConfig(config))
	d.Set("yaml", string(byts))

	return nil
}
