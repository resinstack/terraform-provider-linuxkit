package linuxkit

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
)

type cache struct {
	configs map[string]*moby.Moby
	kernels map[string]*moby.KernelConfig
	inits   map[string][]string
	images  map[string]*moby.Image
	files   map[string]*moby.File

	sync.Mutex
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

func defaultLinuxkitCache() string {
	lktDir := ".linuxkit"

	// Technically this can error if the platform isn't defined,
	// but you can't run linuxkit on any of those platforms, so
	// I'm just discarding the error instead.
	home, _ := os.UserHomeDir()
	return filepath.Join(home, lktDir, "cache")
}
