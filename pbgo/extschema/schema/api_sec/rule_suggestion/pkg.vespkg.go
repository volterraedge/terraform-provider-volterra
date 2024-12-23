// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package rule_suggestion

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.api_sec.rule_suggestion.GetSuggestedAPIEndpointProtectionRuleReq"] = GetSuggestedAPIEndpointProtectionRuleReqValidator()
	vr["ves.io.schema.api_sec.rule_suggestion.GetSuggestedAPIEndpointProtectionRuleRsp"] = GetSuggestedAPIEndpointProtectionRuleRspValidator()
	vr["ves.io.schema.api_sec.rule_suggestion.GetSuggestedOasValidationRuleReq"] = GetSuggestedOasValidationRuleReqValidator()
	vr["ves.io.schema.api_sec.rule_suggestion.GetSuggestedOasValidationRuleRsp"] = GetSuggestedOasValidationRuleRspValidator()
	vr["ves.io.schema.api_sec.rule_suggestion.GetSuggestedRateLimitRuleReq"] = GetSuggestedRateLimitRuleReqValidator()
	vr["ves.io.schema.api_sec.rule_suggestion.GetSuggestedRateLimitRuleRsp"] = GetSuggestedRateLimitRuleRspValidator()
	vr["ves.io.schema.api_sec.rule_suggestion.GetSuggestedSensitiveDataRuleReq"] = GetSuggestedSensitiveDataRuleReqValidator()
	vr["ves.io.schema.api_sec.rule_suggestion.GetSuggestedSensitiveDataRuleRsp"] = GetSuggestedSensitiveDataRuleRspValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.api_sec.rule_suggestion.RuleSuggestionAPI.GetSuggestedAPIEndpointProtectionRule"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "rule.client_matcher.ip_prefix_list.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.api_sec.rule_suggestion.RuleSuggestionAPI.GetSuggestedRateLimitRule"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "rule.client_matcher.ip_prefix_list.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.api_sec.rule_suggestion.RuleSuggestionAPI"] = "config"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {
	sm["config"] = svcfw.P0PolicyInfo{
		Name:            "ves-io-allow-config",
		ServiceSelector: "akar\\.gc.*\\",
	}

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.api_sec.rule_suggestion.Object"] = RuleSuggestionAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.api_sec.rule_suggestion.RuleSuggestionAPI"] = NewRuleSuggestionAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.api_sec.rule_suggestion.RuleSuggestionAPI"] = NewRuleSuggestionAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.api_sec.rule_suggestion.RuleSuggestionAPI"] = RegisterRuleSuggestionAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.api_sec.rule_suggestion.RuleSuggestionAPI"] = RegisterGwRuleSuggestionAPIHandler
		customCSR.ServerRegistry["ves.io.schema.api_sec.rule_suggestion.RuleSuggestionAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewRuleSuggestionAPIServer(svc)
		}

	}()

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
