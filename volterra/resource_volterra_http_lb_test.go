//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_views_http_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
)

// TestAccHTTPLB token creation test
func TestHTTPLB(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTestServer(t, ves_io_schema_views_http_loadbalancer.ObjectType)
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
				Config: testConfigHTTPLB(name, name, systemNS),
			},
		},
	})
}

func testConfigHTTPLB(resourceName, name, namespace string) string {
	return fmt.Sprintf(`
		resource "volterra_http_loadbalancer" "%[1]s" {
		  name = "%[2]s"
		  namespace = "%[3]s"
		  advertise_on_public_default_vip = true
		  no_challenge = true
		  http = true
		  disable_rate_limit = true
		  disable_waf = true
		  no_service_policies = true
		  round_robin = true
		  domains = ["http.helloclouds.app"]
		}
		`, resourceName, name, namespace)
}
