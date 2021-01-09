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

func imageKernelInitrdResource() *schema.Resource {
	return &schema.Resource{
		Description: "`linuxkit_image_kernel_initrd` produces a kernel and initrd suitable for PXE booting.  If you want an archive of the output, use `linuxkit_image_tar_kernel_initrd` instead.",

		Create: imageKernelInitrdCreate,
		Read:   imageKernelInitrdRead,
		Delete: imageKernelInitrdDelete,
		Exists: imageKernelInitrdExists,

		Schema: map[string]*schema.Schema{
			"build": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The build tarball",
				Required:    true,
				ForceNew:    true,
			},

			"destination": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				ForceNew: true,

				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func imageKernelInitrdRead(d *schema.ResourceData, meta interface{}) error {
	id, err := imageKernelInitrdID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageKernelInitrdCreate(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(map[string]interface{})
	build := d.Get("build").(string)

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	defer os.RemoveAll(dir)

	err = moby.Formats(filepath.Join(dir, "base"), build, []string{"kernel+initrd"}, 0, false)
	if err != nil {
		return err
	}

	err = copyFile(filepath.Join(dir, "base-initrd.img"), destination["initrd"].(string))
	if err != nil {
		return err
	}

	err = copyFile(filepath.Join(dir, "base-kernel"), destination["kernel"].(string))
	if err != nil {
		return err
	}

	err = copyFile(filepath.Join(dir, "base-cmdline"), destination["cmdline"].(string))
	if err != nil {
		return err
	}

	id, err := imageKernelInitrdID(d)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func imageKernelInitrdDelete(d *schema.ResourceData, meta interface{}) error {
	destination := d.Get("destination").(map[string]interface{})

	d.SetId("")

	for _, s := range []string{"kernel", "initrd", "cmdline"} {
		_, err := os.Stat(destination[s].(string))

		if os.IsNotExist(err) {
			continue
		}

		if err != nil {
			return err
		}

		err = os.Remove(destination[s].(string))

		if err != nil {
			return err
		}
	}

	return nil
}

func imageKernelInitrdExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	destination := d.Get("destination").(map[string]interface{})

	for _, s := range []string{"kernel", "initrd", "cmdline"} {
		_, err := os.Stat(destination[s].(string))

		if os.IsNotExist(err) {
			return false, nil
		}

		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func imageKernelInitrdID(d *schema.ResourceData) (string, error) {
	destination := d.Get("destination").(map[string]interface{})
	build := d.Get("build").(string)

	h := md5.New()

	f, err := os.Open(build)

	if os.IsNotExist(err) {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	for _, s := range []string{"kernel", "initrd", "cmdline"} {
		f, err := os.Open(destination[s].(string))

		if os.IsNotExist(err) {
			return "", nil
		}

		if err != nil {
			return "", err
		}

		defer f.Close()

		if _, err := io.Copy(h, f); err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
