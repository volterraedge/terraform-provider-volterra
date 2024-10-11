// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_netw_intf "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_interface"
	ves_io_schema_tenant "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/tenant"
)

func TestNetworkInterface(t *testing.T) {
	name := generateResourceName()
	resourceName := "volterra_network_interface." + name
	testURL, stopFunc, f := createTestCustomAPIServer(t, []string{ves_io_schema_netw_intf.ObjectType, ves_io_schema_tenant.ObjectType})
	defer stopFunc()
	tenantName := "ves-io"
	tenantObj := mkDBObjTenant(tenantName, uuid.New().String(), withAddonServices("f5xc-flow-collection", "f5xc-ipv6-standard", "f5xc-waap-standard"))
	f.MustCreateEntry(tenantObj)

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
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testConfigTunnelInterfaceInterface(ves_io_schema_netw_intf.ObjectType, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "tunnel_interface.0.tunnel.0.name", "my-example-label1"),
					resource.TestCheckResourceAttr(resourceName, "tunnel_interface.0.cluster", "false"),
					resource.TestCheckResourceAttr(resourceName, "tunnel_interface.0.node", "master-0"),
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

func testConfigTunnelInterfaceInterface(objType, name string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  namespace = "system"
		  tunnel_interface {
			mtu = "1450"
			site_local_network = true
			cluster  = false
			node = "master-0"
			priority = "42"
			static_ip {
			  node_static_ip {
				default_gw = "192.168.20.1"
				dns_server = "192.168.20.1"
				ip_address = "192.168.20.1/24"
			  }
			}
			tunnel {
			  name      = "my-example-label1"
			  namespace = "system"
			}
		  }
		}
		`, tfResourceName, name, name)
}
