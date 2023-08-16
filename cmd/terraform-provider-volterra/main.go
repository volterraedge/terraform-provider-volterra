// Copyright (c) 2023 F5 Inc. All rights reserved.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/volterraedge/terraform-provider-volterra/volterra"
)

func main() {
	pluginServeOpts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return volterra.Provider()
		},
	}
	var debugMode bool
	flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/volterraedge/volterra", pluginServeOpts)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		plugin.Serve(pluginServeOpts)
	}

}
