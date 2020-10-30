package main

import (
	"flag"
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"

	"github.com/resinstack/terraform-provider-linuxkit/linuxkit"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		err := plugin.Debug(context.Background(), "terraform.resinstack.io/resinstack/linuxkit",
			&plugin.ServeOpts{ProviderFunc: linuxkit.Provider})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		plugin.Serve(&plugin.ServeOpts{ProviderFunc: linuxkit.Provider})
	}
}
