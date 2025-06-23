// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw/test/generic"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_token "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/token"
)

// TestAccToken token creation test
func TestAccToken(t *testing.T) {
	name := generateResourceName()
	f, testURL, stopFunc := createTokenTestServer(t)
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
					resource.TestCheckResourceAttr(resourceName, "name", siteName),
				),
			},
		},
	})
	// validate if token exists
	entry := f.MustFindEntry(ves_io_schema_token.ObjectDefTblName, "", db.WithFindUsingNSIndex("ves-io", systemNS, "token"))
	assert.NotNil(t, entry)
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

func createTokenTestServer(t *testing.T) (*generic.Fixture, string, func()) {

	f, stop := makeTestServer(t, ves_io_schema_token.ObjectType)
	tokenObj := mkDBObjToken(systemNS, "token")
	ctx := client.NewContextWithHeaders(nil, client.WithTenant("ves-io"))
	f.MustCreateEntry(tokenObj, server.WithCtx(ctx))

	return f, fmt.Sprintf("https://localhost:%d", f.Svc.RestServerTLSPort()), stop
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
