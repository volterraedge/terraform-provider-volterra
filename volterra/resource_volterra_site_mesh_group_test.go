//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_smg "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site_mesh_group"
)

// TestSiteMeshGroup object CRUD
func TestSiteMeshGroup(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_smg.ObjectType,
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
				Config: testSiteMeshGrpConfig(name),
			},
		},
	})
}

func testSiteMeshGrpConfig(name string) string {
	return fmt.Sprintf(`
		resource "volterra_site_mesh_group" "smg-full-mesh" {
			name = "%s"
			namespace = "system"
			tunnel_type = "SITE_TO_SITE_TUNNEL_IPSEC"
			type = "SITE_MESH_GROUP_TYPE_FULL_MESH"
			virtual_site {
				name = "test"
				namespace = "system"
			}
		}

	`, name)
}
