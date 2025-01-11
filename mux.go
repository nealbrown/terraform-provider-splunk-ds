package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	splunk "github.com/splunk/terraform-provider-splunk/splunk"
)

func NewMuxProvider(version string) func() providerserver.Provider {
	return func() providerserver.Provider {
		return mux.NewMuxProvider(
			mux.WithProvider("splunk", splunk.New(version)),
			mux.WithProvider("splunk_ds", New(version)),
		)
	}
}
