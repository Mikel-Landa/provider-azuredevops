package git

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuredevops_git_repository", func(r *config.Resource) {

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["project_id"] = config.Reference{
			Type: "github.com/Mikel-Landa/provider-azuredevops/apis/project/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("azuredevops_git_repository_file", func(r *config.Resource) {

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["project_id"] = config.Reference{
			Type: "github.com/Mikel-Landa/provider-azuredevops/apis/azuredevops/v1alpha1.Repository",
		}
	})

}
