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
	ves_io_schema_aws_tgw_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
)

func TestAWSTGWSiteState(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTestCustomAPIServer(t, []string{
		ves_io_schema_cloud_cred.ObjectType,
		ves_io_schema_aws_tgw_site.ObjectType,
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
				Config: testAWSTGWSiteConfig(name, siteName),
			},
		},
	})
}

func testAWSTGWSiteConfig(resourceName, name string) string {
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
		resource "volterra_aws_tgw_site" "%[1]s" {
			name = "%[2]s"
			namespace = "system"
			aws_parameters {
			  aws_certified_hw = "aws-byol-multi-nic-voltmesh"
			  aws_region = "us-east-1"
			  az_nodes {
				  aws_az_name = "us-east-1a"
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
			  aws_cred {
				  name = "%[2]s"
				  namespace = "system"
			  }
			  instance_type = "t3.xlarge"
			  new_vpc {
				  name_tag = "%[2]s"
				  primary_ipv4 = "192.168.0.0/22"
			  }
			  new_tgw {
				  system_generated = true
			  }
			  no_worker_nodes = true
			}
			logs_streaming_disabled = true
			tgw_security {
			  no_network_policy = true
			  no_forward_proxy = true
			  no_east_west_policy = true
			}
			vn_config {
			  no_global_network = true
			  no_inside_static_routes = true
			  no_outside_static_routes = true
			}
			coordinates {
				latitude = 37.404989
				longitude = -121.942300
			}
			lifecycle {
				ignore_changes = [labels]
			}
		  }
		  resource "volterra_tf_params_action" "%[1]s" {
			  site_name = "%[2]s"
			  site_kind = "aws_tgw_site"
			  action = "plan"
			  wait_for_action = false
		  }
		  resource "volterra_cloud_site_labels" "%[1]s" {
			name = volterra_aws_tgw_site.%[1]s.name
			site_type = "aws_tgw_site"
			labels = {
				key1 = "value1"
				key2 = "value2"
			}
			ignore_on_delete = true
		}

		`, resourceName, name)
}
