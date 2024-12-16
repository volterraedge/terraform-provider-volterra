// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
	vh_dns_info "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host_dns_info"
	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/server"
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
	tenantObj := mkDBObjTenant(tenantName, uuid.New().String(), withAddonServices("f5xc-flow-collection", "f5xc-ipv6-standard", "f5xc-waap-standard"))
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
					Http: &http_lb.ProxyTypeHttp{
						PortChoice: &http_lb.ProxyTypeHttp_Port{
							Port: 80,
						},
					},
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
	httpLbResourceName := fmt.Sprintf("volterra_http_loadbalancer.%s", name)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testConfigHTTPLB(name, "app-test", lbName, nsName, 3000),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "cname", "ves-io-082258df-c45e-4163-8884-7878453eeb2a"),
					resource.TestCheckResourceAttr(dataSourceName, "ip_address", "1.1.1.1"),
					resource.TestCheckResourceAttr(httpLbResourceName, "more_option.0.request_headers_to_add.#", "1"),
					resource.TestCheckResourceAttr(httpLbResourceName, "more_option.0.request_headers_to_add.0.append", "false"),
					resource.TestCheckResourceAttr(httpLbResourceName, "more_option.0.request_headers_to_add.0.name", "X-real-ip"),
					resource.TestCheckResourceAttr(httpLbResourceName, "more_option.0.request_headers_to_add.0.value", "$[client_address]"),
				),
			},
			{
				// Test replace of http loadbalancer
				Config: testConfigHTTPLB(name, "app-test", lbName, nsName, 600),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(httpLbResourceName, "more_option.0.request_headers_to_add.#", "1"),
					resource.TestCheckResourceAttr(httpLbResourceName, "more_option.0.request_headers_to_add.0.append", "false"),
					resource.TestCheckResourceAttr(httpLbResourceName, "more_option.0.request_headers_to_add.0.name", "X-real-ip"),
					resource.TestCheckResourceAttr(httpLbResourceName, "more_option.0.request_headers_to_add.0.value", "$[client_address]"),
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
				disable = false
		      }

		      spec {
		        action = "ALLOW"
				any_client           = false
                any_dst_asn          = false
                any_dst_ip           = false
                scheme               = []
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
				waf_in_monitoring_mode = false
                waf_skip_processing    = false
		        }
		      }
		    }
		  }
		}
		resource "volterra_http_loadbalancer" "%[1]s" {
		  name = "%[1]s"
		  default_sensitive_data_policy    = true
		  disable_threat_mesh              = true
		  disable_api_definition           = true
          disable_api_discovery            = true      
          disable_malicious_user_detection = true
          disable_trust_client_ip_headers  = true
		  disable_malware_protection       = true
		  l7_ddos_action_default           = true
		  namespace = volterra_namespace.app.name
		  labels = {
			  "ves.io/app_type" = volterra_app_type.app_type.name
			  "new-test"        = "true"
		  }
		  advertise_on_public_default_vip = true
		  no_challenge = true
		  http {
			dns_volterra_managed = false
			port = 80
		  }
		  disable_rate_limit = false
		  app_firewall {
		    name      = volterra_app_firewall.app_fwd_athena.name
			namespace = "%[2]s"
			tenant    = "tenant1"
		  }
		  no_service_policies = false
		  active_service_policies {
		    policies {
		      name = volterra_service_policy.allow_ns.name
			  namespace = "%[2]s"
			  tenant    = "tenant1"
		    }
		  }
		  round_robin = true
		  domains = ["http.helloclouds.app"]
		  multi_lb_app = true
		  user_identification {
		    name = volterra_user_identification.%[1]s.name
		    namespace = "%[2]s"
			tenant    = "tenant1"
		  }
		  rate_limit {
			no_ip_allowed_list = false
			no_policies        = false
		    ip_allowed_list {
			  ipv6_prefixes = []
		      prefixes = ["8.8.8.8/32"]
		    }
		    policies {
		      policies {
			name = volterra_rate_limiter_policy.%[1]s.name
			namespace = "%[2]s"
			tenant    = "tenant1"
		      }
		    }
		  }
		  more_option {
			disable_default_error_pages         = true
  			disable_path_normalize              = true
  			enable_path_normalize               = false
  			enable_strict_sni_host_header_check = true
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

func TestHTTPLBWithAutoCert(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, f := createTestCustomAPIServer(t, []string{
		ves_io_schema_views_http_loadbalancer.ObjectType,
		ves_io_schema_ns.ObjectType,
		ves_io_schema_tenant.ObjectType,
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
	lbName := "httplb"
	httpLbObj := &http_lb.Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Namespace: nsName,
			Name:      lbName,
		},
		Spec: &http_lb.SpecType{
			GcSpec: &http_lb.GlobalSpecType{
				AutoCertInfo: &virtual_host.AutoCertInfoType{
					DnsRecords: []*virtual_host.DNSRecord{
						{
							Name:  "_acme-challenge.http.helloclouds.app",
							Type:  "CNAME",
							Value: "ba6e11d53ef8559ab57a09413121dee4.autocerts.demo1.volterra.us",
						},
						{
							Name:  "_acme-challenge.http.test.com",
							Type:  "CNAME",
							Value: "785e82a6627b5625819f61d5500dbe89.autocerts.demo1.volterra.us",
						},
						{
							Name:  "_acme-challenge.www.f5cert.com",
							Type:  "CNAME",
							Value: "a33980e7a7915a32b0899f34d881a35e.autocerts.demo1.volterra.us",
						},
					},
				},
				AdvertiseChoice: &http_lb.GlobalSpecType_DoNotAdvertise{
					DoNotAdvertise: &ves_io_schema.Empty{},
				},
				ChallengeType: &http_lb.GlobalSpecType_NoChallenge{
					NoChallenge: &ves_io_schema.Empty{},
				},
				Domains: []string{"http.helloclouds.app", "http.test.com", "www.f5cert.com"},
				LoadbalancerType: &http_lb.GlobalSpecType_HttpsAutoCert{
					HttpsAutoCert: &http_lb.ProxyTypeHttpsAutoCerts{
						AddHsts: true,
						PortChoice: &http_lb.ProxyTypeHttpsAutoCerts_Port{
							Port: 443,
						},
					},
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
	dataSourceName := "data.volterra_http_loadbalancer_state.http_lb"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testConfigHTTPLBWithAutoCert(name, "app-test", lbName, nsName, 3000),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "cname", "ves-io-082258df-c45e-4163-8884-7878453eeb2a"),
					resource.TestCheckResourceAttr(dataSourceName, "ip_address", "1.1.1.1"),
					resource.TestCheckResourceAttr(dataSourceName, "auto_cert_info.0.dns_records.#", "3"),
				),
			},
		},
	})
}

