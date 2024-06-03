// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package ztna_application

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.ztna_application.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.ztna_application.Object"] = ObjectValidator()
	vr["ves.io.schema.ztna_application.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.ztna_application.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.ztna_application.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.ztna_application.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.ztna_application.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.ztna_application.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.ztna_application.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.ztna_application.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.ztna_application.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.ztna_application.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.ztna_application.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.ztna_application.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.ztna_application.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.ztna_application.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.ztna_application.ProxyAdvertisementType"] = ProxyAdvertisementTypeValidator()
	vr["ves.io.schema.ztna_application.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.ztna_application.ServiceDetails"] = ServiceDetailsValidator()
	vr["ves.io.schema.ztna_application.TileAccess"] = TileAccessValidator()
	vr["ves.io.schema.ztna_application.ZTNApolicies"] = ZTNApoliciesValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.ztna_application.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.ztna_application.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.ztna_application.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.ztna_application.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.ztna_application.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.ztna_application.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.ztna_application.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.ztna_application.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.ztna_application.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.ztna_application.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.ztna_application.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.ztna_application.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.proxy_advertisement.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.proxy_advertisement.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.proxy_advertisement.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.ztna_application.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.proxy_advertisement.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.ztna_application.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.ztna_application.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.ztna_application.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.ztna_application.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.ztna_application.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.ztna_application.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.ztna_application.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.ztna_application.Object"] = NewCRUDAPIServer

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
