// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package managedidentity

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures managedidentity group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_user_assigned_identity", func(r *config.Resource) {
		// Make principal_id and client_id sensitive so the can be used in ConnectionDetails
		r.TerraformResource.Schema["principal_id"].Sensitive = true
		r.TerraformResource.Schema["client_id"].Sensitive = true

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return map[string][]byte{
				"principalId": []byte(attr["principal_id"].(string)),
				"clientId":    []byte(attr["client_id"].(string)),
			}, nil
		}
	})
}
