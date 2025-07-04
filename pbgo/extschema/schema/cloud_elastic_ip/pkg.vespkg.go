// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cloud_elastic_ip

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.cloud_elastic_ip.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.cloud_elastic_ip.Object"] = ObjectValidator()
	vr["ves.io.schema.cloud_elastic_ip.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.cloud_elastic_ip.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.cloud_elastic_ip.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.cloud_elastic_ip.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.cloud_elastic_ip.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.cloud_elastic_ip.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.cloud_elastic_ip.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.cloud_elastic_ip.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.cloud_elastic_ip.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.cloud_elastic_ip.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.cloud_elastic_ip.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.cloud_elastic_ip.ForceDeleteCloudElasticIPRequest"] = ForceDeleteCloudElasticIPRequestValidator()
	vr["ves.io.schema.cloud_elastic_ip.ForceDeleteCloudElasticIPResponse"] = ForceDeleteCloudElasticIPResponseValidator()

	vr["ves.io.schema.cloud_elastic_ip.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.cloud_elastic_ip.DeleteSpecType"] = DeleteSpecTypeValidator()
	vr["ves.io.schema.cloud_elastic_ip.ElasticIPInfoType"] = ElasticIPInfoTypeValidator()
	vr["ves.io.schema.cloud_elastic_ip.ElasticIPStatusType"] = ElasticIPStatusTypeValidator()
	vr["ves.io.schema.cloud_elastic_ip.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.cloud_elastic_ip.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.cloud_elastic_ip.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.cloud_elastic_ip.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.cloud_elastic_ip.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.cloud_elastic_ip.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.cloud_elastic_ip.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.cloud_elastic_ip.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.cloud_elastic_ip.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.cloud_elastic_ip.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.cloud_elastic_ip.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.cloud_elastic_ip.API"] = "config"
	sm["ves.io.schema.cloud_elastic_ip.CustomAPI"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.cloud_elastic_ip.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.cloud_elastic_ip.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.cloud_elastic_ip.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.cloud_elastic_ip.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.cloud_elastic_ip.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.cloud_elastic_ip.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.cloud_elastic_ip.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.cloud_elastic_ip.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.cloud_elastic_ip.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.cloud_elastic_ip.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.cloud_elastic_ip.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.cloud_elastic_ip.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.cloud_elastic_ip.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
