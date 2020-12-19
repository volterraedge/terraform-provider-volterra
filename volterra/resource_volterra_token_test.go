//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_token "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/token"
	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/server"
)

// TestAccToken token creation test
func TestAccToken(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTokenTestServer(t)
	resourceName := fmt.Sprintf("volterra_token.%s", name)
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
				Config: testConfigToken(name, siteName, systemNS),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "created_by_tf", "false"),
				),
			},
		},
	})
}

func testConfigToken(resourceName, name, namespace string) string {
	// convert ves.io.schema.virtual_host.Object to volt_virtual_host
	return fmt.Sprintf(`
		resource "volterra_token" "%[1]s" {
		  name = "%[2]s"
		  namespace = "%[3]s"
		}
		`, resourceName, name, namespace)
}

func createTokenTestServer(t *testing.T) (string, func()) {

	f, stop := makeTestServer(t, ves_io_schema_token.ObjectType)
	siteObj := mkDBObjToken(systemNS, "token")
	ctx := client.NewContextWithHeaders(nil, client.WithTenant("ves-io"))
	f.MustCreateEntry(siteObj, server.WithCtx(ctx))

	return fmt.Sprintf("https://localhost:%d", f.Svc.RestServerTLSPort()), stop
}

func mkDBObjToken(namespace, name string) *ves_io_schema_token.DBObject {
	tokenObj := &ves_io_schema_token.Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Namespace: namespace,
			Name:      name,
		},
		Spec: &ves_io_schema_token.SpecType{},
	}
	return ves_io_schema_token.NewDBObject(tokenObj)
}
