// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestResourceTFParamsAction(t *testing.T) {
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{})
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
				Config: testTFParamsAction("site-01", "aws_vpc_site", "plan"),
			},
			{
				Config: testTFParamsAction("site-02", "aws_vpc_site", "apply"),
			},
			{
				Config: testTFParamsAction("site-03", "aws_tgw_site", "plan"),
			},
			{
				Config: testTFParamsAction("site-04", "aws_tgw_site", "apply"),
			},
			{
				Config: testTFParamsAction("site-05", "gcp_vpc_site", "plan"),
			},
			{
				Config: testTFParamsAction("site-06", "gcp_vpc_site", "apply"),
			},
			{
				Config: testTFParamsAction("site-07", "azure_vnet_site", "plan"),
			},
			{
				Config: testTFParamsAction("site-08", "azure_vnet_site", "apply"),
			},
		},
	})
}

func testTFParamsAction(name, kind, action string) string {

	return fmt.Sprintf(`
		resource "%s" "%s" {
		  site_name = "%s"
		  site_kind = "%s"
		  action = "%s"
		  wait_for_action = false
		}
		`, tfParamsAction, name, name, kind, action)
}
