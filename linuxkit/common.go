package linuxkit

import (
	"os"
	"io"
	"crypto/md5"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
)


type linuxkitImageOutput struct {
	lkitfmt string
	size int // Only used for *very* specific output formats
}

func newOutput(lkitfmt string) *linuxkitImageOutput {
	return &linuxkitImageOutput{
		lkitfmt: lkitfmt,
	}
}

func (l *linuxkitImageOutput) read(d *schema.ResourceData, meta interface{}) error {
	id, err := l.id(d)
	if err != nil {
		log.Printf("Error on read: %v", err)
		return err
	}
	d.SetId(id)
	return nil
}

func (l *linuxkitImageOutput) create(d *schema.ResourceData, meta interface{}) error {
	build := d.Get("build").(string)

	sizeattr, ok := d.GetOk("size")
	if ok {
		l.size = sizeattr.(int)
	}

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	err = moby.Formats(filepath.Join(dir, "base"), build, []string{l.lkitfmt}, l.size, defaultLinuxkitCache())
	if err != nil {
		return err
	}

	// Need to handle variable numbers of files here mapping from
	// what linuxkit calls them to what the provider calls them.
	// This likely involves a fancy type switch to handle whether
	// destination is a string or a map of string.
	tnames := l.getTransientArtifactNames(l.lkitfmt)
	dnames := l.getDestination(d)

	// Just a quick sanity check...
	if len(tnames) != len(dnames) {
		log.Println("mistmatch in number of required destinations.  This could be a bug, or you could have missed providing the name of a destination for a multi-dest resource", "got", len(dnames), "expected", len(tnames))
	}

	for i := range tnames {
		if err := copyFile(filepath.Join(dir, tnames[i]), dnames[i]); err != nil {
			return err
		}
	}

	id, err := l.id(d)
	if err != nil {
		return err
	}
	d.SetId(id)

	return nil
}

func (l *linuxkitImageOutput) delete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	for _, path := range l.getDestination(d) {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return nil
}

func (l *linuxkitImageOutput) id(d *schema.ResourceData) (string, error) {
	hasher := md5.New()

	// The ID of this resource changes if the build changes,
	// because that will invalidate everything else in here.
	build, err := os.Open(d.Get("build").(string))
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}
	defer build.Close()
	if _, err := io.Copy(hasher, build); err != nil {
		return "", err
	}

	for _, dest := range l.getDestination(d) {
		f, err := os.Open(dest)
		if os.IsNotExist(err) {
			return "", nil
		} else if err != nil {
			return "", err
		}
		defer f.Close()
		if _, err := io.Copy(hasher, f); err != nil {
			return "", err
		}
		f.Close()
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

// This function and the one below it are order sensitive.  If you add
// fields in this one that are ordered you must make sure that the
// transient artifact names are in the same order to ensure that the
// copy will occur correctly.
func (l *linuxkitImageOutput) getDestination(d *schema.ResourceData) ([]string) {
	switch dest := d.Get("destination").(type) {
	case string:
		return []string{dest}
	case map[string]interface{}:
		// This is only used for the kernel/initrd/cmdline
		// output, if another format gets this type we'll need
		// to do something more intelligent than this.
		return []string{
			dest["initrd"].(string),
			dest["kernel"].(string),
			dest["cmdline"].(string),
		}
	default:
		log.Println("Failure to map destination type, this is a provider bug!")
	}
	return nil
}

// Format names here must match those passed to moby.Formats, and must
// also match the moby.Formats returned resource names for transient
// artifacts.
func (l *linuxkitImageOutput) getTransientArtifactNames(format string) []string {
	switch format {
	case "aws":
		return []string{"base.raw"}
	case "dynamic-vhd":
		return []string{"base.vhd"}
	case "gcp":
		return []string{"base.img.tar.gz"}
	case "iso-bios":
		return []string{"base.iso"}
	case "iso-efi":
		return []string{"base-efi.iso"}
	case "kernel+initrd":
		return []string{"base-initrd.img", "base-kernel", "base-cmdline"}
	case "kernel+squashfs":
		return []string{"base-squashfs.iso"}
	case "qcow2-efi":
		return []string{"base.qcow2"}
	case "raw-bios":
		return []string{"base-bios.img"}
	case "raw-efi":
		return []string{"base-efi.img"}
	case "rpi3":
		return []string{"base.tar"}
	case "tar-kernel-initrd":
		return []string{"base-inird.tar"}
	case "vhd":
		return []string{"base.vhd"}
	case "vmdk":
		return []string{"base.vmdk"}
	default:
		log.Println("Failed to map transient artifact name, this is a bug!")
		return []string{}
	}
}
