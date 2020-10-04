package linuxkit

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// Provider linuxkit
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"linuxkit_kernel":   kernelDataSource(),
			"linuxkit_init":     initDataSource(),
			"linuxkit_image":    imageDataSource(),
			"linuxkit_trust":    trustDataSource(),
			"linuxkit_config":   configDataSource(),
			"linuxkit_file":     fileDataSource(),
			"linuxkit_metadata": metadataDataSource(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"linuxkit_build":                   buildResource(),
			"linuxkit_image_kernel_initrd":     imageKernelInitrdResource(),
			"linuxkit_image_aws":               imageAwsResource(),
			"linuxkit_image_dynamic_vhd":       imageDynamicVhdResource(),
			"linuxkit_image_gcp":               imageGcpResource(),
			"linuxkit_image_iso_bios":          imageIsoBiosResource(),
			"linuxkit_image_iso_efi":           imageIsoEfiResource(),
			"linuxkit_image_kernel_squashfs":   imageKernelSquashfsResource(),
			"linuxkit_image_raw_bios":          imageRawBiosResource(),
			"linuxkit_image_raw_efi":           imageRawEfiResource(),
			"linuxkit_image_tar_kernel_initrd": imageTarKernalInitrdResource(),
			"linuxkit_image_vhd":               imageVhdResource(),
			"linuxkit_image_vmdk":              imageVmdkResource(),
			"linuxkit_image_rpi3":              imageRpi3Resource(),
			"linuxkit_image_qcow2_bios":        imageQcow2BiosResource(),
			"linuxkit_image_qcow2_efi":         imageQcow2EfiResource(),
		},

		ConfigureFunc: configureProvider,
	}
}

// globalCache keeps the instances of the internal types of ignition generated
// by the different data resources with the goal to be reused by the
// ignition_config data resource. The key of the maps are a hash of the types
// calculated on the type serialized to JSON.
var globalCache = &cache{
	configs: make(map[string]*moby.Moby),
	kernels: make(map[string]*moby.KernelConfig),
	inits:   make(map[string][]string),
	images:  make(map[string]*moby.Image),
	files:   make(map[string]*moby.File),
	trust:   make(map[string]*moby.TrustConfig),
}

type cache struct {
	configs map[string]*moby.Moby
	kernels map[string]*moby.KernelConfig
	inits   map[string][]string
	images  map[string]*moby.Image
	files   map[string]*moby.File
	trust   map[string]*moby.TrustConfig

	sync.Mutex
}

func configureProvider(*schema.ResourceData) (meta interface{}, err error) {
	moby.MobyDir, err = defaultMobyConfigDir()
	if err != nil {
		return
	}

	err = errors.Wrap(os.MkdirAll(moby.MobyDir, 0755), "could not create config directory")
	if err != nil {
		return
	}

	err = errors.Wrap(os.MkdirAll(filepath.Join(moby.MobyDir, "tmp"), 0755), "could not create config tmp directory")
	if err != nil {
		return
	}

	return
}

func defaultMobyConfigDir() (string, error) {
	mobyDefaultDir := ".moby"
	home, err := homedir.Dir()
	return filepath.Join(home, mobyDefaultDir), err
}

func (c *cache) addKernel(k *moby.KernelConfig) string {
	c.Lock()
	defer c.Unlock()

	id := id(k)
	c.kernels[id] = k

	return id
}

func (c *cache) addInit(i []string) string {
	c.Lock()
	defer c.Unlock()

	id := id(i)
	c.inits[id] = i

	return id
}

func (c *cache) addImage(i *moby.Image) string {
	c.Lock()
	defer c.Unlock()

	id := id(i)
	c.images[id] = i

	return id
}

func (c *cache) addFile(i *moby.File) string {
	c.Lock()
	defer c.Unlock()

	id := id(i)
	c.files[id] = i

	return id
}

func (c *cache) addConfig(m *moby.Moby) string {
	c.Lock()
	defer c.Unlock()

	id := id(m)
	c.configs[id] = m

	return id
}

func (c *cache) addTrust(t *moby.TrustConfig) string {
	c.Lock()
	defer c.Unlock()

	id := id(t)
	c.trust[id] = t

	return id
}

func id(input interface{}) string {
	b, _ := json.Marshal(input)
	return hash(string(b))
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}

func stringPtr(s string) *string {
	return &s
}

func copyFile(src, dst string) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}

	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer df.Close()

	_, err = io.Copy(df, sf)
	if err != nil {
		return err
	}

	return nil
}
