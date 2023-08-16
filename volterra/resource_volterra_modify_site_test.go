// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/server"
)

const (
	siteName = "site-sb-1"
)

func TestAccSiteUpdate(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createSiteTestServer(t)
	resourceName := fmt.Sprintf("%s.%s", modifySite, name)
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
				Config: testConfgSite(ves_io_schema_site.ObjectType, name, siteName, systemNS),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", siteName),
					resource.TestCheckResourceAttr(resourceName, "namespace", systemNS),
					resource.TestCheckResourceAttr(resourceName, "desired_pool_count", "5"),
				),
			},
		},
	})
}

func testConfgSite(objType, resourceName, name, namespace string) string {
	// convert ves.io.schema.virtual_host.Object to volt_virtual_host
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_modify_%s", parts[3])
	return fmt.Sprintf(`
		resource "%[4]s" "%[1]s" {
		  name = "%[2]s"
		  namespace = "%[3]s"
		  desired_pool_count = 5
		}
		`, resourceName, name, namespace, tfResourceName)
}

func createSiteTestServer(t *testing.T) (string, func()) {

	f, stop := makeTestServer(t, ves_io_schema_site.ObjectType)
	siteObj := mkDBObjSite(systemNS, siteName)
	ctx := client.NewContextWithHeaders(nil, client.WithTenant("ves-io"))
	f.MustCreateEntry(siteObj, server.WithCtx(ctx))

	return fmt.Sprintf("https://localhost:%d", f.Svc.RestServerTLSPort()), stop
}

func mkDBObjSite(namespace, name string) *ves_io_schema_site.DBObject {
	siteObj := &ves_io_schema_site.Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Namespace: namespace,
			Name:      name,
			Labels:    map[string]string{"key1": "value1"},
		},
		Spec: &ves_io_schema_site.SpecType{
			GcSpec: &ves_io_schema_site.GlobalSpecType{
				SiteType: ves_io_schema_site.REGIONAL_EDGE,
				Address:  "This is a simple address",
				Coordinates: &ves_io_schema_site.Coordinates{
					Latitude:  37.3688,
					Longitude: 122.0363,
				},
			},
		},
	}
	return ves_io_schema_site.NewDBObject(siteObj)
}
