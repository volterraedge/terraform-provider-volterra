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
				Config: testAWSCloudCredConfig(name, "aws-cred"),
			},
			{
				Config: testAWSVPCSiteIEConfig(name, "aws-cred", siteName),
			},
			{
				Config: testTFActionParams(name, siteName),
			},
			{
				Config: testAWSVPCSiteIConfig(fmt.Sprintf("%s_two", name), "aws-cred", "aws-i-site"),
			},
			{
				Config: testTFActionParams(fmt.Sprintf("%s_two", name), "aws-i-site"),
			},
		},
	})
}

func testAWSCloudCredConfig(resourceName, name string) string {
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
		}`, resourceName, name)
}

func testAWSVPCSiteIEConfig(resourceName, cloudCred, name string) string {
	return fmt.Sprintf(`
		resource "volterra_aws_vpc_site" "%[1]s" {
		  name = "%[2]s"
		  namespace = "system"
		  aws_region = "us-west-2"
		  aws_cred {
			name = "%[3]s"
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
						ipv4 = "192.168.0.0/25"
					}
				}
				outside_subnet {
					subnet_param {
						ipv4 = "192.168.1.0/25"
					}
				}
			}
			az_nodes {
				aws_az_name = "us-west-2b"
				disk_size = "100"
				inside_subnet {
					subnet_param {
						ipv4 = "192.168.0.128/25"
					}
				}
				outside_subnet {
					subnet_param {
						ipv4 = "192.168.1.128/25"
					}
				}
			}
			az_nodes {
				aws_az_name = "us-west-2c"
				disk_size = "100"
				inside_subnet {
					subnet_param {
						ipv4 = "192.168.2.128/25"
					}
				}
				outside_subnet {
					subnet_param {
						ipv4 = "192.168.3.128/25"
					}
				}
			}
			no_global_network = true
			no_inside_static_routes = true
			no_network_policy = true
			no_outside_static_routes = true
			no_forward_proxy = true
			allowed_vip_port {
				custom_ports {
					port_ranges = "8080-8082,9095"
				}
			}
		  }
		  coordinates {
			latitude = 37.404989
			longitude = -121.942300
		  }
		  logs_streaming_disabled = true
		  ssh_key = "ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAD5sRjpKVBvf5afxhysXd4GyvEFaiDOnPhKQcK8SHNUxkGkjhRV6xMFpBBApNctQ73yaHweV//OhBHurwzUodKOWAEyH+ay0V2BAOpx2aiQHxiMh7b0CGYVxv4lRZ4IPZ1Da9Siz1Sz19RYBjVM7v6Dvo2UlYftUyauKPIDPnd19iN10g=="
		  total_nodes = 2

		}`, resourceName, name, cloudCred)
}

func testAWSVPCSiteIConfig(resourceName, cloudCred, name string) string {
	return fmt.Sprintf(`
	resource "volterra_aws_vpc_site" "%[1]s" {
		name            = "%[2]s"
		namespace       = "system"
		aws_region = "us-east-1"
  		aws_cred {
  		  name      = "%[3]s"
  		  namespace = "system"
  		}
  		instance_type = "t3.xlarge"
  		vpc {
  		  new_vpc {
  		    name_tag = "%[2]s"
  		    primary_ipv4 = "192.168.0.0/22"
  		  }
  		}
		ingress_gw {
		  aws_certified_hw = "aws-byol-voltmesh"
		  az_nodes {
		    aws_az_name = "us-east-1a"
		    disk_size   = 20
		    local_subnet {
		  	  subnet_param {
		  	    ipv4 = "192.168.0.0/24"
		  	  }
		    }
		  }
		}
		logs_streaming_disabled = true
		no_worker_nodes = true
		coordinates {
			latitude = 37.404989
			longitude = -121.942300
		}
	  }`, resourceName, name, cloudCred)
}

func testTFActionParams(resourceName, name string) string {
	return fmt.Sprintf(`
	resource "volterra_tf_params_action" "%[1]s" {
		site_name = "%[2]s"
		site_kind = "aws_vpc_site"
		action = "plan"
		wait_for_action = false
	}`, resourceName, name)
}
