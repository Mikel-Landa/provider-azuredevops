/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/Mikel-Landa/provider-azuredevops/config/git"
	"github.com/Mikel-Landa/provider-azuredevops/config/project"
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "azuredevops"
	modulePath     = "github.com/Mikel-Landa/provider-azuredevops"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("azuredevops.upbound.io"),
		ujconfig.WithShortName("azuredevops"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))
	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		// repository.Configure,
		project.Configure,
		git.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc
}
