package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/skybet/terraform-provider-linuxkit/linuxkit"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: linuxkit.Provider})
}
