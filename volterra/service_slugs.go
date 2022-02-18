//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//
package volterra

import (
	"context"
	"fmt"
)

var (
	// apiGWServiceSlugs is a mapping of public API FQN to the service slug to use on the api gateway
	// this will be replaced by discovery when api gateway supports it. The service slug is the part
	// at https://<xyz>.volterra.us/api/<service-slug>/<resource-uri>
	apiGWServiceSlugs = map[string]string{
		"ves.io.schema.advertise_policy.Object":                                 "config",
		"ves.io.schema.alert.CustomAPI.Alerts":                                  "data",
		"ves.io.schema.alert.CustomAPI.AlertsHistory":                           "data",
		"ves.io.schema.alert.CustomAPI.AlertsHistoryScroll":                     "data",
		"ves.io.schema.alert_policy_set.Object":                                 "config",
		"ves.io.schema.alert_policy.Object":                                     "config",
		"ves.io.schema.alert_receiver.Object":                                   "config",
		"ves.io.schema.alert_receiver.CustomAPI.VerifyAlertReceiver":            "alert",
		"ves.io.schema.alert_receiver.CustomAPI.ConfirmAlertReceiverRequest":    "alert",
		"ves.io.schema.alert_receiver.CustomAPI.TestAlertReceiver":              "alert",
		"ves.io.schema.api_credential.CustomAPI.Create":                         "web",
		"ves.io.schema.api_credential.CustomAPI.Get":                            "web",
		"ves.io.schema.views.api_definition.Object":                             "config",
		"ves.io.schema.app_firewall.Object":                                     "config",
		"ves.io.schema.virtual_host.CustomAPI.GetDnsInfo":                       "config",
		"ves.io.schema.registration.CustomAPI.ListRegistrationsByState":         "register",
		"ves.io.schema.registration.CustomAPI.RegistrationApprove":              "register",
		"ves.io.schema.site.CustomStateAPI.SetState":                            "register",
		"ves.io.schema.address_allocator.Object":                                "web",
		"ves.io.schema.api_group.Object":                                        "web",
		"ves.io.schema.api_group_element.Object":                                "web",
		"ves.io.schema.app_setting.Object":                                      "config",
		"ves.io.schema.app_type.Object":                                         "config",
		"ves.io.schema.bgp.Object":                                              "config",
		"ves.io.schema.bgp_asn_set.Object":                                      "config",
		"ves.io.schema.certified_hardware.Object":                               "config",
		"ves.io.schema.cloud_credentials.Object":                                "config",
		"ves.io.schema.cluster.Object":                                          "config",
		"ves.io.schema.contact.Object":                                          "web",
		"ves.io.schema.customer_support.Object":                                 "web",
		"ves.io.schema.container_registry.Object":                               "config",
		"ves.io.schema.discovery.Object":                                        "config",
		"ves.io.schema.dns_domain.Object":                                       "config",
		"ves.io.schema.endpoint.Object":                                         "config",
		"ves.io.schema.fast_acl.Object":                                         "config",
		"ves.io.schema.fast_acl_rule.Object":                                    "config",
		"ves.io.schema.fleet.Object":                                            "config",
		"ves.io.schema.healthcheck.Object":                                      "config",
		"ves.io.schema.ip_prefix_set.Object":                                    "config",
		"ves.io.schema.jwt.Object":                                              "uam",
		"ves.io.schema.jwt_provider.Object":                                     "uam",
		"ves.io.schema.kms_key.Object":                                          "kms",
		"ves.io.schema.kms_policy.Object":                                       "kms",
		"ves.io.schema.kms_policy_rule.Object":                                  "kms",
		"ves.io.schema.k8s_cluster.Object":                                      "config",
		"ves.io.schema.k8s_cluster_role.Object":                                 "config",
		"ves.io.schema.k8s_cluster_role_binding.Object":                         "config",
		"ves.io.schema.k8s_pod_security_policy.Object":                          "config",
		"ves.io.schema.log_receiver.Object":                                     "config",
		"ves.io.schema.namespace.Object":                                        "web",
		"ves.io.schema.namespace_role.Object":                                   "web",
		"ves.io.schema.namespace.NamespaceCustomAPI.SetFastACLsForInternetVIPs": "config",
		"ves.io.schema.namespace.NamespaceCustomAPI.GetFastACLsForInternetVIPs": "config",
		"ves.io.schema.namespace.NamespaceCustomAPI.SetActiveServicePolicies":   "config",
		"ves.io.schema.namespace.NamespaceCustomAPI.GetActiveServicePolicies":   "config",
		"ves.io.schema.namespace.NamespaceCustomAPI.SetActiveNetworkPolicies":   "config",
		"ves.io.schema.namespace.NamespaceCustomAPI.GetActiveNetworkPolicies":   "config",
		"ves.io.schema.namespace.NamespaceCustomAPI.SetActiveAlertPolicies":     "config",
		"ves.io.schema.namespace.NamespaceCustomAPI.GetActiveAlertPolicies":     "config",
		"ves.io.schema.namespace.CustomAPI.CascadeDelete":                       "web",
		"ves.io.schema.network_connector.Object":                                "config",
		"ves.io.schema.network_firewall.Object":                                 "config",
		"ves.io.schema.network_interface.Object":                                "config",
		"ves.io.schema.network_policy.Object":                                   "config",
		"ves.io.schema.network_policy_rule.Object":                              "config",
		"ves.io.schema.network_policy_set.Object":                               "config",
		"ves.io.schema.policer.Object":                                          "config",
		"ves.io.schema.protocol_policer.Object":                                 "config",
		"ves.io.schema.registration.Object":                                     "register",
		"ves.io.schema.role.Object":                                             "web",
		"ves.io.schema.route.Object":                                            "config",
		"ves.io.schema.rate_limiter.Object":                                     "config",
		"ves.io.schema.secret_management_access.Object":                         "config",
		"ves.io.schema.secret_policy.Object":                                    "secret_management",
		"ves.io.schema.secret_policy_rule.Object":                               "secret_management",
		"ves.io.schema.service_policy.Object":                                   "config",
		"ves.io.schema.service_policy_set.Object":                               "config",
		"ves.io.schema.service_policy_rule.Object":                              "config",
		"ves.io.schema.site.Object":                                             "config",
		"ves.io.schema.site_mesh_group.Object":                                  "config",
		"ves.io.schema.timeseries.Object":                                       "config",
		"ves.io.schema.token.Object":                                            "register",
		"ves.io.schema.tunnel.Object":                                           "register",
		"ves.io.schema.user.Object":                                             "web",
		"ves.io.schema.user_identification.Object":                              "config",
		"ves.io.schema.views.aws_vpc_site.Object":                               "config",
		"ves.io.schema.views.aws_vpc_site.CustomAPI.SetVPCK8SHostnames":         "config",
		"ves.io.schema.views.aws_vpc_site.CustomAPI.SetVIPInfo":                 "config",
		"ves.io.schema.views.aws_tgw_site.Object":                               "config",
		"ves.io.schema.views.aws_tgw_site.CustomAPI.SetTGWInfo":                 "config",
		"ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPCIpPrefixes":           "config",
		"ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPNTunnels":              "config",
		"ves.io.schema.views.aws_tgw_site.CustomAPI.SetVIPInfo":                 "config",
		"ves.io.schema.views.azure_vnet_site.Object":                            "config",
		"ves.io.schema.views.azure_vnet_site.CustomAPI.SetVIPInfo":              "config",
		"ves.io.schema.views.gcp_vpc_site.Object":                               "config",
		"ves.io.schema.views.voltstack_site.Object":                             "config",
		"ves.io.schema.views.forward_proxy_policy.Object":                       "config",
		"ves.io.schema.views.http_loadbalancer.Object":                          "config",
		"ves.io.schema.views.http_loadbalancer.CustomAPI":                       "config",
		"ves.io.schema.views.network_policy_view.Object":                        "config",
		"ves.io.schema.views.origin_pool.Object":                                "config",
		"ves.io.schema.views.rate_limiter_policy.Object":                        "config",
		"ves.io.schema.views.tcp_loadbalancer.Object":                           "config",
		"ves.io.schema.views.tcp_loadbalancer.CustomAPI":                        "config",
		"ves.io.schema.views.terraform_parameters.CustomAPI.GetStatus":          "config",
		"ves.io.schema.views.terraform_parameters.CustomActionAPI.Run":          "terraform",
		"ves.io.schema.views.view_internal.CustomAPI":                           "config",
		"ves.io.schema.virtual_k8s.Object":                                      "config",
		"ves.io.schema.virtual_host.Object":                                     "config",
		"ves.io.schema.virtual_network.Object":                                  "config",
		"ves.io.schema.virtual_site.Object":                                     "config",
		"ves.io.schema.waf.Object":                                              "config",
		"ves.io.schema.waf_rules.Object":                                        "config",
	}
)

// API gateway service slug e.g. "config"
type ctxServiceSlug struct{}

// contextFromServiceSlug sets the api gateway service discriminator in the context
func contextFromServiceSlug(ctx context.Context, slug string) context.Context {
	return context.WithValue(ctx, ctxServiceSlug{}, slug)
}

// addSvcSlugToContext get slugs for given objType
func addSvcSlugToContext(ctx context.Context, objType string) (context.Context, error) {
	if svcSlug, ok := apiGWServiceSlugs[objType]; ok {
		return contextFromServiceSlug(ctx, svcSlug), nil
	}
	return nil, fmt.Errorf("Unknown service slug for object type %s", objType)
}
