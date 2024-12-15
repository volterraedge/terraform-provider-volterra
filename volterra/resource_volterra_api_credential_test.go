// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_api_credential "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/api_credential"
)

func TestAccAPICredentialBasic(t *testing.T) {
	name := generateResourceName()
	resourceName := "volterra_api_credential." + name
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{})
	defer stopFunc()
	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "ves-io")

	timeString := currentTime()
	getCreateResponseAPICredential(name)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testKubeConfigCredential(ves_io_schema_api_credential.ObjectType, name, timeString),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAPICertCredential(ves_io_schema_api_credential.ObjectType, name, timeString),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testKubeConfigCredential(objType, name, timeString string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  api_credential_type = "KUBE_CONFIG"
		  expiry_days = 10
		  created_at = "%s"
		}
		`, tfResourceName, name, name, timeString)
}

func testAPICertCredential(objType, name, timeString string) string {
	// convert ves.io.schema.vnamespace.Object to volt_namespace
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  api_credential_type = "API_CERTIFICATE"
		  created_at = "%s"
		}
		`, tfResourceName, name, name, timeString)
}

func currentTime() string {
	//code to crearte time.Now and RFC366 return a string value
	return time.Now().UTC().Format(time.RFC3339)
}
