//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_service_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy"
)

func TestAccDataSourceServicePolicyBasic(t *testing.T) {
	name := generateResourceName()
	resourceName := "data.volterra_service_policy." + name
	testURL, stopFunc := createTestCustomAPIServer(t, []string{ves_io_schema_service_policy.ObjectType})
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
				Config: testConfigDataSourceServicePolicy(ves_io_schema_service_policy.ObjectType, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testConfigDataSourceServicePolicy(objType, name string) string {
	// convert ves.io.schema.vservice_policy.Object to volt_service_policy
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
	resource "%[1]s" "%[2]s" {
		name = "%[2]s"
		namespace = "system"
	}

	data "%[1]s" "%[2]s" {
		name = "%[2]s"
		namespace = "system"
	}`, tfResourceName, name)
}
