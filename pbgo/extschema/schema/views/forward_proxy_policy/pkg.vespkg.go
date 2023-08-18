// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package forward_proxy_policy

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.forward_proxy_policy.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.forward_proxy_policy.Object"] = ObjectValidator()
	vr["ves.io.schema.views.forward_proxy_policy.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.forward_proxy_policy.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.forward_proxy_policy.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.forward_proxy_policy.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.forward_proxy_policy.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.forward_proxy_policy.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHits"] = ForwardProxyPolicyHitsValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsId"] = ForwardProxyPolicyHitsIdValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsRequest"] = ForwardProxyPolicyHitsRequestValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsResponse"] = ForwardProxyPolicyHitsResponseValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyMetricLabelFilter"] = ForwardProxyPolicyMetricLabelFilterValidator()

	vr["ves.io.schema.views.forward_proxy_policy.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.DomainListType"] = DomainListTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ForwardProxyAdvancedRuleType"] = ForwardProxyAdvancedRuleTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ForwardProxyRuleListType"] = ForwardProxyRuleListTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ForwardProxySimpleRuleType"] = ForwardProxySimpleRuleTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.URLCategoryListType"] = URLCategoryListTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.URLListType"] = URLListTypeValidator()
	vr["ves.io.schema.views.forward_proxy_policy.URLType"] = URLTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.forward_proxy_policy.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.forward_proxy_policy.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.forward_proxy_policy.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.forward_proxy_policy.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.forward_proxy_policy.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.forward_proxy_policy.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.forward_proxy_policy.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.forward_proxy_policy.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.forward_proxy_policy.API.Create"] = []string{
		"spec.rule_list.rules.#.dst_prefix_list.ipv6_prefixes.#",
		"spec.rule_list.rules.#.inside_sources",
		"spec.rule_list.rules.#.interface",
		"spec.rule_list.rules.#.metadata.disable",
		"spec.rule_list.rules.#.namespace",
		"spec.rule_list.rules.#.prefix_list.ipv6_prefixes.#",
		"spec.rule_list.rules.#.rule_description",
		"spec.rule_list.rules.#.rule_name",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.views.forward_proxy_policy.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.forward_proxy_policy.API.Replace"] = []string{
		"spec.rule_list.rules.#.dst_prefix_list.ipv6_prefixes.#",
		"spec.rule_list.rules.#.inside_sources",
		"spec.rule_list.rules.#.interface",
		"spec.rule_list.rules.#.metadata.disable",
		"spec.rule_list.rules.#.namespace",
		"spec.rule_list.rules.#.prefix_list.ipv6_prefixes.#",
		"spec.rule_list.rules.#.rule_description",
		"spec.rule_list.rules.#.rule_name",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.forward_proxy_policy.API"] = "config"
	sm["ves.io.schema.views.forward_proxy_policy.CustomDataAPI"] = "data"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

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
		csr.CRUDSwaggerRegistry["ves.io.schema.views.forward_proxy_policy.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.forward_proxy_policy.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.forward_proxy_policy.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.forward_proxy_policy.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.forward_proxy_policy.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.forward_proxy_policy.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.forward_proxy_policy.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.views.forward_proxy_policy.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.views.forward_proxy_policy.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.views.forward_proxy_policy.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.views.forward_proxy_policy.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.forward_proxy_policy.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.views.forward_proxy_policy.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomDataAPIServer(svc)
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