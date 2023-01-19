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
	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/server"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_app_firewall "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/app_firewall"
	ves_io_schema_app_setting "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/app_setting"
	ves_io_schema_app_type "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/app_type"
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
	ves_io_schema_sp "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy"
	ves_io_schema_tenant "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/tenant"
	ves_io_schema_uid "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/user_identification"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_views_api_definition "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/api_definition"
	http_lb "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
	ves_io_schema_views_http_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/origin_pool"
	ves_io_schema_views_rate_limiter_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/rate_limiter_policy"
	vh_dns_info "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host_dns_info"
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
		origin_pool.ObjectType,
		ves_io_schema_sp.ObjectType,
		ves_io_schema_views_api_definition.ObjectType,
	})
	defer stopFunc()
	tenantName := "tenant1"
	ctx := client.NewContextWithHeaders(f.StartCtx,
		client.WithTenant(tenantName),
		client.WithCreatorClass("test"),
		client.WithCreatorID("test-"+t.Name()))
	tenantObj := mkDBObjTenant(tenantName, uuid.New().String())
	f.MustCreateEntry(tenantObj)

	nsName := "default"
	poolName := "pool-1"
	lbName := "httplb"
	// origin pool creation
	opSpec := &origin_pool.SpecType{
		GcSpec: &origin_pool.GlobalSpecType{
			OriginServers: []*origin_pool.OriginServerType{
				{
					Choice: &origin_pool.OriginServerType_PublicName{
						PublicName: &origin_pool.OriginServerPublicName{
							DnsName: "abc.com",
						},
					},
				},
			},
			PortChoice: &origin_pool.GlobalSpecType_Port{
				Port: 80,
			},
			TlsChoice: &origin_pool.GlobalSpecType_NoTls{
				NoTls: &ves_io_schema.Empty{},
			},
		},
	}
	opObj := &origin_pool.Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Namespace: nsName,
			Name:      poolName,
		},
		Spec: opSpec,
	}
	f.MustCreateEntry(opObj.ToEntry(), server.WithAPIPrivate(), server.WithCtx(ctx))
	httpLbObj := &http_lb.Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Namespace: nsName,
			Name:      lbName,
		},
		Spec: &http_lb.SpecType{
			GcSpec: &http_lb.GlobalSpecType{
				AdvertiseChoice: &http_lb.GlobalSpecType_DoNotAdvertise{
					DoNotAdvertise: &ves_io_schema.Empty{},
				},
				ChallengeType: &http_lb.GlobalSpecType_NoChallenge{
					NoChallenge: &ves_io_schema.Empty{},
				},
				Domains: []string{"abc.com"},
				LoadbalancerType: &http_lb.GlobalSpecType_Http{
					Http: &http_lb.ProxyTypeHttp{},
				},
				DefaultRoutePools: []*ves_io_schema_views.OriginPoolWithWeight{{
					PoolChoice: &ves_io_schema_views.OriginPoolWithWeight_Pool{
						Pool: &ves_io_schema_views.ObjectRefType{
							Name:      poolName,
							Tenant:    tenantName,
							Namespace: nsName,
						},
					},
					Weight: 1,
				}},
				ServicePolicyChoice: &http_lb.GlobalSpecType_ServicePoliciesFromNamespace{
					ServicePoliciesFromNamespace: &ves_io_schema.Empty{},
				},
				UserIdChoice: &http_lb.GlobalSpecType_UserIdClientIp{
					UserIdClientIp: &ves_io_schema.Empty{},
				},
				RateLimitChoice: &http_lb.GlobalSpecType_DisableRateLimit{
					DisableRateLimit: &ves_io_schema.Empty{},
				},
				WafChoice: &http_lb.GlobalSpecType_DisableWaf{
					DisableWaf: &ves_io_schema.Empty{},
				},
				HashPolicyChoice: &http_lb.GlobalSpecType_RoundRobin{
					RoundRobin: &ves_io_schema.Empty{},
				},
				BotDefenseChoice: &http_lb.GlobalSpecType_DisableBotDefense{
					DisableBotDefense: &ves_io_schema.Empty{},
				},
				MlConfigChoice: &http_lb.GlobalSpecType_MultiLbApp{
					MultiLbApp: &ves_io_schema.Empty{},
				},
				HostName: "ves-io-082258df-c45e-4163-8884-7878453eeb2a",
				DnsInfo: []*vh_dns_info.DnsInfo{
					{
						IpAddress: "1.1.1.1",
					},
				},
			},
		},
	}
	f.MustCreateEntry(httpLbObj.ToEntry(), server.WithAPIPrivate(), server.WithCtx(ctx))

	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", testURL)
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "tenant1")
	dataSourceName := "data.volterra_http_loadbalancer_state.exist_http_lb"
	//httpLbResourceName := fmt.Sprintf("volterra_http_loadbalancer.%s", name)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testConfigHTTPLB(name, "app-test", lbName, nsName, 3000),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "cname", "ves-io-082258df-c45e-4163-8884-7878453eeb2a"),
					resource.TestCheckResourceAttr(dataSourceName, "ip_address", "1.1.1.1"),
					//resource.TestCheckResourceAttr(httpLbResourceName, "more_option.648059762.request_headers_to_add.#", "1"),
					//resource.TestCheckResourceAttr(httpLbResourceName, "more_option.648059762.request_headers_to_add.0.append", "false"),
					//resource.TestCheckResourceAttr(httpLbResourceName, "more_option.648059762.request_headers_to_add.0.name", "X-real-ip"),
					//resource.TestCheckResourceAttr(httpLbResourceName, "more_option.648059762.request_headers_to_add.0.value", "$[client_address]"),
				),
			},
			{
				// Test replace of http loadbalancer
				Config: testConfigHTTPLB(name, "app-test", lbName, nsName, 600),
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttr(httpLbResourceName, "more_option.3270546385.request_headers_to_add.#", "1"),
					//resource.TestCheckResourceAttr(httpLbResourceName, "more_option.3270546385.request_headers_to_add.0.append", "false"),
					//resource.TestCheckResourceAttr(httpLbResourceName, "more_option.3270546385.request_headers_to_add.0.name", "X-real-ip"),
					//resource.TestCheckResourceAttr(httpLbResourceName, "more_option.3270546385.request_headers_to_add.0.value", "$[client_address]"),
				),
			},
		},
	})
}

