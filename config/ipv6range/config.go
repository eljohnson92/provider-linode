package ipv6range

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_ipv6_range", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "linode"
		r.Kind = "IPv6range"
	})
}
