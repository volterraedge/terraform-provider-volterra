// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package infraprotect_information

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.infraprotect_information.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.infraprotect_information.Object"] = ObjectValidator()
	vr["ves.io.schema.infraprotect_information.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.infraprotect_information.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.infraprotect_information.GetResponse"] = GetResponseValidator()

	vr["ves.io.schema.infraprotect_information.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.infraprotect_information.GlobalSpecType"] = GlobalSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.infraprotect_information.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.infraprotect_information.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.infraprotect_information.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.infraprotect_information.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.infraprotect_information.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.infraprotect_information.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.infraprotect_information.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.infraprotect_information.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.infraprotect_information.API"] = "infraprotect"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

	sm["infraprotect"] = svcfw.P0PolicyInfo{
		Name:            "ves-io-allow-infraprotect",
		ServiceSelector: "protekti\\.gc.*\\",
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
		csr.CRUDSwaggerRegistry["ves.io.schema.infraprotect_information.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.infraprotect_information.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.infraprotect_information.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.infraprotect_information.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.infraprotect_information.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.infraprotect_information.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.infraprotect_information.Object"] = NewCRUDAPIServer

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
