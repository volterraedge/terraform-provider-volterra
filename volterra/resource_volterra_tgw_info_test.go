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

func TestResourceTGWInfo(t *testing.T) {
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
				Config: testTGWInfo(name),
			},
			{
				Config: testTGWInfoWithDirectConnectInfo(name),
			},
		},
	})
}

func testTGWInfo(name string) string {

	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  tgw_id = "tgw-0d9bf32a5171e5a49"
		  vpc_id = "vpc-02d0d70860b330393"
		  subnet_ids {
			outside_subnet_id = "subnet-12345678"
			inside_subnet_id = "subnet-45678123"
			workload_subnet_id = "subnet-56781234"
			az = "us-east-2a"
			outside_subnet {
				az_name = "us-east-2a"
				id = "subnet-12345678"
				ipv4_prefix = "10.0.0.0/24"
			}
			inside_subnet {
				az_name = "us-east-2a"
				id = "subnet-45678123"
				ipv4_prefix = "10.0.1.0/24"
			}
			workload_subnet {
				az_name = "us-east-2a"
				id = "subnet-56781234"
				ipv4_prefix = "10.0.2.0/24"
			}
		  }
		  public_ips = ["10.0.0.10"]
		  private_ips = ["192.168.0.10"]
		}
		`, setTGWInfo, name, name)
}

func testTGWInfoWithDirectConnectInfo(name string) string {

	return fmt.Sprintf(`
		resource "%s" "%s" {
		  name = "%s"
		  tgw_id = "tgw-0d9bf32a5171e5a49"
		  vpc_id = "vpc-02d0d70860b330393"
		  subnet_ids {
			outside_subnet_id = "subnet-12345678"
			inside_subnet_id = "subnet-45678123"
			workload_subnet_id = "subnet-56781234"
			az = "us-east-2a"
			outside_subnet {
				az_name = "us-east-2a"
				id = "subnet-12345678"
				ipv4_prefix = "10.0.0.0/24"
			}
			inside_subnet {
				az_name = "us-east-2a"
				id = "subnet-45678123"
				ipv4_prefix = "10.0.1.0/24"
			}
			workload_subnet {
				az_name = "us-east-2a"
				id = "subnet-56781234"
				ipv4_prefix = "10.0.2.0/24"
			}
		  }
		  public_ips = ["10.0.0.10"]
		  private_ips = ["192.168.0.10"]
		  direct_connect_info {
			direct_connect_gateway_id = "1234-6789-0987-1234-09876"
			vgw_id = "vgw-12345678"
			asn = 64512
		  }
		}
		`, setTGWInfo, name, name)
}
