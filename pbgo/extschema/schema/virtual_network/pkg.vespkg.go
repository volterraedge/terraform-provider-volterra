// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package virtual_network

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.virtual_network.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.virtual_network.Object"] = ObjectValidator()
	vr["ves.io.schema.virtual_network.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.virtual_network.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.virtual_network.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.virtual_network.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.virtual_network.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.virtual_network.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.virtual_network.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.virtual_network.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.virtual_network.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.virtual_network.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.virtual_network.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.virtual_network.SIDCounterData"] = SIDCounterDataValidator()
	vr["ves.io.schema.virtual_network.SIDCounterRequest"] = SIDCounterRequestValidator()
	vr["ves.io.schema.virtual_network.SIDCounterResponse"] = SIDCounterResponseValidator()
	vr["ves.io.schema.virtual_network.SIDCounterTypeData"] = SIDCounterTypeDataValidator()

	vr["ves.io.schema.virtual_network.ActivePBRPoliciesType"] = ActivePBRPoliciesTypeValidator()
	vr["ves.io.schema.virtual_network.AnyCastVIPFleetType"] = AnyCastVIPFleetTypeValidator()
	vr["ves.io.schema.virtual_network.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.virtual_network.DNSServersList"] = DNSServersListValidator()
	vr["ves.io.schema.virtual_network.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.virtual_network.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.virtual_network.NextHopInterfaceList"] = NextHopInterfaceListValidator()
	vr["ves.io.schema.virtual_network.PerSiteSrv6NetworkType"] = PerSiteSrv6NetworkTypeValidator()
	vr["ves.io.schema.virtual_network.PerTenantVIPType"] = PerTenantVIPTypeValidator()
	vr["ves.io.schema.virtual_network.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.virtual_network.SNATPoolFleetType"] = SNATPoolFleetTypeValidator()
	vr["ves.io.schema.virtual_network.SNATPoolSiteType"] = SNATPoolSiteTypeValidator()
	vr["ves.io.schema.virtual_network.SNATPoolType"] = SNATPoolTypeValidator()
	vr["ves.io.schema.virtual_network.SegmentNetworkType"] = SegmentNetworkTypeValidator()
	vr["ves.io.schema.virtual_network.Srv6NetworkNsParametersType"] = Srv6NetworkNsParametersTypeValidator()
	vr["ves.io.schema.virtual_network.StaticRouteViewType"] = StaticRouteViewTypeValidator()
	vr["ves.io.schema.virtual_network.StaticV6RouteViewType"] = StaticV6RouteViewTypeValidator()
	vr["ves.io.schema.virtual_network.StaticV6RoutesListType"] = StaticV6RoutesListTypeValidator()
	vr["ves.io.schema.virtual_network.VoltADNPrivateNetworkReInfoType"] = VoltADNPrivateNetworkReInfoTypeValidator()
	vr["ves.io.schema.virtual_network.VoltADNPrivateNetworkTenantInfoType"] = VoltADNPrivateNetworkTenantInfoTypeValidator()
	vr["ves.io.schema.virtual_network.VoltADNPrivateNetworkType"] = VoltADNPrivateNetworkTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.virtual_network.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.virtual_network.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.virtual_network.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.virtual_network.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.virtual_network.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.virtual_network.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.virtual_network.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.virtual_network.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.virtual_network.API.Create"] = []string{
		"spec.srv6_network",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.virtual_network.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.srv6_network.site_snat_pool.node_snat_pool.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.static_v6_routes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.virtual_network.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.srv6_network.site_snat_pool.node_snat_pool.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.virtual_network.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.srv6_network.site_snat_pool.node_snat_pool.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.static_v6_routes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.srv6_network.site_snat_pool.node_snat_pool.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.static_v6_routes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.srv6_network.site_snat_pool.node_snat_pool.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.virtual_network.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.srv6_network.site_snat_pool.node_snat_pool.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.virtual_network.API.Replace"] = []string{
		"spec.private_network",
		"spec.srv6_network",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.virtual_network.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.srv6_network.site_snat_pool.node_snat_pool.ipv6_prefixes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.static_v6_routes.#",
			AllowedEnvironments: []string{"crt", "demo1", "softbank_mec", "staging", "test"},
		},
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.virtual_network.API"] = "config"
	sm["ves.io.schema.virtual_network.CustomDataAPI"] = "data"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.virtual_network.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.virtual_network.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.virtual_network.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.virtual_network.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.virtual_network.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.virtual_network.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.virtual_network.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.virtual_network.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.virtual_network.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.virtual_network.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.virtual_network.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.virtual_network.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.virtual_network.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
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
