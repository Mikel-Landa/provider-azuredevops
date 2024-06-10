/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"azuredevops_project":             config.IdentifierFromProvider,
	"azuredevops_git_repository":      config.IdentifierFromProvider,
	"azuredevops_git_repository_file": config.TemplatedStringAsIdentifier("file", "{{ .parameters.repository_id }}/{{ .external_name }}:{{ .parameters.branch }}"),
}

// func gitRepositoryFileIDConf() config.ExternalName {
// 	e := config.IdentifierFromProvider
// 	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
// 		id, ok := tfstate["id"]
// 		if !ok {
// 			return "", errors.New("cannot get id")
// 		}
// 		val := strings.Join(strings.Split(id.(string), "/")[1:], "/")
// 		return val, nil
// 	}
// 	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
// 		keyVaultID, ok := parameters["key_vault_id"]
// 		if !ok {
// 			return "", errors.New("cannot get key_vault_id")
// 		}
// 		words := strings.Split(keyVaultID.(string), "/")
// 		keyVaultName := words[len(words)-1]

// 		if len(externalName) == 0 || len(strings.Split(externalName, "/")) < 2 {
// 			if parameters["version"] == nil || parameters["version"] == "" {
// 				return "", nil
// 			}
// 		}

// 		return fmt.Sprintf("https://%s.vault.azure.net/%s/%s",
// 			keyVaultName, resourceType, externalName), nil
// 	}
// 	return e
// }

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
