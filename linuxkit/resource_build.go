package linuxkit

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"

	"github.com/moby/tool/src/moby"

	"github.com/hashicorp/terraform/helper/schema"
)

func buildResource() *schema.Resource {
	return &schema.Resource{
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

	config, err := moby.NewConfig([]byte(d.Get("config").(string)))
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

	h.Write([]byte(d.Get("config").(string)))
	h.Write([]byte(d.Get("type").(string)))

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
