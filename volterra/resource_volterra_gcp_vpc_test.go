//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_cloud_cred "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/cloud_credentials"
	ves_io_schema_gcp_vpc_site "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/views/gcp_vpc_site"
)

func TestGCPVPCSiteState(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTestCustomAPIServer(t, []string{
		ves_io_schema_cloud_cred.ObjectType,
		ves_io_schema_gcp_vpc_site.ObjectType,
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
				Config: testGCPVPCSiteConfig(name, siteName),
			},
		},
	})
}

func testGCPVPCSiteConfig(resourceName, name string) string {
	return fmt.Sprintf(`
		resource "volterra_cloud_credentials" "%[1]s" {
		  name = "%[2]s"
		  namespace = "system"
		  gcp_cred_file {
			credential_file {
				clear_secret_info {
					url = "string:///aXRodWIuY29tL2hhc2hpY29ycC8K"
				}
			}
		  }
		}
		resource "volterra_gcp_vpc_site" "%[1]s" {
		  name = "%[2]s"
		  namespace = "system"
		  cloud_credentials {
			name = "%[2]s"
			namespace = "system"
		  }
		  gcp_region = "us-east1"
		  instance_type = "n1-standard-4"

		  ingress_egress_gw {
			gcp_certified_hw = "gcp-byol-multi-nic-voltmesh"
			gcp_zone_names = ["us-east1-b"]
			inside_network {
				new_network {
					name = "inside-network"
				}
			}
			inside_subnet{
				new_subnet {
					primary_ipv4 = "192.168.0.0/24"
					subnet_name = "inside-subnet"
				}
			}
			outside_network {
				new_network {
					name = "outside-network"
				}
			}
			outside_subnet{
				new_subnet {
					primary_ipv4 = "192.168.1.0/24"
					subnet_name = "outside-subnet"
				}
			}
			no_global_network = true
			no_inside_static_routes = true
			no_network_policy = true
			no_outside_static_routes = true
			no_forward_proxy = true
		  }
		  ssh_key = "ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAD5sRjpKVBvf5afxhysXd4GyvEFaiDOnPhKQcK8SHNUxkGkjhRV6xMFpBBApNctQ73yaHweV//OhBHurwzUodKOWAEyH+ay0V2BAOpx2aiQHxiMh7b0CGYVxv4lRZ4IPZ1Da9Siz1Sz19RYBjVM7v6Dvo2UlYftUyauKPIDPnd19iN10g=="

		}
		resource "volterra_tf_params_action" "%[1]s" {
			site_name = "%[2]s"
			site_kind = "gcp_vpc_site"
			action = "plan"
			wait_for_action = false
		}
		`, resourceName, name)
}
