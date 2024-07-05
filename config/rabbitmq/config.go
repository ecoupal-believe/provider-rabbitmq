package rabbitmq

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rabbitmq_user", func(r *config.Resource) {
		r.ShortGroup = "User"
		r.Kind = "User"
	})

	p.AddResourceConfigurator("rabbitmq_vhost", func(r *config.Resource) {
		r.ShortGroup = "Vhost"
		r.Kind = "Vhost"
	})

	p.AddResourceConfigurator("rabbitmq_permission", func(r *config.Resource) {
		r.ShortGroup = "Permission"
		r.Kind = "Permission"

		r.References["user"] = config.Reference{
			TerraformName: "rabbitmq_user",
		}
		r.References["vhost "] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})

	p.AddResourceConfigurator("rabbitmq_topic-permission", func(r *config.Resource) {
		r.ShortGroup = "TopicPermission"
		r.Kind = "TopicPermission"

		r.References["user"] = config.Reference{
			TerraformName: "rabbitmq_user",
		}
		r.References["vhost "] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})

	p.AddResourceConfigurator("rabbitmq_exchange", func(r *config.Resource) {
		r.ShortGroup = "Exchange"
		r.Kind = "Exchange"

		r.References["vhost "] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})

	p.AddResourceConfigurator("rabbitmq_queue", func(r *config.Resource) {
		r.ShortGroup = "Queue"
		r.Kind = "Queue"

		r.References["vhost "] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})

	p.AddResourceConfigurator("rabbitmq_binding", func(r *config.Resource) {
		r.ShortGroup = "Binding"
		r.Kind = "Binding"

		r.References["source"] = config.Reference{
			TerraformName: "rabbitmq_exchange",
		}
		r.References["vhost "] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
		r.References["destination "] = config.Reference{
			TerraformName: "rabbitmq_queue",
		}
	})

	p.AddResourceConfigurator("rabbitmq_policy", func(r *config.Resource) {
		r.ShortGroup = "Policy"
		r.Kind = "Policy"

		r.References["vhost "] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})

	p.AddResourceConfigurator("rabbitmq_operator-policy", func(r *config.Resource) {
		r.ShortGroup = "OperatorPolicy"
		r.Kind = "OperatorPolicy"

		r.References["vhost "] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})

	p.AddResourceConfigurator("rabbitmq_shovel", func(r *config.Resource) {
		r.ShortGroup = "Shovel"
		r.Kind = "Shovel"

		r.References["vhost "] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})

	p.AddResourceConfigurator("rabbitmq_federation-upstream", func(r *config.Resource) {
		r.ShortGroup = "FederationUpstream"
		r.Kind = "FederationUpstream"

		r.References["vhost "] = config.Reference{
			TerraformName: "rabbitmq_vhost",
		}
	})
}
