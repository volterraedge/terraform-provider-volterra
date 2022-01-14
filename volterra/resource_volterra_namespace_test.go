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
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
)

func TestAccNamespaceBasic(t *testing.T) {
	name := generateResourceName()
	resourceName := "volterra_namespace." + name
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
				Config: testConfigNS(ves_io_schema_ns.ObjectType, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testConfigNS(objType, name string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		}
		`, tfResourceName, name, name)
}
