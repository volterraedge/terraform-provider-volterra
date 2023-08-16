// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
)

func TestAccDataSourceNamespaceBasic(t *testing.T) {
	t.Skipf("Skipping: TBD debug terraform provider issue and fix it later")
	name := generateResourceName()
	//resourceName := "volterra_namespace." + name
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{ves_io_schema_ns.ObjectType})
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
				Config: testConfigDataSourceNamespace(ves_io_schema_ns.ObjectType, name),
				//Check: resource.ComposeTestCheckFunc(
				//	resource.TestCheckResourceAttr(resourceName, "name", name),
				//),
			},
		},
	})
}

func testConfigDataSourceNamespace(objType, name string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
	resource "%[1]s" "%[2]s" {
		name = "%[2]s"
	}

	data "%[1]s" "%[2]s" {
		name = "%[2]s"
	}`, tfResourceName, name)
}
