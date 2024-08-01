package driftdetection

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policy"
	ves_io_schema_service_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy"
	ves_io_schema_service_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy_rule"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	"gopkg.volterra.us/stdlib/client/vesapi"
)

func FlattenPrefixList(f *ves_io_schema_views.PrefixStringListType) []interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"prefixes": f.GetPrefixes(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenTlsFingerPrintClasses(f []ves_io_schema_policy.KnownTlsFingerprintClass) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range f {
		rslt = append(rslt, val.String())
	}
	return rslt
}

func FlattenAllowList(f *ves_io_schema_service_policy.SourceList) []interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"asn_list":                   FlattenAsnList(f.GetAsnList()),
			"asn_set":                    FlattenVObjectRefTypeList(f.GetAsnSet()),
			"country_list":               FlattenCountryList(f.GetCountryList()),
			"default_action_allow":       isEmpty(f.GetDefaultActionAllow()),
			"default_action_deny":        isEmpty(f.GetDefaultActionDeny()),
			"default_action_next_policy": isEmpty(f.GetDefaultActionNextPolicy()),
			"ip_prefix_set":              FlattenVObjectRefTypeList(f.GetIpPrefixSet()),
			"prefix_list":                FlattenPrefixList(f.GetPrefixList()),
			"tls_fingerprint_classes":    FlattenTlsFingerPrintClasses(f.GetTlsFingerprintClasses()),
			"tls_fingerprint_values":     f.GetTlsFingerprintValues(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenPortMatcher(f *ves_io_schema_policy.PortMatcherType) []interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"ports":          f.GetPorts(),
			"invert_matcher": f.GetInvertMatcher(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenRulesServicePolicy(f []*ves_io_schema.ObjectRefType) interface{} {
	rstl := make([]interface{}, 0)
	for _, fval := range f {
		val := map[string]interface{}{
			"kind":      fval.GetKind(),
			"name":      fval.GetName(),
			"namespace": fval.GetNamespace(),
			"tenant":    fval.GetTenant(),
		}
		rstl = append(rstl, val)
	}
	return rstl
}

func FlattenLegacyRuleList(f *ves_io_schema_service_policy.LegacyRuleList) []interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"rules": FlattenRulesServicePolicy(f.GetRules()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenApiGroupMatcher(f *ves_io_schema_policy.StringMatcherType) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"invert_matcher": f.GetInvertMatcher(),
			"match":          f.GetMatch(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenBotAction(f *ves_io_schema_policy.BotAction) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"bot_skip_processing": isEmpty(f.GetBotSkipProcessing()),
			"none":                isEmpty(f.GetNone()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenIpThreatCategoryListServicePolicy(f *ves_io_schema_service_policy_rule.IPThreatCategoryListType) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"ip_threat_categories": FlattenIpThreatCategories(f.GetIpThreatCategories()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenClientRole(f *ves_io_schema_policy.RoleMatcherType) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"match": f.GetMatch(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenContentRewriteAction(f *ves_io_schema_policy.ContentRewriteAction) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"element_selector": f.GetElementSelector(),
			"insert_content":   f.GetInsertContent(),
			"position":         f.GetPosition().String(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenDstAsnList(f *ves_io_schema_policy.AsnMatchList) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"as_numbers": FlattenAsNumbers(f.GetAsNumbers()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenL4Dests(f []*ves_io_schema.L4DestType) interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range f {
		t := map[string]interface{}{
			"port_ranges": val.GetPortRanges(),
			"prefixes":    val.GetPrefixes(),
		}
		rslt = append(rslt, t)
	}
	return rslt
}

func FlattenL4DestMatcher(f *ves_io_schema_policy.L4DestMatcherType) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"invert_matcher": f.GetInvertMatcher(),
			"l4_dests":       FlattenL4Dests(f.GetL4Dests()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenLabelMatcher(f *ves_io_schema.LabelMatcherType) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"keys": f.GetKeys(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenPathServicePolicy(f *ves_io_schema_policy.PathMatcherType) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"exact_values":  f.GetExactValues(),
			"prefix_values": f.GetPrefixValues(),
			"regex_values":  f.GetRegexValues(),
			"suffix_values": f.GetSuffixValues(),
			"transformers":  FlattenTransformers(f.GetTransformers()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenServerSelector(f *ves_io_schema.LabelSelectorType) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"expressions": f.GetExpressions(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenShapeProtectedEndpointAction(f *ves_io_schema_policy.ShapeProtectedEndpointAction) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"app_traffic_type": f.GetAppTrafficType().String(),
			"mitigation":       FlattenSPMitigation(f.GetMitigation()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenUrlItems(f []*ves_io_schema_policy.URLItem) interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range f {
		t := map[string]interface{}{
			"domain_regex": val.GetDomainRegex(),
			"domain_value": val.GetDomainValue(),
			"path_prefix":  val.GetPathPrefix(),
			"path_regex":   val.GetPathRegex(),
			"path_value":   val.GetPathValue(),
		}
		rslt = append(rslt, t)
	}
	return rslt
}

func FlattenUrlMatcher(f *ves_io_schema_policy.URLMatcherType) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"invert_matcher": f.GetInvertMatcher(),
			"url_items":      FlattenUrlItems(f.GetUrlItems()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenWafAction(f *ves_io_schema_policy.WafAction) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"app_firewall_detection_control": FlattenSPAppFirewallDetectionControl(f.GetAppFirewallDetectionControl()),
			"none":                           isEmpty(f.GetNone()),
			"waf_in_monitoring_mode":         isEmpty(f.GetWafInMonitoringMode()),
			"waf_skip_processing":            isEmpty(f.GetWafSkipProcessing()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenSpecServicePolicy(f *ves_io_schema_service_policy_rule.GlobalSpecType) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"action":                          f.GetAction().String(),
			"api_group_matcher":               FlattenApiGroupMatcher(f.GetApiGroupMatcher()),
			"arg_matchers":                    FlattenArgMatchers(f.GetArgMatchers()),
			"any_asn":                         f.GetAnyAsn() != nil,
			"asn_list":                        FlattenAsnList(f.GetAsnList()),
			"asn_matcher":                     FlattenAsnMatcher(f.GetAsnMatcher()),
			"body_matcher":                    FlattenItem(f.GetBodyMatcher()),
			"bot_action":                      FlattenBotAction(f.GetBotAction()),
			"challenge_action":                f.GetChallengeAction().String(),
			"any_client":                      f.GetAnyClient() != nil,
			"client_name":                     f.GetClientName(),
			"client_name_matcher":             FlattenItem(f.GetClientNameMatcher()),
			"client_selector":                 FlattenServerSelector(f.GetClientSelector()),
			"ip_threat_category_list":         FlattenIpThreatCategoryListServicePolicy(f.GetIpThreatCategoryList()),
			"client_role":                     FlattenClientRole(f.GetClientRole()),
			"content_rewrite_action":          FlattenContentRewriteAction(f.GetContentRewriteAction()),
			"cookie_matchers":                 FlattenCookieMatchers(f.GetCookieMatchers()),
			"domain_matcher":                  FlattenItem(f.GetDomainMatcher()),
			"any_dst_asn":                     isEmpty(f.GetAnyDstAsn()),
			"dst_asn_list":                    FlattenDstAsnList(f.GetDstAsnList()),
			"dst_asn_matcher":                 FlattenAsnMatcher(f.GetDstAsnMatcher()),
			"any_dst_ip":                      f.GetAnyDstIp() != nil,
			"dst_ip_matcher":                  FlattenIpMatcher(f.GetDstIpMatcher()),
			"dst_ip_prefix_list":              FlattenSPIpPrefixList(f.GetDstIpPrefixList()),
			"expiration_timestamp":            FlattenExpirationTimestamp(f.GetExpirationTimestamp()),
			"forwarding_class":                FlattenObjectRefTypeList(f.GetForwardingClass()),
			"goto_policy":                     FlattenObjectRefTypeList(f.GetGotoPolicy()),
			"headers":                         FlattenHeaders(f.GetHeaders()),
			"http_method":                     FlattenHttpMethod(f.GetHttpMethod()),
			"any_ip":                          f.GetAnyIp() != nil,
			"ip_matcher":                      FlattenIpMatcher(f.GetIpMatcher()),
			"ip_prefix_list":                  FlattenSPIpPrefixList(f.GetIpPrefixList()),
			"l4_dest_matcher":                 FlattenL4DestMatcher(f.GetL4DestMatcher()),
			"label_matcher":                   FlattenLabelMatcher(f.GetLabelMatcher()),
			"path":                            FlattenPathServicePolicy(f.GetPath()),
			"port_matcher":                    FlattenPortMatcher(f.GetPortMatcher()),
			"query_params":                    FlattenQueryParams(f.GetQueryParams()),
			"rate_limiter":                    FlattenObjectRefTypeList(f.GetRateLimiter()),
			"scheme":                          f.GetScheme(),
			"server_selector":                 FlattenServerSelector(f.GetServerSelector()),
			"shape_protected_endpoint_action": FlattenShapeProtectedEndpointAction(f.GetShapeProtectedEndpointAction()),
			"tls_fingerprint_matcher":         FlattenTlsFingerPrintMatcher(f.GetTlsFingerprintMatcher()),
			"url_matcher":                     FlattenUrlMatcher(f.GetUrlMatcher()),
			"virtual_host_matcher":            FlattenItem(f.GetVirtualHostMatcher()),
			"waf_action":                      FlattenWafAction(f.GetWafAction()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenRulesRuleList(f []*ves_io_schema_service_policy.Rule) interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range f {
		temp := map[string]interface{}{
			"metadata": FlattenMetadata(val.GetMetadata()),
			"spec":     FlattenSpecServicePolicy(val.GetSpec()),
		}
		rslt = append(rslt, temp)
	}
	return rslt
}

func FlattenRuleListServicePolicy(f *ves_io_schema_service_policy.RuleList) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"rules": FlattenRulesRuleList(f.GetRules()),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenServerNameMatcher(f *ves_io_schema_policy.MatcherTypeBasic) interface{} {
	rslt := make([]interface{}, 0)
	if f != nil {
		val := map[string]interface{}{
			"exact_values": f.GetExactValues(),
			"regex_values": f.GetRegexValues(),
		}
		rslt = append(rslt, val)
	}
	return rslt
}

func FlattenSPIpPrefixList(x *ves_io_schema_policy.PrefixMatchList) []interface{} {
	ipPLValue := make([]interface{}, 0)
	if x != nil {
		ipPLVal := map[string]interface{}{
			"invert_match": x.GetInvertMatch(),
			"ip_prefixes":  x.GetIpPrefixes(),
		}
		ipPLValue = append(ipPLValue, ipPLVal)
	}
	return ipPLValue
}

func FlattenSPMitigation(x *ves_io_schema_policy.ShapeBotMitigationAction) []interface{} {
	mValue := make([]interface{}, 0)
	if x != nil {
		mVal := map[string]interface{}{
			"block":    FlattenSPBlock(x.GetBlock()),
			"none":     isEmpty(x.GetNone()),
			"redirect": FlattenRedirect(x.GetRedirect()),
		}
		mValue = append(mValue, mVal)
	}
	return mValue
}
func FlattenSPBlock(x *ves_io_schema_policy.ShapeBotBlockMitigationActionType) []interface{} {
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

func FlattenSPAppFirewallDetectionControl(x *ves_io_schema_policy.AppFirewallDetectionControl) []interface{} {
	afdcValue := make([]interface{}, 0)
	if x != nil {
		afdcVal := map[string]interface{}{
			"exclude_attack_type_contexts": FlattenSPExcludeAttackTypeContexts(x.GetExcludeAttackTypeContexts()),
			"exclude_signature_contexts":   FlattenSPExcludeSignatureContexts(x.GetExcludeSignatureContexts()),
			"exclude_violation_contexts":   FlattenSPExcludeViolationContexts(x.GetExcludeViolationContexts()),
		}
		afdcValue = append(afdcValue, afdcVal)
	}
	return afdcValue
}

func FlattenSPExcludeAttackTypeContexts(x []*ves_io_schema_policy.AppFirewallAttackTypeContext) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"exclude_attack_type": val.GetExcludeAttackType().String(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}
func FlattenSPExcludeSignatureContexts(x []*ves_io_schema_policy.AppFirewallSignatureContext) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"signature_id": val.GetSignatureId(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func FlattenSPExcludeViolationContexts(x []*ves_io_schema_policy.AppFirewallViolationContext) []interface{} {
	rslt := make([]interface{}, 0)
	for _, val := range x {
		mapValue := map[string]interface{}{
			"exclude_violation": val.GetExcludeViolation().String(),
		}
		rslt = append(rslt, mapValue)
	}
	return rslt
}

func DriftDetectionSpecServicePolicy(d *schema.ResourceData, resp vesapi.GetObjectResponse) {
	spec := resp.GetObjSpec().(*ves_io_schema_service_policy.SpecType)

	d.Set("algo", spec.GcSpec.GetAlgo().String())
	d.Set("port_matcher", FlattenPortMatcher(spec.GcSpec.GetPortMatcher()))
	d.Set("allow_all_requests", isEmpty(spec.GcSpec.GetAllowAllRequests()))
	d.Set("allow_list", FlattenAllowList(spec.GcSpec.GetAllowList()))
	d.Set("deny_all_requests", isEmpty(spec.GcSpec.GetDenyAllRequests()))
	d.Set("deny_list", FlattenAllowList(spec.GcSpec.GetDenyList()))
	d.Set("internally_generated", isEmpty(spec.GcSpec.GetInternallyGenerated()))
	d.Set("legacy_rule_list", FlattenLegacyRuleList(spec.GcSpec.GetLegacyRuleList()))
	d.Set("rule_list", FlattenRuleListServicePolicy(spec.GcSpec.GetRuleList()))
	d.Set("any_server", isEmpty(spec.GcSpec.GetAnyServer()))
	d.Set("server_name", spec.GcSpec.GetServerName())
	d.Set("server_name_matcher", FlattenServerNameMatcher(spec.GcSpec.GetServerNameMatcher()))
	d.Set("server_selector", FlattenServerSelector(spec.GetGcSpec().GetServerSelector()))
}
