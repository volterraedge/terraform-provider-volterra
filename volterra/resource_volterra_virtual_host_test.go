// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema_tenant "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/tenant"
	ves_io_schema_vh "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
)

func TestAccVirtualHostBasic(t *testing.T) {
	name := generateResourceName()
	resourceName := "volterra_virtual_host." + name
	testURL, stopFunc, f := createTestCustomAPIServer(t, []string{
		ves_io_schema_tenant.ObjectType,
		ves_io_schema_vh.ObjectType,
	})
	defer stopFunc()
	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "ves-io")
	tenantName := "ves-io"
	tenantObj := mkDBObjTenant(tenantName, uuid.New().String())
	f.MustCreateEntry(tenantObj)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testConfigVH(ves_io_schema_vh.ObjectType, name, "ns1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "namespace", "ns1"),
				),
			},
		},
	})
}

func testConfigVH(objType, name, namespace string) string {
	// convert ves.io.schema.virtual_host.Object to volt_virtual_host
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
		resource "%[3]s" "%[1]s" {
		  name = "%[1]s"
		  namespace = "%[2]s"
		  domains = ["terraform.test.io"]
		  proxy = "UDP_PROXY"
		  no_challenge = true
		}
		`, name, namespace, tfResourceName)
}
