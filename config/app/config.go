package app

import (
	"github.com/crossplane-contrib/provider-jet-okta/config/common"
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("okta_app_oauth", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["okta_app_oauth_id"] = []byte(a)
			}
			if a, ok := attr["client_id"].(string); ok {
				conn["okta_app_oauth_client_id"] = []byte(a)
			}
			if a, ok := attr["client_secret"].(string); ok {
				conn["okta_app_oauth_client_secret"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("okta_app_oauth_redirect_uri", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["app_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-okta/apis/app/v1alpha2.Oauth",
		}
	})

	p.AddResourceConfigurator("okta_app_group_assignments", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.ExternalName = config.IdentifierFromProvider
		r.References["app_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-okta/apis/app/v1alpha2.Oauth",
		}
	})
}
