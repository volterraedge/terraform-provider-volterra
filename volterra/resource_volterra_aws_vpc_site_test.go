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
	ves_io_schema_aws_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
)

func TestAWSVPCSiteState(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTestCustomAPIServer(t, []string{
		ves_io_schema_cloud_cred.ObjectType,
		ves_io_schema_aws_vpc_site.ObjectType,
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
				Config: testAWSVPCSiteConfig(name, siteName),
			},
		},
	})
}

func testAWSVPCSiteConfig(resourceName, name string) string {
	return fmt.Sprintf(`
		resource "volterra_cloud_credentials" "%[1]s" {
		  name = "%[2]s"
		  namespace = "system"
		  aws_secret_key {
			access_key = "HFJGTOLITYSX"
			secret_key {
				clear_secret_info {
					url = "string:///aXRodWIuY29tL2hhc2hpY29ycC8K"
				}
			}
		  }
		}
		resource "volterra_aws_vpc_site" "%[1]s" {
		  name = "%[2]s"
		  namespace = "system"
		  aws_region = "us-west-2"
		  aws_cred {
			name = "%[2]s"
			namespace = "system"
		  }
		  vpc {
			new_vpc {
				name_tag = "%[2]s"
				primary_ipv4 = "192.168.0.0/22"
			}
		  }
		  disk_size = 80
		  instance_type = "t3.xlarge"

		  ingress_egress_gw {
			aws_certified_hw = "aws-byol-multi-nic-voltmesh"
			az_nodes {
				aws_az_name = "us-west-2a"
				disk_size = "100"
				inside_subnet {
					subnet_param {
						ipv4 = "192.168.0.0/24"
					}
				}
				outside_subnet {
					subnet_param {
						ipv4 = "192.168.1.0/24"
					}
				}
			}
			no_global_network = true
			no_inside_static_routes = true
			no_network_policy = true
			no_outside_static_routes = true
			no_forward_proxy_policy = true
		  }
		  ssh_key = "ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAD5sRjpKVBvf5afxhysXd4GyvEFaiDOnPhKQcK8SHNUxkGkjhRV6xMFpBBApNctQ73yaHweV//OhBHurwzUodKOWAEyH+ay0V2BAOpx2aiQHxiMh7b0CGYVxv4lRZ4IPZ1Da9Siz1Sz19RYBjVM7v6Dvo2UlYftUyauKPIDPnd19iN10g=="

		}
		resource "volterra_tf_params_action" "%[1]s" {
			site_name = "%[2]s"
			site_kind = "aws_vpc_site"
			action = "plan"
			wait_for_action = false
		}
		`, resourceName, name)
}