func testConfigHTTPLBWithAutoCert(name, namespace, existLbName, existNsName string, timeout int32) string {
	return fmt.Sprintf(`
		resource "volterra_namespace" "app" {
		  name = "%[2]s"
		}
		resource "volterra_http_loadbalancer" "%[1]s" {
		  name = "%[1]s"
		  default_sensitive_data_policy    = true
		  disable_threat_mesh              = true
		  disable_api_definition           = true
          disable_api_discovery            = true       
		  disable_malware_protection       = true       
          disable_malicious_user_detection = true
          disable_trust_client_ip_headers  = true
		  l7_ddos_action_default           = true
		  disable_waf                      = true 
		  user_id_client_ip                = true
		  namespace = volterra_namespace.app.name
		  labels = {
			  "new-test"        = "true"
		  }
		  advertise_on_public_default_vip = true
		  no_challenge = true
		  https_auto_cert {
			add_hsts                 = true 
			connection_idle_timeout  = 120000 
			default_header           = false
			default_loadbalancer     = false
			disable_path_normalize   = false
			enable_path_normalize    = true
			http_redirect            = true
			no_mtls                  = true
			non_default_loadbalancer = true
			pass_through             = false
			port = "443"
		  }
		  disable_rate_limit = true
		  no_service_policies = true
		  round_robin = true
		  domains = ["http.helloclouds.app", "http.test.com", "www.f5cert.com"]
		  multi_lb_app = true
		}
		data "volterra_http_loadbalancer_state" "http_lb" {
			name = "%[3]s"
			namespace = "%[4]s"
		}
		`, name, namespace, existLbName, existNsName, timeout)
}
