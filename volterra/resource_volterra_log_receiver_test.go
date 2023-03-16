// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_lr "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/log_receiver"
)

// TestLogReceiver object CRUD
func TestLogReceiver(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_lr.ObjectType,
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
				Config: testLogReceiver(name),
			},
		},
	})
}

func testLogReceiver(name string) string {
	return fmt.Sprintf(`
		resource "volterra_log_receiver" "lr" {
			name = "%s"
			namespace = "system"
			syslog {
				tcp_server {
					port = 2000
					server_name = "server.test.this"
				}
				syslog_rfc5424 = 500

			}
			site_local = true
		}

	`, name)
}
