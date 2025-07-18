// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package udp_loadbalancer

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.udp_loadbalancer.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.udp_loadbalancer.Object"] = ObjectValidator()
	vr["ves.io.schema.views.udp_loadbalancer.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.udp_loadbalancer.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.udp_loadbalancer.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.udp_loadbalancer.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.udp_loadbalancer.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.udp_loadbalancer.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.udp_loadbalancer.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.udp_loadbalancer.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.udp_loadbalancer.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.udp_loadbalancer.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.udp_loadbalancer.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.udp_loadbalancer.GetDnsInfoRequest"] = GetDnsInfoRequestValidator()
	vr["ves.io.schema.views.udp_loadbalancer.GetDnsInfoResponse"] = GetDnsInfoResponseValidator()

	vr["ves.io.schema.views.udp_loadbalancer.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.udp_loadbalancer.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.udp_loadbalancer.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.udp_loadbalancer.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.udp_loadbalancer.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.udp_loadbalancer.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.udp_loadbalancer.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.udp_loadbalancer.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.udp_loadbalancer.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.udp_loadbalancer.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.udp_loadbalancer.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.udp_loadbalancer.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.views.udp_loadbalancer.API.Create"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.CreateRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.cloud_edge_segment.ipv6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.CreateRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.site.ipv6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.CreateRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.site_segment.ipv6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.CreateRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.virtual_network.v6_vip_choice.specific_v6_vip",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.CreateRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.virtual_site_segment.ipv6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.udp_loadbalancer.API.Create"] = []string{
		"spec.do_not_retract_cluster",
		"spec.hash_policy_choice_least_active",
		"spec.port_ranges",
		"spec.retract_cluster",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.views.udp_loadbalancer.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_site_with_vip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.udp_loadbalancer.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_site_with_vip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.udp_loadbalancer.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.advertise_custom.advertise_where.#.virtual_site_with_vip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.advertise_custom.advertise_where.#.virtual_site_with_vip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_site_with_vip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.udp_loadbalancer.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.advertise_custom.advertise_where.#.virtual_site_with_vip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.views.udp_loadbalancer.API.Replace"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.ReplaceRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.cloud_edge_segment.ipv6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.ReplaceRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.site.ipv6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.ReplaceRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.site_segment.ipv6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.ReplaceRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.virtual_network.v6_vip_choice.specific_v6_vip",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.udp_loadbalancer.ReplaceRequest.spec.advertise_choice.advertise_custom.advertise_where.choice.virtual_site_segment.ipv6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.udp_loadbalancer.API.Replace"] = []string{
		"spec.do_not_retract_cluster",
		"spec.hash_policy_choice_least_active",
		"spec.port_ranges",
		"spec.retract_cluster",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.views.udp_loadbalancer.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.cloud_edge_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.site.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_network.v6_vip_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_site_segment.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.advertise_custom.advertise_where.#.virtual_site_with_vip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.udp_loadbalancer.API"] = "config"
	sm["ves.io.schema.views.udp_loadbalancer.CustomAPI"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.views.udp_loadbalancer.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.udp_loadbalancer.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.udp_loadbalancer.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.udp_loadbalancer.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.udp_loadbalancer.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.udp_loadbalancer.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.udp_loadbalancer.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.views.udp_loadbalancer.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.views.udp_loadbalancer.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.views.udp_loadbalancer.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.views.udp_loadbalancer.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.udp_loadbalancer.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.views.udp_loadbalancer.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
