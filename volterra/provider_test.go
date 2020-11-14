//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
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
