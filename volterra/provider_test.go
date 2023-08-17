// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"volterra": testAccProvider,
	}
}

func testAccPreCheck() {

	if v := os.Getenv("VOLT_TENANT"); v == "" {
		os.Setenv("VOLT_TENANT", "test")
	}
	if v := os.Getenv("VOLT_ENV"); v == "" {
		os.Setenv("VOLT_ENV", "test")
	}
}
