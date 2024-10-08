// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package dns_lb_health_check

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.dns_lb_health_check.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.dns_lb_health_check.Object"] = ObjectValidator()
	vr["ves.io.schema.dns_lb_health_check.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.dns_lb_health_check.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.dns_lb_health_check.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.dns_lb_health_check.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.dns_lb_health_check.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.dns_lb_health_check.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.dns_lb_health_check.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.dns_lb_health_check.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.dns_lb_health_check.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.dns_lb_health_check.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.dns_lb_health_check.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.dns_lb_health_check.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.dns_lb_health_check.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.dns_lb_health_check.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.dns_lb_health_check.HttpHealthCheck"] = HttpHealthCheckValidator()
	vr["ves.io.schema.dns_lb_health_check.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.dns_lb_health_check.TcpHealthCheck"] = TcpHealthCheckValidator()
	vr["ves.io.schema.dns_lb_health_check.TcpHexHealthCheck"] = TcpHexHealthCheckValidator()
	vr["ves.io.schema.dns_lb_health_check.UdpHealthCheck"] = UdpHealthCheckValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.dns_lb_health_check.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.dns_lb_health_check.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dns_lb_health_check.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.dns_lb_health_check.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.dns_lb_health_check.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.dns_lb_health_check.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dns_lb_health_check.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.dns_lb_health_check.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.dns_lb_health_check.API"] = "config/dns"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

	sm["dns"] = svcfw.P0PolicyInfo{
		Name:            "ves-io-allow-config-dns",
		ServiceSelector: "bifrost\\.gc.*\\",
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
		csr.CRUDSwaggerRegistry["ves.io.schema.dns_lb_health_check.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.dns_lb_health_check.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.dns_lb_health_check.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.dns_lb_health_check.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.dns_lb_health_check.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.dns_lb_health_check.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.dns_lb_health_check.Object"] = NewCRUDAPIServer

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
