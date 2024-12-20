// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package uztna_application_tag

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.uztna.application.uztna_application_tag.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.uztna.application.uztna_application_tag.Object"] = ObjectValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.uztna.application.uztna_application_tag.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.uztna.application.uztna_application_tag.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.uztna.application.uztna_application_tag.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.uztna.application.uztna_application_tag.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.uztna.application.uztna_application_tag.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.uztna.application.uztna_application_tag.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.uztna.application.uztna_application_tag.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.uztna.application.uztna_application_tag.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.uztna.application.uztna_application_tag.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.uztna.application.uztna_application_tag.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.uztna.application.uztna_application_tag.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.uztna.application.uztna_application_tag.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.uztna.application.uztna_application_tag.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.uztna.application.uztna_application_tag.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.uztna.application.uztna_application_tag.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.uztna.application.uztna_application_tag.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.uztna.application.uztna_application_tag.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.uztna.application.uztna_application_tag.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.uztna.application.uztna_application_tag.Object"] = NewCRUDAPIServer

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