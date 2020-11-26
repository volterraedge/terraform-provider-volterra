//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_ap "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/advertise_policy"
	ves_io_schema_ns "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/namespace"
	ves_io_schema_vh "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
	ves_io_schema_vn "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/virtual_network"
)

func TestAccDataSourceVHDNSInfoBasic(t *testing.T) {
	name := generateResourceName()
	ns := "ns1"
	getTestObj := func() []string {
		return []string{
			ves_io_schema_vh.ObjectType, ves_io_schema_ns.ObjectType,
			ves_io_schema_vn.ObjectType, ves_io_schema_ap.ObjectType,
		}
	}
	testURL, stopFunc := createTestCustomAPIServer(t, getTestObj())
	defer stopFunc()
	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "ves-io")
	os.Setenv("TF_ACC", "true")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testConfigDataSourceVHDNSInfo(name, ns),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("volterra_namespace.%s", ns), "name", ns),
					resource.TestCheckResourceAttr("volterra_virtual_network.public", "name", "public"),
					resource.TestCheckResourceAttr(fmt.Sprintf("volterra_advertise_policy.%s", name), "name", name),
				),
			},
		},
	})
}

func testConfigDataSourceVHDNSInfo(name, ns string) string {
	return fmt.Sprintf(`
	resource "volterra_namespace" "shared" {
		name = "shared"
	}

	resource "volterra_namespace" "system" {
		name = "system"
	}

	resource "volterra_namespace" "%[2]s" {
		name = "%[2]s"
	}

	resource "volterra_virtual_network" "public" {
		name = "public"
		namespace = "shared"
		global_network = true
	}

	resource "volterra_advertise_policy" "%[1]s" {
		name = "%[1]s"
		namespace = "%[2]s"
		depends_on = ["volterra_virtual_network.public"]
		where {
			virtual_network {
				ref {
					name = "public"
					namespace = "shared"
					tenant = "ves-io"
				}
			}
		}
		port = 80
	}

	resource "volterra_virtual_host" "%[1]s" {
		name = "%[1]s"
		namespace = "%[2]s"
		advertise_policies {
			name = "%[1]s"
			namespace = "%[2]s"
			tenant = "ves-io"
		}
		domains = ["terraform.test.io"]
		proxy = "http_proxy"
		no_challenge = true
	}

	data "volterra_virtual_host_dns_info" "%[1]s" {
		name = "%[1]s"
		namespace = "%[2]s"
	}
	`, name, ns)
}
