//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_cr "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/container_registry"
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
)

// TestContainerRegistry object CRUD
func TestContainerRegistry(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_ns.ObjectType,
		ves_io_schema_cr.ObjectType,
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
				Config: testContainerRegistry(name, "app-test"),
			},
		},
	})
}

func testContainerRegistry(name, namespace string) string {
	return fmt.Sprintf(`
		resource "volterra_namespace" "app" {
			name = "%[2]s"
		}
		resource "volterra_container_registry" "cr" {
			name = "%[1]s"
			namespace = volterra_namespace.app.name
			email = "user@example.com"
			password {
				clear_secret_info {
					provider = "test"
					url = "string:///dGVzdAo="
				}
			}
			registry = "docker.io"
			user_name = "user"
		}

	`, name, namespace)
}
