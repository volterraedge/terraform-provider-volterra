// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_cloud_cred "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cloud_credentials"
	ves_io_schema_azure_vnet_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/azure_vnet_site"
)

func TestAzureVNETSiteState(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_cloud_cred.ObjectType,
		ves_io_schema_azure_vnet_site.ObjectType,
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
				Config: testAzureVNETSiteConfig(name, siteName),
			},
		},
	})
}

func testAzureVNETSiteConfig(resourceName, name string) string {
	return fmt.Sprintf(`
		resource "volterra_cloud_credentials" "%[1]s" {
		  name = "%[2]s"
		  namespace = "system"
		  azure_client_secret {
			client_id = "HFJGTOLITYSX"
			subscription_id = "HFJGTOLnetdsfssITYSX"
			tenant_id = "123456789"
			client_secret {
				clear_secret_info {
					url = "string:///aXRodWIuY29tL2hhc2hpY29ycC8K"
				}
			}
		  }
		}
		resource "volterra_azure_vnet_site" "%[1]s" {
		  name = "%[2]s"
		  namespace = "system"
		  azure_region = "eastus"
		  resource_group = "%[2]s"
		  azure_cred {
			name = "%[2]s"
			namespace = "system"
		  }
		  vnet {
			new_vnet {
				name = "%[2]s"
				primary_ipv4 = "192.168.0.0/22"
			}
		  }
		  disk_size = 80
		  machine_type = "Standard_D4_v2"

		  ingress_egress_gw {
			azure_certified_hw = "azure-byol-multi-nic-voltmesh"
			az_nodes {
				azure_az = "1"
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
				azure_az = "2"
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
				azure_az = "3"
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
		  }
		  logs_streaming_disabled = true
		  ssh_key = "ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAD5sRjpKVBvf5afxhysXd4GyvEFaiDOnPhKQcK8SHNUxkGkjhRV6xMFpBBApNctQ73yaHweV//OhBHurwzUodKOWAEyH+ay0V2BAOpx2aiQHxiMh7b0CGYVxv4lRZ4IPZ1Da9Siz1Sz19RYBjVM7v6Dvo2UlYftUyauKPIDPnd19iN10g=="
		  nodes_per_az = 2
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
			site_kind = "azure_vnet_site"
			action = "plan"
			wait_for_action = false
		}
		resource "volterra_cloud_site_labels" "%[1]s" {
			name = volterra_azure_vnet_site.%[1]s.name
			site_type = "azure_vnet_site"
			labels = {
				key1 = "value1"
				key2 = "value2"
			}
			ignore_on_delete = true
		}
		resource "volterra_set_cloud_site_info" "%[1]s" {
			name = volterra_azure_vnet_site.%[1]s.name
			site_type = "azure_vnet_site"
			public_ips = ["10.0.0.1"]
			private_ips = ["192.168.0.1"]
			spoke_vnet_prefix_info {
				vnet_name = "spoke-vnet-1"
				resource_group = "spoke-rg"
				prefixes = ["192.168.0.0/22"]
			}
			express_route_info {
				route_server_ips = ["10.1.1.1", "10.1.1.2"]
				route_server_asn = 64512
			}
			node_info {
				node_instance_name = "name1"
				node_id = "id1"
			}
			node_info {
				node_instance_name = "name2"
				node_id = "id2"
			}
			vnet {
				vnet_name = "net1"
				resource_id = "account1/test1/net1"
			}

		}`, resourceName, name)
}
