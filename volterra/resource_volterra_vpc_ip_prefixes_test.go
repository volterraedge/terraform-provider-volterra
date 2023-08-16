// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_aws_tgw_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
)

func TestResourceVPCIPPrefixes(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
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
				Config: testVPCIPPrefixes(name),
			},
		},
	})
}

func testVPCIPPrefixes(name string) string {

	return fmt.Sprintf(`
		resource "volterra_aws_tgw_site" "%[2]s" {
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
			vpc_attachments {
				vpc_list {
					labels = {
						"env" = "dev"
					}
					vpc_id = "vpc-0b79db47b7529876d"
				}
			}
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
		  }
		resource "%[1]s" "%[2]s" {
		  name = volterra_aws_tgw_site.%[2]s.name
		  namespace = "system"
		  vpc_ip_prefixes {
			  name = "vpc-0b79db47b7529876d"
			  value = ["192.168.0.0/22"]
		  }
		}
		`, setVPCIPPrefixes, name)
}
