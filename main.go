package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
}

// Imported SkyBet main.
// 
// 	"github.com/hashicorp/terraform/plugin"
// 	"github.com/skybet/terraform-provider-linuxkit/linuxkit"
// )

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{ProviderFunc: linuxkit.Provider})
// >>>>>>> f76c96debd2b18d68ae3c75c8a6b6075fb52e786
// }
