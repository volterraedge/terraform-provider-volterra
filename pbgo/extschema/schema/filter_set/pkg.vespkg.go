// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package filter_set

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.filter_set.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.filter_set.Object"] = ObjectValidator()
	vr["ves.io.schema.filter_set.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.filter_set.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.filter_set.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.filter_set.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.filter_set.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.filter_set.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.filter_set.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.filter_set.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.filter_set.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.filter_set.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.filter_set.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.filter_set.FindFilterSetsReq"] = FindFilterSetsReqValidator()
	vr["ves.io.schema.filter_set.FindFilterSetsRsp"] = FindFilterSetsRspValidator()

	vr["ves.io.schema.filter_set.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.filter_set.FilterExpressionField"] = FilterExpressionFieldValidator()
	vr["ves.io.schema.filter_set.FilterSetField"] = FilterSetFieldValidator()
	vr["ves.io.schema.filter_set.FilterStringField"] = FilterStringFieldValidator()
	vr["ves.io.schema.filter_set.FilterTimeRangeField"] = FilterTimeRangeFieldValidator()
	vr["ves.io.schema.filter_set.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.filter_set.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.filter_set.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.filter_set.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.filter_set.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.filter_set.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.filter_set.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.filter_set.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.filter_set.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.filter_set.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.filter_set.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedRequestFieldsRegistry["ves.io.schema.filter_set.API.Create"] = []string{
		"spec.filter_fields.#.label_selector_field",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.filter_set.API.Create"] = []string{
		"spec.filter_fields.#.label_selector_field",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.filter_set.API.Get"] = []string{
		"create_form.spec.filter_fields.#.label_selector_field",
		"replace_form.spec.filter_fields.#.label_selector_field",
		"spec.filter_fields.#.label_selector_field",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.filter_set.API.List"] = []string{
		"items.#.get_spec.filter_fields.#.label_selector_field",
	}

	mdr.RPCDeprecatedRequestFieldsRegistry["ves.io.schema.filter_set.API.Replace"] = []string{
		"spec.filter_fields.#.label_selector_field",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.filter_set.CustomAPI.FindFilterSets"] = []string{
		"filter_sets.#.spec.gc_spec.filter_fields.#.label_selector_field",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.filter_set.API"] = "config"
	sm["ves.io.schema.filter_set.CustomAPI"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.filter_set.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.filter_set.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.filter_set.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.filter_set.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.filter_set.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.filter_set.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.filter_set.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.filter_set.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.filter_set.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.filter_set.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.filter_set.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.filter_set.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.filter_set.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
