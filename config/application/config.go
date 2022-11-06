package application

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_application_segment", func(r *config.Resource) {
		r.ShortGroup = "application"
		r.References = config.References{
			"server_groups.id": {
				Type: "github.com/dkb-bank/provider-zscaler-zpa/apis/server/v1alpha1.Group",
			},
			"segment_group_id": {
				Type: "github.com/dkb-bank/provider-zscaler-zpa/apis/segment/v1alpha1.Group",
			},
		}
	})

	p.AddResourceConfigurator("zpa_application_server", func(r *config.Resource) {
		r.ShortGroup = "application"
		r.References = config.References{
			"app_server_group_ids": {
				Type:              "github.com/dkb-bank/provider-zscaler-zpa/apis/server/v1alpha1.Group",
				RefFieldName:      "AppServerGroupIDRefs",
				SelectorFieldName: "AppServerGroupIDSelector",
			},
		}
	})
}
