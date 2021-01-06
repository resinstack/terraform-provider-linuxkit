package linuxkit

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
)

func imageRpi3Resource() *schema.Resource {
	return &schema.Resource{
		Description: "`linuxkit_image_rpi3` produces a filesystem image suitable for booting a Raspberry Pi model 3, and by extension compatible boards.",

		Create: imageRpi3Create,
		Read:   imageRpi3Read,
		Delete: imageRpi3Delete,
		Exists: imageRpi3Exists,

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

func imageRpi3Read(d *schema.ResourceData, meta interface{}) error {
	id, err := imageRpi3ID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageRpi3Create(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)
	build := d.Get("build").(string)

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	err = moby.Formats(filepath.Join(dir, "base"), build, []string{"rpi3"}, 0, false)
	if err != nil {
		return err
	}

	err = copyFile(filepath.Join(dir, "base.tar"), destination)
	if err != nil {
		return err
	}

	id, err := imageRpi3ID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageRpi3Delete(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)

	d.SetId("")

	err := os.Remove(destination)

	if err != nil {
		return err
	}

	return nil
}

func imageRpi3Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
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

func imageRpi3ID(d *schema.ResourceData) (string, error) {
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
