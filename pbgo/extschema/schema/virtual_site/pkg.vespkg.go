// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package virtual_site

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.virtual_site.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.virtual_site.Object"] = ObjectValidator()
	vr["ves.io.schema.virtual_site.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.virtual_site.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.virtual_site.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.virtual_site.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.virtual_site.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.virtual_site.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.virtual_site.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.virtual_site.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.virtual_site.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.virtual_site.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.virtual_site.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.virtual_site.SelecteeItemType"] = SelecteeItemTypeValidator()
	vr["ves.io.schema.virtual_site.SelecteeRequest"] = SelecteeRequestValidator()
	vr["ves.io.schema.virtual_site.SelecteeResponse"] = SelecteeResponseValidator()

	vr["ves.io.schema.virtual_site.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.virtual_site.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.virtual_site.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.virtual_site.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.virtual_site.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.virtual_site.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.virtual_site.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.virtual_site.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.virtual_site.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.virtual_site.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.virtual_site.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.virtual_site.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.virtual_site.API"] = "config"
	sm["ves.io.schema.virtual_site.CustomAPI"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.virtual_site.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.virtual_site.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.virtual_site.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.virtual_site.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.virtual_site.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.virtual_site.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.virtual_site.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.virtual_site.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.virtual_site.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.virtual_site.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.virtual_site.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.virtual_site.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.virtual_site.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
