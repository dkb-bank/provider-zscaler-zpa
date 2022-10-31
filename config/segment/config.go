package segment

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_segment_group", func(r *config.Resource) {
		r.ShortGroup = "segment"
	})
}
