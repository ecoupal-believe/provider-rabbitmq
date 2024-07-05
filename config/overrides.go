// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"regexp"
	"strings"

	tjconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/name"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// VersionV1Beta1 is used to signify that the resource has been tested and external name configured
const VersionV1Beta1 = "v1beta1"

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

func externalNameConfig() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
	}
}

func groupOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		for k, v := range groupMap {
			ok, err := regexp.MatchString(k, r.Name)
			if err != nil {
				panic(errors.Wrap(err, "cannot match regular expression"))
			}
			if ok {
				r.ShortGroup, r.Kind = v(r.Name)
				return
			}
		}
	}
}

var groupMap = map[string]GroupKindCalculator{}

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		// "google_cloud_run_domain_mapping": "cloudrun" -> (cloudrun, DomainMapping)
		words := strings.Split(strings.TrimPrefix(resource, "google_"), "_")
		if group == "" {
			group = strings.Join(words[:count], "")
		}
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

func defaultVersion() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		r.Version = VersionV1Beta1
	}
}

func descriptionOverrides() tjconfig.ResourceOption {
	return func(r *tjconfig.Resource) {
		tjconfig.ManipulateEveryField(r.TerraformResource, func(sch *schema.Schema) {
			sch.Description = ""
		})
	}
}
