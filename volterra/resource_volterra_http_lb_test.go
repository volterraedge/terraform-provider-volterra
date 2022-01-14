//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_app_firewall "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/app_firewall"
	ves_io_schema_app_setting "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/app_setting"
	ves_io_schema_app_type "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/app_type"
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
	ves_io_schema_tenant "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/tenant"
	ves_io_schema_uid "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/user_identification"
	ves_io_schema_views_http_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
	ves_io_schema_views_rate_limiter_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/rate_limiter_policy"
)

// TestAccHTTPLB token creation test
func TestHTTPLB(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, f := createTestCustomAPIServer(t, []string{
		ves_io_schema_views_http_loadbalancer.ObjectType,
		ves_io_schema_uid.ObjectType,
		ves_io_schema_ns.ObjectType,
		ves_io_schema_views_rate_limiter_policy.ObjectType,
		ves_io_schema_app_firewall.ObjectType,
		ves_io_schema_app_type.ObjectType,
		ves_io_schema_app_setting.ObjectType,
		ves_io_schema_tenant.ObjectType,
	})
	tenantName := "tenant1"
	tenantObj := mkDBObjTenant(tenantName, uuid.New().String())
	f.MustCreateEntry(tenantObj)
	defer stopFunc()
	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "tenant1")
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
		resource "volterra_namespace" "shared" {
			name = "shared"
		}
		resource "volterra_app_type" "app_type" {
		  name = "%[3]s"
		  namespace = volterra_namespace.shared.name
		  business_logic_markup_setting {
		    enable = true
		  }
		  features {
		    type = "BUSINESS_LOGIC_MARKUP"
		  }
		  features {
		    type = "PER_REQ_ANOMALY_DETECTION"
		  }
		  features {
		    type = "USER_BEHAVIOR_ANALYSIS"
		  }
		}
		resource "volterra_app_setting" "app_setting" {
		  name = "%[2]s"
		  namespace = volterra_namespace.app.name
		  app_type_settings {
		    app_type_ref {
			name = volterra_app_type.app_type.name
			namespace = "shared"
		    }
		    timeseries_analyses_setting {
			metric_selectors {
				metrics_source = "VIRTUAL_HOSTS"
				metric = [
					"REQUEST_RATE",
					"ERROR_RATE",
					"LATENCY",
					"THROUGHPUT"
				]
			}
		    }
		    business_logic_markup_setting {
			    enable = true
		    }
		    user_behavior_analysis_setting {
			enable_learning = true
			enable_detection {
				include_forbidden_activity {
					forbidden_requests_threshold = 10
				}
				exclude_failed_login_activity = true
				exclude_waf_activity = true
				cooling_off_period = 20
			}
		    }
		  }
		}
		resource "volterra_app_firewall" "app_fwd_athena" {
		  name = "%[2]s"
		  namespace = volterra_namespace.app.name
		  blocking                  = true
		  allow_all_response_codes  = true
		  default_anonymization     = true
		  use_default_blocking_page = true

		  detection_settings  {
		    signature_selection_setting {
		      default_attack_type_settings        = true
		      high_medium_low_accuracy_signatures = true
		    }

		    enable_suppression         = true
		    enable_threat_campaigns    = true
		    default_violation_settings = true
		  }

		  bot_protection_setting {
		    malicious_bot_action  = "BLOCK"
		    suspicious_bot_action = "REPORT"
		    good_bot_action       = "REPORT"
		  }
		}
		resource "volterra_http_loadbalancer" "%[1]s" {
		  name = "%[2]s"
		  namespace = volterra_namespace.app.name
		  labels = {
			  "ves.io/app_type" = volterra_app_type.app_type.name
		  }
		  advertise_on_public_default_vip = true
		  no_challenge = true
		  http {
			dns_volterra_managed = false
		  }
		  disable_rate_limit = false
		  app_firewall {
		    name      = volterra_app_firewall.app_fwd_athena.name
		  }
		  no_service_policies = true
		  round_robin = true
		  domains = ["http.helloclouds.app"]
		  multi_lb_app = true
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
