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

func imageQcow2EfiResource() *schema.Resource {
	return &schema.Resource{
		Description: "`linuxkit_image_qcow_efi` produces a qcow2 filesystem bootable on EFI systems.",

		Create: imageQcow2EfiCreate,
		Read:   imageQcow2EfiRead,
		Delete: imageQcow2EfiDelete,
		Exists: imageQcow2EfiExists,

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

func imageQcow2EfiRead(d *schema.ResourceData, meta interface{}) error {
	id, err := imageQcow2EfiID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageQcow2EfiCreate(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)
	build := d.Get("build").(string)

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	err = moby.Formats(filepath.Join(dir, "base"), build, []string{"qcow2-efi"}, 0, defaultLinuxkitCache())
	if err != nil {
		return err
	}

	err = copyFile(filepath.Join(dir, "base.qcow2"), destination)
	if err != nil {
		return err
	}

	id, err := imageQcow2EfiID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageQcow2EfiDelete(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)

	d.SetId("")

	err := os.Remove(destination)

	if err != nil {
		return err
	}

	return nil
}

func imageQcow2EfiExists(d *schema.ResourceData, meta interface{}) (bool, error) {
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

func imageQcow2EfiID(d *schema.ResourceData) (string, error) {
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
