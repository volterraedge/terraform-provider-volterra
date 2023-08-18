// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package policy

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.policy.AppFirewallAttackTypeContext"] = AppFirewallAttackTypeContextValidator()
	vr["ves.io.schema.policy.AppFirewallDetectionControl"] = AppFirewallDetectionControlValidator()
	vr["ves.io.schema.policy.AppFirewallSignatureContext"] = AppFirewallSignatureContextValidator()
	vr["ves.io.schema.policy.AppFirewallViolationContext"] = AppFirewallViolationContextValidator()
	vr["ves.io.schema.policy.ArgMatcherType"] = ArgMatcherTypeValidator()
	vr["ves.io.schema.policy.AsnMatchList"] = AsnMatchListValidator()
	vr["ves.io.schema.policy.AsnMatcherType"] = AsnMatcherTypeValidator()
	vr["ves.io.schema.policy.BotAction"] = BotActionValidator()
	vr["ves.io.schema.policy.BotNameContext"] = BotNameContextValidator()
	vr["ves.io.schema.policy.ClientMatcher"] = ClientMatcherValidator()
	vr["ves.io.schema.policy.ContentRewriteAction"] = ContentRewriteActionValidator()
	vr["ves.io.schema.policy.CookieMatcherType"] = CookieMatcherTypeValidator()
	vr["ves.io.schema.policy.CountryCodeList"] = CountryCodeListValidator()
	vr["ves.io.schema.policy.DataGuardControl"] = DataGuardControlValidator()
	vr["ves.io.schema.policy.DenyInformation"] = DenyInformationValidator()
	vr["ves.io.schema.policy.GraphQLRule"] = GraphQLRuleValidator()
	vr["ves.io.schema.policy.GraphQLSettingsType"] = GraphQLSettingsTypeValidator()
	vr["ves.io.schema.policy.HeaderMatcherType"] = HeaderMatcherTypeValidator()
	vr["ves.io.schema.policy.HeaderMatcherTypeBasic"] = HeaderMatcherTypeBasicValidator()
	vr["ves.io.schema.policy.HttpCookieName"] = HttpCookieNameValidator()
	vr["ves.io.schema.policy.HttpHeaderName"] = HttpHeaderNameValidator()
	vr["ves.io.schema.policy.HttpMethodMatcherType"] = HttpMethodMatcherTypeValidator()
	vr["ves.io.schema.policy.HttpQueryParameterName"] = HttpQueryParameterNameValidator()
	vr["ves.io.schema.policy.IPThreatCategoryListType"] = IPThreatCategoryListTypeValidator()
	vr["ves.io.schema.policy.IpMatcherType"] = IpMatcherTypeValidator()
	vr["ves.io.schema.policy.JwtTokenAuthOptions"] = JwtTokenAuthOptionsValidator()
	vr["ves.io.schema.policy.L4DestMatcherType"] = L4DestMatcherTypeValidator()
	vr["ves.io.schema.policy.MatcherType"] = MatcherTypeValidator()
	vr["ves.io.schema.policy.MatcherTypeBasic"] = MatcherTypeBasicValidator()
	vr["ves.io.schema.policy.ModifyAction"] = ModifyActionValidator()
	vr["ves.io.schema.policy.OpenApiValidationAction"] = OpenApiValidationActionValidator()
	vr["ves.io.schema.policy.OriginServerSubsetRule"] = OriginServerSubsetRuleValidator()
	vr["ves.io.schema.policy.PathMatcherType"] = PathMatcherTypeValidator()
	vr["ves.io.schema.policy.PortMatcherType"] = PortMatcherTypeValidator()
	vr["ves.io.schema.policy.PrefixMatchList"] = PrefixMatchListValidator()
	vr["ves.io.schema.policy.QueryParameterMatcherType"] = QueryParameterMatcherTypeValidator()
	vr["ves.io.schema.policy.RequestConstraintType"] = RequestConstraintTypeValidator()
	vr["ves.io.schema.policy.RequestMatcher"] = RequestMatcherValidator()
	vr["ves.io.schema.policy.RoleMatcherType"] = RoleMatcherTypeValidator()
	vr["ves.io.schema.policy.ShapeBotBlockMitigationActionType"] = ShapeBotBlockMitigationActionTypeValidator()
	vr["ves.io.schema.policy.ShapeBotFlagMitigationActionChoiceType"] = ShapeBotFlagMitigationActionChoiceTypeValidator()
	vr["ves.io.schema.policy.ShapeBotFlagMitigationActionType"] = ShapeBotFlagMitigationActionTypeValidator()
	vr["ves.io.schema.policy.ShapeBotMitigationAction"] = ShapeBotMitigationActionValidator()
	vr["ves.io.schema.policy.ShapeBotRedirectMitigationActionType"] = ShapeBotRedirectMitigationActionTypeValidator()
	vr["ves.io.schema.policy.ShapeProtectedEndpointAction"] = ShapeProtectedEndpointActionValidator()
	vr["ves.io.schema.policy.SimpleDataGuardRule"] = SimpleDataGuardRuleValidator()
	vr["ves.io.schema.policy.SimpleWafExclusionRule"] = SimpleWafExclusionRuleValidator()
	vr["ves.io.schema.policy.StringMatcherType"] = StringMatcherTypeValidator()
	vr["ves.io.schema.policy.TlsFingerprintMatcherType"] = TlsFingerprintMatcherTypeValidator()
	vr["ves.io.schema.policy.URLItem"] = URLItemValidator()
	vr["ves.io.schema.policy.URLMatcherType"] = URLMatcherTypeValidator()
	vr["ves.io.schema.policy.WafAction"] = WafActionValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

}

func InitializeMDRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	initializeEntryRegistry(mdr)
	initializeValidatorRegistry(mdr.ValidatorRegistry)

	initializeCRUDServiceRegistry(mdr, isExternal)
	if isExternal {
		return
	}

	initializeRPCRegistry(mdr)
	initializeAPIGwServiceSlugsRegistry(mdr.APIGwServiceSlugs)
	initializeP0PolicyRegistry(mdr.P0PolicyRegistry)

}