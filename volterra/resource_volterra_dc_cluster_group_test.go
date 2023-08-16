// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_dcg "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dc_cluster_group"
)

// TestDCClusterGrp object CRUD
func TestDCClusterGrp(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_dcg.ObjectType,
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
				Config: testDCClusterGrp(name),
			},
		},
	})
}

func testDCClusterGrp(name string) string {
	return fmt.Sprintf(`
		resource "volterra_dc_cluster_group" "dd" {
			name = "%s"
			namespace = "system"
		}

	`, name)
}
