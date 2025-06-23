package driftdetection

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"
	ves_io_schema_views_tcp_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/tcp_loadbalancer"
)

func FlattenTlsTcp(x *ves_io_schema_views_tcp_loadbalancer.ProxyTypeTLSTCP) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"tls_cert_params": FlattenTlsCertParams(x.GetTlsCertParams()),
			"tls_parameters":  FlattenTlsParameters(x.GetTlsParameters()),
		}
		rslt = append(rslt, test)
	}
	return rslt
}

func FlattenTlsTcpAutoCert(x *ves_io_schema_views_tcp_loadbalancer.ProxyTypeTLSTCPAutoCerts) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"no_mtls":    isEmpty(x.GetNoMtls()),
			"use_mtls":   FlattenUseMtls(x.GetUseMtls()),
			"tls_config": FlattenTlsConfig(x.GetTlsConfig()),
		}
		rslt = append(rslt, test)
	}
	return rslt
}

func FlattenActiveServicePoliciesTcp(x *ves_io_schema_views_tcp_loadbalancer.ServicePolicyList) []interface{} {
	val := make([]interface{}, 0)

	if x != nil {
		test := map[string]interface{}{
			"policies": FlattenVObjectRefTypeList(x.GetPolicies()),
		}
		val = append(val, test)
	}
	return val
}

func DriftDetectionTcpLoadbalancer(d *schema.ResourceData, resp vesapi.GetObjectResponse) {
	spec := resp.GetObjSpec().(*ves_io_schema_views_tcp_loadbalancer.SpecType)
	d.Set("advertise_custom", FlattenAdvertiseCustom(spec.GcSpec.GetAdvertiseCustom()))
	d.Set("advertise_on_public", FlattenAdvertiseOnPublic(spec.GcSpec.GetAdvertiseOnPublic()))
	d.Set("advertise_on_public_default_vip", isEmpty(spec.GcSpec.GetAdvertiseOnPublicDefaultVip()))
	d.Set("do_not_advertise", isEmpty(spec.GcSpec.GetDoNotAdvertise()))
	d.Set("do_not_retract_cluster", isEmpty(spec.GcSpec.GetDoNotRetractCluster()))
	d.Set("retract_cluster", isEmpty(spec.GcSpec.GetRetractCluster()))
	d.Set("dns_volterra_managed", spec.GcSpec.GetDnsVolterraManaged())
	d.Set("domains", spec.GcSpec.GetDomains())
	d.Set("hash_policy_choice_least_active", isEmpty(spec.GcSpec.GetHashPolicyChoiceLeastActive()))
	d.Set("hash_policy_choice_random", isEmpty(spec.GcSpec.GetHashPolicyChoiceRandom()))
	d.Set("hash_policy_choice_round_robin", isEmpty(spec.GcSpec.GetHashPolicyChoiceRoundRobin()))
	d.Set("hash_policy_choice_source_ip_stickiness", isEmpty(spec.GcSpec.GetHashPolicyChoiceSourceIpStickiness()))
	d.Set("idle_timeout", spec.GcSpec.GetIdleTimeout())
	d.Set("tcp", isEmpty(spec.GcSpec.GetTcp()))
	d.Set("tls_tcp", FlattenTlsTcp(spec.GcSpec.GetTlsTcp()))
	d.Set("tls_tcp_auto_cert", FlattenTlsTcpAutoCert(spec.GcSpec.GetTlsTcpAutoCert()))
	d.Set("origin_pools_weights", FlattenDefaultRoutePools(spec.GcSpec.GetOriginPoolsWeights()))
	d.Set("listen_port", spec.GcSpec.GetListenPort())
	d.Set("port_ranges", spec.GcSpec.GetPortRanges())
	d.Set("active_service_policies", FlattenActiveServicePoliciesTcp(spec.GcSpec.GetActiveServicePolicies()))
	d.Set("no_service_policies", isEmpty(spec.GcSpec.GetNoServicePolicies()))
	d.Set("service_policies_from_namespace", isEmpty(spec.GcSpec.GetServicePoliciesFromNamespace()))
	d.Set("default_lb_with_sni", isEmpty(spec.GcSpec.GetDefaultLbWithSni()))
	d.Set("no_sni", isEmpty(spec.GcSpec.GetNoSni()))
	d.Set("sni", isEmpty(spec.GcSpec.GetSni()))

}
