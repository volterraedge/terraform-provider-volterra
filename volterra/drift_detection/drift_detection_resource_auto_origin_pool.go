package driftdetection

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ves_io_schema_cluster "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cluster"
	ves_io_schema_views_origin_pool "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/origin_pool"
	"gopkg.volterra.us/stdlib/client/vesapi"
)

func FlattenHTTP1Config(x *ves_io_schema_cluster.Http1ProtocolOptions) []interface{} {
	ruiValue := make([]interface{}, 0)
	if x != nil {
		ruiVal := map[string]interface{}{
			"header_transformation": FlattenHeaderTransformationType(x.GetHeaderTransformation()),
		}
		ruiValue = append(ruiValue, ruiVal)
	}
	return ruiValue
}

func FlattenAdvancedOptions(x *ves_io_schema_views_origin_pool.OriginPoolAdvancedOptions) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mpValue := map[string]interface{}{
			"circuit_breaker":                  FlattenCircuitBreaker(x.GetCircuitBreaker()),
			"default_circuit_breaker":          isEmpty(x.GetDefaultCircuitBreaker()),
			"disable_circuit_breaker":          isEmpty(x.GetDisableCircuitBreaker()),
			"connection_timeout":               x.GetConnectionTimeout(),
			"header_transformation_type":       FlattenHeaderTransformationType(x.GetHeaderTransformationType()),
			"http_idle_timeout":                x.GetHttpIdleTimeout(),
			"auto_http_config":                 isEmpty(x.GetAutoHttpConfig()),
			"http1_config":                     FlattenHTTP1Config(x.GetHttp1Config()),
			"http2_options":                    FlattenHttp2Options(x.GetHttp2Options()),
			"disable_lb_source_ip_persistance": isEmpty(x.GetDisableLbSourceIpPersistance()),
			"enable_lb_source_ip_persistance":  isEmpty(x.GetEnableLbSourceIpPersistance()),
			"disable_outlier_detection":        isEmpty(x.GetDisableOutlierDetection()),
			"outlier_detection":                FlattenOutlierDetection(x.GetOutlierDetection()),
			"no_panic_threshold":               isEmpty(x.GetNoPanicThreshold()),
			"panic_threshold":                  x.GetPanicThreshold(),
			"disable_proxy_protocol":           isEmpty(x.GetDisableProxyProtocol()),
			"proxy_protocol_v1":                isEmpty(x.GetProxyProtocolV1()),
			"proxy_protocol_v2":                isEmpty(x.GetProxyProtocolV2()),
			"disable_subsets":                  isEmpty(x.GetDisableSubsets()),
			"enable_subsets":                   FlattenEnableSubsets(x.GetEnableSubsets()),
		}
		rslt = append(rslt, mpValue)
	}
	return rslt
}

func DriftDetectionSpec_OriginPool(d *schema.ResourceData, resp vesapi.GetObjectResponse) {
	spec := resp.GetObjSpec().(*ves_io_schema_views_origin_pool.SpecType)

	d.Set("advanced_options", FlattenAdvancedOptions(spec.GcSpec.GetAdvancedOptions()))

	d.Set("endpoint_selection", spec.GcSpec.GetEndpointSelection().String())

	d.Set("health_check_port", spec.GcSpec.GetHealthCheckPort())

	d.Set("same_as_endpoint_port", isEmpty(spec.GcSpec.GetSameAsEndpointPort()))

	d.Set("healthcheck", FlattenVObjectRefTypeList(spec.GcSpec.GetHealthcheck()))

	d.Set("loadbalancer_algorithm", spec.GcSpec.GetLoadbalancerAlgorithm().String())

	d.Set("origin_servers", FlattenOriginServers(spec.GcSpec.OriginServers))

	d.Set("automatic_port", isEmpty(spec.GcSpec.GetAutomaticPort()))

	d.Set("lb_port", isEmpty(spec.GcSpec.GetLbPort()))

	d.Set("Port", spec.GcSpec.GetPort())

	d.Set("no_tls", isEmpty(spec.GcSpec.GetNoTls()))

	d.Set("use_tls", FlattenUseTls(spec.GcSpec.GetUseTls()))

}
