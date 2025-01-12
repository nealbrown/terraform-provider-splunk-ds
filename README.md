# terraform-provider-splunk-ds

New splunk provider using terraform plugin framework

This is an attempt to implement muxing per https://developer.hashicorp.com/terraform/plugin/framework/migrating/mux to allow me to utilize the resources and auth function in the old splunk provider (https://github.com/splunk/terraform-provider-splunk) while building new resources (specifically server classes and deployment apps) in the new terraform plugin framework.
