//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_network_firewall "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/network_firewall"
)

// TestAccNetworkFirewall token creation test
func TestAccNetworkFirewall(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTestServer(t, ves_io_schema_network_firewall.ObjectType)
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
				Config: testConfigNetworkFirewall(name, name, systemNS),
			},
		},
	})
}

func testConfigNetworkFirewall(resourceName, name, namespace string) string {
	return fmt.Sprintf(`
		resource "volterra_network_firewall" "%[1]s" {
		  name = "%[2]s"
		  namespace = "%[3]s"
		  active_forward_proxy_policies {
		    forward_proxy_policies {
		  	name      = "test1"
		  	namespace = "system"
		  	tenant    = "ves-io"
		    }
		  }
		  active_network_policies {
		    network_policies {
		  	name      = "test1"
		  	namespace = "system"
		  	tenant    = "ves-io"
		    }
		  }
		  disable_fast_acl = true
		}
		`, resourceName, name, namespace)
}
