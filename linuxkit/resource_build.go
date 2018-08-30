package linuxkit

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/moby/tool/src/moby"
	"github.com/pkg/errors"

	"github.com/hashicorp/terraform/helper/schema"
)

func buildResource() *schema.Resource {
	resource := &schema.Resource{
		Create: buildCreate,
		Read:   buildRead,
		Update: buildUpdate,
		Delete: buildDelete,
		Exists: buildExists,

		Schema: map[string]*schema.Schema{
			"destination": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The destination of the generated OS image",
				Required:    true,
				ForceNew:    true,
			},

			"config": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The linuxkit config id",
				Required:    true,
			},
			"config_yaml": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The linuxkit config yaml",
				Required:    true,
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "tar",
				Description: "Type of build, can be tar or docker",
				Optional:    true,
			},
		},
	}

	withConfigSchema(resource.Schema)

	return resource
}

func buildRead(d *schema.ResourceData, meta interface{}) error {
	id, err := buildID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func buildUpdate(d *schema.ResourceData, meta interface{}) error {
	return buildCreate(d, meta)
}

func buildCreate(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)
	typ := d.Get("type").(string)

	configData, err := buildConfig(d)
	if err != nil {
		return err
	}

	config, err := moby.NewConfig(configData)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	err = moby.Build(config, outputFile, true, typ)
	if err != nil {
		return err
	}

	id, err := buildID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func buildDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")

	destination := d.Get("destination").(string)

	return os.Remove(destination)
}

func buildExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	destination := d.Get("destination").(string)

	if _, err := os.Stat(destination); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func buildID(d *schema.ResourceData) (string, error) {
	destination := d.Get("destination").(string)

	f, err := os.Open(destination)

	if os.IsNotExist(err) {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	config, err := buildConfig(d)
	if err != nil {
		return "", err
	}

	h.Write(config)
	h.Write([]byte(d.Get("type").(string)))

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func buildConfig(d *schema.ResourceData) ([]byte, error) {
	var err error
	var bts []byte

	if v, ok := d.GetOk("config"); ok {
		id := v.(string)

		if config, ok := globalCache.configs[id]; ok {
			bts, err = yaml.Marshal(config)
		} else {
			return nil, errors.New("unknown config id")
		}
	} else if v, ok := d.GetOk("config_yaml"); ok {
		bts = []byte(v.(string))
	} else {
		_, bts, err = fromConfigSchema(d)
	}

	return bts, err
}
