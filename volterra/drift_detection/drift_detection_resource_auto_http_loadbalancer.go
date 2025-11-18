package driftdetection

import (
	"strconv"

	"github.com/gogo/protobuf/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_app_type "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/app_type"
	ves_io_schema_cluster "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cluster"
	ves_io_schema_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policy"
	ves_io_schema_rate_limiter "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/rate_limiter"
	ves_io_schema_route "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/route"
	ves_io_schema_service_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy_rule"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/common_cache_rule"
	ves_io_schema_views_common_security "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/common_security"
	ves_io_schema_views_common_waf "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/common_waf"
	ves_io_schema_views_http_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
	ves_io_schema_views_origin_pool "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/origin_pool"
	ves_io_schema_rate_limiter_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/rate_limiter_policy"
	ves_io_schema_virtual_host "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
	"gopkg.volterra.us/stdlib/client/vesapi"
)

func FlattenVk8s(x *ves_io_schema_views.WhereVK8SService) []interface{} {
	vk8sValue := make([]interface{}, 0)
	if x != nil {
		vk8sVal := map[string]interface{}{
			"site":         FlattenObjectRefTypeSet(x.GetSite()),
			"virtual_site": FlattenObjectRefTypeSet(x.GetVirtualSite()),
		}
		vk8sValue = append(vk8sValue, vk8sVal)
	}
	return vk8sValue
}

func FlattenObjectRefTypeSet(x *ves_io_schema_views.ObjectRefType) []interface{} {
	res := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"name":      x.GetName(),
			"namespace": x.GetNamespace(),
			"tenant":    x.GetTenant(),
		}
		res = append(res, val)
	}
	return res
}

func FlattenObjectRefTypeList(x []*ves_io_schema.ObjectRefType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		t := map[string]interface{}{
			"kind":      v.GetKind(),
			"name":      v.GetName(),
			"namespace": v.GetNamespace(),
			"tenant":    v.GetTenant(),
		}
		rslt = append(rslt, t)
	}
	return rslt
}

func FlattenVObjectRefTypeList(x []*ves_io_schema_views.ObjectRefType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		t := map[string]interface{}{
			"name":      v.GetName(),
			"namespace": v.GetNamespace(),
			"tenant":    v.GetTenant(),
		}
		rslt = append(rslt, t)
	}
	return rslt
}

func FlattenVirtualSite(x *ves_io_schema_views.WhereVirtualSite) []interface{} {
	vSValue := make([]interface{}, 0)
	if x != nil {
		vSVal := map[string]interface{}{
			"network":      x.GetNetwork().String(),
			"virtual_site": FlattenObjectRefTypeSet(x.GetVirtualSite()),
		}
		vSValue = append(vSValue, vSVal)
	}
	return vSValue
}

func FlattenVirtualNetwork(x *ves_io_schema_views.WhereVirtualNetwork) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		vNVal := map[string]interface{}{
			"default_v6_vip":  isEmpty(x.GetDefaultV6Vip()),
			"specific_v6_vip": x.GetSpecificV6Vip(),
			"default_vip":     isEmpty(x.GetDefaultVip()),
			"specific_vip":    x.GetSpecificVip(),
			"virtual_network": FlattenObjectRefTypeSet(x.GetVirtualNetwork()),
		}
		rslt = append(rslt, vNVal)
	}
	return rslt
}

func FlattenSite(x *ves_io_schema_views.WhereSite) []interface{} {
	siteValue := make([]interface{}, 0)
	if x != nil {
		siteVal := map[string]interface{}{
			"ip":      x.GetIp(),
			"ipv6":    x.GetIpv6(),
			"network": x.GetNetwork().String(),
			"site":    FlattenObjectRefTypeSet(x.GetSite()),
		}
		siteValue = append(siteValue, siteVal)
	}
	return siteValue
}

func FlattenSiteSegment(x *ves_io_schema_views.WhereSiteSegment) []interface{} {
	siteSegmentValue := make([]interface{}, 0)
	if x != nil {
		siteSegmentVal := map[string]interface{}{
			"ip":      x.GetIp(),
			"ipv6":    x.GetIpv6(),
			"segment": FlattenObjectRefTypeSet(x.GetSegment()),
			"site":    FlattenObjectRefTypeSet(x.GetSite()),
		}
		siteSegmentValue = append(siteSegmentValue, siteSegmentVal)
	}
	return siteSegmentValue
}

func FlattenVirtualSiteSegment(x *ves_io_schema_views.WhereVirtualSiteSegment) []interface{} {
	virtualSiteSegmentValue := make([]interface{}, 0)
	if x != nil {
		virtualSiteSegmentVal := map[string]interface{}{
			"ip":           x.GetIp(),
			"ipv6":         x.GetIpv6(),
			"segment":      FlattenObjectRefTypeSet(x.GetSegment()),
			"virtual_site": FlattenObjectRefTypeSet(x.GetVirtualSite()),
		}
		virtualSiteSegmentValue = append(virtualSiteSegmentValue, virtualSiteSegmentVal)
	}
	return virtualSiteSegmentValue
}

func FlattenAdvertiseWhere(x []*ves_io_schema_views.WhereType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"advertise_on_public":   FlattenAdvertiseOnPublic(val.GetAdvertiseOnPublic()),
			"site":                  FlattenSite(val.GetSite()),
			"site_segment":          FlattenSiteSegment(val.GetSiteSegment()),
			"virtual_network":       FlattenVirtualNetwork(val.GetVirtualNetwork()),
			"virtual_site":          FlattenVirtualSite(val.GetVirtualSite()),
			"virtual_site_with_vip": FlattenVirtualSiteWithVip(val.GetVirtualSiteWithVip()),
			"virtual_site_segment":  FlattenVirtualSiteSegment(val.GetVirtualSiteSegment()),
			"vk8s_service":          FlattenVk8s(val.GetVk8SService()),
			"port":                  val.GetPort(),
			"port_ranges":           val.GetPortRanges(),
			"use_default_port":      isEmpty(val.GetUseDefaultPort()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenVirtualSiteWithVip(x *ves_io_schema_views.WhereVirtualSiteSpecifiedVIP) []interface{} {
	virtualSiteWithVipValue := make([]interface{}, 0)
	if x != nil {
		virtualSiteWithVipVal := map[string]interface{}{
			"ip":           x.GetIp(),
			"ipv6":         x.GetIpv6(),
			"network":      x.GetNetwork().String(),
			"virtual_site": FlattenObjectRefTypeSet(x.GetVirtualSite()),
		}
		virtualSiteWithVipValue = append(virtualSiteWithVipValue, virtualSiteWithVipVal)
	}
	return virtualSiteWithVipValue
}

func FlattenWhSegment(x *ves_io_schema_views.WhereSegment) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"ipv4_vip": x.GetIpv4Vip(),
			"ipv6_vip": x.GetIpv6Vip(),
			"segment":  FlattenObjectRefTypeSet(x.GetSegment()),
		}
		rslt = append(rslt, test)
	}
	return rslt
}

func FlattenCloudEdgeSegment(x *ves_io_schema_views.WhereCloudEdgeSegment) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"cloud_edge": FlattenObjectRefTypeSet(x.GetCloudEdge()),
			"ip":         x.GetIp(),
			"ipv6":       x.GetIpv6(),
			"segment":    FlattenObjectRefTypeSet(x.GetSegment()),
		}
		rslt = append(rslt, test)
	}
	return rslt
}

func FlattenAdvertiseCustom(x *ves_io_schema_views.AdvertiseCustom) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"advertise_where": FlattenAdvertiseWhere(x.GetAdvertiseWhere()),
		}
		rslt = append(rslt, test)
	}
	return rslt
}

func FlattenAdvertiseOnPublic(x *ves_io_schema_views.AdvertisePublic) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"public_ip": FlattenObjectRefTypeSet(x.GetPublicIp()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenApiDefinitions(x *ves_io_schema_views_common_waf.ApiDefinitionList) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"api_definitions": FlattenVObjectRefTypeList(x.GetApiDefinitions()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenMetadata(x *ves_io_schema.MessageMetaType) []interface{} {
	metadataValue := make([]interface{}, 0)
	if x != nil {
		mdVal := map[string]interface{}{
			"description": x.GetDescription(),
			"name":        x.GetName(),
		}
		metadataValue = append(metadataValue, mdVal)
	}
	return metadataValue
}

func FlattenApiEndPoint(x *ves_io_schema_views_common_waf.ApiEndpointDetails) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"methods": FlattenMethods(x.GetMethods()),
			"path":    x.GetPath(),
		}

		rslt = append(rslt, test)
	}
	return rslt
}