func testConfigHTTPLB(name, namespace, existLbName, existNsName string, timeout int32) string {
	return fmt.Sprintf(`
		resource "volterra_namespace" "app" {
		  name = "%[2]s"
		}
		resource "volterra_namespace" "shared" {
			name = "shared"
		}
		resource "volterra_app_type" "app_type" {
		  name = "%[2]s"
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
		  name = "%[1]s"
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
		  name = "%[1]s"
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
		resource "volterra_service_policy" "allow_ns" {
		  name       = "allow-ns"
		  namespace  = volterra_namespace.app.name
		  algo       = "FIRST_MATCH"
		  any_server = true
		  rule_list {
		    rules {
		      metadata {
		        name = "allow-ns"
		      }

		      spec {
		        action = "ALLOW"
		        client_selector {
		          expressions = [
		           "name.ves.io/namespace in (foo)"
		          ]
		        }
		        any_ip           = true
		        any_asn          = true
		        challenge_action = "DEFAULT_CHALLENGE"
		        waf_action {
		        none = true
		        }
		      }
		    }
		  }
		}
		resource "volterra_http_loadbalancer" "%[1]s" {
		  name = "%[1]s"
		  namespace = volterra_namespace.app.name
		  labels = {
			  "ves.io/app_type" = volterra_app_type.app_type.name
			  "new-test"        = "true"
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
		  no_service_policies = false
		  active_service_policies {
		    policies {
		      name = volterra_service_policy.allow_ns.name
		    }
		  }
		  round_robin = true
		  domains = ["http.helloclouds.app"]
		  multi_lb_app = true
		  user_identification {
		    name = volterra_user_identification.%[1]s.name
		    namespace = "%[2]s"
		  }
		  rate_limit {
		    ip_allowed_list {
		      prefixes = ["8.8.8.8/32"]
		    }
		    policies {
		      policies {
			name = volterra_rate_limiter_policy.%[1]s.name
			namespace = "%[2]s"
		      }
		    }
		  }
		  more_option {
		    disable_default_error_pages         = false
		    disable_path_normalize              = true
		    enable_path_normalize               = false
		    enable_strict_sni_host_header_check = false
		    idle_timeout                        = %[5]d
		    max_request_header_size             = 80
		    request_headers_to_remove           = []
		    response_headers_to_remove          = []

		    request_headers_to_add {
		      append = false
		      name   = "X-real-ip"
		      value  = "$[client_address]"
		    }
		  }
		}
		resource "volterra_api_definition" "example" {
			name      = "api-def-boutique"
			namespace = volterra_namespace.app.name

			swagger_specs = [ "https://console.volterra.ves.io/api/object_store/namespaces/ws-namespace/stored_objects/swagger/api-swagger-boutique/v2-22-02-09" ]
		      }

		data "volterra_http_loadbalancer_state" "http_lb" {
			name = volterra_http_loadbalancer.%[1]s.name
			namespace = volterra_namespace.app.name
		}
		resource "volterra_user_identification" "%[1]s" {
		  name = "%[1]s"
		  namespace = volterra_namespace.app.name
		  rules {
		    client_ip = true
		  }
		}
		resource "volterra_rate_limiter_policy" "%[1]s" {
		  name = "%[1]s"
		  namespace = volterra_namespace.app.name
		  rules {
		    metadata {
		      name = "%[1]s"
		    }
		    spec {
		      apply_rate_limiter = true
		      domain_matcher {
		        exact_values = ["test.io"]
		      }
		    }
		  }
		}
		data "volterra_http_loadbalancer_state" "exist_http_lb" {
			name = "%[3]s"
			namespace = "%[4]s"
		}
		`, name, namespace, existLbName, existNsName, timeout)
}
