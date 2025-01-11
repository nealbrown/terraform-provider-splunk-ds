package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/nealbrown/terraform-provider-splunk-ds/provider"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/nealbrown/splunk",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.NewMuxProvider("0.1.0"), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
