package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	splunk "github.com/splunk/terraform-provider-splunk/splunk"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &splunkProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &splunkProvider{
			version: version,
		}
	}
}

// splunkProvider is the provider implementation.
type splunkProvider struct {
	version string
}

// Metadata returns the provider type name.
func (p *splunkProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "splunk"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *splunkProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// Configure prepares a Splunk API client for data sources and resources.
func (p *splunkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config struct {
		Host     string `tfsdk:"host"`
		Username string `tfsdk:"username"`
		Password string `tfsdk:"password"`
	}

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new Splunk client using the configuration values
	client, err := splunk.NewClient(&splunk.Config{
		Host:     config.Host,
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Splunk API Client",
			"An unexpected error occurred when creating the Splunk API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Splunk Client Error: "+err.Error(),
		)
		return
	}

	// Make the Splunk client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// Resources returns the resource implementations supported by the provider.
func (p *splunkProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// Add resource functions here
	}
}

// DataSources returns the data source implementations supported by the provider.
func (p *splunkProvider) DataSources(_ context.Context) []func() resource.DataSource {
	return []func() resource.DataSource{
		// Add data source functions here
	}
}
