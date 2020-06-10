package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/kevinvalk/terraform-provider-transip/transip"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: transip.Provider,
	})
}
