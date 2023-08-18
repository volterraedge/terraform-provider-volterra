// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package user_identification

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.user_identification.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.user_identification.Object"] = ObjectValidator()
	vr["ves.io.schema.user_identification.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.user_identification.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.user_identification.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.user_identification.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.user_identification.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.user_identification.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.user_identification.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.user_identification.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.user_identification.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.user_identification.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.user_identification.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.user_identification.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.user_identification.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.user_identification.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.user_identification.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.user_identification.UserIdentificationRule"] = UserIdentificationRuleValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.user_identification.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.user_identification.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.user_identification.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.user_identification.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.user_identification.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.user_identification.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.user_identification.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.user_identification.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.user_identification.API.Get"] = []string{
		"object",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.user_identification.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.user_identification.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.user_identification.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.user_identification.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.user_identification.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.user_identification.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.user_identification.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.user_identification.Object"] = NewCRUDAPIServer

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