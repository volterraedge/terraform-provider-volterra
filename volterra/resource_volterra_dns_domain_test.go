// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_dns_domain "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dns_domain"
)

// TestDNSDomain object CRUD
func TestDNSDomain(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_dns_domain.ObjectType,
	})
	defer stopFunc()
	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "ves-io")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testDNSDomain(name),
			},
		},
	})
}

func testDNSDomain(name string) string {
	return fmt.Sprintf(`
		resource "volterra_dns_domain" "dd" {
			name = "%s"
			namespace = "system"
			dnssec_mode = "DNSSEC_DISABLE"
			volterra_managed = true
			verification_only = false
		}

	`, name)
}
