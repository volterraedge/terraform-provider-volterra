// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package api_discovery

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.api_sec.api_discovery.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.api_sec.api_discovery.Object"] = ObjectValidator()
	vr["ves.io.schema.api_sec.api_discovery.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.api_sec.api_discovery.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.api_sec.api_discovery.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.api_sec.api_discovery.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.api_sec.api_discovery.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.api_sec.api_discovery.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.api_sec.api_discovery.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.api_sec.api_discovery.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.api_sec.api_discovery.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.api_sec.api_discovery.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.api_sec.api_discovery.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.api_sec.api_discovery.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.api_sec.api_discovery.CustomAuthType"] = CustomAuthTypeValidator()
	vr["ves.io.schema.api_sec.api_discovery.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.api_sec.api_discovery.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.api_sec.api_discovery.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.api_sec.api_discovery.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.api_sec.api_discovery.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.api_sec.api_discovery.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.api_sec.api_discovery.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.api_sec.api_discovery.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.api_sec.api_discovery.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.api_sec.api_discovery.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.api_sec.api_discovery.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.api_sec.api_discovery.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.api_sec.api_discovery.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.api_sec.api_discovery.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.api_sec.api_discovery.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.api_sec.api_discovery.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.api_sec.api_discovery.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.api_sec.api_discovery.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.api_sec.api_discovery.Object"] = NewCRUDAPIServer

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
