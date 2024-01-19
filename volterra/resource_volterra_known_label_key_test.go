// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_known_label_key "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/known_label_key"
)

func TestAccKnownLabelKeyBasic(t *testing.T) {
	name := generateResourceName()
	resourceName := "volterra_known_label_key." + name
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{ves_io_schema_known_label_key.ObjectType})
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
				Config: testConfigKnownLabelKey(ves_io_schema_known_label_key.ObjectType, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "key", name),
				),
			},
		},
	})
}

func testConfigKnownLabelKey(objType, name string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  key = "%s"
		  namespace = "shared"
		}
		`, tfResourceName, name, name)
}
