//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ves_io_schema_k8s_cluster "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/k8s_cluster"
	ves_io_schema_k8s_cluster_role "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/k8s_cluster_role"
	ves_io_schema_k8s_cluster_role_binding "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/k8s_cluster_role_binding"
	ves_io_schema_k8s_pod_security_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/k8s_pod_security_policy"
	ves_io_schema_voltstack_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/voltstack_site"
)

func TestPK8sSite(t *testing.T) {
	name := generateResourceName()
	testURL, stopFunc := createTestCustomAPIServer(t, []string{
		ves_io_schema_voltstack_site.ObjectType,
		ves_io_schema_k8s_cluster.ObjectType,
		ves_io_schema_k8s_cluster_role.ObjectType,
		ves_io_schema_k8s_cluster_role_binding.ObjectType,
		ves_io_schema_k8s_pod_security_policy.ObjectType,
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
				Config: testPk8sSiteConfig(name),
			},
		},
	})
}

func testPk8sSiteConfig(name string) string {
	return fmt.Sprintf(`
		resource "volterra_k8s_pod_security_policy" "this" {
		  name = "ves-io-psp-permissive"
		  namespace = "system"
		  psp_spec {
			privileged = true
			allow_privilege_escalation = true
			no_default_capabilities = true
			allowed_capabilities {
				capabilities = [ "*" ]
			}
			no_drop_capabilities = true
			volumes = ["*"]
			allowed_flex_volumes = []
			allowed_proc_mounts = []
			read_only_root_filesystem = false
			allowed_csi_drivers = []
			host_network = true
			host_port_ranges = "0-65535"
			host_ipc = true
			host_pid = true
			allowed_unsafe_sysctls = []
			forbidden_sysctls = []
			no_run_as_user = true
			no_run_as_group = true
			no_supplemental_groups = true
			no_fs_groups = true
			no_se_linux_options = true
			no_runtime_class = true
		  }
		}
		resource "volterra_k8s_cluster_role" "this" {
			name = "ves-io-psp-permissive"
			namespace = "system"
			policy_rule_list {
				policy_rule {
					resource_list {
						api_groups = ["extensions"]
						resource_types = ["podsecuritypolicies"]
						resource_instances = []
						verbs = ["use"]
					}
				}
			}
		}
		resource "volterra_k8s_cluster_role_binding" "this" {
			name = "ves-io-psp-permissive"
			namespace = "system"
			k8s_cluster_role {
				name = volterra_k8s_cluster_role.this.name
				namespace = "system"
			}
			subjects {
				group = "system:authenticated"
			}
		}
		resource "volterra_k8s_cluster" "this" {
			name = "%[1]s"
			namespace = "system"
			local_access_config {
				local_domain = "test.example.com"
				default_port = true
			}
			no_local_access = false
			global_access_enable = true
			no_insecure_registries = true
			use_custom_cluster_role_bindings {
				cluster_role_bindings {
					name = volterra_k8s_cluster_role_binding.this.name
					namespace = "system"
				}
			}
			use_custom_psp_list {
				pod_security_policies {
					name = volterra_k8s_pod_security_policy.this.name
					namespace = "system"
				}
			}
			use_default_psp = false
			use_default_cluster_roles = true
			no_cluster_wide_apps = true
		}
		resource "volterra_voltstack_site" "this" {
			name = "%[1]s"
			namespace = "system"
			no_bond_devices = true
			disable_gpu = true
			no_k8s_cluster = false
			k8s_cluster {
				name = volterra_k8s_cluster.this.name
				namespace = "system"
			}
			master_nodes = [ "master-0" ]
			custom_network_config {
			  forward_proxy_allow_all = true
			  no_global_network = true
			  default_interface_config = true
			  no_network_policy = true
			  slo_config {
				  no_dc_cluster_group = true
				  no_static_routes = true
			  }
			}
			default_network_config = false
			default_storage_config = true
			volterra_certified_hw = "generic-single-nic-volstack-combo"
			logs_streaming_disabled = true
			allow_all_usb = true
			no_local_control_plane = true
		  }`, name)
}
