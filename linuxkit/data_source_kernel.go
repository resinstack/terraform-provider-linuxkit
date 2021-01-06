package linuxkit

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/linuxkit/linuxkit/src/cmd/linuxkit/moby"
)

func kernelDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "`linuxkit_kernel` is a single kernel to be included in a `linuxkit_config`.",

		Read: kernelRead,

		Schema: map[string]*schema.Schema{
			"image": {
				Type:        schema.TypeString,
				Description: "The Docker image which should contain a kernel file that will be booted",
				Required:    true,
				ForceNew:    true,
			},
			"cmdline": {
				Type:        schema.TypeString,
				Description: "Kernel command line options if required.",
				Optional:    true,
				ForceNew:    true,
			},
			"binary": {
				Type:        schema.TypeString,
				Description: "Name of the kernel file that will be booted",
				Optional:    true,
				ForceNew:    true,
			},
			"tar": {
				Type:        schema.TypeString,
				Description: "Name of tarball that unpacked into the root.",
				Optional:    true,
				ForceNew:    true,
			},
			"ucode": {
				Type:        schema.TypeString,
				Description: "Name of cpio archive containing CPU microcode which needs prepending to the initrd.",
				Optional:    true,
				ForceNew:    true,
			},
		},
	}
}

func kernelRead(d *schema.ResourceData, meta interface{}) error {
	kernel := &moby.KernelConfig{
		Image:   d.Get("image").(string),
		Cmdline: d.Get("cmdline").(string),
		Binary:  d.Get("binary").(string),
	}

	if v, ok := d.GetOk("tar"); ok {
		s := v.(string)
		kernel.Tar = &s
	}

	if v, ok := d.GetOk("ucode"); ok {
		s := v.(string)
		kernel.UCode = &s
	}

	d.SetId(globalCache.addKernel(kernel))

	return nil
}
