// data_source_volterra_address_allocator_test.go
//
// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/server"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_addr_alloc "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/address_allocator"
)

func TestDataAddressAllocator(t *testing.T) {
	name := generateResourceName()
	dataName := "data.volterra_address_allocator." + name
	testURL, stopFunc := createAddrAllocTestServer(t, name)
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
				Config: testConfigDataAddressAllocator(ves_io_schema_addr_alloc.ObjectType, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "mode", "GLOBAL_PER_SITE_NODE"),
				),
			},
		},
	})
}

func testConfigDataAddressAllocator(objType, name string) string {
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
	data "%[1]s" "%[2]s" {
		name = "%[2]s"
		namespace = "system"
	}`, tfResourceName, name)
}

func createAddrAllocTestServer(t *testing.T, name string) (string, func()) {

	f, stop := makeTestServer(t, ves_io_schema_addr_alloc.ObjectType)
	pbObj := &ves_io_schema_addr_alloc.Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Namespace: "system",
			Name:      name,
			Uid:       uuid.New().String(),
		},
		Spec: &ves_io_schema_addr_alloc.SpecType{
			GcSpec: &ves_io_schema_addr_alloc.GlobalSpecType{
				Mode:        ves_io_schema_addr_alloc.GLOBAL_PER_SITE_NODE,
				AddressPool: []string{"129.254.192.0/22"},
				AllocationMap: &ves_io_schema_addr_alloc.NodePrefixMapType{
					Endpoints: map[string]*ves_io_schema_addr_alloc.NodePrefixType{
						"ves-node-id-68fc55dd7b": &ves_io_schema_addr_alloc.NodePrefixType{Prefix: "169.254.192.0/30"},
					},
				},
				AddressAllocationScheme: &ves_io_schema_addr_alloc.AllocationScheme{},
			},
		},
	}
	addrAllocObj := ves_io_schema_addr_alloc.NewDBObject(pbObj)
	ctx := client.NewContextWithHeaders(nil, client.WithTenant("ves-io"))
	f.MustCreateEntry(addrAllocObj, server.WithCtx(ctx))

	return fmt.Sprintf("https://localhost:%d", f.Svc.RestServerTLSPort()), stop
}
