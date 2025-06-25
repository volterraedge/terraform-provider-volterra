// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_cloud_cred "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cloud_credentials"
	ves_io_schema_tenant "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/tenant"
	ves_io_schema_aws_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
)

func TestAWSVPCSiteState(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, f := createTestCustomAPIServer(t, []string{
		ves_io_schema_cloud_cred.ObjectType,
		ves_io_schema_aws_vpc_site.ObjectType,
		ves_io_schema_tenant.ObjectType,
	})
	defer stopFunc()
	tenantName := "ves-io"
	tenantObj := mkDBObjTenant(tenantName, uuid.New().String(), withAddonServices("f5xc-flow-collection", "f5xc-ipv6-standard", "f5xc-waap-standard"))
	f.MustCreateEntry(tenantObj)

	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("TF_LOG", "")
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
		  lifecycle {
			ignore_changes = [labels]
		  }

		}
		resource "volterra_cloud_site_labels" "%[1]s" {
			name = volterra_aws_vpc_site.%[1]s.name
			site_type = "aws_vpc_site"
			labels = {
				key1 = "value1"
				key2 = "value2"
			}
			ignore_on_delete = true
		}
		resource "volterra_set_cloud_site_info" "%[1]s" {
			name = volterra_aws_vpc_site.%[1]s.name
			site_type = "aws_vpc_site"
			public_ips = ["10.0.0.1"]
			private_ips = ["192.168.0.1"]
			vpc_id = "vpc-1a2b3c4d"
			subnet_ids {
				outside_subnet_id = "subnet-12345678"
				inside_subnet_id = "subnet-45678123"
				workload_subnet_id = "subnet-56781234"
				az = "us-east-2a"
				outside_subnet {
					az_name = "us-east-2a"
					id = "subnet-12345678"
					ipv4_prefix = "10.0.0.0/24"
				}
				inside_subnet {
					az_name = "us-east-2a"
					id = "subnet-45678123"
					ipv4_prefix = "10.0.1.0/24"
				}
				workload_subnet {
					az_name = "us-east-2a"
					id = "subnet-56781234"
					ipv4_prefix = "10.0.2.0/24"
				}
			}
			direct_connect_info {
				direct_connect_gateway_id = "1234-6789-0987-1234-09876"
				vgw_id = "vgw-12345678"
				asn = 64512
			}

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
		    local_subnet {
		  	  subnet_param {
		  	    ipv4 = "192.168.0.0/24"
		  	  }
		    }
		  }
		}
		ssh_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDWf4iSk7KTSYwWYlGq+1PyPRZ2soqMfWztm90shv66s8y4N9CSJzpHx1FXbMb+iIT9jxQ2Qt8upovaM1d3Y0lTNC0o9eKKwGMLcg5HG3dIu1r8RcakbL9XlbS+XOapx8BNTCNa7kx2X1BD8Zibzr7JpBokrd18A4GeWtX8EjLfCzvLRMl1UKIGmaSGm49yYWLONZwFdQJkvUb0eogM4LZ5B7/KgUh/a8+j9LmflflkyvkcoEOIJ847iNzswWPUMwTcSmaG174DSnhFRus+a5yixnXjywvwu0zOGoFFITSmsG+VxmI42EwS58TRNDC1PDB3l25LSulzXgoBcopz7Y1lMPfoqxuubZ2DPKDYGvmuUURXej0HV1+cb33IxOMwOluPzXmXvBpNPIpVJvwZUhwVHQYhnK07IDp8m47Bw/7DIUm2hb6cvBQ/rMkCMgeFNbvPEgxKm73aZkA0zLUEUo32BltbGEJVQwouuf5Ggon7i/cn+OosF3FiPtkL0Oklqo8="
		logs_streaming_disabled = true
		no_worker_nodes = true
		coordinates {
			latitude = 37.404989
			longitude = -121.942300
		}
	  }
	  resource "volterra_set_cloud_site_info" "%[1]s" {
		name = volterra_aws_vpc_site.%[1]s.name
		site_type = "aws_vpc_site"
		public_ips = ["10.0.0.1"]
		private_ips = ["192.168.0.1"]

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
