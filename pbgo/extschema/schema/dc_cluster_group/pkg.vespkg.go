// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package dc_cluster_group

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.dc_cluster_group.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.dc_cluster_group.Object"] = ObjectValidator()
	vr["ves.io.schema.dc_cluster_group.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.dc_cluster_group.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.dc_cluster_group.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.dc_cluster_group.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.dc_cluster_group.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.dc_cluster_group.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.dc_cluster_group.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.dc_cluster_group.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.dc_cluster_group.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.dc_cluster_group.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.dc_cluster_group.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.dc_cluster_group.MetricData"] = MetricDataValidator()
	vr["ves.io.schema.dc_cluster_group.MetricTypeData"] = MetricTypeDataValidator()
	vr["ves.io.schema.dc_cluster_group.MetricsRequest"] = MetricsRequestValidator()
	vr["ves.io.schema.dc_cluster_group.MetricsResponse"] = MetricsResponseValidator()

	vr["ves.io.schema.dc_cluster_group.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.dc_cluster_group.DCClusterGroupMeshType"] = DCClusterGroupMeshTypeValidator()
	vr["ves.io.schema.dc_cluster_group.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.dc_cluster_group.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.dc_cluster_group.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.dc_cluster_group.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.dc_cluster_group.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dc_cluster_group.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.dc_cluster_group.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.dc_cluster_group.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.dc_cluster_group.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dc_cluster_group.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.dc_cluster_group.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.dc_cluster_group.API"] = "config"
	sm["ves.io.schema.dc_cluster_group.CustomDataAPI"] = "data"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.dc_cluster_group.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.dc_cluster_group.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.dc_cluster_group.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.dc_cluster_group.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.dc_cluster_group.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.dc_cluster_group.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.dc_cluster_group.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.dc_cluster_group.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.dc_cluster_group.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.dc_cluster_group.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.dc_cluster_group.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.dc_cluster_group.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.dc_cluster_group.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomDataAPIServer(svc)
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
