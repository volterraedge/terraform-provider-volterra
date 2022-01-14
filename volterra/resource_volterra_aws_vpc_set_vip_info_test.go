//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestResourceAWSVPCSiteVipInfo(t *testing.T) {
	name := generateResourceName()
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
				Config: testAWSVPCSiteVipInfo(name),
			},
		},
	})
}

func testAWSVPCSiteVipInfo(name string) string {
	return fmt.Sprintf(`
		resource "%[1]s" "aws" {
		  name = "%[2]s"
		  namespace = "system"
		  site_type = "aws_vpc_site"
		  vip_params_per_az {
			outside_vip = ["1.1.1.1"]
			az_name = "ap-northeast-1a"
		  }
		  vip_params_per_az {
			outside_vip = ["2.2.2.2"]
			az_name = "ap-northeast-1d"
		  }
		  vip_params_per_az {
			outside_vip = ["3.3.3.3"]
			az_name = "ap-northeast-1c"
		  }
		}
		resource "%[1]s" "azure" {
			name = "%[2]s"
			namespace = "system"
			site_type = "azure_vnet_site"
			vip_params_per_az {
			  outside_vip = ["1.1.1.1"]
			  az_name = "1"
			}
			vip_params_per_az {
			  outside_vip = ["2.2.2.2"]
			  az_name = "2"
			}
			vip_params_per_az {
			  outside_vip = ["3.3.3.3"]
			  az_name = "3"
			}
		  }

		resource "%[1]s" "gvp" {
			name = "%[2]s"
			namespace = "system"
			site_type = "azure_vnet_site"
			vip_params_per_az {
			  outside_vip = ["1.1.1.1"]
			  az_name = "northamerica-northeast1-a"
			}
			vip_params_per_az {
			  outside_vip = ["2.2.2.2"]
			  az_name = "northamerica-northeast1-b"
			}
			vip_params_per_az {
			  outside_vip = ["3.3.3.3"]
			  az_name = "northamerica-northeast1-c"
			}
		  }
		`, setSiteVipInfo, name)
}