func FlattenVASEOpenApiValidationRules(x []*ves_io_schema_views_common_waf.FallThroughRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"action_block":  isEmpty(val.GetActionBlock()),
			"action_report": isEmpty(val.GetActionReport()),
			"action_skip":   isEmpty(val.GetActionSkip()),
			"api_endpoint":  FlattenApiEndPoint(val.GetApiEndpoint()),
			"api_group":     val.GetApiGroup(),
			"base_path":     val.GetBasePath(),
			"metadata":      FlattenMetadata(val.GetMetadata()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenVCLFTMOpenApiValidationRules(x []*ves_io_schema_views_common_waf.FallThroughRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"action_block":  isEmpty(val.GetActionBlock()),
			"action_report": isEmpty(val.GetActionReport()),
			"action_skip":   isEmpty(val.GetActionSkip()),
			"api_endpoint":  FlattenApiEndPoint(val.GetApiEndpoint()),
			"api_group":     val.GetApiGroup(),
			"base_path":     val.GetBasePath(),
			"metadata":      FlattenMetadata(val.GetMetadata()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenRequestValidationProperties(x []ves_io_schema.OpenApiValidationProperties) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenVMA(x *ves_io_schema_views_common_waf.OpenApiValidationModeActive) []interface{} {
	vMAValue := make([]interface{}, 0)
	if x != nil {
		vMAVal := map[string]interface{}{
			"request_validation_properties": FlattenRequestValidationProperties(x.GetRequestValidationProperties()),
			"enforcement_block":             isEmpty(x.GetEnforcementBlock()),
			"enforcement_report":            isEmpty(x.GetEnforcementReport()),
		}
		vMAValue = append(vMAValue, vMAVal)
	}
	return vMAValue
}

func FlattenVCLOpenApiValidationRules(x []*ves_io_schema_views_common_waf.OpenApiValidationRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"api_endpoint":    FlattenApiEndPoint(val.GetApiEndpoint()),
			"api_group":       val.GetApiGroup(),
			"base_path":       val.GetBasePath(),
			"any_domain":      isEmpty(val.GetAnyDomain()),
			"specific_domain": val.GetSpecificDomain(),
			"metadata":        FlattenMetadata(val.GetMetadata()),
			"validation_mode": FlattenValidationMode(val.GetValidationMode()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenFallThroughModeCustom(x *ves_io_schema_views_common_waf.CustomFallThroughMode) []interface{} {
	ftmcValue := make([]interface{}, 0)
	if x != nil {
		ftmcVal := map[string]interface{}{
			"open_api_validation_rules": FlattenVASEOpenApiValidationRules(x.GetOpenApiValidationRules()),
		}
		ftmcValue = append(ftmcValue, ftmcVal)
	}
	return ftmcValue
}

func FlattenFallThroughMode(x *ves_io_schema_views_common_waf.OpenApiFallThroughMode) []interface{} {
	ftmValue := make([]interface{}, 0)
	if x != nil {
		ftmVal := map[string]interface{}{
			"fall_through_mode_allow":  isEmpty(x.GetFallThroughModeAllow()),
			"fall_through_mode_custom": FlattenFallThroughModeCustom(x.GetFallThroughModeCustom()),
		}
		ftmValue = append(ftmValue, ftmVal)
	}
	return ftmValue
}

func FlattenQueryParameters(x *ves_io_schema_views_common_waf.ValidationSettingForQueryParameters) []interface{} {
	qpValue := make([]interface{}, 0)
	if x != nil {
		qpeValue := map[string]interface{}{
			"allow_additional_headers":    isEmpty(x.GetAllowAdditionalParameters()),
			"disallow_additional_headers": isEmpty(x.GetDisallowAdditionalParameters()),
		}
		qpValue = append(qpValue, qpeValue)
	}
	return qpValue
}

func FlattenPVSCHeaders(x *ves_io_schema_views_common_waf.ValidationSettingForHeaders) []interface{} {
	hValue := make([]interface{}, 0)
	if x != nil {
		heValue := map[string]interface{}{
			"allow_additional_headers":    x.GetAllowAdditionalHeaders() != nil,
			"disallow_additional_headers": x.GetDisallowAdditionalHeaders() != nil,
		}
		hValue = append(hValue, heValue)
	}
	return hValue
}

func FlattenPropertyValidationSettingsCustom(x *ves_io_schema_views_common_waf.ValidationPropertySetting) []interface{} {
	pvscValue := make([]interface{}, 0)
	if x != nil {
		pvsceValue := map[string]interface{}{
			"query_parameters": FlattenQueryParameters(x.GetQueryParameters()),
		}
		pvscValue = append(pvscValue, pvsceValue)
	}
	return pvscValue
}

func FlattenSettings(x *ves_io_schema_views_common_waf.OpenApiValidationCommonSettings) []interface{} {
	sValue := make([]interface{}, 0)
	if x != nil {
		seVal := map[string]interface{}{
			"oversized_body_fail_validation":       x.GetOversizedBodyFailValidation() != nil,
			"oversized_body_skip_validation":       x.GetOversizedBodySkipValidation() != nil,
			"property_validation_settings_custom":  FlattenPropertyValidationSettingsCustom(x.GetPropertyValidationSettingsCustom()),
			"property_validation_settings_default": x.GetPropertyValidationSettingsDefault() != nil,
		}
		sValue = append(sValue, seVal)
	}
	return sValue
}

func FlattenResponseValidationModeActive(x *ves_io_schema_views_common_waf.OpenApiValidationModeActiveResponse) []interface{} {
	rvmaValue := make([]interface{}, 0)
	if x != nil {
		rvmaeValue := map[string]interface{}{
			"response_validation_properties": FlattenRequestValidationProperties(x.GetResponseValidationProperties()),
			"enforcement_block":              x.GetEnforcementBlock() != nil,
			"enforcement_report":             x.GetEnforcementReport() != nil,
		}
		rvmaValue = append(rvmaValue, rvmaeValue)
	}
	return rvmaValue
}

func FlattenValidationModeActive(x *ves_io_schema_views_common_waf.OpenApiValidationModeActive) []interface{} {
	vmaValue := make([]interface{}, 0)
	if x != nil {
		vmaVal := map[string]interface{}{
			"request_validation_properties": FlattenRequestValidationProperties(x.GetRequestValidationProperties()),
			"enforcement_block":             isEmpty(x.GetEnforcementBlock()),
			"enforcement_report":            isEmpty(x.GetEnforcementReport()),
		}
		vmaValue = append(vmaValue, vmaVal)
	}
	return vmaValue
}

func FlattenValidationMode(x *ves_io_schema_views_common_waf.OpenApiValidationMode) []interface{} {
	vmValue := make([]interface{}, 0)
	if x != nil {
		vmVal := map[string]interface{}{
			"response_validation_mode_active": FlattenResponseValidationModeActive(x.GetResponseValidationModeActive()),
			"skip_response_validation":        x.GetSkipResponseValidation() != nil,
			"skip_validation":                 isEmpty(x.GetSkipValidation()),
			"validation_mode_active":          FlattenValidationModeActive(x.GetValidationModeActive()),
		}
		vmValue = append(vmValue, vmVal)
	}
	return vmValue
}

func FlattenValidationAllSpecEndpoints(x *ves_io_schema_views_common_waf.OpenApiValidationAllSpecEndpointsSettings) []interface{} {
	valseValue := make([]interface{}, 0)
	if x != nil {
		valseVal := map[string]interface{}{
			"fall_through_mode": FlattenFallThroughMode(x.GetFallThroughMode()),
			"settings":          FlattenSettings(x.GetSettings()),
			"validation_mode":   FlattenValidationMode(x.GetValidationMode()),
		}
		valseValue = append(valseValue, valseVal)
	}
	return valseValue
}

func FlattenOAVR(x []*ves_io_schema_views_common_waf.OpenApiValidationRule) []interface{} {
	oavrValue := make([]interface{}, 0)
	for _, v := range x {
		t := map[string]interface{}{
			"api_endpoint":    FlattenApiEndPoint(v.GetApiEndpoint()),
			"api_group":       v.GetApiGroup(),
			"base_path":       v.GetBasePath(),
			"any_domain":      isEmpty(v.GetAnyDomain()),
			"specific_domain": v.GetSpecificDomain(),
			"metadata":        FlattenMetadata(v.GetMetadata()),
			"validation_mode": FlattenValidationMode(v.GetValidationMode()),
		}
		oavrValue = append(oavrValue, t)
	}
	return oavrValue
}

func FlattenValidationCustomList(x *ves_io_schema_views_common_waf.ValidateApiBySpecRule) []interface{} {
	vclValue := make([]interface{}, 0)
	if x != nil {
		vclVal := map[string]interface{}{
			"fall_through_mode":              FlattenFallThroughMode(x.GetFallThroughMode()),
			"open_api_validation_rules":      FlattenOAVR(x.GetOpenApiValidationRules()),
			"oversized_body_fail_validation": x.GetOversizedBodyFailValidation(),
			"oversized_body_skip_validation": x.GetOversizedBodySkipValidation(),
			"settings":                       FlattenSettings(x.GetSettings()),
		}
		vclValue = append(vclValue, vclVal)
	}
	return vclValue
}

func FlattenApiSpecification(x *ves_io_schema_views_common_waf.APISpecificationSettings) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mpValue := map[string]interface{}{
			"api_definition":                FlattenObjectRefTypeSet(x.GetApiDefinition()),
			"validation_all_spec_endpoints": FlattenValidationAllSpecEndpoints(x.GetValidationAllSpecEndpoints()),
			"validation_custom_list":        FlattenValidationCustomList(x.GetValidationCustomList()),
			"validation_disabled":           isEmpty(x.GetValidationDisabled()),
		}
		rslt = append(rslt, mpValue)
	}
	return rslt
}

func FlattenDiscoveredApiSettings(x *ves_io_schema_app_type.DiscoveredAPISettings) []interface{} {
	dasValue := make([]interface{}, 0)
	if x != nil {
		dasVal := map[string]interface{}{
			"purge_duration_for_inactive_discovered_apis": x.GetPurgeDurationForInactiveDiscoveredApis(),
		}
		dasValue = append(dasValue, dasVal)
	}
	return dasValue
}

func FlattenEnableApiDiscovery(x *ves_io_schema_views_common_waf.ApiDiscoverySetting) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"api_crawler":                         FlattenApiCrawler(x.GetApiCrawler()),
			"api_discovery_from_code_scan":        FlattenApiDiscoveryFromCodeScan(x.GetApiDiscoveryFromCodeScan()),
			"custom_api_auth_discovery":           FlattenCustomApiAuthDiscovery(x.GetCustomApiAuthDiscovery()),
			"default_api_auth_discovery":          isEmpty(x.GetDefaultApiAuthDiscovery()),
			"discovered_api_settings":             FlattenDiscoveredApiSettings(x.GetDiscoveredApiSettings()),
			"disable_learn_from_redirect_traffic": isEmpty(x.GetDisableLearnFromRedirectTraffic()),
			"enable_learn_from_redirect_traffic":  isEmpty(x.GetEnableLearnFromRedirectTraffic()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenCustomApiAuthDiscovery(x *ves_io_schema_views_common_waf.ApiDiscoveryAdvancedSettings) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"api_discovery_ref": FlattenObjectRefTypeSet(x.GetApiDiscoveryRef()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenApiCrawler(x *ves_io_schema_views_common_waf.ApiCrawler) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"api_crawler_config":  FlattenApiCrawlerConfig(x.GetApiCrawlerConfig()),
			"disable_api_crawler": isEmpty(x.GetDisableApiCrawler()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenApiCrawlerConfig(x *ves_io_schema_views_common_waf.ApiCrawlerConfiguration) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"domains": FlattenDomains(x.GetDomains()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenDomains(x []*ves_io_schema_views_common_waf.DomainConfiguration) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		val := map[string]interface{}{
			"domain":       v.GetDomain(),
			"simple_login": FlattenSimpleLogin(v.GetSimpleLogin()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenSimpleLogin(x *ves_io_schema_views_common_waf.SimpleLogin) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"password": FlattenPassword(x.GetPassword()),
			"user":     x.GetUser(),
		}
		val = append(val, test)
	}
	return val

}

func FlattenPassword(x *ves_io_schema.SecretType) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"blindfold_secret_info": FlattenBlindFoldSecretInfo(x.GetBlindfoldSecretInfo()),
			"clear_secret_info":     FlattenClearSecretInfo(x.GetClearSecretInfo()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenApiDiscoveryFromCodeScan(f *ves_io_schema_views_common_waf.ApiDiscoveryFromCodeScan) []interface{} {
	val := make([]interface{}, 0)
	if f != nil {
		test := map[string]interface{}{
			"code_base_integrations": FlattenCodeBaseIntegrations(f.GetCodeBaseIntegrations()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenCodeBaseIntegrations(x []*ves_io_schema_views_common_waf.CodeBaseIntegrationSelection) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		val := map[string]interface{}{
			"all_repos":             isEmpty(v.GetAllRepos()),
			"selected_repos":        FlattenSelectedRepos(v.GetSelectedRepos()),
			"code_base_integration": FlattenObjectRefTypeSet(v.GetCodeBaseIntegration()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenSelectedRepos(x *ves_io_schema_views_common_waf.ApiCodeRepos) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"api_code_repo": x.GetApiCodeRepo(),
		}
		val = append(val, test)
	}
	return val
}

func FlattenSensitiveDataDetectionRules(x *ves_io_schema_app_type.SensitiveDataDetectionRules) []interface{} {
	sddrValue := make([]interface{}, 0)
	if x != nil {
		sddrVal := map[string]interface{}{
			"custom_sensitive_data_detection_rules": FlattenCustomSensitiveDataDetectionRules(x.GetCustomSensitiveDataDetectionRules()),
			"disabled_built_in_rules":               FlattenDisabledBuiltInRules(x.GetDisabledBuiltInRules()),
		}
		sddrValue = append(sddrValue, sddrVal)
	}
	return sddrValue
}

func FlattenCustomSensitiveDataDetectionRules(x []*ves_io_schema_app_type.CustomSensitiveDataDetectionRule) []interface{} {
	csddrValue := make([]interface{}, 0)
	for _, v := range x {
		csddrVal := map[string]interface{}{
			"metadata":                        FlattenMetadata(v.GetMetadata()),
			"sensitive_data_detection_config": FlattenSensitiveDataDetectionConfig(v.GetSensitiveDataDetectionConfig()),
		}
		csddrValue = append(csddrValue, csddrVal)
	}
	return csddrValue
}

func FlattenSensitiveDataDetectionConfig(v *ves_io_schema_app_type.CustomDataDetectionConfig) []interface{} {
	sddcValue := make([]interface{}, 0)
	if v != nil {
		sddcVal := map[string]interface{}{
			"any_domain":            isEmpty(v.GetAnyDomain()),
			"specific_domain":       v.GetSpecificDomain(),
			"key_pattern":           FlattenKeyPattern(v.GetKeyPattern()),
			"key_value_pattern":     FlattenKeyValuePattern(v.GetKeyValuePattern()),
			"value_pattern":         FlattenValuePattern(v.GetValuePattern()),
			"all_request_sections":  v.GetAllRequestSections() != nil,
			"all_response_sections": v.GetAllResponseSections() != nil,
			"all_sections":          v.GetAllSections() != nil,
			"custom_sections":       FlattenCustomSections(v.GetCustomSections()),
			"any_target":            v.GetAnyTarget() != nil,
			"api_endpoint_target":   FlattenApiEndpointTarget(v.GetApiEndpointTarget()),
			"api_group":             v.GetApiGroup(),
			"base_path":             v.GetBasePath(),
		}
		sddcValue = append(sddcValue, sddcVal)
	}
	return sddcValue
}

func FlattenApiEndpointTarget(v *ves_io_schema_app_type.APIEndpoint) []interface{} {
	aetValue := make([]interface{}, 0)
	if v != nil {
		aetVal := map[string]interface{}{
			"api_endpoint_path": v.GetApiEndpointPath(),
			"methods":           FlattenMethods(v.GetMethods()),
		}
		aetValue = append(aetValue, aetVal)
	}
	return aetValue
}

func FlattenCustomSectionsSliceToString(x []ves_io_schema.HttpSections) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenCustomSections(v *ves_io_schema_app_type.CustomSections) []interface{} {
	csValue := make([]interface{}, 0)
	if v != nil {
		csVal := map[string]interface{}{
			"custom_sections": FlattenCustomSectionsSliceToString(v.GetCustomSections()),
		}
		csValue = append(csValue, csVal)
	}
	return csValue
}

func FlattenKeyPattern(v *ves_io_schema_app_type.KeyPattern) []interface{} {
	kpValue := make([]interface{}, 0)
	if v != nil {
		kpVal := map[string]interface{}{
			"exact_value": v.GetExactValue(),
			"regex_value": v.GetRegexValue(),
		}
		kpValue = append(kpValue, kpVal)
	}
	return kpValue
}

func FlattenValuePattern(v *ves_io_schema_app_type.ValuePattern) []interface{} {
	vpValue := make([]interface{}, 0)
	if v != nil {
		vpVal := map[string]interface{}{
			"exact_value": v.GetExactValue(),
			"regex_value": v.GetRegexValue(),
		}
		vpValue = append(vpValue, vpVal)
	}
	return vpValue
}

func FlattenKeyValuePattern(v *ves_io_schema_app_type.KeyValuePattern) []interface{} {
	kvpValue := make([]interface{}, 0)
	if v != nil {
		kvpVal := map[string]interface{}{
			"key_pattern":   FlattenKeyPattern(v.GetKeyPattern()),
			"value_pattern": FlattenValuePattern(v.GetValuePattern()),
		}
		kvpValue = append(kvpValue, kvpVal)
	}
	return kvpValue
}

func FlattenMethods(x []ves_io_schema.HttpMethod) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenExpression(x []string) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenIpThreatCategories(x []ves_io_schema_policy.IPThreatCategory) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenAsNumbers(x []uint32) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenIpPrefixes(x []string) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenClasses(x []ves_io_schema_policy.KnownTlsFingerprintClass) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenExactValues(x []string) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenExcludedvalues(x []string) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		rslt = append(rslt, val)
	}
	return rslt
}

func ConvertSliceStringToInterface(x []string) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenTransformers(x []ves_io_schema_policy.Transformer) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenItem(x *ves_io_schema_policy.MatcherType) []interface{} {
	itemValue := make([]interface{}, 0)
	if x != nil {
		itemVal := map[string]interface{}{
			"exact_values": x.GetExactValues(),
			"regex_values": x.GetRegexValues(),
			"transformers": FlattenTransformers(x.GetTransformers()),
		}
		itemValue = append(itemValue, itemVal)
	}
	return itemValue
}

func FlattenCookieMatchers(x []*ves_io_schema_policy.CookieMatcherType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"invert_matcher":    val.GetInvertMatcher(),
			"check_not_present": isEmpty(val.GetCheckNotPresent()),
			"check_present":     isEmpty(val.GetCheckPresent()),
			"item":              FlattenItem(val.GetItem()),
			"name":              val.GetName(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenHeaders(x []*ves_io_schema_policy.HeaderMatcherType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"invert_matcher":    val.GetInvertMatcher(),
			"check_not_present": isEmpty(val.GetCheckNotPresent()),
			"check_present":     isEmpty(val.GetCheckPresent()),
			"item":              FlattenItem(val.GetItem()),
			"name":              val.GetName(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenQueryParams(x []*ves_io_schema_policy.QueryParameterMatcherType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"invert_matcher":    val.GetInvertMatcher(),
			"key":               val.GetKey(),
			"check_not_present": isEmpty(val.GetCheckNotPresent()),
			"check_present":     isEmpty(val.GetCheckPresent()),
			"item":              FlattenItem(val.GetItem()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenApiEndpointRules(x []*ves_io_schema_views_common_waf.APIEndpointProtectionRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"action":              FlattenAction(val.GetAction()),
			"api_endpoint_method": FlattenHttpMethod(val.GetApiEndpointMethod()),
			"api_endpoint_path":   val.GetApiEndpointPath(),
			"client_matcher":      FlattenClientMatcher(val.GetClientMatcher()),
			"any_domain":          isEmpty(val.GetAnyDomain()),
			"specific_domain":     val.GetSpecificDomain(),
			"metadata":            FlattenMetadata(val.GetMetadata()),
			"request_matcher":     FlattenRequestMatcher(val.GetRequestMatcher()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenJwtClaims(x []*ves_io_schema_policy.JWTClaimMatcherType) []interface{} {
	jwtClaimsValue := make([]interface{}, 0)
	for _, v := range x {
		jwtClaimsVal := map[string]interface{}{
			"invert_matcher":    v.GetInvertMatcher(),
			"check_not_present": isEmpty(v.GetCheckNotPresent()),
			"check_present":     isEmpty(v.GetCheckPresent()),
			"item":              FlattenItem(v.GetItem()),
			"name":              v.GetName(),
		}
		jwtClaimsValue = append(jwtClaimsValue, jwtClaimsVal)
	}
	return jwtClaimsValue
}

func FlattenRequestMatcher(x *ves_io_schema_policy.RequestMatcher) []interface{} {
	rmValue := make([]interface{}, 0)
	if x != nil {
		rmVal := map[string]interface{}{
			"cookie_matchers": FlattenCookieMatchers(x.GetCookieMatchers()),
			"headers":         FlattenHeaders(x.GetHeaders()),
			"jwt_claims":      FlattenJwtClaims(x.GetJwtClaims()),
			"query_params":    FlattenQueryParams(x.GetQueryParams()),
		}
		rmValue = append(rmValue, rmVal)
	}
	return rmValue
}

func FlattenClientMatcher(x *ves_io_schema_policy.ClientMatcher) []interface{} {
	cmValue := make([]interface{}, 0)
	if x != nil {
		cmVal := map[string]interface{}{
			"any_client":              isEmpty(x.GetAnyClient()),
			"client_selector":         FlattenServerSelector(x.GetClientSelector()),
			"ip_threat_category_list": FlattenIpThreatCategoryList(x.GetIpThreatCategoryList()),
			"any_ip":                  isEmpty(x.GetAnyIp()),
			"asn_list":                FlattenAsnList(x.GetAsnList()),
			"asn_matcher":             FlattenAsnMatcher(x.GetAsnMatcher()),
			"ip_matcher":              FlattenIpMatcher(x.GetIpMatcher()),
			"ip_prefix_list":          FlattenIpPrefixList(x.GetIpPrefixList()),
			"tls_fingerprint_matcher": FlattenTlsFingerPrintMatcher(x.GetTlsFingerprintMatcher()),
		}
		cmValue = append(cmValue, cmVal)
	}
	return cmValue
}

func FlattenTlsFingerPrintMatcher(x *ves_io_schema_policy.TlsFingerprintMatcherType) []interface{} {
	tlsFMValue := make([]interface{}, 0)
	if x != nil {
		tlsFMVal := map[string]interface{}{
			"classes":         FlattenClasses(x.GetClasses()),
			"exact_values":    x.GetExactValues(),
			"excluded_values": x.GetExcludedValues(),
		}
		tlsFMValue = append(tlsFMValue, tlsFMVal)
	}
	return tlsFMValue
}

func FlattenIpPrefixList(x *ves_io_schema_policy.PrefixMatchList) []interface{} {
	ipPLValue := make([]interface{}, 0)
	if x != nil {
		ipPLVal := map[string]interface{}{
			"invert_match":  x.GetInvertMatch(),
			"ip_prefixes":   x.GetIpPrefixes(),
			"ipv6_prefixes": x.GetIpv6Prefixes(),
		}
		ipPLValue = append(ipPLValue, ipPLVal)
	}
	return ipPLValue
}

func FlattenIpMatcher(x *ves_io_schema_policy.IpMatcherType) []interface{} {
	ipMValue := make([]interface{}, 0)
	if x != nil {
		ipMVal := map[string]interface{}{
			"invert_matcher": x.GetInvertMatcher(),
			"prefix_sets":    FlattenObjectRefTypeList(x.GetPrefixSets()),
		}
		ipMValue = append(ipMValue, ipMVal)
	}
	return ipMValue
}

func FlattenAsnMatcher(x *ves_io_schema_policy.AsnMatcherType) []interface{} {
	asnMValue := make([]interface{}, 0)
	if x != nil {
		asnMVal := map[string]interface{}{
			"asn_sets": FlattenObjectRefTypeList(x.GetAsnSets()),
		}
		asnMValue = append(asnMValue, asnMVal)
	}
	return asnMValue
}

func FlattenAsnList(x *ves_io_schema_policy.AsnMatchList) []interface{} {
	asnLValue := make([]interface{}, 0)
	if x != nil {
		asnLVal := map[string]interface{}{
			"as_numbers": x.GetAsNumbers(),
		}
		asnLValue = append(asnLValue, asnLVal)
	}
	return asnLValue
}

func FlattenIpThreatCategoryList(x *ves_io_schema_policy.IPThreatCategoryListType) []interface{} {
	ipTCLValue := make([]interface{}, 0)
	if x != nil {
		ipTCLVal := map[string]interface{}{
			"ip_threat_categories": FlattenIpThreatCategories(x.GetIpThreatCategories()),
		}
		ipTCLValue = append(ipTCLValue, ipTCLVal)
	}
	return ipTCLValue
}

func FlattenAction(x *ves_io_schema_views_common_waf.APIProtectionRuleAction) []interface{} {
	actionValue := make([]interface{}, 0)
	if x != nil {
		actionVal := map[string]interface{}{
			"allow": isEmpty(x.GetAllow()),
			"deny":  x.GetDeny() != nil,
		}
		actionValue = append(actionValue, actionVal)
	}
	return actionValue
}

func FlattenFAPGIpThreatCategories(x []ves_io_schema_policy.IPThreatCategory) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func ConvertSliceUint32ToInterface(x []uint32) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenApiGroupRules(x []*ves_io_schema_views_common_waf.APIGroupProtectionRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"action":          FlattenAction(val.GetAction()),
			"api_group":       val.GetApiGroup(),
			"base_path":       val.GetBasePath(),
			"client_matcher":  FlattenClientMatcher(val.GetClientMatcher()),
			"any_domain":      isEmpty(val.GetAnyDomain()),
			"specific_domain": val.GetSpecificDomain(),
			"metadata":        FlattenMetadata(val.GetMetadata()),
			"request_matcher": FlattenRequestMatcher(val.GetRequestMatcher()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenApiProtectionRules(x *ves_io_schema_views_common_waf.APIProtectionRules) []interface{} {
	rslt := make([]interface{}, 0)

	if x != nil {
		mapValue := map[string]interface{}{
			"api_endpoint_rules": FlattenApiEndpointRules(x.GetApiEndpointRules()),
			"api_groups_rules":   FlattenApiGroupRules(x.GetApiGroupsRules()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenActions(x []ves_io_schema_views_common_waf.ClientSrcRuleAction) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenHTTPHeaders(x []*ves_io_schema.HeaderMatcherType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"invert_match": val.GetInvertMatch(),
			"name":         val.GetName(),
			"exact":        val.GetExact(),
			"presence":     val.GetPresence(),
			"regex":        val.GetRegex(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenHTTPHeader(x *ves_io_schema_views_common_waf.HttpHeaderMatcherList) []interface{} {
	httpHValue := make([]interface{}, 0)
	if x != nil {
		httpHVal := map[string]interface{}{
			"headers": FlattenHTTPHeaders(x.GetHeaders()),
		}
		httpHValue = append(httpHValue, httpHVal)
	}
	return httpHValue
}

func FlattenBlockedClients(x []*ves_io_schema_views_common_waf.SimpleClientSrcRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"actions":              FlattenActions(val.GetActions()),
			"as_number":            val.GetAsNumber(),
			"http_header":          FlattenHTTPHeader(val.GetHttpHeader()),
			"ip_prefix":            val.GetIpPrefix(),
			"ipv6_prefix":          val.GetIpv6Prefix(),
			"user_identifier":      val.GetUserIdentifier(),
			"expiration_timestamp": FlattenExpirationTimestamp(val.GetExpirationTimestamp()),
			"metadata":             FlattenMetadata(val.GetMetadata()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenDomain(x *ves_io_schema.DomainType) []interface{} {
	domainValue := make([]interface{}, 0)
	if x != nil {
		domainVal := map[string]interface{}{
			"exact_value":  x.GetExactValue(),
			"regex_value":  x.GetRegexValue(),
			"suffix_value": x.GetSuffixValue(),
		}
		domainValue = append(domainValue, domainVal)
	}
	return domainValue
}

func FlattenPath(x *ves_io_schema.PathMatcherType) []interface{} {
	pathValue := make([]interface{}, 0)
	if x != nil {
		pathVal := map[string]interface{}{
			"path":   x.GetPath(),
			"prefix": x.GetPrefix(),
			"regex":  x.GetRegex(),
		}
		pathValue = append(pathValue, pathVal)
	}
	return pathValue
}

func FlattenExcludeList(x []*ves_io_schema_views_common_security.ShapeJavaScriptExclusionRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"any_domain": isEmpty(val.GetAnyDomain()),
			"domain":     FlattenDomain(val.GetDomain()),
			"metadata":   FlattenMetadata(val.GetMetadata()),
			"path":       FlattenPath(val.GetPath()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenRules(x []*ves_io_schema_views_common_security.ShapeJavaScriptInsertionRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"any_domain":          val.GetAnyDomain() != nil,
			"domain":              FlattenDomain(val.GetDomain()),
			"javascript_location": val.GetJavascriptLocation().String(),
			"metadata":            FlattenMetadata(val.GetMetadata()),
			"path":                FlattenPath(val.GetPath()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenMIHeaders(x []*ves_io_schema_policy.HeaderMatcherTypeBasic) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"check_not_present": isEmpty(val.GetCheckNotPresent()),
			"check_present":     isEmpty(val.GetCheckPresent()),
			"item":              FlattenItem(val.GetItem()),
			"name":              val.GetName(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenFailureConditions(x []*ves_io_schema.BotDefenseTransactionResultCondition) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"name":         val.GetName(),
			"regex_values": val.GetRegexValues(),
			"status":       val.GetStatus().String(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenHttpMethods(x []ves_io_schema.BotHttpMethod) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenProtectedAppEndpoints(x []*ves_io_schema_views_common_security.AppEndpointType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"mobile":               isEmpty(val.GetMobile()),
			"web":                  isEmpty(val.GetWeb()),
			"web_mobile":           FlattenWebMobile(val.GetWebMobile()),
			"any_domain":           isEmpty(val.GetAnyDomain()),
			"domain":               FlattenDomain(val.GetDomain()),
			"flow_label":           FlattenFlowLabel(val.GetFlowLabel()),
			"undefined_flow_label": isEmpty(val.GetUndefinedFlowLabel()),
			"allow_good_bots":      isEmpty(val.GetAllowGoodBots()),
			"mitigate_good_bots":   isEmpty(val.GetMitigateGoodBots()),
			"headers":              FlattenHeaders(val.GetHeaders()),
			"http_methods":         FlattenHttpMethods(val.GetHttpMethods()),
			"metadata":             FlattenMetadata(val.GetMetadata()),
			"mitigation":           FlattenMitigation(val.GetMitigation()),
			"path":                 FlattenPath(val.GetPath()),
			"protocol":             val.GetProtocol().String(),
			"query_params":         FlattenQueryParams(val.GetQueryParams()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenMitigation(x *ves_io_schema_policy.ShapeBotMitigationAction) []interface{} {
	mValue := make([]interface{}, 0)
	if x != nil {
		mVal := map[string]interface{}{
			"block":    FlattenBlock(x.GetBlock()),
			"flag":     FlattenFlag(x.GetFlag()),
			"redirect": FlattenRedirect(x.GetRedirect()),
		}
		mValue = append(mValue, mVal)
	}
	return mValue
}

func FlattenRedirect(x *ves_io_schema_policy.ShapeBotRedirectMitigationActionType) []interface{} {
	rValue := make([]interface{}, 0)
	if x != nil {
		rVal := map[string]interface{}{
			"uri": x.GetUri(),
		}
		rValue = append(rValue, rVal)
	}
	return rValue
}

func FlattenFlag(x *ves_io_schema_policy.ShapeBotFlagMitigationActionChoiceType) []interface{} {
	flagValue := make([]interface{}, 0)
	if x != nil {
		flagVal := map[string]interface{}{
			"append_headers": FlattenAppendHeaders(x.GetAppendHeaders()),
			"no_headers":     isEmpty(x.GetNoHeaders()),
		}
		flagValue = append(flagValue, flagVal)
	}
	return flagValue
}

func FlattenAppendHeaders(x *ves_io_schema_policy.ShapeBotFlagMitigationActionType) []interface{} {
	ahValue := make([]interface{}, 0)
	if x != nil {
		ahVal := map[string]interface{}{
			"auto_type_header_name": x.GetAutoTypeHeaderName(),
			"inference_header_name": x.GetInferenceHeaderName(),
		}
		ahValue = append(ahValue, ahVal)
	}
	return ahValue
}

func FlattenBlock(x *ves_io_schema_policy.ShapeBotBlockMitigationActionType) []interface{} {
	bValue := make([]interface{}, 0)
	if x != nil {
		bVal := map[string]interface{}{
			"body":   x.GetBody(),
			"status": x.GetStatus().String(),
		}
		bValue = append(bValue, bVal)
	}
	return bValue
}

func FlattenFlowLabel(x *ves_io_schema.BotDefenseFlowLabelCategoriesChoiceType) []interface{} {
	flValue := make([]interface{}, 0)
	if x != nil {
		flVal := map[string]interface{}{
			"account_management":  FlattenAccountManagement(x.GetAccountManagement()),
			"authentication":      FlattenAuthentication(x.GetAuthentication()),
			"financial_services":  FlattenFinancialServices(x.GetFinancialServices()),
			"flight":              FlattenFlight(x.GetFlight()),
			"profile_management":  FlattenProfileManagement(x.GetProfileManagement()),
			"search":              FlattenSearch(x.GetSearch()),
			"shopping_gift_cards": FlattenShoppingGiftCards(x.GetShoppingGiftCards()),
		}
		flValue = append(flValue, flVal)
	}
	return flValue
}

func FlattenShoppingGiftCards(x *ves_io_schema.BotDefenseFlowLabelShoppingGiftCardsChoiceType) []interface{} {
	sgcValue := make([]interface{}, 0)
	if x != nil {
		sgcVal := map[string]interface{}{
			"gift_card_make_purchase_with_gift_card": isEmpty(x.GetGiftCardMakePurchaseWithGiftCard()),
			"gift_card_validation":                   isEmpty(x.GetGiftCardValidation()),
			"shop_add_to_cart":                       isEmpty(x.GetShopAddToCart()),
			"shop_checkout":                          isEmpty(x.GetShopCheckout()),
			"shop_choose_seat":                       isEmpty(x.GetShopChooseSeat()),
			"shop_enter_drawing_submission":          isEmpty(x.GetShopEnterDrawingSubmission()),
			"shop_make_payment":                      isEmpty(x.GetShopMakePayment()),
			"shop_order":                             isEmpty(x.GetShopOrder()),
			"shop_price_inquiry":                     isEmpty(x.GetShopPriceInquiry()),
			"shop_promo_code_validation":             isEmpty(x.GetShopPromoCodeValidation()),
			"shop_purchase_gift_card":                isEmpty(x.GetShopPurchaseGiftCard()),
			"shop_update_quantity":                   isEmpty(x.GetShopUpdateQuantity()),
		}
		sgcValue = append(sgcValue, sgcVal)
	}
	return sgcValue
}

func FlattenSearch(x *ves_io_schema.BotDefenseFlowLabelSearchChoiceType) []interface{} {
	sValue := make([]interface{}, 0)
	if x != nil {
		sVal := map[string]interface{}{
			"flight_search":      isEmpty(x.GetFlightSearch()),
			"product_search":     isEmpty(x.GetProductSearch()),
			"reservation_search": isEmpty(x.GetReservationSearch()),
			"room_search":        isEmpty(x.GetRoomSearch()),
		}
		sValue = append(sValue, sVal)
	}
	return sValue
}

func FlattenProfileManagement(x *ves_io_schema.BotDefenseFlowLabelProfileManagementChoiceType) []interface{} {
	pmValue := make([]interface{}, 0)
	if x != nil {
		pmVal := map[string]interface{}{
			"create": isEmpty(x.GetCreate()),
			"update": isEmpty(x.GetUpdate()),
			"view":   isEmpty(x.GetView()),
		}
		pmValue = append(pmValue, pmVal)
	}
	return pmValue
}

func FlattenFlight(x *ves_io_schema.BotDefenseFlowLabelFlightChoiceType) []interface{} {
	fValue := make([]interface{}, 0)
	if x != nil {
		fVal := map[string]interface{}{
			"checkin": isEmpty(x.GetCheckin()),
		}
		fValue = append(fValue, fVal)
	}
	return fValue
}

func FlattenFinancialServices(x *ves_io_schema.BotDefenseFlowLabelFinancialServicesChoiceType) []interface{} {
	fsValue := make([]interface{}, 0)
	if x != nil {
		fsVal := map[string]interface{}{
			"apply":          isEmpty(x.GetApply()),
			"money_transfer": isEmpty(x.GetMoneyTransfer()),
		}
		fsValue = append(fsValue, fsVal)
	}
	return fsValue
}

func FlattenAuthentication(x *ves_io_schema.BotDefenseFlowLabelAuthenticationChoiceType) []interface{} {
	authValue := make([]interface{}, 0)
	if x != nil {
		authVal := map[string]interface{}{
			"login_mfa":     isEmpty(x.GetLoginMfa()),
			"login_partner": isEmpty(x.GetLoginPartner()),
			"logout":        isEmpty(x.GetLogout()),
			"token_refresh": isEmpty(x.GetTokenRefresh()),
		}
		authValue = append(authValue, authVal)
	}
	return authValue
}

func FlattenLogin(x *ves_io_schema.BotDefenseTransactionResult) []interface{} {
	lValue := make([]interface{}, 0)
	if x != nil {
		lVal := map[string]interface{}{
			"disable_transaction_result": isEmpty(x.GetDisableTransactionResult()),
			"transaction_result":         FlattenTransactionResult(x.GetTransactionResult()),
		}
		lValue = append(lValue, lVal)
	}
	return lValue
}

func FlattenTransactionResult(x *ves_io_schema.BotDefenseTransactionResultType) []interface{} {
	trValue := make([]interface{}, 0)
	if x != nil {
		trVal := map[string]interface{}{
			"failure_conditions": FlattenFailureConditions(x.GetFailureConditions()),
			"success_conditions": FlattenFailureConditions(x.GetSuccessConditions()),
		}
		trValue = append(trValue, trVal)
	}
	return trValue
}

func FlattenAccountManagement(x *ves_io_schema.BotDefenseFlowLabelAccountManagementChoiceType) []interface{} {
	amValue := make([]interface{}, 0)
	if x != nil {
		amVal := map[string]interface{}{
			"create":         isEmpty(x.GetCreate()),
			"password_reset": isEmpty(x.GetPasswordReset()),
		}
		amValue = append(amValue, amVal)
	}
	return amValue
}

func FlattenWebMobile(x *ves_io_schema_views_common_security.WebMobileTrafficType) []interface{} {
	wmValue := make([]interface{}, 0)
	if x != nil {
		wmVal := map[string]interface{}{
			"mobile_identifier": x.GetMobileIdentifier().String(),
		}
		wmValue = append(wmValue, wmVal)
	}
	return wmValue
}

func FlattenHeader(x *ves_io_schema_policy.HeaderMatcherTypeBasic) []interface{} {
	hValue := make([]interface{}, 0)
	if x != nil {
		hVal := map[string]interface{}{
			"check_not_present": isEmpty(x.GetCheckNotPresent()),
			"check_present":     isEmpty(x.GetCheckPresent()),
			"item":              FlattenItem(x.GetItem()),
			"name":              x.GetName(),
		}
		hValue = append(hValue, hVal)
	}
	return hValue
}

func FlattenBotDefense(x *ves_io_schema_views_common_security.ShapeBotDefenseType) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"policy":            FlattenPolicy(x.GetPolicy()),
			"regional_endpoint": x.GetRegionalEndpoint().String(),
			"timeout":           x.GetTimeout(),
		}
		val = append(val, test)
	}
	return val
}

func FlattenPolicy(x *ves_io_schema_views_common_security.ShapeBotDefensePolicyType) []interface{} {
	pValue := make([]interface{}, 0)
	if x != nil {
		pVal := map[string]interface{}{
			"disable_js_insert":          isEmpty(x.GetDisableJsInsert()),
			"js_insert_all_pages":        FlattenJsInsertAllPages(x.GetJsInsertAllPages()),
			"js_insert_all_pages_except": FlattenJsInsertAllPagesExcept(x.GetJsInsertAllPagesExcept()),
			"js_insertion_rules":         FlattenJsInsertionRules(x.GetJsInsertionRules()),
			"javascript_mode":            x.GetJavascriptMode().String(),
			"js_download_path":           x.GetJsDownloadPath(),
			"disable_mobile_sdk":         isEmpty(x.GetDisableMobileSdk()),
			"mobile_sdk_config":          FlattenMobileSdkConfig(x.GetMobileSdkConfig()),
			"protected_app_endpoints":    FlattenProtectedAppEndpoints(x.GetProtectedAppEndpoints()),
		}
		pValue = append(pValue, pVal)
	}
	return pValue
}

func FlattenMobileSdkConfig(x *ves_io_schema_views_common_security.MobileSDKConfigType) []interface{} {
	mscValue := make([]interface{}, 0)
	if x != nil {
		mscVal := map[string]interface{}{
			"mobile_identifier": FlattenMobileIdentifier(x.GetMobileIdentifier()),
		}
		mscValue = append(mscValue, mscVal)
	}
	return mscValue
}

func FlattenMobileIdentifier(x *ves_io_schema_views_common_security.MobileTrafficIdentifierType) []interface{} {
	miValue := make([]interface{}, 0)
	if x != nil {
		miVal := map[string]interface{}{
			"headers": FlattenMIHeaders(x.GetHeaders()),
		}
		miValue = append(miValue, miVal)
	}
	return miValue
}

func FlattenJsInsertionRules(x *ves_io_schema_views_common_security.ShapeJavaScriptInsertType) []interface{} {
	jirValue := make([]interface{}, 0)
	if x != nil {
		jirVal := map[string]interface{}{
			"exclude_list": FlattenExcludeList(x.GetExcludeList()),
			"rules":        FlattenRules(x.GetRules()),
		}
		jirValue = append(jirValue, jirVal)
	}
	return jirValue
}

func FlattenJsInsertAllPagesExcept(x *ves_io_schema_views_common_security.ShapeJavaScriptInsertAllWithExceptionsType) []interface{} {
	jiapeValue := make([]interface{}, 0)
	if x != nil {
		jiapeVal := map[string]interface{}{
			"exclude_list":        FlattenExcludeList(x.GetExcludeList()),
			"javascript_location": x.GetJavascriptLocation().String(),
		}
		jiapeValue = append(jiapeValue, jiapeVal)
	}
	return jiapeValue
}

func FlattenJsInsertAllPages(x *ves_io_schema_views_common_security.ShapeJavaScriptInsertAllType) []interface{} {
	jiapValue := make([]interface{}, 0)
	if x != nil {
		jiapVal := map[string]interface{}{
			"javascript_location": x.GetJavascriptLocation().String(),
		}
		jiapValue = append(jiapValue, jiapVal)
	}
	return jiapValue
}

func FlattenCaptchaChallenge(x *ves_io_schema_virtual_host.CaptchaChallengeType) []interface{} {
	rslt := make([]interface{}, 0)

	if x != nil {
		mapValue := map[string]interface{}{
			"cookie_expiry": x.GetCookieExpiry(),
			"custom_page":   x.GetCustomPage(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenEnableChallenge(x *ves_io_schema_views_common_waf.EnableChallenge) []interface{} {
	rslt := make([]interface{}, 0)

	if x != nil {
		mapValue := map[string]interface{}{
			"captcha_challenge_parameters":         FlattenCaptchaChallenge(x.GetCaptchaChallengeParameters()),
			"default_captcha_challenge_parameters": isEmpty(x.GetDefaultCaptchaChallengeParameters()),
			"default_js_challenge_parameters":      isEmpty(x.GetDefaultJsChallengeParameters()),
			"js_challenge_parameters":              FlattenJSChallenge(x.GetJsChallengeParameters()),
			"default_mitigation_settings":          isEmpty(x.GetDefaultMitigationSettings()),
			"malicious_user_mitigation":            FlattenObjectRefTypeSet(x.GetMaliciousUserMitigation()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenJSChallenge(x *ves_io_schema_virtual_host.JavascriptChallengeType) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"cookie_expiry":   x.GetCookieExpiry(),
			"custom_page":     x.GetCustomPage(),
			"js_script_delay": x.GetJsScriptDelay(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenArgMatchers(x []*ves_io_schema_policy.ArgMatcherType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"invert_matcher":    val.GetInvertMatcher(),
			"check_not_present": isEmpty(val.GetCheckNotPresent()),
			"check_present":     isEmpty(val.GetCheckPresent()),
			"item":              FlattenItem(val.GetItem()),
			"presence":          val.GetPresence(),
			"name":              val.GetName(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenPBCRules(x []*ves_io_schema_views_common_waf.ChallengeRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"metadata": FlattenMetadata(val.GetMetadata()),
			"spec":     FlattenSpec(val.GetSpec()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenSpec(x *ves_io_schema_service_policy_rule.ChallengeRuleSpec) []interface{} {
	sValue := make([]interface{}, 0)
	if x != nil {
		sVal := map[string]interface{}{
			"arg_matchers":                FlattenArgMatchers(x.GetArgMatchers()),
			"any_asn":                     isEmpty(x.GetAnyAsn()),
			"asn_list":                    FlattenAsnList(x.GetAsnList()),
			"asn_matcher":                 FlattenAsnMatcher(x.GetAsnMatcher()),
			"body_matcher":                FlattenItem(x.GetBodyMatcher()),
			"disable_challenge":           isEmpty(x.GetDisableChallenge()),
			"enable_captcha_challenge":    isEmpty(x.GetEnableCaptchaChallenge()),
			"enable_javascript_challenge": isEmpty(x.GetEnableJavascriptChallenge()),
			"any_client":                  isEmpty(x.GetAnyClient()),
			"client_selector":             FlattenServerSelector(x.GetClientSelector()),
			"cookie_matchers":             FlattenCookieMatchers(x.GetCookieMatchers()),
			"domain_matcher":              FlattenDomainMatcher(x.GetDomainMatcher()),
			"expiration_timestamp":        FlattenExpirationTimestamp(x.GetExpirationTimestamp()),
			"headers":                     FlattenHeaders(x.GetHeaders()),
			"http_method":                 FlattenHttpMethod(x.GetHttpMethod()),
			"any_ip":                      isEmpty(x.GetAnyIp()),
			"ip_matcher":                  FlattenIpMatcher(x.GetIpMatcher()),
			"ip_prefix_list":              FlattenIpPrefixList(x.GetIpPrefixList()),
			"path":                        FlattenSpecPath(x.GetPath()),
			"query_params":                FlattenQueryParams(x.GetQueryParams()),
			"tls_fingerprint_matcher":     FlattenTlsFingerPrintMatcher(x.GetTlsFingerprintMatcher()),
		}
		sValue = append(sValue, sVal)
	}
	return sValue
}

func FlattenJa4TlsFingerprint(x *ves_io_schema_policy.JA4TlsFingerprintMatcherType) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"exact_values": x.GetExactValues(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenSpecPath(x *ves_io_schema_policy.PathMatcherType) []interface{} {
	pathValue := make([]interface{}, 0)
	if x != nil {
		pathVal := map[string]interface{}{
			"exact_values":   x.GetExactValues(),
			"invert_matcher": x.GetInvertMatcher(),
			"prefix_values":  x.GetPrefixValues(),
			"regex_values":   x.GetRegexValues(),
			"suffix_values":  x.GetSuffixValues(),
			"transformers":   FlattenTransformers(x.GetTransformers()),
		}
		pathValue = append(pathValue, pathVal)
	}
	return pathValue
}

func FlattenHttpMethod(x *ves_io_schema_policy.HttpMethodMatcherType) []interface{} {
	httpmValue := make([]interface{}, 0)
	if x != nil {
		httpmVal := map[string]interface{}{
			"invert_matcher": x.GetInvertMatcher(),
			"methods":        FlattenMethods(x.GetMethods()),
		}
		httpmValue = append(httpmValue, httpmVal)
	}
	return httpmValue
}

func FlattenDomainMatcher(x *ves_io_schema_policy.MatcherTypeBasic) []interface{} {
	domainValue := make([]interface{}, 0)
	if x != nil {
		domainVal := map[string]interface{}{
			"exact_values": x.GetExactValues(),
			"regex_values": x.GetRegexValues(),
		}
		domainValue = append(domainValue, domainVal)
	}
	return domainValue
}

func FlattenPolicyBasedChallenge(x *ves_io_schema_views_common_waf.PolicyBasedChallenge) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"captcha_challenge_parameters":         FlattenCaptchaChallenge(x.GetCaptchaChallengeParameters()),
			"default_captcha_challenge_parameters": isEmpty(x.GetDefaultCaptchaChallengeParameters()),
			"always_enable_captcha_challenge":      isEmpty(x.GetAlwaysEnableCaptchaChallenge()),
			"always_enable_js_challenge":           isEmpty(x.GetAlwaysEnableJsChallenge()),
			"no_challenge":                         isEmpty(x.GetNoChallenge()),
			"default_js_challenge_parameters":      isEmpty(x.GetDefaultJsChallengeParameters()),
			"js_challenge_parameters":              FlattenJSChallenge(x.GetJsChallengeParameters()),
			"default_mitigation_settings":          isEmpty(x.GetDefaultMitigationSettings()),
			"malicious_user_mitigation":            FlattenObjectRefTypeSet(x.GetMaliciousUserMitigation()),
			"rule_list":                            FlattenRuleList(x.GetRuleList()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenTemporaryUserBlocking(x *ves_io_schema_virtual_host.TemporaryUserBlockingType) []interface{} {
	tubValue := make([]interface{}, 0)
	if x != nil {
		tubVal := map[string]interface{}{
			"custom_page": x.GetCustomPage(),
		}
		tubValue = append(tubValue, tubVal)
	}
	return tubValue
}

func FlattenRuleList(x *ves_io_schema_views_common_waf.ChallengeRuleList) []interface{} {
	rlValue := make([]interface{}, 0)
	if x != nil {
		rlVal := map[string]interface{}{
			"rules": FlattenPBCRules(x.GetRules()),
		}
		rlValue = append(rlValue, rlVal)
	}
	return rlValue
}

func FlattenCSDRules(x []*ves_io_schema_views_common_security.CSDJavaScriptInsertionRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"any_domain": val.GetAnyDomain() != nil,
			"domain":     FlattenDomain(val.GetDomain()),
			"metadata":   FlattenMetadata(val.GetMetadata()),
			"path":       FlattenPath(val.GetPath()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenClientSideDefense(x *ves_io_schema_views_common_security.ClientSideDefenseType) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"policy": FlattenCSDPolicy(x.GetPolicy()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenCSDPolicy(x *ves_io_schema_views_common_security.ClientSideDefensePolicyType) []interface{} {
	pValue := make([]interface{}, 0)
	if x != nil {
		pVal := map[string]interface{}{
			"disable_js_insert":          isEmpty(x.GetDisableJsInsert()),
			"js_insert_all_pages":        isEmpty(x.GetJsInsertAllPages()),
			"js_insert_all_pages_except": FlattenCSDJsInsertAllPagesExcept(x.GetJsInsertAllPagesExcept()),
			"js_insertion_rules":         FlattenCSDJsInsertionRules(x.GetJsInsertionRules()),
		}
		pValue = append(pValue, pVal)
	}
	return pValue
}

func FlattenCSDJsInsertionRules(x *ves_io_schema_views_common_security.CSDJavaScriptInsertType) []interface{} {
	jirValue := make([]interface{}, 0)
	if x != nil {
		jirVal := map[string]interface{}{
			"exclude_list": FlattenExcludeList(x.GetExcludeList()),
			"rules":        FlattenCSDRules(x.GetRules()),
		}
		jirValue = append(jirValue, jirVal)
	}
	return jirValue
}

func FlattenCSDJsInsertAllPagesExcept(x *ves_io_schema_views_common_security.CSDJavaScriptInsertAllWithExceptionsType) []interface{} {
	jiapeValue := make([]interface{}, 0)
	if x != nil {
		jiapeVal := map[string]interface{}{
			"exclude_list": FlattenExcludeList(x.GetExcludeList()),
		}
		jiapeValue = append(jiapeValue, jiapeVal)
	}
	return jiapeValue
}

func FlattenCorsPolicy(x *ves_io_schema.CorsPolicy) []interface{} {
	rslt := make([]interface{}, 0)

	if x != nil {
		mapValue := map[string]interface{}{
			"allow_credentials":  x.GetAllowCredentials(),
			"allow_headers":      x.GetAllowHeaders(),
			"allow_methods":      x.GetAllowMethods(),
			"allow_origin":       ConvertSliceStringToInterface(x.GetAllowOrigin()),
			"allow_origin_regex": ConvertSliceStringToInterface(x.GetAllowOriginRegex()),
			"disabled":           x.GetDisabled(),
			"expose_headers":     x.GetExposeHeaders(),
			"maximum_age":        x.GetMaximumAge(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenCsrfPolicy(x *ves_io_schema.CsrfPolicy) []interface{} {
	rslt := make([]interface{}, 0)

	if x != nil {
		mapValue := map[string]interface{}{
			"all_load_balancer_domains": x.GetAllLoadBalancerDomains() != nil,
			"custom_domain_list":        FlattenCustomDomainList(x.GetCustomDomainList()),
			"disabled":                  isEmpty(x.GetDisabled()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenCustomDomainList(x *ves_io_schema.DomainNameList) []interface{} {
	cdlValue := make([]interface{}, 0)
	if x != nil {
		cdlVal := map[string]interface{}{
			"domains": ConvertSliceStringToInterface(x.GetDomains()),
		}
		cdlValue = append(cdlValue, cdlVal)
	}
	return cdlValue
}

func FlattenDataGuardRules(x []*ves_io_schema_policy.SimpleDataGuardRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"apply_data_guard": isEmpty(val.GetApplyDataGuard()),
			"skip_data_guard":  isEmpty(val.GetSkipDataGuard()),
			"any_domain":       isEmpty(val.GetAnyDomain()),
			"exact_value":      val.GetExactValue(),
			"suffix_value":     val.GetSuffixValue(),
			"metadata":         FlattenMetadata(val.GetMetadata()),
			"path":             FlattenPath(val.GetPath()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenCountryList(x []ves_io_schema_policy.CountryCode) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenExpirationTimestamp(x *types.Timestamp) string {
	if x == nil {
		return ""
	}
	return x.String()
}

func FlattenDDOSMitigationRules(x []*ves_io_schema_views_common_security.DDoSMitigationRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"expiration_timestamp": FlattenExpirationTimestamp(val.GetExpirationTimestamp()),
			"metadata":             FlattenMetadata(val.GetMetadata()),
			"block":                isEmpty(val.GetBlock()),
			"ddos_client_source":   FlattenDdosClientSource(val.GetDdosClientSource()),
			"ip_prefix_list":       FlattenIpPrefixList(val.GetIpPrefixList()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenDdosClientSource(x *ves_io_schema_views_common_security.DDoSClientSource) []interface{} {
	dcsValue := make([]interface{}, 0)
	if x != nil {
		dcsVal := map[string]interface{}{
			"asn_list":                    FlattenAsnList(x.GetAsnList()),
			"country_list":                FlattenCountryList(x.GetCountryList()),
			"ja4_tls_fingerprint_matcher": FlattenJa4TlsFingerprint(x.GetJa4TlsFingerprintMatcher()),
			"tls_fingerprint_matcher":     FlattenTlsFingerPrintMatcher(x.GetTlsFingerprintMatcher()),
		}
		dcsValue = append(dcsValue, dcsVal)
	}
	return dcsValue
}

func FlattenEndpointSubsets(x map[string]string) map[string]interface{} {
	rslt := make(map[string]interface{})
	for k, val := range x {
		rslt[k] = val
	}
	return rslt
}

func FlattenDefaultRoutePools(x []*ves_io_schema_views.OriginPoolWithWeight) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		nestedStruct := map[string]interface{}{
			"endpoint_subsets": FlattenEndpointSubsets(val.GetEndpointSubsets()),
			"cluster":          FlattenObjectRefTypeSet(val.GetCluster()),
			"pool":             FlattenObjectRefTypeSet(val.GetPool()),
			"priority":         val.GetPriority(),
			"weight":           val.GetWeight(),
		}
		rslt = append(rslt, nestedStruct)
	}
	return rslt
}

func FlattenGraphRules(x []*ves_io_schema_policy.GraphQLRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"any_domain":       isEmpty(val.GetAnyDomain()),
			"exact_value":      val.GetExactValue(),
			"suffix_value":     val.GetSuffixValue(),
			"exact_path":       val.GetExactPath(),
			"graphql_settings": FlattenGraphqlSettings(val.GetGraphqlSettings()),
			"metadata":         FlattenMetadata(val.GetMetadata()),
			"method_get":       isEmpty(val.GetMethodGet()),
			"method_post":      isEmpty(val.GetMethodPost()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenGraphqlSettings(x *ves_io_schema_policy.GraphQLSettingsType) []interface{} {
	gsValue := make([]interface{}, 0)
	if x != nil {
		gsVal := map[string]interface{}{
			"disable_introspection": isEmpty(x.GetDisableIntrospection()),
			"enable_introspection":  isEmpty(x.GetEnableIntrospection()),
			"max_batched_queries":   x.GetMaxBatchedQueries(),
			"max_depth":             x.GetMaxDepth(),
			"max_total_length":      x.GetMaxTotalLength(),
		}
		gsValue = append(gsValue, gsVal)
	}
	return gsValue
}

func FlattenHashPolicy(x []*ves_io_schema_route.HashPolicyType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"cookie":      FlattenCookie(val.GetCookie()),
			"header_name": val.GetHeaderName(),
			"source_ip":   val.GetSourceIp(),
			"terminal":    val.GetTerminal(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenCookie(x *ves_io_schema_route.CookieForHashing) []interface{} {
	cValue := make([]interface{}, 0)
	if x != nil {
		cVal := map[string]interface{}{
			"add_httponly":    isEmpty(x.GetAddHttponly()),
			"ignore_httponly": isEmpty(x.GetIgnoreHttponly()),
			"name":            x.GetName(),
			"path":            x.GetPath(),
			"ignore_samesite": isEmpty(x.GetIgnoreSamesite()),
			"samesite_lax":    isEmpty(x.GetSamesiteLax()),
			"samesite_none":   isEmpty(x.GetSamesiteNone()),
			"samesite_strict": isEmpty(x.GetSamesiteStrict()),
			"add_secure":      isEmpty(x.GetAddSecure()),
			"ignore_secure":   isEmpty(x.GetIgnoreSecure()),
			"ttl":             x.GetTtl(),
		}
		cValue = append(cValue, cVal)
	}
	return cValue
}

func FlattenEnableIpReputation(x *ves_io_schema_views_common_waf.IPThreatCategoryListType) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"ip_threat_categories": FlattenIpThreatCategories(x.GetIpThreatCategories()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenJwtValidation(x *ves_io_schema_views_common_waf.JWTValidation) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		rs := map[string]interface{}{
			"action":           FlattenJWTAction(x.GetAction()),
			"jwks_config":      FlattenJwksConfig(x.GetJwksConfig()),
			"mandatory_claims": FlattenMandatoryClaims(x.GetMandatoryClaims()),
			"reserved_claims":  FlattenReservedClaims(x.GetReservedClaims()),
			"target":           FlattenTarget(x.GetTarget()),
			"token_location":   FlattenTokenLocation(x.GetTokenLocation()),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenTokenLocation(x *ves_io_schema_views_common_waf.TokenLocation) []interface{} {
	tlValue := make([]interface{}, 0)
	if x != nil {
		tlVal := map[string]interface{}{
			"bearer_token": isEmpty(x.GetBearerToken()),
		}
		tlValue = append(tlValue, tlVal)
	}
	return tlValue
}

func FlattenTarget(x *ves_io_schema_views_common_waf.Target) []interface{} {
	tValue := make([]interface{}, 0)
	if x != nil {
		tVal := map[string]interface{}{
			"all_endpoint": isEmpty(x.GetAllEndpoint()),
			"api_groups":   FlattenApiGroups(x.GetApiGroups()),
			"base_paths":   FlattenBasePaths(x.GetBasePaths()),
		}
		tValue = append(tValue, tVal)
	}
	return tValue
}

func FlattenBasePaths(x *ves_io_schema_views_common_waf.BasePathsType) []interface{} {
	bpValue := make([]interface{}, 0)
	if x != nil {
		bpVal := map[string]interface{}{
			"base_paths": x.GetBasePaths(),
		}
		bpValue = append(bpValue, bpVal)
	}
	return bpValue
}

func FlattenApiGroups(x *ves_io_schema_views_common_waf.APIGroups) []interface{} {
	agValue := make([]interface{}, 0)
	if x != nil {
		agVal := map[string]interface{}{
			"api_groups": x.GetApiGroups(),
		}
		agValue = append(agValue, agVal)
	}
	return agValue
}

func FlattenReservedClaims(x *ves_io_schema_views_common_waf.ReservedClaims) []interface{} {
	rcValue := make([]interface{}, 0)
	if x != nil {
		rcVal := map[string]interface{}{
			"audience":                FlattenAudience(x.GetAudience()),
			"audience_disable":        isEmpty(x.GetAudienceDisable()),
			"issuer":                  x.GetIssuer(),
			"issuer_disable":          isEmpty(x.GetIssuerDisable()),
			"validate_period_disable": isEmpty(x.GetValidatePeriodDisable()),
			"validate_period_enable":  isEmpty(x.GetValidatePeriodEnable()),
		}
		rcValue = append(rcValue, rcVal)
	}
	return rcValue
}

func FlattenAudience(x *ves_io_schema_views_common_waf.Audiences) []interface{} {
	aValue := make([]interface{}, 0)
	if x != nil {
		aVal := map[string]interface{}{
			"audiences": x.GetAudiences(),
		}
		aValue = append(aValue, aVal)
	}
	return aValue
}

func FlattenMandatoryClaims(x *ves_io_schema_views_common_waf.MandatoryClaims) []interface{} {
	mcValue := make([]interface{}, 0)
	if x != nil {
		mcVal := map[string]interface{}{
			"claim_names": x.GetClaimNames(),
		}
		mcValue = append(mcValue, mcVal)
	}
	return mcValue
}

func FlattenJwksConfig(x *ves_io_schema_views_common_waf.JWKS) []interface{} {
	jcValue := make([]interface{}, 0)
	if x != nil {
		jcVal := map[string]interface{}{
			"cleartext": x.GetCleartext(),
		}
		jcValue = append(jcValue, jcVal)
	}
	return jcValue
}

func FlattenJWTAction(x *ves_io_schema.Action) []interface{} {
	aValue := make([]interface{}, 0)
	if x != nil {
		aVal := map[string]interface{}{
			"block":  isEmpty(x.GetBlock()),
			"report": isEmpty(x.GetReport()),
		}
		aValue = append(aValue, aVal)
	}
	return aValue
}

func FlattenHttp(x *ves_io_schema_views_http_loadbalancer.ProxyTypeHttp) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"dns_volterra_managed": x.GetDnsVolterraManaged(),
			"port":                 x.GetPort(),
			"port_ranges":          x.GetPortRanges(),
		}
		val = append(val, test)
	}
	return val
}

func FlattenHashAlgorithms(x []ves_io_schema.HashAlgorithm) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenTlsCertificates(x []*ves_io_schema.TlsCertificateType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"certificate_url":        val.GetCertificateUrl(),
			"description":            val.GetDescription(),
			"custom_hash_algorithms": FlattenCustomHashAlgorithms(val.GetCustomHashAlgorithms()),
			"private_key":            FlattenPrivateKey(val.GetPrivateKey()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenPrivateKey(x *ves_io_schema.SecretType) []interface{} {
	pkValue := make([]interface{}, 0)
	if x != nil {
		pkVal := map[string]interface{}{
			"blindfold_secret_info": FlattenBlindFoldSecretInfo(x.GetBlindfoldSecretInfo()),
			"clear_secret_info":     FlattenClearSecretInfo(x.GetClearSecretInfo()),
		}
		pkValue = append(pkValue, pkVal)
	}
	return pkValue
}

func FlattenWingmanSecretInfo(x *ves_io_schema.WingmanSecretInfoType) []interface{} {
	wsiValue := make([]interface{}, 0)
	if x != nil {
		wsiVal := map[string]interface{}{
			"name": x.GetName(),
		}
		wsiValue = append(wsiValue, wsiVal)
	}
	return wsiValue
}

func FlattenVaultSecretInfo(x *ves_io_schema.VaultSecretInfoType) []interface{} {
	vsiValue := make([]interface{}, 0)
	if x != nil {
		vsiVal := map[string]interface{}{
			"key":             x.GetKey(),
			"location":        x.GetLocation(),
			"provider":        x.GetProvider(),
			"secret_encoding": x.GetSecretEncoding().String(),
			"version":         x.GetVersion(),
		}
		vsiValue = append(vsiValue, vsiVal)
	}
	return vsiValue
}

func FlattenClearSecretInfo(x *ves_io_schema.ClearSecretInfoType) []interface{} {
	csiValue := make([]interface{}, 0)
	if x != nil {
		csiVal := map[string]interface{}{
			"provider": x.GetProvider(),
			"url":      x.GetUrl(),
		}
		csiValue = append(csiValue, csiVal)
	}
	return csiValue
}

func FlattenBlindFoldSecretInfo(x *ves_io_schema.BlindfoldSecretInfoType) []interface{} {
	bsiValue := make([]interface{}, 0)
	if x != nil {
		bsiVal := map[string]interface{}{
			"decryption_provider": x.GetDecryptionProvider(),
			"location":            x.GetLocation(),
			"store_provider":      x.GetStoreProvider(),
		}
		bsiValue = append(bsiValue, bsiVal)
	}
	return bsiValue
}

func FlattenCustomHashAlgorithms(x *ves_io_schema.HashAlgorithms) []interface{} {
	chaValue := make([]interface{}, 0)
	if x != nil {
		chaVal := map[string]interface{}{
			"hash_algorithms": FlattenHashAlgorithms(x.GetHashAlgorithms()),
		}
		chaValue = append(chaValue, chaVal)
	}
	return chaValue
}

func FlattenHttps(x *ves_io_schema_views_http_loadbalancer.ProxyTypeHttps) []interface{} {
	val := make([]interface{}, 0)

	if x != nil {
		test := map[string]interface{}{
			"add_hsts":                 x.GetAddHsts(),
			"coalescing_options":       FlattenCoalescingOptions(x.GetCoalescingOptions()),
			"connection_idle_timeout":  x.GetConnectionIdleTimeout(),
			"default_loadbalancer":     isEmpty(x.GetDefaultLoadbalancer()),
			"non_default_loadbalancer": isEmpty(x.GetNonDefaultLoadbalancer()),
			"http_protocol_options":    FlattenHttpProtocolOptions(x.GetHttpProtocolOptions()),
			"http_redirect":            x.GetHttpRedirect(),
			"disable_path_normalize":   isEmpty(x.GetDisablePathNormalize()),
			"enable_path_normalize":    isEmpty(x.GetEnablePathNormalize()),
			"port":                     x.GetPort(),
			"port_ranges":              x.GetPortRanges(),
			"append_server_name":       x.GetAppendServerName(),
			"default_header":           isEmpty(x.GetDefaultHeader()),
			"pass_through":             isEmpty(x.GetPassThrough()),
			"server_name":              x.GetServerName(),
			"tls_cert_params":          FlattenTlsCertParams(x.GetTlsCertParams()),
			"tls_parameters":           FlattenTlsParameters(x.GetTlsParameters()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenCoalescingOptions(x *ves_io_schema.TLSCoalescingOptions) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"default_coalescing": isEmpty(x.GetDefaultCoalescing()),
			"strict_coalescing":  isEmpty(x.GetStrictCoalescing()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenTlsParameters(x *ves_io_schema_views.DownstreamTlsParamsType) []interface{} {
	tpValue := make([]interface{}, 0)
	if x != nil {
		tpVal := map[string]interface{}{
			"no_mtls":          isEmpty(x.GetNoMtls()),
			"use_mtls":         FlattenUseMtls(x.GetUseMtls()),
			"tls_certificates": FlattenTlsCertificates(x.GetTlsCertificates()),
			"tls_config":       FlattenTlsConfig(x.GetTlsConfig()),
		}
		tpValue = append(tpValue, tpVal)
	}
	return tpValue
}

func FlattenTlsCertParams(x *ves_io_schema_views.DownstreamTLSCertsParams) []interface{} {
	tcpValue := make([]interface{}, 0)
	if x != nil {
		tcpVal := map[string]interface{}{
			"certificates": FlattenVObjectRefTypeList(x.GetCertificates()),
			"no_mtls":      isEmpty(x.GetNoMtls()),
			"use_mtls":     FlattenUseMtls(x.GetUseMtls()),
			"tls_config":   FlattenTlsConfig(x.GetTlsConfig()),
		}
		tcpValue = append(tcpValue, tcpVal)
	}
	return tcpValue
}

func FlattenTlsConfig(x *ves_io_schema_views.TlsConfig) []interface{} {
	tcValue := make([]interface{}, 0)
	if x != nil {
		tcVal := map[string]interface{}{
			"custom_security":  FlattenCustomSecurity(x.GetCustomSecurity()),
			"default_security": isEmpty(x.GetDefaultSecurity()),
			"low_security":     isEmpty(x.GetLowSecurity()),
			"medium_security":  isEmpty(x.GetMediumSecurity()),
		}
		tcValue = append(tcValue, tcVal)
	}
	return tcValue
}

func FlattenCustomSecurity(x *ves_io_schema_views.CustomCiphers) []interface{} {
	csValue := make([]interface{}, 0)
	if x != nil {
		csVal := map[string]interface{}{
			"cipher_suites": x.GetCipherSuites(),
			"max_version":   x.GetMaxVersion().String(),
			"min_version":   x.GetMinVersion().String(),
		}
		csValue = append(csValue, csVal)
	}
	return csValue
}

func FlattenUseMtls(x *ves_io_schema_views.DownstreamTlsValidationContext) []interface{} {
	umValue := make([]interface{}, 0)
	if x != nil {
		umVal := map[string]interface{}{
			"client_certificate_optional": x.GetClientCertificateOptional(),
			"crl":                         FlattenObjectRefTypeSet(x.GetCrl()),
			"no_crl":                      isEmpty(x.GetNoCrl()),
			"trusted_ca":                  FlattenObjectRefTypeSet(x.GetTrustedCa()),
			"trusted_ca_url":              x.GetTrustedCaUrl(),
			"xfcc_disabled":               isEmpty(x.GetXfccDisabled()),
			"xfcc_options":                FlattenXfccOptions(x.GetXfccOptions()),
		}
		umValue = append(umValue, umVal)
	}
	return umValue
}

func FlattenXfccHeaderElements(x []ves_io_schema.XfccElement) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenXfccOptions(x *ves_io_schema_views.XfccHeaderKeys) []interface{} {
	xoValue := make([]interface{}, 0)
	if x != nil {
		xoVal := map[string]interface{}{
			"xfcc_header_elements": FlattenXfccHeaderElements(x.GetXfccHeaderElements()),
		}
		xoValue = append(xoValue, xoVal)
	}
	return xoValue
}

func FlattenHttpProtocolOptions(x *ves_io_schema_virtual_host.HttpProtocolOptions) []interface{} {
	hpoValue := make([]interface{}, 0)
	if x != nil {
		hpoVal := map[string]interface{}{
			"http_protocol_enable_v1_only": FlattenHttpProtocolEnableV1Only(x.GetHttpProtocolEnableV1Only()),
			"http_protocol_enable_v1_v2":   isEmpty(x.GetHttpProtocolEnableV1V2()),
			"http_protocol_enable_v2_only": isEmpty(x.GetHttpProtocolEnableV2Only()),
		}
		hpoValue = append(hpoValue, hpoVal)
	}
	return hpoValue
}

func FlattenHttpProtocolEnableV1Only(x *ves_io_schema_virtual_host.Http1ProtocolOptions) []interface{} {
	hpoValue := make([]interface{}, 0)
	if x != nil {
		hpoVal := map[string]interface{}{
			"header_transformation": FlattenHeaderTransformationType(x.GetHeaderTransformation()),
		}
		hpoValue = append(hpoValue, hpoVal)
	}
	return hpoValue
}

func FlattenHeaderTransformationType(x *ves_io_schema.HeaderTransformationType) []interface{} {
	httValue := make([]interface{}, 0)
	if x != nil {
		httVal := map[string]interface{}{
			"default_header_transformation":       isEmpty(x.GetDefaultHeaderTransformation()),
			"legacy_header_transformation":        isEmpty(x.GetLegacyHeaderTransformation()),
			"preserve_case_header_transformation": isEmpty(x.GetPreserveCaseHeaderTransformation()),
			"proper_case_header_transformation":   isEmpty(x.GetProperCaseHeaderTransformation()),
		}
		httValue = append(httValue, httVal)
	}
	return httValue
}

func FlattenHttpsAutoCert(x *ves_io_schema_views_http_loadbalancer.ProxyTypeHttpsAutoCerts) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"add_hsts":                 x.GetAddHsts(),
			"coalescing_options":       FlattenCoalescingOptions(x.GetCoalescingOptions()),
			"connection_idle_timeout":  x.GetConnectionIdleTimeout(),
			"default_loadbalancer":     isEmpty(x.GetDefaultLoadbalancer()),
			"non_default_loadbalancer": isEmpty(x.GetNonDefaultLoadbalancer()),
			"http_protocol_options":    FlattenHttpProtocolOptions(x.GetHttpProtocolOptions()),
			"http_redirect":            x.GetHttpRedirect(),
			"no_mtls":                  isEmpty(x.GetNoMtls()),
			"use_mtls":                 FlattenUseMtls(x.GetUseMtls()),
			"disable_path_normalize":   isEmpty(x.GetDisablePathNormalize()),
			"enable_path_normalize":    isEmpty(x.GetEnablePathNormalize()),
			"port":                     x.GetPort(),
			"port_ranges":              x.GetPortRanges(),
			"append_server_name":       x.GetAppendServerName(),
			"default_header":           isEmpty(x.GetDefaultHeader()),
			"pass_through":             isEmpty(x.GetPassThrough()),
			"server_name":              x.GetServerName(),
			"tls_config":               FlattenTlsConfig(x.GetTlsConfig()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenCSDDR(x []*ves_io_schema_app_type.CustomSensitiveDataDetectionRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		rs := map[string]interface{}{
			"metadata":                        FlattenMetadata(val.GetMetadata()),
			"sensitive_data_detection_config": FlattenSensitiveDataDetectionConfig(val.GetSensitiveDataDetectionConfig()),
			"sensitive_data_type":             FlattenSensitiveDataType(val.GetSensitiveDataType()),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenSensitiveDataType(x *ves_io_schema_app_type.CustomSensitiveDataType) []interface{} {
	sdtValue := make([]interface{}, 0)
	if x != nil {
		sdtVal := map[string]interface{}{
			"type": x.GetType(),
		}
		sdtValue = append(sdtValue, sdtVal)
	}
	return sdtValue
}

func FlattenDisabledBuiltInRules(x []*ves_io_schema_app_type.BuiltInSensitiveDataType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		rs := map[string]interface{}{
			"name": v.GetName(),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenSingleLbApp(x *ves_io_schema_views_http_loadbalancer.SingleLoadBalancerAppSetting) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"disable_discovery":                isEmpty(x.GetDisableDiscovery()),
			"enable_discovery":                 FlattenEnableApiDiscovery(x.GetEnableDiscovery()),
			"disable_malicious_user_detection": isEmpty(x.GetDisableMaliciousUserDetection()),
			"enable_malicious_user_detection":  isEmpty(x.GetEnableMaliciousUserDetection()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenCookiesToModify(x []*ves_io_schema.CookieManipulationOptionType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"disable_tampering_protection": isEmpty(val.GetDisableTamperingProtection()),
			"enable_tampering_protection":  isEmpty(val.GetEnableTamperingProtection()),
			"add_httponly":                 isEmpty(val.GetAddHttponly()),
			"ignore_httponly":              isEmpty(val.GetIgnoreHttponly()),
			"name":                         val.GetName(),
			"ignore_samesite":              isEmpty(val.GetIgnoreSamesite()),
			"samesite_lax":                 isEmpty(val.GetSamesiteLax()),
			"samesite_none":                isEmpty(val.GetSamesiteNone()),
			"samesite_strict":              isEmpty(val.GetSamesiteStrict()),
			"add_secure":                   isEmpty(val.GetAddSecure()),
			"ignore_secure":                isEmpty(val.GetIgnoreSecure()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenCustomError(x map[uint32]string) map[string]interface{} {
	mpValue := make(map[string]interface{})
	for key, val := range x {
		mpValue[strconv.Itoa(int(key))] = val
	}
	return mpValue
}

func FlattenRequestHeadersToAdd(x []*ves_io_schema.HeaderManipulationOptionType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"append":       val.GetAppend(),
			"name":         val.GetName(),
			"secret_value": FlattenPrivateKey(val.GetSecretValue()),
			"value":        val.GetValue(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenReqestCookiesToAdd(x []*ves_io_schema.CookieValueOption) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"overwrite":    val.GetOverwrite(),
			"name":         val.GetName(),
			"secret_value": FlattenPrivateKey(val.GetSecretValue()),
			"value":        val.GetValue(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenResponseCookiesToAdd(x []*ves_io_schema.SetCookieValueOption) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"add_domain":         val.GetAddDomain(),
			"ignore_domain":      isEmpty(val.GetIgnoreDomain()),
			"add_expiry":         val.GetAddExpiry(),
			"ignore_expiry":      isEmpty(val.GetIgnoreExpiry()),
			"add_httponly":       isEmpty(val.GetAddHttponly()),
			"ignore_httponly":    isEmpty(val.GetIgnoreHttponly()),
			"ignore_max_age":     isEmpty(val.GetIgnoreMaxAge()),
			"max_age_value":      val.GetMaxAgeValue(),
			"name":               val.GetName(),
			"overwrite":          val.GetOverwrite(),
			"add_partitioned":    isEmpty(val.GetAddPartitioned()),
			"ignore_partitioned": isEmpty(val.GetIgnorePartitioned()),
			"add_path":           val.GetAddPath(),
			"ignore_path":        isEmpty(val.GetIgnoreSamesite()),
			"ignore_samesite":    isEmpty(val.GetIgnoreSamesite()),
			"samesite_lax":       isEmpty(val.GetSamesiteLax()),
			"samesite_none":      isEmpty(val.GetSamesiteNone()),
			"samesite_strict":    isEmpty(val.GetSamesiteStrict()),
			"add_secure":         isEmpty(val.GetAddSecure()),
			"ignore_secure":      isEmpty(val.GetIgnoreSecure()),
			"ignore_value":       isEmpty(val.GetIgnoreValue()),
			"secret_value":       FlattenPrivateKey(val.GetSecretValue()),
			"value":              val.GetValue(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenMoreOption(x *ves_io_schema_views_http_loadbalancer.AdvancedOptionsType) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"buffer_policy":               FlattenBufferPolicy(x.GetBufferPolicy()),
			"compression_params":          FlattenCompressionParams(x.GetCompressionParams()),
			"custom_errors":               FlattenCustomError(x.GetCustomErrors()),
			"disable_default_error_pages": x.GetDisableDefaultErrorPages(),
			"idle_timeout":                x.GetIdleTimeout(),
			"max_request_header_size":     x.GetMaxRequestHeaderSize(),
			"request_cookies_to_add":      FlattenReqestCookiesToAdd(x.GetRequestCookiesToAdd()),
			"request_cookies_to_remove":   x.GetRequestCookiesToRemove(),
			"request_headers_to_add":      FlattenRequestHeadersToAdd(x.GetRequestHeadersToAdd()),
			"request_headers_to_remove":   x.GetRequestHeadersToRemove(),
			"response_cookies_to_add":     FlattenResponseCookiesToAdd(x.GetResponseCookiesToAdd()),
			"response_cookies_to_remove":  x.GetResponseCookiesToRemove(),
			"response_headers_to_add":     FlattenRequestHeadersToAdd(x.GetResponseHeadersToAdd()),
			"response_headers_to_remove":  x.GetResponseHeadersToRemove(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenAdditionalDomains(x *ves_io_schema.DomainNameList) []interface{} {
	adValue := make([]interface{}, 0)
	if x != nil {
		adVal := map[string]interface{}{
			"domains": x.GetDomains(),
		}
		adValue = append(adValue, adVal)
	}
	return adValue
}

func FlattenJavascriptInfo(x *ves_io_schema_virtual_host.JavaScriptConfigType) []interface{} {
	jiValue := make([]interface{}, 0)
	if x != nil {
		jiVal := map[string]interface{}{
			"cache_prefix":      x.GetCachePrefix(),
			"custom_script_url": x.GetCustomScriptUrl(),
		}
		jiValue = append(jiValue, jiVal)
	}
	return jiValue
}

func FlattenCompressionParams(x *ves_io_schema_virtual_host.CompressionType) []interface{} {
	cpValue := make([]interface{}, 0)
	if x != nil {
		cpVal := map[string]interface{}{
			"content_length":                x.GetContentLength(),
			"content_type":                  x.GetContentType(),
			"disable_on_etag_header":        x.GetDisableOnEtagHeader(),
			"remove_accept_encoding_header": x.GetRemoveAcceptEncodingHeader(),
		}
		cpValue = append(cpValue, cpVal)
	}
	return cpValue
}

func FlattenBufferPolicy(x *ves_io_schema.BufferConfigType) []interface{} {
	bpValue := make([]interface{}, 0)
	if x != nil {
		bpVal := map[string]interface{}{
			"disabled":          x.GetDisabled(),
			"max_request_bytes": x.GetMaxRequestBytes(),
		}
		bpValue = append(bpValue, bpVal)
	}
	return bpValue
}

func FlattenOriginServers(x []*ves_io_schema_views_origin_pool.OriginServerType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {

		mapValue := map[string]interface{}{
			"cbip_service":           FlattenCbipService(val.GetCbipService()),
			"consul_service":         FlattenConsulService(val.GetConsulService()),
			"custom_endpoint_object": FlattenCustomEndpointObject(val.GetCustomEndpointObject()),
			"k8s_service":            FlattenK8sService(val.GetK8SService()),
			"private_ip":             FlattenPrivateIp(val.GetPrivateIp()),
			"private_name":           FlattenPrivateName(val.GetPrivateName()),
			"public_ip":              FlattenOSPublicIp(val.GetPublicIp()),
			"public_name":            FlattenPublicName(val.GetPublicName()),
			"vn_private_ip":          FlattenVnPrivateIp(val.GetVnPrivateIp()),
			"vn_private_name":        FlattenVnPrivateName(val.GetVnPrivateName()),
			"labels":                 FlattenEndpointSubsets(val.GetLabels()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenCbipService(x *ves_io_schema_views_origin_pool.OriginServerCBIPService) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"service_name": x.GetServiceName(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenVnPrivateName(x *ves_io_schema_views_origin_pool.OriginServerVirtualNetworkName) []interface{} {
	vnPrivateNValue := make([]interface{}, 0)
	if x != nil {
		vnPrivateNVal := map[string]interface{}{
			"dns_name":        x.GetDnsName(),
			"private_network": FlattenObjectRefTypeSet(x.GetPrivateNetwork()),
		}
		vnPrivateNValue = append(vnPrivateNValue, vnPrivateNVal)
	}
	return vnPrivateNValue
}

func FlattenVnPrivateIp(x *ves_io_schema_views_origin_pool.OriginServerVirtualNetworkIP) []interface{} {
	vnPrivateIpValue := make([]interface{}, 0)
	if x != nil {
		vnPrivateIpVal := map[string]interface{}{
			"ip":              x.GetIp(),
			"virtual_network": FlattenObjectRefTypeSet(x.GetVirtualNetwork()),
			"ipv6":            x.GetIpv6(),
		}
		vnPrivateIpValue = append(vnPrivateIpValue, vnPrivateIpVal)
	}
	return vnPrivateIpValue
}

func FlattenPublicName(x *ves_io_schema_views_origin_pool.OriginServerPublicName) []interface{} {
	publicNValue := make([]interface{}, 0)
	if x != nil {
		publicNVal := map[string]interface{}{
			"dns_name":         x.GetDnsName(),
			"refresh_interval": x.GetRefreshInterval(),
		}
		publicNValue = append(publicNValue, publicNVal)
	}
	return publicNValue
}

func FlattenOSPublicIp(x *ves_io_schema_views_origin_pool.OriginServerPublicIP) []interface{} {
	publicIpValue := make([]interface{}, 0)
	if x != nil {
		publicIpVal := map[string]interface{}{
			"ip":   x.GetIp(),
			"ipv6": x.GetIpv6(),
		}
		publicIpValue = append(publicIpValue, publicIpVal)
	}
	return publicIpValue
}

func FlattenPrivateName(x *ves_io_schema_views_origin_pool.OriginServerPrivateName) []interface{} {
	pnValue := make([]interface{}, 0)
	if x != nil {
		pnVal := map[string]interface{}{
			"dns_name":         x.GetDnsName(),
			"inside_network":   isEmpty(x.GetInsideNetwork()),
			"outside_network":  isEmpty(x.GetOutsideNetwork()),
			"segment":          FlattenObjectRefTypeSet(x.GetSegment()),
			"refresh_interval": x.GetRefreshInterval(),
			"site_locator":     FlattenSiteLocator(x.GetSiteLocator()),
			"snat_pool":        FlattenSnatPool(x.GetSnatPool()),
		}
		pnValue = append(pnValue, pnVal)
	}
	return pnValue
}

func FlattenPrivateIp(x *ves_io_schema_views_origin_pool.OriginServerPrivateIP) []interface{} {
	piValue := make([]interface{}, 0)
	if x != nil {
		piVal := map[string]interface{}{
			"inside_network":  isEmpty(x.GetInsideNetwork()),
			"outside_network": isEmpty(x.GetOutsideNetwork()),
			"segment":         FlattenObjectRefTypeSet(x.GetSegment()),
			"ip":              x.GetIp(),
			"ipv6":            x.GetIpv6(),
			"site_locator":    FlattenSiteLocator(x.GetSiteLocator()),
			"snat_pool":       FlattenSnatPool(x.GetSnatPool()),
		}
		piValue = append(piValue, piVal)
	}
	return piValue
}

func FlattenK8sService(x *ves_io_schema_views_origin_pool.OriginServerK8SService) []interface{} {
	k8sValue := make([]interface{}, 0)
	if x != nil {
		k8sVal := map[string]interface{}{
			"inside_network":  isEmpty(x.GetInsideNetwork()),
			"outside_network": isEmpty(x.GetOutsideNetwork()),
			"vk8s_networks":   isEmpty(x.GetVk8SNetworks()),
			"protocol":        x.GetProtocol(),
			"service_name":    x.GetServiceName(),
			"site_locator":    FlattenSiteLocator(x.GetSiteLocator()),
			"snat_pool":       FlattenSnatPool(x.GetSnatPool()),
		}
		k8sValue = append(k8sValue, k8sVal)
	}
	return k8sValue
}

func FlattenCustomEndpointObject(x *ves_io_schema_views_origin_pool.OriginServerCustomEndpoint) []interface{} {
	ceoValue := make([]interface{}, 0)
	if x != nil {
		ceoVal := map[string]interface{}{
			"endpoint": FlattenObjectRefTypeSet(x.GetEndpoint()),
		}
		ceoValue = append(ceoValue, ceoVal)
	}
	return ceoValue
}

func FlattenSnatPool(x *ves_io_schema_views.SnatPoolConfiguration) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		rs := map[string]interface{}{
			"no_snat_pool": isEmpty(x.GetNoSnatPool()),
			"snat_pool":    FlattenIpAllowedList(x.GetSnatPool()),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenConsulService(x *ves_io_schema_views_origin_pool.OriginServerConsulService) []interface{} {
	csValue := make([]interface{}, 0)
	if x != nil {
		csVal := map[string]interface{}{
			"inside_network":  isEmpty(x.GetInsideNetwork()),
			"outside_network": isEmpty(x.GetOutsideNetwork()),
			"service_name":    x.GetServiceName(),
			"site_locator":    FlattenSiteLocator(x.GetSiteLocator()),
			"snat_pool":       FlattenSnatPool(x.GetSnatPool()),
		}
		csValue = append(csValue, csVal)
	}
	return csValue
}

func FlattenSiteLocator(x *ves_io_schema_views.SiteLocator) []interface{} {
	slValue := make([]interface{}, 0)
	if x != nil {
		slVal := map[string]interface{}{
			"site":         FlattenObjectRefTypeSet(x.GetSite()),
			"virtual_site": FlattenObjectRefTypeSet(x.GetVirtualSite()),
		}
		slValue = append(slValue, slVal)
	}
	return slValue
}

func FlattenCircuitBreaker(x *ves_io_schema_cluster.CircuitBreaker) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		rs := map[string]interface{}{
			"connection_limit": x.GetConnectionLimit(),
			"max_requests":     x.GetMaxRequests(),
			"pending_requests": x.GetPendingRequests(),
			"priority":         x.GetPriority().String(),
			"retries":          x.GetRetries(),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenDefaultPool(x *ves_io_schema_views_origin_pool.GlobalSpecType) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"advanced_options":       FlattenDPAdvancedOption(x.GetAdvancedOptions()),
			"endpoint_selection":     x.GetEndpointSelection().String(),
			"health_check_port":      x.GetHealthCheckPort(),
			"same_as_endpoint_port":  x.GetSameAsEndpointPort() != nil,
			"healthcheck":            FlattenVObjectRefTypeList(x.GetHealthcheck()),
			"loadbalancer_algorithm": x.GetLoadbalancerAlgorithm().String(),
			"origin_servers":         FlattenOriginServers(x.GetOriginServers()),
			"automatic_port":         isEmpty(x.GetAutomaticPort()),
			"lb_port":                isEmpty(x.GetLbPort()),
			"port":                   x.GetPort(),
			"no_tls":                 isEmpty(x.GetNoTls()),
			"use_tls":                FlattenUseTls(x.GetUseTls()),
			"view_internal":          FlattenObjectRefTypeSet(x.GetViewInternal()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenUseTls(x *ves_io_schema_views_origin_pool.UpstreamTlsParameters) []interface{} {
	useTlsValue := make([]interface{}, 0)
	if x != nil {
		useTlsVal := map[string]interface{}{
			"default_session_key_caching": isEmpty(x.GetDefaultSessionKeyCaching()),
			"disable_session_key_caching": isEmpty(x.GetDisableSessionKeyCaching()),
			"max_session_keys":            x.GetMaxSessionKeys(),
			"no_mtls":                     isEmpty(x.GetNoMtls()),
			"use_mtls":                    FlattenOPUseMtls(x.GetUseMtls()),
			"use_mtls_obj":                FlattenObjectRefTypeSet(x.GetUseMtlsObj()),
			"skip_server_verification":    isEmpty(x.GetSkipServerVerification()),
			"use_server_verification":     FlattenUseServerVerification(x.GetUseServerVerification()),
			"volterra_trusted_ca":         isEmpty(x.GetVolterraTrustedCa()),
			"disable_sni":                 isEmpty(x.GetDisableSni()),
			"sni":                         x.GetSni(),
			"use_host_header_as_sni":      isEmpty(x.GetUseHostHeaderAsSni()),
			"tls_config":                  FlattenTlsConfig(x.GetTlsConfig()),
		}
		useTlsValue = append(useTlsValue, useTlsVal)
	}
	return useTlsValue
}

func FlattenUseServerVerification(x *ves_io_schema_views_origin_pool.UpstreamTlsValidationContext) []interface{} {
	usvValue := make([]interface{}, 0)
	if x != nil {
		usvVal := map[string]interface{}{
			"trusted_ca":     FlattenObjectRefTypeSet(x.GetTrustedCa()),
			"trusted_ca_url": x.GetTrustedCaUrl(),
		}
		usvValue = append(usvValue, usvVal)
	}
	return usvValue
}

func FlattenOPUseMtls(x *ves_io_schema_views_origin_pool.TlsCertificatesType) []interface{} {
	useMtlsValue := make([]interface{}, 0)
	if x != nil {
		useMtlsVal := map[string]interface{}{
			"tls_certificates": FlattenTlsCertificates(x.GetTlsCertificates()),
		}
		useMtlsValue = append(useMtlsValue, useMtlsVal)
	}
	return useMtlsValue
}

func FlattenDPAdvancedOption(x *ves_io_schema_views_origin_pool.OriginPoolAdvancedOptions) []interface{} {
	aoValue := make([]interface{}, 0)
	if x != nil {
		aoVal := map[string]interface{}{
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
			"disable_proxy_protocol":           x.GetDisableProxyProtocol(),
			"proxy_protocol_v1":                x.GetProxyProtocolV1(),
			"proxy_protocol_v2":                x.GetProxyProtocolV2(),
			"disable_subsets":                  isEmpty(x.GetDisableSubsets()),
			"enable_subsets":                   FlattenEnableSubsets(x.GetEnableSubsets()),
		}
		aoValue = append(aoValue, aoVal)
	}
	return aoValue
}

func FlattenDPAOEndpointSubsets(x []*ves_io_schema_cluster.EndpointSubsetSelectorType) []interface{} {
	res := make([]interface{}, 0)
	for _, v := range x {
		val := map[string]interface{}{
			"keys": v.GetKeys(),
		}
		res = append(res, val)
	}
	return res
}
func FlattenEnableSubsets(x *ves_io_schema_views_origin_pool.OriginPoolSubsets) []interface{} {
	esValue := make([]interface{}, 0)
	if x != nil {
		esVal := map[string]interface{}{
			"endpoint_subsets": FlattenDPAOEndpointSubsets(x.GetEndpointSubsets()),
			"any_endpoint":     isEmpty(x.GetAnyEndpoint()),
			"default_subset":   FlattenDefaultSubset(x.GetDefaultSubset()),
			"fail_request":     isEmpty(x.GetFailRequest()),
		}
		esValue = append(esValue, esVal)
	}
	return esValue
}

func FlattenDefaultSubset(x *ves_io_schema_views_origin_pool.OriginPoolDefaultSubset) []interface{} {
	dsValue := make([]interface{}, 0)
	if x != nil {
		dsVal := map[string]interface{}{
			"default_subset": FlattenEndpointSubsets(x.GetDefaultSubset()),
		}
		dsValue = append(dsValue, dsVal)
	}
	return dsValue
}

func FlattenOutlierDetection(x *ves_io_schema_cluster.OutlierDetectionType) []interface{} {
	odValue := make([]interface{}, 0)
	if x != nil {
		odVal := map[string]interface{}{
			"base_ejection_time":          x.GetBaseEjectionTime(),
			"consecutive_5xx":             x.GetConsecutive_5Xx(),
			"consecutive_gateway_failure": x.GetConsecutiveGatewayFailure(),
			"interval":                    x.GetInterval(),
			"max_ejection_percent":        x.GetMaxEjectionPercent(),
		}
		odValue = append(odValue, odVal)
	}
	return odValue
}

func FlattenHttp2Options(x *ves_io_schema_cluster.Http2ProtocolOptions) []interface{} {
	hoValue := make([]interface{}, 0)
	if x != nil {
		hoVal := map[string]interface{}{
			"enabled": x.GetEnabled(),
		}
		hoValue = append(hoValue, hoVal)
	}
	return hoValue
}

func FlattenDefaultPoolList(x *ves_io_schema_views.OriginPoolListType) []interface{} {
	rslt := make([]interface{}, 0)

	if x != nil {
		mapValue := map[string]interface{}{
			"pools": FlattenDefaultRoutePools(x.GetPools()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenCountryCode(x []ves_io_schema_policy.CountryCode) []string {
	rslt := make([]string, 0)
	for _, val := range x {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenOSSR(x []*ves_io_schema_policy.OriginServerSubsetRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		rs := map[string]interface{}{
			"any_asn":                      isEmpty(v.GetAnyAsn()),
			"asn_list":                     FlattenAsnList(v.GetAsnList()),
			"asn_matcher":                  FlattenAsnMatcher(v.GetAsnMatcher()),
			"country_codes":                FlattenCountryCode(v.GetCountryCodes()),
			"any_ip":                       isEmpty(v.GetAnyIp()),
			"ip_matcher":                   FlattenIpMatcher(v.GetIpMatcher()),
			"ip_prefix_list":               FlattenIpPrefixList(v.GetIpPrefixList()),
			"metadata":                     FlattenMetadata(v.GetMetadata()),
			"origin_server_subsets_action": FlattenEndpointSubsets(v.GetOriginServerSubsetsAction()),
			"re_name_list":                 v.GetReNameList(),
			"client_selector":              FlattenServerSelector(v.GetClientSelector()),
			"none":                         isEmpty(v.GetNone()),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenOSSRL(x *ves_io_schema_views_http_loadbalancer.OriginServerSubsetRuleListType) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		rs := map[string]interface{}{
			"origin_server_subset_rules": FlattenOSSR(x.GetOriginServerSubsetRules()),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenARLApiEndpointRules(x []*ves_io_schema_views_common_waf.ApiEndpointRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"api_endpoint_method": FlattenHttpMethod(val.GetApiEndpointMethod()),
			"api_endpoint_path":   val.GetApiEndpointPath(),
			// "base_path":           val.GetBasePath(),
			"client_matcher":      FlattenClientMatcher(val.GetClientMatcher()),
			"any_domain":          val.GetAnyDomain() != nil,
			"specific_domain":     val.GetSpecificDomain(),
			"inline_rate_limiter": FlattenInlineRateLimiter(val.GetInlineRateLimiter()),
			"ref_rate_limiter":    FlattenObjectRefTypeSet(val.GetRefRateLimiter()),
			"request_matcher":     FlattenRequestMatcher(val.GetRequestMatcher()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenInlineRateLimiter(x *ves_io_schema_views_common_waf.InlineRateLimiter) []interface{} {
	irlValue := make([]interface{}, 0)
	if x != nil {
		irlVal := map[string]interface{}{
			"ref_user_id":         FlattenObjectRefTypeSet(x.GetRefUserId()),
			"use_http_lb_user_id": isEmpty(x.GetUseHttpLbUserId()),
			"threshold":           x.GetThreshold(),
			"unit":                x.GetUnit().String(),
		}
		irlValue = append(irlValue, irlVal)
	}
	return irlValue
}

func FlattenServerURLRules(x []*ves_io_schema_views_common_waf.ServerUrlRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"api_group":           val.GetApiGroup(),
			"base_path":           val.GetBasePath(),
			"client_matcher":      FlattenClientMatcher(val.GetClientMatcher()),
			"any_domain":          isEmpty(val.GetAnyDomain()),
			"specific_domain":     val.GetSpecificDomain(),
			"inline_rate_limiter": FlattenInlineRateLimiter(val.GetInlineRateLimiter()),
			"ref_rate_limiter":    FlattenObjectRefTypeSet(val.GetRefRateLimiter()),
			"request_matcher":     FlattenRequestMatcher(val.GetRequestMatcher()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenBRLR(x []*ves_io_schema_views_common_waf.BypassRateLimitingRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		rs := map[string]interface{}{
			"client_matcher":  FlattenClientMatcher(val.GetClientMatcher()),
			"any_url":         isEmpty(val.GetAnyUrl()),
			"api_endpoint":    FlattenApiEndPoint(val.GetApiEndpoint()),
			"api_groups":      FlattenApiGroups(val.GetApiGroups()),
			"base_path":       val.GetBasePath(),
			"any_domain":      isEmpty(val.GetAnyDomain()),
			"specific_domain": val.GetSpecificDomain(),
			"request_matcher": FlattenRequestMatcher(val.GetRequestMatcher()),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenApiRateLimit(x *ves_io_schema_views_common_waf.APIRateLimit) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"api_endpoint_rules":         FlattenARLApiEndpointRules(x.GetApiEndpointRules()),
			"bypass_rate_limiting_rules": FlattenBypassRateLimitingRules(x.GetBypassRateLimitingRules()),
			"custom_ip_allowed_list":     FlattenCustomIpAllowedList(x.GetCustomIpAllowedList()),
			"ip_allowed_list":            FlattenIpAllowedList(x.GetIpAllowedList()),
			"no_ip_allowed_list":         isEmpty(x.GetNoIpAllowedList()),
			"server_url_rules":           FlattenServerURLRules(x.GetServerUrlRules()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenIpAllowedList(x *ves_io_schema_views.PrefixStringListType) []interface{} {
	ialValue := make([]interface{}, 0)
	if x != nil {
		ialVal := map[string]interface{}{
			"ipv6_prefixes": x.GetIpv6Prefixes(),
			"prefixes":      x.GetPrefixes(),
		}
		ialValue = append(ialValue, ialVal)
	}
	return ialValue
}

func FlattenCustomIpAllowedList(x *ves_io_schema_views_common_waf.CustomIpAllowedList) []interface{} {
	cialValue := make([]interface{}, 0)
	if x != nil {
		cialVal := map[string]interface{}{
			"rate_limiter_allowed_prefixes": FlattenVObjectRefTypeList(x.GetRateLimiterAllowedPrefixes()),
		}
		cialValue = append(cialValue, cialVal)
	}
	return cialValue
}

func FlattenBypassRateLimitingRules(x *ves_io_schema_views_common_waf.BypassRateLimitingRules) []interface{} {
	brlrValue := make([]interface{}, 0)
	if x != nil {
		brlrVal := map[string]interface{}{
			"bypass_rate_limiting_rules": FlattenBRLR(x.GetBypassRateLimitingRules()),
		}
		brlrValue = append(brlrValue, brlrVal)
	}
	return brlrValue
}

func FlattenRateLimit(x *ves_io_schema_views_common_waf.RateLimitConfigType) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"custom_ip_allowed_list": FlattenCustomIpAllowedList(x.GetCustomIpAllowedList()),
			"ip_allowed_list":        FlattenIpAllowedList(x.GetIpAllowedList()),
			"no_ip_allowed_list":     isEmpty(x.GetNoIpAllowedList()),
			"no_policies":            isEmpty(x.GetNoPolicies()),
			"policies":               FlattenPolicies(x.GetPolicies()),
			"rate_limiter":           FlattenRateLimiter(x.GetRateLimiter()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenHours(x *ves_io_schema_rate_limiter.InputHours) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"duration": x.GetDuration(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenMinutes(x *ves_io_schema_rate_limiter.InputMinutes) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"duration": x.GetDuration(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenSeconds(x *ves_io_schema_rate_limiter.InputSeconds) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"duration": x.GetDuration(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenActionBlock(x *ves_io_schema_rate_limiter.RateLimitBlockAction) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"hours":   FlattenHours(x.GetHours()),
			"minutes": FlattenMinutes(x.GetMinutes()),
			"seconds": FlattenSeconds(x.GetSeconds()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenRateLimiter(x *ves_io_schema_rate_limiter.RateLimitValue) []interface{} {
	rlValue := make([]interface{}, 0)
	if x != nil {
		rlVal := map[string]interface{}{
			"action_block":      FlattenActionBlock(x.GetActionBlock()),
			"burst_multiplier":  x.GetBurstMultiplier(),
			"period_multiplier": x.GetPeriodMultiplier(),
			"total_number":      x.GetTotalNumber(),
			"unit":              x.GetUnit().String(),
		}
		rlValue = append(rlValue, rlVal)
	}
	return rlValue
}

func FlattenPolicies(x *ves_io_schema_rate_limiter_policy.PolicyList) []interface{} {
	pValue := make([]interface{}, 0)
	if x.GetPolicies() != nil {
		pVal := map[string]interface{}{
			"policies": FlattenVObjectRefTypeList(x.GetPolicies()),
		}
		pValue = append(pValue, pVal)
	}
	return pValue
}

func FlattenTA(x []*ves_io_schema_route.TagAttribute) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		rs := map[string]interface{}{
			"javascript_tag": v.GetJavascriptTag().String(),
			"tag_value":      v.GetTagValue(),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenJTags(x []*ves_io_schema_route.JavaScriptTag) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		rs := map[string]interface{}{
			"javascript_url": v.GetJavascriptUrl(),
			"tag_attributes": FlattenTA(v.GetTagAttributes()),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenRoutes(x []*ves_io_schema_views_http_loadbalancer.RouteType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"custom_route_object":   FlattenCustomRouteObject(val.GetCustomRouteObject()),
			"direct_response_route": FlattenDirectResponseRoute(val.GetDirectResponseRoute()),
			"redirect_route":        FlattenRedirectRoute(val.GetRedirectRoute()),
			"simple_route":          FlattenSimpleRoute(val.GetSimpleRoute()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenSimpleRoute(x *ves_io_schema_views_http_loadbalancer.RouteTypeSimple) []interface{} {
	simpleRouteValue := make([]interface{}, 0)
	if x != nil {
		simpleRouteVal := map[string]interface{}{
			"advanced_options":     FlattenSRAdvancedOptions(x.GetAdvancedOptions()),
			"headers":              FlattenHTTPHeaders(x.GetHeaders()),
			"auto_host_rewrite":    isEmpty(x.GetAutoHostRewrite()),
			"disable_host_rewrite": isEmpty(x.GetDisableHostRewrite()),
			"host_rewrite":         x.GetHostRewrite(),
			"http_method":          x.GetHttpMethod().String(),
			"incoming_port":        FlattenIncomingPort(x.GetIncomingPort()),
			"origin_pools":         FlattenDefaultRoutePools(x.GetOriginPools()),
			"path":                 FlattenPath(x.GetPath()),
			"query_params":         FlattenSRQueryParams(x.GetQueryParams()),
		}
		simpleRouteValue = append(simpleRouteValue, simpleRouteVal)
	}
	return simpleRouteValue
}

func FlattenSRQueryParams(x *ves_io_schema_route.QueryParamsSimpleRoute) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"remove_all_params": isEmpty(x.GetRemoveAllParams()),
			"replace_params":    x.GetReplaceParams(),
			"retain_all_params": isEmpty(x.GetRetainAllParams()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenSRAdvancedOptions(x *ves_io_schema_views_http_loadbalancer.RouteSimpleAdvancedOptions) []interface{} {
	adOptionValue := make([]interface{}, 0)
	if x != nil {
		adOptionVal := map[string]interface{}{
			"bot_defense_javascript_injection":           FlattenBotDefenseJavascriptInjection(x.GetBotDefenseJavascriptInjection()),
			"inherited_bot_defense_javascript_injection": x.GetInheritedBotDefenseJavascriptInjection() != nil,
			"buffer_policy":              FlattenBufferPolicy(x.GetBufferPolicy()),
			"common_buffering":           isEmpty(x.GetCommonBuffering()),
			"do_not_retract_cluster":     isEmpty(x.GetDoNotRetractCluster()),
			"retract_cluster":            isEmpty(x.GetRetractCluster()),
			"cors_policy":                FlattenCorsPolicy(x.GetCorsPolicy()),
			"csrf_policy":                FlattenCsrfPolicy(x.GetCsrfPolicy()),
			"disable_location_add":       x.GetDisableLocationAdd(),
			"endpoint_subsets":           FlattenEndpointSubsets(x.GetEndpointSubsets()),
			"common_hash_policy":         x.GetCommonHashPolicy() != nil,
			"specific_hash_policy":       FlattenSpecificHashPolicy(x.GetSpecificHashPolicy()),
			"disable_mirroring":          isEmpty(x.GetDisableMirroring()),
			"mirror_policy":              FlattenMirrorPolicy(x.GetMirrorPolicy()),
			"priority":                   x.GetPriority().String(),
			"request_cookies_to_add":     FlattenReqestCookiesToAdd(x.GetRequestCookiesToAdd()),
			"request_cookies_to_remove":  x.GetRequestCookiesToRemove(),
			"request_headers_to_add":     FlattenRequestHeadersToAdd(x.GetRequestHeadersToAdd()),
			"request_headers_to_remove":  x.GetRequestHeadersToRemove(),
			"response_cookies_to_add":    FlattenResponseCookiesToAdd(x.GetResponseCookiesToAdd()),
			"response_cookies_to_remove": x.GetResponseCookiesToRemove(),
			"response_headers_to_add":    FlattenRequestHeadersToAdd(x.GetResponseHeadersToAdd()),
			"response_headers_to_remove": x.GetResponseHeadersToRemove(),
			"default_retry_policy":       isEmpty(x.GetDefaultRetryPolicy()),
			"no_retry_policy":            isEmpty(x.GetNoRetryPolicy()),
			"retry_policy":               FlattenRetryPolicy(x.GetRetryPolicy()),
			"disable_prefix_rewrite":     isEmpty(x.GetDisablePrefixRewrite()),
			"prefix_rewrite":             x.GetPrefixRewrite(),
			"regex_rewrite":              FlattenRegexRewrite(x.GetRegexRewrite()),
			"disable_spdy":               isEmpty(x.GetDisableSpdy()),
			"enable_spdy":                isEmpty(x.GetEnableSpdy()),
			"timeout":                    x.GetTimeout(),
			"app_firewall":               FlattenObjectRefTypeSet(x.GetAppFirewall()),
			"disable_waf":                isEmpty(x.GetDisableWaf()),
			"inherited_waf":              isEmpty(x.GetInheritedWaf()),
			"disable_web_socket_config":  isEmpty(x.GetDisableWebSocketConfig()),
			"web_socket_config":          FlattenWebSocketConfig(x.GetWebSocketConfig()),
		}
		adOptionValue = append(adOptionValue, adOptionVal)
	}
	return adOptionValue
}

func FlattenWebSocketConfig(x *ves_io_schema_route.WebsocketConfigType) []interface{} {
	wscValue := make([]interface{}, 0)
	if x != nil {
		wscVal := map[string]interface{}{
			"use_websocket": x.GetUseWebsocket(),
		}
		wscValue = append(wscValue, wscVal)
	}
	return wscValue
}

func FlattenRegexRewrite(x *ves_io_schema.RegexMatchRewrite) []interface{} {
	rrValue := make([]interface{}, 0)
	if x != nil {
		rrVal := map[string]interface{}{
			"pattern":      x.GetPattern(),
			"substitution": x.GetSubstitution(),
		}
		rrValue = append(rrValue, rrVal)
	}
	return rrValue
}

func FlattenRetryPolicy(x *ves_io_schema.RetryPolicyType) []interface{} {
	retryPolicyValue := make([]interface{}, 0)
	if x != nil {
		retryPolicyVal := map[string]interface{}{
			"back_off":               FlattenBackOff(x.GetBackOff()),
			"num_retries":            x.GetNumRetries(),
			"per_try_timeout":        x.GetPerTryTimeout(),
			"retriable_status_codes": x.GetRetriableStatusCodes(),
			"retry_condition":        x.GetRetryCondition(),
			"retry_on":               x.GetRetryOn(),
		}
		retryPolicyValue = append(retryPolicyValue, retryPolicyVal)
	}
	return retryPolicyValue
}

func FlattenBackOff(x *ves_io_schema.RetryBackOff) []interface{} {
	backOffValue := make([]interface{}, 0)
	if x != nil {
		backOffVal := map[string]interface{}{
			"base_interval": x.GetBaseInterval(),
			"max_interval":  x.GetMaxInterval(),
		}
		backOffValue = append(backOffValue, backOffVal)
	}
	return backOffValue
}

func FlattenPercent(x *ves_io_schema.FractionalPercent) []interface{} {
	percentValue := make([]interface{}, 0)
	if x != nil {
		percentVal := map[string]interface{}{
			"denominator": x.GetDenominator().String(),
			"numerator":   x.GetNumerator(),
		}
		percentValue = append(percentValue, percentVal)
	}
	return percentValue
}

func FlattenMirrorPolicy(x *ves_io_schema_views_http_loadbalancer.MirrorPolicyType) []interface{} {
	mirrorPolicyValue := make([]interface{}, 0)
	if x != nil {
		mirrorPolicyVal := map[string]interface{}{
			"origin_pool": FlattenObjectRefTypeSet(x.GetOriginPool()),
			"percent":     FlattenPercent(x.GetPercent()),
		}
		mirrorPolicyValue = append(mirrorPolicyValue, mirrorPolicyVal)
	}
	return mirrorPolicyValue
}

func FlattenSpecificHashPolicy(x *ves_io_schema_views_http_loadbalancer.HashPolicyListType) []interface{} {
	shpValue := make([]interface{}, 0)
	if x != nil {
		shpVal := map[string]interface{}{
			"hash_policy": FlattenHashPolicy(x.GetHashPolicy()),
		}
		shpValue = append(shpValue, shpVal)
	}
	return shpValue
}

func FlattenBotDefenseJavascriptInjection(x *ves_io_schema_route.BotDefenseJavascriptInjectionType) []interface{} {
	bdjiValue := make([]interface{}, 0)
	if x != nil {
		bdjiVal := map[string]interface{}{
			"javascript_location": x.GetJavascriptLocation().String(),
			"javascript_tags":     FlattenJTags(x.GetJavascriptTags()),
		}
		bdjiValue = append(bdjiValue, bdjiVal)
	}
	return bdjiValue
}

func FlattenRedirectRoute(x *ves_io_schema_views_http_loadbalancer.RouteTypeRedirect) []interface{} {
	redirectValue := make([]interface{}, 0)
	if x != nil {
		redirectVal := map[string]interface{}{
			"headers":        FlattenHTTPHeaders(x.GetHeaders()),
			"http_method":    x.GetHttpMethod().String(),
			"incoming_port":  FlattenIncomingPort(x.GetIncomingPort()),
			"path":           FlattenPath(x.GetPath()),
			"route_redirect": FlattenRouteRedirect(x.GetRouteRedirect()),
		}
		redirectValue = append(redirectValue, redirectVal)
	}
	return redirectValue
}

func FlattenRouteRedirect(x *ves_io_schema_route.RouteRedirect) []interface{} {
	routeRValue := make([]interface{}, 0)
	if x != nil {
		routeRVal := map[string]interface{}{
			"host_redirect":      x.GetHostRedirect(),
			"port_redirect":      x.GetPortRedirect(),
			"proto_redirect":     x.GetProtoRedirect(),
			"all_params":         x.GetAllParams(),
			"remove_all_params":  isEmpty(x.GetRemoveAllParams()),
			"replace_params":     x.GetReplaceParams(),
			"retain_all_params":  isEmpty(x.GetRetainAllParams()),
			"strip_query_params": FlattenStripQueryParams(x.GetStripQueryParams()),
			"path_redirect":      x.GetPathRedirect(),
			"prefix_rewrite":     x.GetPrefixRewrite(),
			"response_code":      x.GetResponseCode(),
		}
		routeRValue = append(routeRValue, routeRVal)
	}
	return routeRValue
}

func FlattenStripQueryParams(x *ves_io_schema_route.RouteQueryParams) []interface{} {
	sqpValue := make([]interface{}, 0)
	if x != nil {
		sqpVal := map[string]interface{}{
			"query_params": x.GetQueryParams(),
		}
		sqpValue = append(sqpValue, sqpVal)
	}
	return sqpValue
}

func FlattenDirectResponseRoute(x *ves_io_schema_views_http_loadbalancer.RouteTypeDirectResponse) []interface{} {
	drrValue := make([]interface{}, 0)
	if x != nil {
		drrVal := map[string]interface{}{
			"headers":               FlattenHTTPHeaders(x.GetHeaders()),
			"http_method":           x.GetHttpMethod().String(),
			"incoming_port":         FlattenIncomingPort(x.GetIncomingPort()),
			"path":                  FlattenPath(x.GetPath()),
			"route_direct_response": FlattenRouteDirectResponse(x.GetRouteDirectResponse()),
		}
		drrValue = append(drrValue, drrVal)
	}
	return drrValue
}

func FlattenRouteDirectResponse(x *ves_io_schema_route.RouteDirectResponse) []interface{} {
	rdrValue := make([]interface{}, 0)
	if x != nil {
		rdrVal := map[string]interface{}{
			"response_body": x.GetResponseBody(),
			"response_code": x.GetResponseCode(),
		}
		rdrValue = append(rdrValue, rdrVal)
	}
	return rdrValue
}

func FlattenIncomingPort(x *ves_io_schema.PortMatcherType) []interface{} {
	ipValue := make([]interface{}, 0)
	if x != nil {
		ipVal := map[string]interface{}{
			"no_port_match": isEmpty(x.GetNoPortMatch()),
			"port":          x.GetPort(),
			"port_ranges":   x.GetPortRanges(),
		}
		ipValue = append(ipValue, ipVal)
	}
	return ipValue
}

func FlattenCustomRouteObject(x *ves_io_schema_views_http_loadbalancer.RouteTypeCustomRoute) []interface{} {
	croValue := make([]interface{}, 0)
	if x != nil {
		croVal := map[string]interface{}{
			"route_ref": FlattenObjectRefTypeSet(x.GetRouteRef()),
		}
		croValue = append(croValue, croVal)
	}
	return croValue
}

func FlattenActiveServicePolicies(x *ves_io_schema_views_common_waf.ServicePolicyList) []interface{} {
	val := make([]interface{}, 0)

	if x != nil {
		test := map[string]interface{}{
			"policies": FlattenVObjectRefTypeList(x.GetPolicies()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenSlowDDOSMitigation(x *ves_io_schema_virtual_host.SlowDDoSMitigation) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"request_headers_timeout": x.GetRequestHeadersTimeout(),
			"disable_request_timeout": isEmpty(x.GetDisableRequestTimeout()),
			"request_timeout":         x.GetRequestTimeout(),
		}
		val = append(val, test)
	}
	return val
}

func FlattenEnableTrustClientIpHeaders(x *ves_io_schema_virtual_host.ClientIPHeaders) []interface{} {
	val := make([]interface{}, 0)
	if x != nil {
		test := map[string]interface{}{
			"client_ip_headers": ConvertSliceStringToInterface(x.GetClientIpHeaders()),
		}
		val = append(val, test)
	}
	return val
}

func FlattenExcludeAttackTypeContexts(x []*ves_io_schema_policy.AppFirewallAttackTypeContext) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"exclude_attack_type": val.GetExcludeAttackType().String(),
			"context":             val.GetContext(),
			"context_name":        val.GetContextName(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenExcludeBotNameContexts(x []*ves_io_schema_policy.BotNameContext) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"bot_name": val.GetBotName(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenExcludeSignatureContexts(x []*ves_io_schema_policy.AppFirewallSignatureContext) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"context":      val.GetContext().String(),
			"context_name": val.GetContextName(),
			"signature_id": val.GetSignatureId(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenExcludeViolationContexts(x []*ves_io_schema_policy.AppFirewallViolationContext) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"context":           val.GetContext().String(),
			"context_name":      val.GetContextName(),
			"exclude_violation": val.GetExcludeViolation().String(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenWafExclusionRules(x []*ves_io_schema_policy.SimpleWafExclusionRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"any_domain":                     val.GetAnyDomain() != nil,
			"exact_value":                    val.GetExactValue(),
			"suffix_value":                   val.GetSuffixValue(),
			"expiration_timestamp":           FlattenExpirationTimestamp(val.GetExpirationTimestamp()),
			"metadata":                       FlattenMetadata(val.GetMetadata()),
			"methods":                        FlattenMethods(val.GetMethods()),
			"any_path":                       val.GetAnyPath() != nil,
			"path_prefix":                    val.GetPathPrefix(),
			"path_regex":                     val.GetPathRegex(),
			"app_firewall_detection_control": FlattenAppFirewallDetectionControl(val.GetAppFirewallDetectionControl()),
			"waf_skip_processing":            val.GetWafSkipProcessing() != nil,
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenAppFirewallDetectionControl(x *ves_io_schema_policy.AppFirewallDetectionControl) []interface{} {
	afdcValue := make([]interface{}, 0)
	if x != nil {
		afdcVal := map[string]interface{}{
			"exclude_attack_type_contexts": FlattenExcludeAttackTypeContexts(x.GetExcludeAttackTypeContexts()),
			"exclude_bot_name_contexts":    FlattenExcludeBotNameContexts(x.GetExcludeBotNameContexts()),
			"exclude_signature_contexts":   FlattenExcludeSignatureContexts(x.GetExcludeSignatureContexts()),
			"exclude_violation_contexts":   FlattenExcludeViolationContexts(x.GetExcludeViolationContexts()),
		}
		afdcValue = append(afdcValue, afdcVal)
	}
	return afdcValue
}

func isEmpty(x *ves_io_schema.Empty) bool {
	return x != nil
}

func FlattenApiRateLimitLegacy(x *ves_io_schema_views_common_waf.APIRateLimitLegacy) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"api_endpoint_rules":         FlattenARLApiEndpointRules(x.GetApiEndpointRules()),
			"bypass_rate_limiting_rules": FlattenBypassRateLimitingRules(x.GetBypassRateLimitingRules()),
			"custom_ip_allowed_list":     FlattenCustomIpAllowedList(x.GetCustomIpAllowedList()),
			"ip_allowed_list":            FlattenIpAllowedList(x.GetIpAllowedList()),
			"no_ip_allowed_list":         isEmpty(x.GetNoIpAllowedList()),
			"server_url_rules":           FlattenServerURLRules(x.GetServerUrlRules()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenBDAPMobileSdkConfig(x *ves_io_schema_views_common_security.BotAdvancedMobileSDKConfigType) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"mobile_identifier": FlattenMobileIdentifier(x.GetMobileIdentifier()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenQuery(x []*ves_io_schema_views_common_security.Query) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"name":           val.GetName(),
			"check_presence": isEmpty(val.GetCheckPresence()),
			"exact_value":    val.GetExactValue(),
			"regex_value":    val.GetRegexValue(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenRequestBody(x []*ves_io_schema_views_common_security.RequestBody) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"name":        val.GetName(),
			"exact_value": val.GetExactValue(),
			"regex_value": val.GetRegexValue(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenBDAPMProtectedAppEndpoints(x []*ves_io_schema_views_common_security.ProtectedAppEndpointType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"mobile_client":        isEmpty(val.GetMobileClient()),
			"web_client":           isEmpty(val.GetWebClient()),
			"web_mobile_client":    FlattenWebMobile(val.GetWebMobileClient()),
			"any_domain":           isEmpty(val.GetAnyDomain()),
			"domain":               FlattenDomain(val.GetDomain()),
			"flow_label":           FlattenFlowLabel(val.GetFlowLabel()),
			"undefined_flow_label": isEmpty(val.GetUndefinedFlowLabel()),
			"http_methods":         FlattenHttpMethods(val.GetHttpMethods()),
			"metadata":             FlattenMetadata(val.GetMetadata()),
			"path":                 FlattenPath(val.GetPath()),
			"query":                FlattenQuery(val.GetQuery()),
			"request_body":         FlattenRequestBody(val.GetRequestBody()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenBotDefenseAdvancedPolicy(x *ves_io_schema_views_common_security.BotDefenseAdvancedPolicyType) []interface{} {
	pValue := make([]interface{}, 0)
	if x != nil {
		pVal := map[string]interface{}{
			"js_download_path":        x.GetJsDownloadPath(),
			"disable_mobile_sdk":      isEmpty(x.GetDisableMobileSdk()),
			"mobile_sdk_config":       FlattenBDAPMobileSdkConfig(x.GetMobileSdkConfig()),
			"protected_app_endpoints": FlattenBDAPMProtectedAppEndpoints(x.GetProtectedAppEndpoints()),
		}
		pValue = append(pValue, pVal)
	}
	return pValue
}

func FlattenBody(x *ves_io_schema_views_http_loadbalancer.BodySectionMaskingOptions) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"fields": x.GetFields(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenSensitiveDataTypesInResponse(x []*ves_io_schema_views_http_loadbalancer.SensitiveDataTypes) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"body":         FlattenBody(val.GetBody()),
			"api_endpoint": FlattenApiEndPoint(val.GetApiEndpoint()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenSensitiveDataDisclosureRules(x *ves_io_schema_views_http_loadbalancer.SensitiveDataDisclosureRules) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"sensitive_data_types_in_response": FlattenSensitiveDataTypesInResponse(x.GetSensitiveDataTypesInResponse()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenSensitiveDataPolicy(x *ves_io_schema_views_common_security.SensitiveDataPolicySettings) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		mapValue := map[string]interface{}{
			"sensitive_data_policy_ref": FlattenObjectRefTypeSet(x.GetSensitiveDataPolicyRef()),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenMalwareProtectionSettings(x *ves_io_schema_views_common_security.MalwareProtectionPolicy) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"malware_protection_rules": FlattenMalwareProtectionRules(x.GetMalwareProtectionRules()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenMalwareProtectionRules(x []*ves_io_schema_views_common_security.MalwareProtectionRule) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"action":       FlattenMPRAction(val.GetAction()),
			"domain":       FlattenMPRDomain(val.GetDomain()),
			"http_methods": FlattenMethods(val.GetHttpMethods()),
			"metadata":     FlattenMetadata(val.GetMetadata()),
			"path":         FlattenPath(val.GetPath()),
			// "protocol":     val.GetProtocol(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenMPRDomain(x *ves_io_schema.DomainMatcherType) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"domain": FlattenDomain(x.GetDomain()),
		}
		rslt = append(rslt, val)
	}
	return rslt

}

func FlattenMPRAction(x *ves_io_schema.Action) []interface{} {
	actionValue := make([]interface{}, 0)
	if x != nil {
		actionVal := map[string]interface{}{
			"block":  isEmpty(x.GetBlock()),
			"report": isEmpty(x.GetReport()),
		}
		actionValue = append(actionValue, actionVal)
	}
	return actionValue
}

func FlattenL7DdosProtection(x *ves_io_schema_views_http_loadbalancer.L7DDoSProtectionSettings) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"ddos_policy_custom":      FlattenObjectRefTypeSet(x.GetDdosPolicyCustom()),
			"ddos_policy_none":        isEmpty(x.GetDdosPolicyNone()),
			"mitigation_block":        isEmpty(x.GetMitigationBlock()),
			"mitigation_js_challenge": FlattenJSChallenge(x.GetMitigationJsChallenge()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenApiKey(x *ves_io_schema_views_http_loadbalancer.ApiKey) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"key":   x.GetKey(),
			"value": FlattenPrivateKey(x.GetValue()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenBasicAuth(x *ves_io_schema_views_http_loadbalancer.BasicAuthentication) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"password": FlattenPrivateKey(x.GetPassword()),
			"user":     x.GetUser(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenBearerToken(x *ves_io_schema_views_http_loadbalancer.Bearer) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"token": FlattenPrivateKey(x.GetToken()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenLoginEndpoint(x *ves_io_schema_views_http_loadbalancer.LoginEndpoint) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"json_payload":       FlattenPrivateKey(x.GetJsonPayload()),
			"method":             x.GetMethod(),
			"path":               x.GetPath(),
			"token_response_key": x.GetTokenResponseKey(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenCredentials(x []*ves_io_schema_views_http_loadbalancer.Credentials) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		val := map[string]interface{}{
			"credential_name": v.GetCredentialName(),
			"api_key":         FlattenApiKey(v.GetApiKey()),
			"basic_auth":      FlattenBasicAuth(v.GetBasicAuth()),
			"bearer_token":    FlattenBearerToken(v.GetBearerToken()),
			"login_endpoint":  FlattenLoginEndpoint(v.GetLoginEndpoint()),
			"admin":           isEmpty(v.GetAdmin()),
			"standard":        isEmpty(v.GetStandard()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenDomainsApiTesting(x []*ves_io_schema_views_http_loadbalancer.DomainConfiguration) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		val := map[string]interface{}{
			"allow_destructive_methods": v.GetAllowDestructiveMethods(),
			"credentials":               FlattenCredentials(v.GetCredentials()),
			"domain":                    v.GetDomain(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenApiTesting(x *ves_io_schema_views_http_loadbalancer.ApiTesting) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"custom_header_value": x.GetCustomHeaderValue(),
			"domains":             FlattenDomainsApiTesting(x.GetDomains()),
			"every_day":           x.GetEveryDay(),
			"every_month":         x.GetEveryMonth(),
			"every_week":          x.GetEveryWeek(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenCdnCacheRules(x []*ves_io_schema_views.ObjectRefType) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		t := map[string]interface{}{
			"name":      v.GetName(),
			"namespace": v.GetNamespace(),
			"tenant":    v.GetTenant(),
		}
		rslt = append(rslt, t)
	}
	return rslt
}

func FlattenCustomCacheRules(x *common_cache_rule.CustomCacheRule) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"cdn_cache_rules": FlattenCdnCacheRules(x.GetCdnCacheRules()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenDefaultCacheAction(x *common_cache_rule.DefaultCacheAction) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"cache_disabled":     x.GetCacheDisabled(),
			"cache_ttl_default":  x.GetCacheTtlDefault(),
			"cache_ttl_override": x.GetCacheTtlOverride(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenCachingPolicy(x *ves_io_schema_views_http_loadbalancer.CachingPolicy) []interface{} {
	rslt := make([]interface{}, 0)
	if x != nil {
		val := map[string]interface{}{
			"custom_cache_rule":    FlattenCustomCacheRules(x.GetCustomCacheRule()),
			"default_cache_action": FlattenDefaultCacheAction(x.GetDefaultCacheAction()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func DriftDetectionSpec(d *schema.ResourceData, resp vesapi.GetObjectResponse) {
	spec := resp.GetObjSpec().(*ves_io_schema_views_http_loadbalancer.SpecType)

	d.Set("add_location", spec.GcSpec.GetAddLocation())

	d.Set("advertise_custom", FlattenAdvertiseCustom(spec.GcSpec.GetAdvertiseCustom()))

	d.Set("advertise_on_public", FlattenAdvertiseOnPublic(spec.GcSpec.GetAdvertiseOnPublic()))

	d.Set("advertise_on_public_default_vip", isEmpty(spec.GcSpec.GetAdvertiseOnPublicDefaultVip()))

	d.Set("do_not_advertise", isEmpty(spec.GcSpec.GetDoNotAdvertise()))

	d.Set("api_specification", FlattenApiSpecification(spec.GcSpec.GetApiSpecification()))

	d.Set("disable_api_definition", isEmpty(spec.GcSpec.GetDisableApiDefinition()))

	d.Set("disable_api_discovery", isEmpty(spec.GcSpec.GetDisableApiDiscovery()))

	d.Set("enable_api_discovery", FlattenEnableApiDiscovery(spec.GcSpec.GetEnableApiDiscovery()))

	d.Set("api_protection_rules", FlattenApiProtectionRules(spec.GcSpec.GetApiProtectionRules()))

	d.Set("api_testing", FlattenApiTesting(spec.GcSpec.GetApiTesting()))

	d.Set("disable_api_testing", isEmpty(spec.GcSpec.GetDisableApiTesting()))

	d.Set("blocked_clients", FlattenBlockedClients(spec.GcSpec.GetBlockedClients()))

	d.Set("bot_defense", FlattenBotDefense(spec.GcSpec.GetBotDefense()))

	d.Set("disable_bot_defense", isEmpty(spec.GcSpec.GetDisableBotDefense()))

	d.Set("captcha_challenge", FlattenCaptchaChallenge(spec.GcSpec.GetCaptchaChallenge()))

	d.Set("enable_challenge", FlattenEnableChallenge(spec.GcSpec.GetEnableChallenge()))

	d.Set("js_challenge", FlattenJSChallenge(spec.GcSpec.GetJsChallenge()))

	d.Set("no_challenge", isEmpty(spec.GcSpec.GetNoChallenge()))

	d.Set("policy_based_challenge", FlattenPolicyBasedChallenge(spec.GcSpec.GetPolicyBasedChallenge()))

	d.Set("client_side_defense", FlattenClientSideDefense(spec.GcSpec.GetClientSideDefense()))

	d.Set("disable_client_side_defense", spec.GcSpec.GetDisableClientSideDefense() != nil)

	d.Set("cors_policy", FlattenCorsPolicy(spec.GcSpec.GetCorsPolicy()))

	d.Set("csrf_policy", FlattenCsrfPolicy(spec.GcSpec.GetCsrfPolicy()))

	d.Set("data_guard_rules", FlattenDataGuardRules(spec.GcSpec.GetDataGuardRules()))

	d.Set("ddos_mitigation_rules", FlattenDDOSMitigationRules(spec.GcSpec.GetDdosMitigationRules()))

	d.Set("default_route_pools", FlattenDefaultRoutePools(spec.GcSpec.GetDefaultRoutePools()))

	d.Set("domains", spec.GcSpec.GetDomains())

	d.Set("graphql_rules", FlattenGraphRules(spec.GcSpec.GetGraphqlRules()))

	d.Set("cookie_stickiness", FlattenCookie(spec.GcSpec.GetCookieStickiness()))

	d.Set("least_active", isEmpty(spec.GcSpec.GetLeastActive()))

	d.Set("random", isEmpty(spec.GcSpec.GetRandom()))

	d.Set("ring_hash", FlattenSpecificHashPolicy(spec.GcSpec.GetRingHash()))

	d.Set("round_robin", isEmpty(spec.GcSpec.GetRoundRobin()))

	d.Set("source_ip_stickiness", isEmpty(spec.GcSpec.GetSourceIpStickiness()))

	d.Set("disable_ip_reputation", isEmpty(spec.GcSpec.GetDisableIpReputation()))

	d.Set("enable_ip_reputation", FlattenEnableIpReputation(spec.GcSpec.GetEnableIpReputation()))

	d.Set("jwt_validation", FlattenJwtValidation(spec.GcSpec.GetJwtValidation()))

	d.Set("l7_ddos_protection", FlattenL7DdosProtection(spec.GcSpec.GetL7DdosProtection()))

	d.Set("http", FlattenHttp(spec.GcSpec.GetHttp()))

	d.Set("https", FlattenHttps(spec.GcSpec.GetHttps()))

	d.Set("https_auto_cert", FlattenHttpsAutoCert(spec.GcSpec.GetHttpsAutoCert()))

	d.Set("disable_malicious_user_detection", isEmpty(spec.GcSpec.GetDisableMaliciousUserDetection()))

	d.Set("enable_malicious_user_detection", isEmpty(spec.GcSpec.GetEnableMaliciousUserDetection()))

	d.Set("disable_malware_protection", isEmpty(spec.GcSpec.GetDisableMalwareProtection()))

	d.Set("malware_protection_settings", FlattenMalwareProtectionSettings(spec.GcSpec.GetMalwareProtectionSettings()))

	d.Set("more_option", FlattenMoreOption(spec.GcSpec.GetMoreOption()))

	d.Set("origin_server_subset_rule_list", FlattenOSSRL(spec.GcSpec.GetOriginServerSubsetRuleList()))

	d.Set("protected_cookies", FlattenCookiesToModify(spec.GcSpec.GetProtectedCookies()))

	d.Set("api_rate_limit", FlattenApiRateLimit(spec.GcSpec.GetApiRateLimit()))

	d.Set("disable_rate_limit", isEmpty(spec.GcSpec.GetDisableRateLimit()))

	d.Set("rate_limit", FlattenRateLimit(spec.GcSpec.GetRateLimit()))

	d.Set("routes", FlattenRoutes(spec.GcSpec.GetRoutes()))

	d.Set("sensitive_data_disclosure_rules", FlattenSensitiveDataDisclosureRules(spec.GcSpec.GetSensitiveDataDisclosureRules()))

	d.Set("default_sensitive_data_policy", isEmpty(spec.GcSpec.GetDefaultSensitiveDataPolicy()))

	d.Set("sensitive_data_policy", FlattenSensitiveDataPolicy(spec.GcSpec.GetSensitiveDataPolicy()))

	d.Set("active_service_policies", FlattenActiveServicePolicies(spec.GcSpec.GetActiveServicePolicies()))

	d.Set("no_service_policies", isEmpty(spec.GcSpec.GetNoServicePolicies()))

	d.Set("service_policies_from_namespace", isEmpty(spec.GcSpec.GetServicePoliciesFromNamespace()))

	d.Set("slow_ddos_mitigation", FlattenSlowDDOSMitigation(spec.GcSpec.GetSlowDdosMitigation()))

	d.Set("system_default_timeouts", isEmpty(spec.GcSpec.GetSystemDefaultTimeouts()))

	d.Set("disable_threat_mesh", isEmpty(spec.GcSpec.GetDisableThreatMesh()))

	d.Set("enable_threat_mesh", isEmpty(spec.GcSpec.GetEnableThreatMesh()))

	d.Set("disable_trust_client_ip_headers", isEmpty(spec.GcSpec.GetDisableTrustClientIpHeaders()))

	d.Set("enable_trust_client_ip_headers", FlattenEnableTrustClientIpHeaders(spec.GcSpec.GetEnableTrustClientIpHeaders()))

	d.Set("trusted_clients", FlattenBlockedClients(spec.GcSpec.GetTrustedClients()))

	d.Set("user_id_client_ip", isEmpty(spec.GcSpec.GetUserIdClientIp()))

	d.Set("user_identification", FlattenObjectRefTypeSet(spec.GcSpec.GetUserIdentification()))

	d.Set("app_firewall", FlattenObjectRefTypeSet(spec.GcSpec.GetAppFirewall()))

	d.Set("disable_waf", isEmpty(spec.GcSpec.GetDisableWaf()))

	d.Set("waf_exclusion_rules", FlattenWafExclusionRules(spec.GcSpec.GetWafExclusionRules()))

	d.Set("caching_policy", FlattenCachingPolicy(spec.GcSpec.GetCachingPolicy()))

	d.Set("disable_caching", isEmpty(spec.GcSpec.GetDisableCaching()))
}
