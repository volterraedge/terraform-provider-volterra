//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestResourceVPCTunnels(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTestCustomAPIServer(t, []string{})
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
				Config: testVPNTunnels(name),
			},
		},
	})
}

func testVPNTunnels(name string) string {

	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  namespace = "system"
		  vpn_tunnel_config {
			node_name = "master-0"
			tunnel_remote_ips = ["192.168.0.1"]
			node_id = "ves-node-id-xxxxxx"
			type = "SERVICES"
		  }
		}
		`, setVPNTunnels, name, name)
}
