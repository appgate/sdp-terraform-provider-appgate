package appgate

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppgateIdentityProviderDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: `
					data "appgatesdp_identity_provider" "test_identity_provider_ds" {
                        identity_provider_name = "local"
					}
                `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.appgatesdp_identity_provider.test_identity_provider_ds", "identity_provider_name"),
					resource.TestCheckResourceAttrSet("data.appgatesdp_identity_provider.test_identity_provider_ds", "identity_provider_id"),
				),
			},
		},
	})
}
