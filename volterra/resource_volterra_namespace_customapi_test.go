//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_ap "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/alert_policy"
	ves_io_schema_ap_receiver "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/alert_receiver"
	ves_io_schema_facl "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fast_acl"
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
	ves_io_schema_np "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_policy"
	ves_io_schema_policer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policer"
	ves_io_schema_pp "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/protocol_policer"
	ves_io_schema_sp "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy"
)

// TestNamespaceCustomAPI namespace custom api test
func TestNamespaceCustomAPI(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc, _ := createTestCustomAPIServer(t, []string{
		ves_io_schema_sp.ObjectType,
		ves_io_schema_np.ObjectType,
		ves_io_schema_pp.ObjectType,
		ves_io_schema_policer.ObjectType,
		ves_io_schema_facl.ObjectType,
		ves_io_schema_ns.ObjectType,
		ves_io_schema_ap.ObjectType,
		ves_io_schema_ap_receiver.ObjectType,
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
				Config: testActiveServicePolicy(name, "test"),
			},
			{
				Config: testActiveNetworkPolicy(name, "test"),
			},
			{
				Config: testFastACLsForInternetVIPs(name, "system"),
			},
			{
				Config: testActiveAlertPolicy(name, "system"),
			},
		},
	})
}

func testActiveServicePolicy(name, namespace string) string {
	return fmt.Sprintf(`
		resource "volterra_namespace" "system" {
			name = "%[2]s"
		}
		resource "volterra_service_policy" "%[1]s" {
			name        = "%[1]s-test"
			namespace   = volterra_namespace.system.name
			algo        = "FIRST_MATCH"
			any_server  = true

			allow_all_requests = true
			deny_all_requests  = false
		}

		resource "volterra_active_service_policies" "%[1]s" {
			depends_on = [volterra_service_policy.%[1]s]
			namespace = volterra_namespace.system.name
			policies {
				name = "%[1]s-test"
				namespace = volterra_namespace.system.name
				tenant = "ves-io"
			}
		}
		`, name, namespace)
}

func testActiveNetworkPolicy(name, namespace string) string {
	return fmt.Sprintf(`
		resource "volterra_namespace" "system" {
			name = "%[2]s"
		}
		resource "volterra_network_policy" "%[1]s" {
			name        = "%[1]s-test"
			namespace   = volterra_namespace.system.name
			endpoint {
				any = true
			}
			rules {
				egress_rules {
					action = "ALLOW"
					metadata {
						name = "allow-egress-all"
					}
					any = true
					all_traffic = true
				}
				ingress_rules {
					action = "ALLOW"
					metadata {
						name = "allow-ingress-all"
					}
					any = true
					all_traffic = true
				}
			}

		}

		resource "volterra_active_network_policies" "%[1]s" {
			depends_on = [volterra_network_policy.%[1]s]
			namespace = volterra_namespace.system.name
			policies {
				name = "%[1]s-test"
				namespace = volterra_namespace.system.name
				tenant = "ves-io"
			}
		}
		`, name, namespace)
}

func testFastACLsForInternetVIPs(name, namespace string) string {
	return fmt.Sprintf(`
		resource "volterra_policer" "this" {
			name = "test-policer"
			namespace = "%[2]s"
			policer_type = "POLICER_SINGLE_RATE_TWO_COLOR"
			committed_information_rate = 10000
			burst_size = 5000
			policer_mode = "POLICER_MODE_SHARED"
		}
		resource "volterra_protocol_policer" "this" {
			name = "test-proto-policer"
			namespace = "%[2]s"
			protocol_policer {
				policer {
					name = volterra_policer.this.name
					namespace = "%[2]s"
				}
				protocol {
					tcp {
						flags = ["ALL_TCP_FLAGS"]
					}
				}
			}
		}
		resource "volterra_fast_acl" "this" {
			name        = "%[1]s-test"
			namespace   = "%[2]s"
			re_acl {
				all_public_vips = true
				fast_acl_rules {
					action {
						simple_action = "DENY"
					}
					metadata {
						name = "deny-certain-ip"
					}
					prefix {
						prefix = ["1.1.1.1/32"]
					}
				}
			}
		}

		resource "volterra_fast_acl_for_internet_vips" "%[1]s" {
			depends_on = [volterra_fast_acl.this]
			fast_acls {
				name = "%[1]s-test"
				tenant = "ves-io"
			}
		}
		`, name, namespace)
}

func testActiveAlertPolicy(name, namespace string) string {
	return fmt.Sprintf(`
		resource "volterra_namespace" "system" {
			name = "%[2]s"
		}
		resource "volterra_alert_receiver" "email"{
			name = "email-alert"
			namespace   = volterra_namespace.system.name
			email {
				email = "support@ves.io"
			}
		}
		resource "volterra_alert_policy" "%[1]s"{
			name = "dev-policy"
			namespace   = volterra_namespace.system.name
			receivers {
				name = volterra_alert_receiver.email.name
				namespace = volterra_namespace.system.name
			}
			routes {
				send = true
				any = true
			}
		}

		resource "volterra_active_alert_policies" "%[1]s" {
			depends_on = [volterra_alert_policy.%[1]s]
			namespace = volterra_namespace.system.name
			policies {
				name = "dev-policy"
				namespace = volterra_namespace.system.name
				tenant = "ves-io"
			}
		}
		`, name, namespace)
}
