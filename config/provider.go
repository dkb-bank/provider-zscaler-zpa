/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	application "github.com/dkb-bank/provider-zscaler-zpa/config/application"
	segment "github.com/dkb-bank/provider-zscaler-zpa/config/segment"
	server "github.com/dkb-bank/provider-zscaler-zpa/config/server"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "zscaler-zpa"
	modulePath     = "github.com/dkb-bank/provider-zscaler-zpa"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithShortName("zscaler-zpa"),
		ujconfig.WithRootGroup("zscaler-zpa.crossplane.io"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		application.Configure,
		segment.Configure,
		server.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
