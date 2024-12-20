// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package nat_policy

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.nat_policy.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.nat_policy.Object"] = ObjectValidator()
	vr["ves.io.schema.nat_policy.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.nat_policy.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.nat_policy.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.nat_policy.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.nat_policy.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.nat_policy.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.nat_policy.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.nat_policy.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.nat_policy.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.nat_policy.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.nat_policy.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.nat_policy.ActionType"] = ActionTypeValidator()
	vr["ves.io.schema.nat_policy.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.nat_policy.DynamicPool"] = DynamicPoolValidator()
	vr["ves.io.schema.nat_policy.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.nat_policy.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.nat_policy.MatchCriteriaType"] = MatchCriteriaTypeValidator()
	vr["ves.io.schema.nat_policy.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.nat_policy.RuleType"] = RuleTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.nat_policy.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.nat_policy.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.nat_policy.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.nat_policy.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.nat_policy.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.nat_policy.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.nat_policy.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.nat_policy.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.nat_policy.API.Create"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.nat_policy.CreateRequest.spec.rules.action.source_nat_choice.dynamic.pool_choice.pools.ipv6_prefixes",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.nat_policy.API.Create"] = []string{
		"spec.rules.#.criteria.segment.virtual_networks.#",
		"spec.rules.#.segment.virtual_networks.#",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.nat_policy.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.rules.#.action.dynamic.pools.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.nat_policy.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.rules.#.action.dynamic.pools.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.nat_policy.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.rules.#.action.dynamic.pools.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.rules.#.action.dynamic.pools.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.rules.#.action.dynamic.pools.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.nat_policy.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.rules.#.action.dynamic.pools.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.nat_policy.API.Replace"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.nat_policy.ReplaceRequest.spec.rules.action.source_nat_choice.dynamic.pool_choice.pools.ipv6_prefixes",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.nat_policy.API.Replace"] = []string{
		"spec.rules.#.criteria.segment.virtual_networks.#",
		"spec.rules.#.segment.virtual_networks.#",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.nat_policy.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.rules.#.action.dynamic.pools.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.nat_policy.API"] = "config"

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

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.nat_policy.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.nat_policy.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.nat_policy.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.nat_policy.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.nat_policy.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.nat_policy.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.nat_policy.Object"] = NewCRUDAPIServer

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