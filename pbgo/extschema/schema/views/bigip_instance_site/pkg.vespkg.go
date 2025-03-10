// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package bigip_instance_site

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.bigip_instance_site.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.bigip_instance_site.Object"] = ObjectValidator()
	vr["ves.io.schema.views.bigip_instance_site.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.bigip_instance_site.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.bigip_instance_site.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.bigip_instance_site.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.bigip_instance_site.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.bigip_instance_site.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.bigip_instance_site.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.bigip_instance_site.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.bigip_instance_site.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.bigip_instance_site.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.bigip_instance_site.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.bigip_instance_site.CentralManagerList"] = CentralManagerListValidator()
	vr["ves.io.schema.views.bigip_instance_site.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.bigip_instance_site.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.bigip_instance_site.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.bigip_instance_site.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.bigip_instance_site.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.bigip_instance_site.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.bigip_instance_site.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.bigip_instance_site.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.bigip_instance_site.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.bigip_instance_site.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.bigip_instance_site.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.bigip_instance_site.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.views.bigip_instance_site.API.Create"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.views.bigip_instance_site.CreateRequest.spec.node_list.interface_list.ipv6_address_choice.ipv6_auto_config",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.bigip_instance_site.CreateRequest.spec.node_list.interface_list.ipv6_address_choice.static_ipv6_address",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.bigip_instance_site.API.Create"] = []string{
		"spec.node_list.#.interface_list.#.dhcp_server.dhcp_networks.#.network_prefix_allocator",
		"spec.node_list.#.interface_list.#.dhcp_server.dhcp_networks.#.pools.#.exclude",
		"spec.node_list.#.interface_list.#.dhcp_server.dhcp_option82_tag",
		"spec.node_list.#.interface_list.#.ipv6_auto_config.router.stateful.dhcp_networks.#.network_prefix_allocator",
		"spec.node_list.#.interface_list.#.ipv6_auto_config.router.stateful.dhcp_networks.#.pools.#.exclude",
		"spec.node_list.#.interface_list.#.static_ipv6_address.fleet_static_ip",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.views.bigip_instance_site.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.node_list.#.interface_list.#.ipv6_address_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.bigip_instance_site.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.node_list.#.interface_list.#.ipv6_address_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.bigip_instance_site.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.node_list.#.interface_list.#.ipv6_address_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.node_list.#.interface_list.#.ipv6_address_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.node_list.#.interface_list.#.ipv6_address_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.bigip_instance_site.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.node_list.#.interface_list.#.ipv6_address_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.views.bigip_instance_site.API.Replace"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.views.bigip_instance_site.ReplaceRequest.spec.node_list.interface_list.ipv6_address_choice.ipv6_auto_config",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.views.bigip_instance_site.ReplaceRequest.spec.node_list.interface_list.ipv6_address_choice.static_ipv6_address",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.bigip_instance_site.API.Replace"] = []string{
		"spec.node_list.#.interface_list.#.dhcp_server.dhcp_networks.#.network_prefix_allocator",
		"spec.node_list.#.interface_list.#.dhcp_server.dhcp_networks.#.pools.#.exclude",
		"spec.node_list.#.interface_list.#.dhcp_server.dhcp_option82_tag",
		"spec.node_list.#.interface_list.#.ipv6_auto_config.router.stateful.dhcp_networks.#.network_prefix_allocator",
		"spec.node_list.#.interface_list.#.ipv6_auto_config.router.stateful.dhcp_networks.#.pools.#.exclude",
		"spec.node_list.#.interface_list.#.static_ipv6_address.fleet_static_ip",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.views.bigip_instance_site.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.node_list.#.interface_list.#.ipv6_address_choice",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.bigip_instance_site.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.views.bigip_instance_site.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.bigip_instance_site.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.bigip_instance_site.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.bigip_instance_site.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.bigip_instance_site.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.bigip_instance_site.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.bigip_instance_site.Object"] = NewCRUDAPIServer

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
