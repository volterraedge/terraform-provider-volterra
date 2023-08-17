// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_voltstack_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/voltstack_site"
)

func TestVoltstackSite(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_voltstack_site.ObjectType,
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
				Config: testVoltStackSiteConfig(name),
			},
		},
	})
}

func testVoltStackSiteConfig(name string) string {
	return fmt.Sprintf(`
		resource "volterra_voltstack_site" "this" {
		  name = "%[1]s"
		  namespace = "system"
		  no_bond_devices = true
		  disable_gpu = true
		  no_k8s_cluster = true
		  master_nodes = [ "master-0" ]
		  custom_network_config {
			forward_proxy_allow_all = true
			no_global_network = true
			default_interface_config = true
			no_network_policy = true
			slo_config {
				no_dc_cluster_group = true
				no_static_routes = true
			}
		  }
		  enable_vm = true
		  default_network_config = false
		  default_storage_config = true
		  volterra_certified_hw = "generic-single-nic-volstack-combo"
		  logs_streaming_disabled = true
		  allow_all_usb = true
		  no_local_control_plane = true
		}`, name)
}
