// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package secret_policy

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.secret_policy.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.secret_policy.Object"] = ObjectValidator()
	vr["ves.io.schema.secret_policy.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.secret_policy.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.secret_policy.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.secret_policy.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.secret_policy.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.secret_policy.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.secret_policy.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.secret_policy.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.secret_policy.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.secret_policy.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.secret_policy.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.secret_policy.ListPolicyRequest"] = ListPolicyRequestValidator()
	vr["ves.io.schema.secret_policy.ListPolicyResponse"] = ListPolicyResponseValidator()
	vr["ves.io.schema.secret_policy.ListPolicyResponseItem"] = ListPolicyResponseItemValidator()
	vr["ves.io.schema.secret_policy.RecoverRequest"] = RecoverRequestValidator()
	vr["ves.io.schema.secret_policy.RecoverResponse"] = RecoverResponseValidator()
	vr["ves.io.schema.secret_policy.SoftDeleteRequest"] = SoftDeleteRequestValidator()
	vr["ves.io.schema.secret_policy.SoftDeleteResponse"] = SoftDeleteResponseValidator()

	vr["ves.io.schema.secret_policy.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.secret_policy.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.secret_policy.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.secret_policy.LegacyRuleList"] = LegacyRuleListValidator()
	vr["ves.io.schema.secret_policy.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.secret_policy.Rule"] = RuleValidator()
	vr["ves.io.schema.secret_policy.RuleList"] = RuleListValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.secret_policy.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.secret_policy.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.secret_policy.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.secret_policy.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.secret_policy.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.secret_policy.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.secret_policy.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.secret_policy.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.secret_policy.API.Create"] = []string{
		"spec.algo",
		"spec.legacy_rule_list",
		"spec.rule_list.rules.#.metadata.disable",
		"spec.rule_list.rules.#.spec.label_matcher",
		"spec.rules.#",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.secret_policy.API.Replace"] = []string{
		"spec.algo",
		"spec.rule_list.rules.#.metadata.disable",
		"spec.rule_list.rules.#.spec.label_matcher",
		"spec.rules.#",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.secret_policy.API"] = "secret_management"
	sm["ves.io.schema.secret_policy.CustomAPI"] = "secret_management"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.secret_policy.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.secret_policy.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.secret_policy.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.secret_policy.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.secret_policy.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.secret_policy.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.secret_policy.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.secret_policy.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.secret_policy.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.secret_policy.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.secret_policy.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.secret_policy.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.secret_policy.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
