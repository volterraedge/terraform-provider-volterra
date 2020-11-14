//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/server"
)

func TestAccSiteState(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createSiteTestCustomServer(t)
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
				Config: testConfgSiteState(name, siteName),
			},
		},
	})
}

func testConfgSiteState(resourceName, name string) string {
	return fmt.Sprintf(`
		resource "volterra_site_state" "%[1]s" {
		  name = "%[2]s"
		  when = "delete"
		  state = "DECOMMISSIONING"
		}
		`, resourceName, name)
}

func createSiteTestCustomServer(t *testing.T) (string, func()) {

	f, stop := makeCustomTestServer(t, []string{ves_io_schema_site.ObjectType})
	siteObj := mkDBObjSite("system", siteName)
	ctx := client.NewContextWithHeaders(nil, client.WithTenant("ves-io"))
	f.MustCreateEntry(siteObj, server.WithCtx(ctx))

	return fmt.Sprintf("https://localhost:%d", f.Svc.RestServerTLSPort()), stop
}
