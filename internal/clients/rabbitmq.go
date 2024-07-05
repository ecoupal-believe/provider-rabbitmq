/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/believe/provider-rabbitmq/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal rabbitmq credentials as JSON"

	username        = "username"
	password        = "password"
	endpoint        = "endpoint"
	insecure        = "insecure"
	cacert_file     = "cacert_file"
	clientcert_file = "clientcert_file"
	clientkey_file  = "clientkey_file"
	proxy           = "proxy"
)

func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]any{}
		if v, ok := creds[username]; ok {
			ps.Configuration[username] = v
		}
		if v, ok := creds[password]; ok {
			ps.Configuration[password] = v
		}
		if v, ok := creds[endpoint]; ok {
			ps.Configuration[endpoint] = v
		}
		if v, ok := creds[insecure]; ok {
			ps.Configuration[insecure] = v
		}
		if v, ok := creds[cacert_file]; ok {
			ps.Configuration[cacert_file] = v
		}
		if v, ok := creds[clientcert_file]; ok {
			ps.Configuration[clientcert_file] = v
		}
		if v, ok := creds[clientkey_file]; ok {
			ps.Configuration[clientkey_file] = v
		}
		if v, ok := creds[proxy]; ok {
			ps.Configuration[proxy] = v
		}

		return ps, nil
	}
}
