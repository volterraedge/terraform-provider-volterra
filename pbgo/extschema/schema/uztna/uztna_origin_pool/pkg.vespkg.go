// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package uztna_origin_pool

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.uztna.uztna_origin_pool.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.uztna.uztna_origin_pool.Object"] = ObjectValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.uztna.uztna_origin_pool.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.uztna.uztna_origin_pool.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.LoadbalancerAlgorithm"] = LoadbalancerAlgorithmValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.OriginServerPrivateIP"] = OriginServerPrivateIPValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.OriginServerPrivateName"] = OriginServerPrivateNameValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.OriginServerType"] = OriginServerTypeValidator()
	vr["ves.io.schema.uztna.uztna_origin_pool.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.uztna.uztna_origin_pool.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.uztna.uztna_origin_pool.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.uztna.uztna_origin_pool.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.uztna.uztna_origin_pool.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.uztna.uztna_origin_pool.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.uztna.uztna_origin_pool.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.uztna.uztna_origin_pool.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.uztna.uztna_origin_pool.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.uztna.uztna_origin_pool.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.uztna.uztna_origin_pool.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.uztna.uztna_origin_pool.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.uztna.uztna_origin_pool.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.uztna.uztna_origin_pool.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.uztna.uztna_origin_pool.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.uztna.uztna_origin_pool.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.uztna.uztna_origin_pool.Object"] = NewCRUDAPIServer

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
