// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_known_label "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/known_label"
	ves_io_schema_known_label_key "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/known_label_key"
)

func TestAccKnownLabelBasic(t *testing.T) {
	key := generateResourceName()
	value := generateResourceName()
	getTestObj := func() []string {
		return []string{
			ves_io_schema_known_label.ObjectType, ves_io_schema_known_label_key.ObjectType,
		}
	}
	testURL, stopFunc, _ := createTestCustomAPIServer(t, getTestObj())
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
				Config: testConfigKnownLabel(key, value),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("volterra_known_label_key.%s", key), "key", key),
					resource.TestCheckResourceAttr(fmt.Sprintf("volterra_known_label.%s", value), "value", value),
				),
			},
		},
	})
}

func testConfigKnownLabel(key, value string) string {
	return fmt.Sprintf(`
	resource "volterra_known_label_key" "%[1]s" {
		key = "%[1]s"
		namespace = "shared"
	}
	resource "volterra_known_label" "%[2]s" {
		key = "%[1]s"
		namespace = "shared"
		value       = "%[2]s"
	}
	`, key, value)
}
