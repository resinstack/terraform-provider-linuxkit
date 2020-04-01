package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func metadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: metadataRead,

		Schema: map[string]*schema.Schema{
			"base_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The base path to use when building an archive.",
				Required:    true,
				ForceNew:    true,
			},
			"json": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The packed version of the input file tree",
				Computed:    true,
			},
		},
	}
}

func metadataRead(d *schema.ResourceData, meta interface{}) error {
	type pathnode struct {
		Content string               `json:"content,omitempty"`
		Entries map[string]*pathnode `json:"entries,omitempty"`
		Perm    string               `json:"perm,omitempty"`
	}

	nodes := make(map[string]*pathnode)

	bd, ok := d.GetOk("base_path")
	if !ok {
		return errors.New("missing base_path")
	}
	baseDir := bd.(string)

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// The metadata module needs all the paths to be
		// relative to the base directory.  Once we have rpath
		// we write it into the struct and don't mess with it
		// anymore.
		rpath, err := filepath.Rel(baseDir, path)
		if err != nil {
			return nil
		}

		parent := filepath.Dir(rpath)
		base := filepath.Base(rpath)
		tmp := new(pathnode)
		if info.IsDir() {
			tmp.Entries = make(map[string]*pathnode)
			nodes[rpath] = tmp
		} else {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			tmp.Content = string(b[:])
		}

		// If the parent isn't the base, i.e. this isn't the
		// root node, then add this node to the list of
		// entries that this node maintains.
		if parent != base {
			nodes[parent].Entries[base] = tmp
		}

		return nil
	})
	if err != nil {
		return err
	}

	out, err := json.Marshal(nodes["."].Entries)
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())
	d.Set("json", string(out[:]))

	return nil
}
