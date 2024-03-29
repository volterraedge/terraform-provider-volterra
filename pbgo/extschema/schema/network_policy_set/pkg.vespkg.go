// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package network_policy_set

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.network_policy_set.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.network_policy_set.Object"] = ObjectValidator()
	vr["ves.io.schema.network_policy_set.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.network_policy_set.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.network_policy_set.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.network_policy_set.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.network_policy_set.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.network_policy_set.ListResponseItem"] = ListResponseItemValidator()

	vr["ves.io.schema.network_policy_set.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.network_policy_set.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.network_policy_set.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.network_policy_set.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.network_policy_set.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.network_policy_set.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.network_policy_set.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.network_policy_set.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.network_policy_set.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.network_policy_set.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.network_policy_set.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.network_policy_set.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.network_policy_set.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.network_policy_set.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.network_policy_set.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.network_policy_set.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.network_policy_set.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.network_policy_set.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.network_policy_set.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.network_policy_set.Object"] = NewCRUDAPIServer

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
