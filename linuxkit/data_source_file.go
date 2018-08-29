package linuxkit

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/moby/tool/src/moby"
)

func fileDataSource() *schema.Resource {
	return &schema.Resource{
		Read: fileRead,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The path to create the file or directory",
				Required:    true,
				ForceNew:    true,
			},
			"directory": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If true a directory is created",
				Optional:    true,
				ForceNew:    true,
			},
			"symlink": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The path to link to",
				Optional:    true,
				ForceNew:    true,
			},
			"contents": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The contents of the file",
				Optional:    true,
				ForceNew:    true,
			},
			"source": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The path to the source of the file",
				Optional:    true,
				ForceNew:    true,
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Format to write image metadata, only yaml is currently supported",
				Optional:    true,
				ForceNew:    true,
			},
			"optional": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "File is optional, dont fail if source does not exist",
				Optional:    true,
				ForceNew:    true,
			},
			"mode": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The mode to create the file or directory",
				Optional:    true,
				ForceNew:    true,
			},
			"uid": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The uid of the file/directory owner",
				Optional:    true,
				ForceNew:    true,
			},
			"gid": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The gid of the file/directory owner",
				Optional:    true,
				ForceNew:    true,
			},
			"uid_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The uid name of the file/directory owner",
				Optional:    true,
				ForceNew:    true,
			},
			"gid_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The gid name of the file/directory owner",
				Optional:    true,
				ForceNew:    true,
			},
		},
	}
}

func fileRead(d *schema.ResourceData, meta interface{}) error {
	file := &moby.File{}

	if v, ok := d.GetOk("path"); ok {
		file.Path = v.(string)
	}

	if v, ok := d.GetOk("directory"); ok {
		file.Directory = v.(bool)
	}

	if v, ok := d.GetOk("symlink"); ok {
		file.Symlink = v.(string)
	}

	if v, ok := d.GetOk("contents"); ok {
		file.Contents = stringPtr(v.(string))
	}

	if v, ok := d.GetOk("source"); ok {
		file.Source = v.(string)
	}

	if v, ok := d.GetOk("metadata"); ok {
		file.Metadata = v.(string)
	}

	if v, ok := d.GetOk("optional"); ok {
		file.Optional = v.(bool)
	}

	if v, ok := d.GetOk("mode"); ok {
		file.Mode = v.(string)
	}

	if v, ok := d.GetOk("uid"); ok {
		file.UID = v
	}

	if v, ok := d.GetOk("gid"); ok {
		file.GID = v
	}

	if v, ok := d.GetOk("uid_name"); ok {
		file.UID = v
	}

	if v, ok := d.GetOk("uid_name"); ok {
		file.GID = v
	}

	d.SetId(globalCache.addFile(file))

	return nil
}
