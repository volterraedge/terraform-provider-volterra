// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_dns_zone "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dns_zone"
)

func TestAccDNSZoneRecord(t *testing.T) {
	name := generateResourceName()
	resourceName := "volterra_dns_zone_record." + name
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{ves_io_schema_dns_zone.ObjectType})
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
				Config: testConfigDNSZoneRecord(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "group_name", "vestest-services"),
					resource.TestCheckResourceAttr(resourceName, "rrset.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rrset.0.a_record.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rrset.0.a_record.0.values.#", "1"),
				),
			},
		},
	})
}

func testConfigDNSZoneRecord(name string) string {
	return fmt.Sprintf(`
	resource "volterra_dns_zone" "system_yblviux3bvh7-co-jp" {
		primary {
		  rr_set_group {
		    metadata {
		      description = ""
		      name        = "vestest-services"
		    }
			}
		}
		name      = "%[1]s"
		namespace = "system"
	}
	resource "volterra_dns_zone_record" "%[1]s" {
		dns_zone_name = "%[1]s"
		group_name = "vestest-services"
		rrset {
			description = "test is written1234"
			ttl = "68"
			a_record {
				name = "%[1]s"
				values = ["1.5.1.1"]
			}
		}
	}
	`, name)
}
