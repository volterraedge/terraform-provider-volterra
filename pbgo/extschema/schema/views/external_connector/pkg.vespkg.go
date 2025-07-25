// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package external_connector

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.external_connector.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.external_connector.Object"] = ObjectValidator()
	vr["ves.io.schema.views.external_connector.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.external_connector.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.external_connector.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.external_connector.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.external_connector.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.external_connector.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.external_connector.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.external_connector.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.external_connector.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.external_connector.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.external_connector.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.external_connector.ConnectionTypeDirect"] = ConnectionTypeDirectValidator()
	vr["ves.io.schema.views.external_connector.ConnectionTypeGRE"] = ConnectionTypeGREValidator()
	vr["ves.io.schema.views.external_connector.ConnectionTypeIPSec"] = ConnectionTypeIPSecValidator()
	vr["ves.io.schema.views.external_connector.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.external_connector.DpdKeepAliveTimer"] = DpdKeepAliveTimerValidator()
	vr["ves.io.schema.views.external_connector.GRETunnelParameters"] = GRETunnelParametersValidator()
	vr["ves.io.schema.views.external_connector.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.external_connector.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.external_connector.IkeParameters"] = IkeParametersValidator()
	vr["ves.io.schema.views.external_connector.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.views.external_connector.TunnelEndpoint"] = TunnelEndpointValidator()
	vr["ves.io.schema.views.external_connector.TunnelParameters"] = TunnelParametersValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.external_connector.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.external_connector.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.external_connector.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.external_connector.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.external_connector.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.external_connector.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.external_connector.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.external_connector.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.external_connector.API.Create"] = []string{
		"spec.gre.gre_parameters.segment.virtual_networks.#",
		"spec.ipsec.ike_parameters.lc_hostname",
		"spec.ipsec.ike_parameters.lc_ip_address",
		"spec.ipsec.ipsec_tunnel_parameters.segment.virtual_networks.#",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.views.external_connector.API.Create"] = "ves.io.schema.views.external_connector.CreateRequest"

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.external_connector.API.Replace"] = []string{
		"spec.gre.gre_parameters.segment.virtual_networks.#",
		"spec.ipsec.ike_parameters.lc_hostname",
		"spec.ipsec.ike_parameters.lc_ip_address",
		"spec.ipsec.ipsec_tunnel_parameters.segment.virtual_networks.#",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.views.external_connector.API.Replace"] = "ves.io.schema.views.external_connector.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.external_connector.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.views.external_connector.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.external_connector.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.external_connector.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.external_connector.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.external_connector.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.external_connector.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.external_connector.Object"] = NewCRUDAPIServer

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
