// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_fleet "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fleet"
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
	ves_io_schema_nc "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_connector"
	ves_io_schema_vn "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_network"
)

// TestFleet token creation test
func TestFleet(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_fleet.ObjectType,
		ves_io_schema_vn.ObjectType,
		ves_io_schema_nc.ObjectType,
		ves_io_schema_ns.ObjectType,
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
				Config: testConfigFleet(name, name, systemNS),
			},
		},
	})
}

func testConfigFleet(resourceName, name, namespace string) string {
	return fmt.Sprintf(`
		resource "volterra_namespace" "system" {
			name = "%[3]s"
		}
		resource "volterra_virtual_network" "sli_%[1]s" {
			name        = "%[2]s-sli"
			namespace = "%[3]s"
			site_local_inside_network = true
			static_routes {
				attrs = ["ROUTE_ATTR_INSTALL_HOST", "ROUTE_ATTR_INSTALL_FORWARDING"]
				ip_address = "1.1.1.1"
				ip_prefixes = ["1.1.1.0/24"]
			}
		}

		resource "volterra_virtual_network" "slo_%[1]s" {
			name        = "%[2]s-slo"
			namespace = "%[3]s"
			site_local_network = true
		}

		resource "volterra_network_connector" "connector_%[1]s" {
			name = "%[2]s"
			namespace = "%[3]s"
			sli_to_slo_snat {
				default_gw_snat = true
				interface_ip = true
			}
			disable_forward_proxy = true
		}

		resource "volterra_fleet" "fleet_%[1]s" {
		  name = "%[2]s"
		  fleet_label = "%[2]s"
		  namespace = "%[3]s"
		  no_bond_devices = true
		  no_dc_cluster_group = true
		  disable_gpu = true
		  default_config = true
		  default_storage_class = true
		  no_storage_device = true
		  no_storage_static_routes = true
		  no_storage_interfaces = true
		  logs_streaming_disabled = true
		  allow_all_usb = true
		  inside_virtual_network {
			name = volterra_virtual_network.sli_%[1]s.name
			namespace = volterra_namespace.system.name
			tenant    = volterra_namespace.system.tenant_name
		  }

		  outside_virtual_network {
			name = volterra_virtual_network.slo_%[1]s.name
			namespace = volterra_namespace.system.name
			tenant    = volterra_namespace.system.tenant_name
		  }
		  network_connectors {
			name      = volterra_network_connector.connector_%[1]s.name
			namespace = volterra_namespace.system.name
			tenant    = volterra_namespace.system.tenant_name
		  }
		}

		resource "volterra_fleet" "fleet_storage_test" {
			name = "fleet-storage-test"
			namespace = "system"
			fleet_label = "fleet-storage-test"
			disable_gpu = true
			logs_streaming_disabled = true
			deny_all_usb = true
			default_config = true
			no_storage_interfaces = true
			no_storage_static_routes = true
			no_bond_devices = true
			no_dc_cluster_group = true
			storage_class_list {
			  storage_classes {
				storage_class_name = "mayastor"
				storage_device = "sd-test"
				custom_storage {
					yaml = "https://example-yaml.com/custom-storeag.yaml"
				}
			  }
			}
			storage_device_list {
			  storage_devices {
				storage_device = "sd-test"
				custom_storage = true
			  }
			}
		  }
		`, resourceName, name, namespace)
}
