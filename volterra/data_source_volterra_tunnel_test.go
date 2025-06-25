// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/server"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_tunnel "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/tunnel"
)

func TestDataTunnel(t *testing.T) {
	name := generateResourceName()
	dataName := "data.volterra_tunnel." + name
	testURL, stopFunc := createTunnelTestServer(t, name)
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
				Config: testConfigDataTunnel(ves_io_schema_tunnel.ObjectType, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "type", "IPSEC_PSK"),
				),
			},
		},
	})
}

func testConfigDataTunnel(objType, name string) string {
	parts := strings.Split(objType, ".")
	tfResourceName := fmt.Sprintf("volterra_%s", parts[3])
	return fmt.Sprintf(`
	data "%[1]s" "%[2]s" {
		name = "%[2]s"
		namespace = "system"
	}`, tfResourceName, name)
}

func createTunnelTestServer(t *testing.T, name string) (string, func()) {

	f, stop := makeTestServer(t, ves_io_schema_tunnel.ObjectType)
	pbObj := &ves_io_schema_tunnel.Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Namespace: "system",
			Name:      name,
			Uid:       uuid.New().String(),
		},
		Spec: &ves_io_schema_tunnel.SpecType{
			GcSpec: &ves_io_schema_tunnel.GlobalSpecType{
				TunnelType: ves_io_schema_tunnel.IPSEC_PSK,
				LocalIp: &ves_io_schema_tunnel.LocalIpAddressSelector{
					Type: &ves_io_schema_tunnel.LocalIpAddressSelector_IpAddress{
						IpAddress: &ves_io_schema_tunnel.LocalIpAddressType{},
					},
				},
				RemoteIp: &ves_io_schema_tunnel.RemoteIpAddressSelector{
					Type: &ves_io_schema_tunnel.RemoteIpAddressSelector_Ip{
						Ip: &ves_io_schema.IpAddressType{},
					},
				},
			},
		},
	}
	tunnelObj := ves_io_schema_tunnel.NewDBObject(pbObj)
	ctx := client.NewContextWithHeaders(nil, client.WithTenant("ves-io"))
	f.MustCreateEntry(tunnelObj, server.WithCtx(ctx))

	return fmt.Sprintf("https://localhost:%d", f.Svc.RestServerTLSPort()), stop
}
