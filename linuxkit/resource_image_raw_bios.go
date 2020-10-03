package linuxkit

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
)

func imageRawBiosResource() *schema.Resource {
	return &schema.Resource{
		Create: imageRawBiosCreate,
		Read:   imageRawBiosRead,
		Delete: imageRawBiosDelete,
		Exists: imageRawBiosExists,

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

func imageRawBiosRead(d *schema.ResourceData, meta interface{}) error {
	id, err := imageRawBiosID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageRawBiosCreate(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)
	build := d.Get("build").(string)

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	err = moby.Formats(filepath.Join(dir, "base"), build, []string{"raw-bios"}, 0, false)
	if err != nil {
		return err
	}

	err = copyFile(filepath.Join(dir, "base-bios.iso"), destination)
	if err != nil {
		return err
	}

	id, err := imageRawBiosID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageRawBiosDelete(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)

	d.SetId("")

	err := os.Remove(destination)

	if err != nil {
		return err
	}

	return nil
}

func imageRawBiosExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	destination := d.Get("destination").(string)

	_, err := os.Stat(destination)

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func imageRawBiosID(d *schema.ResourceData) (string, error) {
	destination := d.Get("destination").(string)
	build := d.Get("build").(string)

	h := md5.New()

	f1, err := os.Open(destination)

	if os.IsNotExist(err) {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	defer f1.Close()

	f2, err := os.Open(build)

	if os.IsNotExist(err) {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	defer f2.Close()

	if _, err := io.Copy(h, f1); err != nil {
		return "", err
	}

	if _, err := io.Copy(h, f2); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
