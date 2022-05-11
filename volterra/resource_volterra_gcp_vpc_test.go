//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_cloud_cred "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cloud_credentials"
	ves_io_schema_gcp_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/gcp_vpc_site"
)

func TestGCPVPCSiteState(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
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
			node_number = 1
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
		  coordinates {
			latitude = 37.404989
			longitude = -121.942300
		  }
		  logs_streaming_disabled = true
		  ssh_key = "ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAD5sRjpKVBvf5afxhysXd4GyvEFaiDOnPhKQcK8SHNUxkGkjhRV6xMFpBBApNctQ73yaHweV//OhBHurwzUodKOWAEyH+ay0V2BAOpx2aiQHxiMh7b0CGYVxv4lRZ4IPZ1Da9Siz1Sz19RYBjVM7v6Dvo2UlYftUyauKPIDPnd19iN10g=="
		  lifecycle {
			ignore_changes = [labels]
		  }

		}
		resource "volterra_tf_params_action" "%[1]s" {
			site_name = "%[2]s"
			site_kind = "gcp_vpc_site"
			action = "plan"
			wait_for_action = false
		}
		resource "volterra_cloud_site_labels" "%[1]s" {
			name = volterra_gcp_vpc_site.%[1]s.name
			site_type = "gcp_vpc_site"
			labels = {
				key1 = "value1"
				key2 = "value2"
			}
			ignore_on_delete = true
		}
		resource "volterra_set_cloud_site_info" "%[1]s" {
			name = volterra_gcp_vpc_site.%[1]s.name
			site_type = "gcp_vpc_site"
			public_ips = ["10.0.0.1"]
			private_ips = ["192.168.0.1"]

		}
		`, resourceName, name)
}
