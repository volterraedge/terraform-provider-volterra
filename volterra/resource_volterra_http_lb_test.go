//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
	ves_io_schema_uid "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/user_identification"
	ves_io_schema_views_http_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
	ves_io_schema_views_rate_limiter_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/rate_limiter_policy"
)

// TestAccHTTPLB token creation test
func TestHTTPLB(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTestCustomAPIServer(t, []string{
		ves_io_schema_views_http_loadbalancer.ObjectType,
		ves_io_schema_uid.ObjectType,
		ves_io_schema_ns.ObjectType,
		ves_io_schema_views_rate_limiter_policy.ObjectType,
	})
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
				Config: testConfigHTTPLB(name, name, "app-test"),
			},
		},
	})
}

func testConfigHTTPLB(resourceName, name, namespace string) string {
	return fmt.Sprintf(`
		resource "volterra_namespace" "app" {
		  name = "%[3]s"
		}
		resource "volterra_http_loadbalancer" "%[1]s" {
		  name = "%[2]s"
		  namespace = volterra_namespace.app.name
		  advertise_on_public_default_vip = true
		  no_challenge = true
		  http {
			dns_volterra_managed = false
		  }
		  disable_rate_limit = false
		  disable_waf = true
		  no_service_policies = true
		  round_robin = true
		  domains = ["http.helloclouds.app"]
		  user_identification {
		    name = volterra_user_identification.%[1]s.name
		    namespace = "%[3]s"
		  }
		  rate_limit {
		    ip_allowed_list {
		      prefixes = ["8.8.8.8/32"]
		    }
		    policies {
		      policies {
			name = volterra_rate_limiter_policy.%[1]s.name
			namespace = "%[3]s"
		      }
		    }
		  }
		}
		resource "volterra_user_identification" "%[1]s" {
		  name = "%[2]s"
		  namespace = volterra_namespace.app.name
		  rules {
		    client_ip = true
		  }
		}
		resource "volterra_rate_limiter_policy" "%[1]s" {
		  name = "%[2]s"
		  namespace = volterra_namespace.app.name
		  rules {
		    metadata {
		      name = "%[2]s"
		    }
		    spec {
		      apply_rate_limiter = true
		      domain_matcher {
		        exact_values = ["test.io"]
		      }
		    }
		  }
		}
		`, resourceName, name, namespace)
}
