package linuxkit

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/moby/tool/src/moby"
)

func imageQcow2BiosResource() *schema.Resource {
	return &schema.Resource{
		Create: imageQcow2BiosCreate,
		Read:   imageQcow2BiosRead,
		Delete: imageQcow2BiosDelete,
		Exists: imageQcow2BiosExists,

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

func imageQcow2BiosRead(d *schema.ResourceData, meta interface{}) error {
	id, err := imageQcow2BiosID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageQcow2BiosCreate(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)
	size := d.Get("size").(int)
	build := d.Get("build").(string)

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	err = moby.Formats(filepath.Join(dir, "base"), build, []string{"qcow2-bios"}, size)
	if err != nil {
		return err
	}

	err = copyFile(filepath.Join(dir, "base.qcow2"), destination)
	if err != nil {
		return err
	}

	id, err := imageQcow2BiosID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageQcow2BiosDelete(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(string)

	d.SetId("")

	err := os.Remove(destination)

	if err != nil {
		return err
	}

	return nil
}

func imageQcow2BiosExists(d *schema.ResourceData, meta interface{}) (bool, error) {
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

func imageQcow2BiosID(d *schema.ResourceData) (string, error) {
	destination := d.Get("destination").(string)
	build := d.Get("build").(string)
	size := d.Get("size").(int)

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

	if err := binary.Write(h, binary.BigEndian, int64(size)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
