// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package dns_load_balancer

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.dns_load_balancer.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.dns_load_balancer.Object"] = ObjectValidator()
	vr["ves.io.schema.dns_load_balancer.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.dns_load_balancer.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.dns_load_balancer.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.dns_load_balancer.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.dns_load_balancer.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.dns_load_balancer.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.dns_load_balancer.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.dns_load_balancer.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.dns_load_balancer.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.dns_load_balancer.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.dns_load_balancer.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.dns_load_balancer.DNSLBHealthStatusListRequest"] = DNSLBHealthStatusListRequestValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBHealthStatusListResponse"] = DNSLBHealthStatusListResponseValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBHealthStatusListResponseItem"] = DNSLBHealthStatusListResponseItemValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBHealthStatusRequest"] = DNSLBHealthStatusRequestValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBHealthStatusResponse"] = DNSLBHealthStatusResponseValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBPoolHealthStatusListResponseItem"] = DNSLBPoolHealthStatusListResponseItemValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBPoolHealthStatusRequest"] = DNSLBPoolHealthStatusRequestValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBPoolHealthStatusResponse"] = DNSLBPoolHealthStatusResponseValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBPoolMemberHealthStatusEvent"] = DNSLBPoolMemberHealthStatusEventValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBPoolMemberHealthStatusListRequest"] = DNSLBPoolMemberHealthStatusListRequestValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBPoolMemberHealthStatusListResponse"] = DNSLBPoolMemberHealthStatusListResponseValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBPoolMemberHealthStatusListResponseItem"] = DNSLBPoolMemberHealthStatusListResponseItemValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBPoolMemberHealthStatusRequest"] = DNSLBPoolMemberHealthStatusRequestValidator()
	vr["ves.io.schema.dns_load_balancer.DNSLBPoolMemberHealthStatusResponse"] = DNSLBPoolMemberHealthStatusResponseValidator()
	vr["ves.io.schema.dns_load_balancer.HealthStatusSummary"] = HealthStatusSummaryValidator()

	vr["ves.io.schema.dns_load_balancer.SuggestValuesReq"] = SuggestValuesReqValidator()
	vr["ves.io.schema.dns_load_balancer.SuggestValuesResp"] = SuggestValuesRespValidator()
	vr["ves.io.schema.dns_load_balancer.SuggestedItem"] = SuggestedItemValidator()

	vr["ves.io.schema.dns_load_balancer.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.dns_load_balancer.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.dns_load_balancer.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.dns_load_balancer.LoadBalancingRule"] = LoadBalancingRuleValidator()
	vr["ves.io.schema.dns_load_balancer.LoadBalancingRuleList"] = LoadBalancingRuleListValidator()
	vr["ves.io.schema.dns_load_balancer.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.dns_load_balancer.ResponseCache"] = ResponseCacheValidator()
	vr["ves.io.schema.dns_load_balancer.ResponseCacheParameters"] = ResponseCacheParametersValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.dns_load_balancer.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.dns_load_balancer.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dns_load_balancer.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.dns_load_balancer.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.dns_load_balancer.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.dns_load_balancer.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dns_load_balancer.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.dns_load_balancer.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.dns_load_balancer.API.Create"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.dns_load_balancer.CreateRequest.spec.rule_list.rules.client_choice.ip_prefix_list.ipv6_prefixes",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.dns_load_balancer.API.Create"] = []string{
		"spec.rule_list.rules.#.nxdomain",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.dns_load_balancer.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.rule_list.rules.#.ip_prefix_list.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.dns_load_balancer.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.rule_list.rules.#.ip_prefix_list.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.dns_load_balancer.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.rule_list.rules.#.ip_prefix_list.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.rule_list.rules.#.ip_prefix_list.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.rule_list.rules.#.ip_prefix_list.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.dns_load_balancer.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.rule_list.rules.#.ip_prefix_list.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.dns_load_balancer.API.Replace"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.dns_load_balancer.ReplaceRequest.spec.rule_list.rules.client_choice.ip_prefix_list.ipv6_prefixes",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.dns_load_balancer.API.Replace"] = []string{
		"spec.rule_list.rules.#.nxdomain",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.dns_load_balancer.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.rule_list.rules.#.ip_prefix_list.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.dns_load_balancer.API"] = "config/dns"
	sm["ves.io.schema.dns_load_balancer.CustomDataAPI"] = "data"
	sm["ves.io.schema.dns_load_balancer.CustomAPI"] = "config/dns"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

	sm["dns"] = svcfw.P0PolicyInfo{
		Name:            "ves-io-allow-config-dns",
		ServiceSelector: "bifrost\\.gc.*\\",
	}

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.dns_load_balancer.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.dns_load_balancer.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.dns_load_balancer.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.dns_load_balancer.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.dns_load_balancer.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.dns_load_balancer.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.dns_load_balancer.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.dns_load_balancer.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.dns_load_balancer.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.dns_load_balancer.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.dns_load_balancer.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.dns_load_balancer.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.dns_load_balancer.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomDataAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.dns_load_balancer.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.dns_load_balancer.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.dns_load_balancer.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.dns_load_balancer.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.dns_load_balancer.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.dns_load_balancer.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomAPIServer(svc)
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
