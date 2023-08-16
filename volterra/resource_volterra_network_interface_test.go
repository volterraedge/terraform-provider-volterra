// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_netw_intf "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_interface"
)

func TestNetworkInterface(t *testing.T) {
	name := generateResourceName()
	resourceName := "volterra_network_interface." + name
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{ves_io_schema_netw_intf.ObjectType})
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
				Config: testConfigSLIInterface(ves_io_schema_netw_intf.ObjectType, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testConfigSLOInterface(ves_io_schema_netw_intf.ObjectType, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testConfigSLIInterface(objType, name string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  namespace = "system"
		  ethernet_interface {
			static_ip {
			  cluster_static_ip {
				interface_ip_map {
				  name     = "master-0"
				  value {
				    ip_address = "10.0.0.5/24"
				  }
				}
			  }
			}
			device                    = "eth1"
			cluster                   = true
			site_local_inside_network = true
			not_primary               = true
			untagged                  = true
		  }
		}
		`, tfResourceName, name, name)
}

func testConfigSLOInterface(objType, name string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  namespace = "system"
		  ethernet_interface {
			device             = "eth0"
			cluster            = true
			site_local_network = true
			dhcp_client        = true
			is_primary         = true
			untagged           = true
		  }
		}
		`, tfResourceName, name, name)
}
