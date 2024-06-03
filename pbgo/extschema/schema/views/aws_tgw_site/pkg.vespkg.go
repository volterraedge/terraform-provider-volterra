// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package aws_tgw_site

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.aws_tgw_site.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.aws_tgw_site.Object"] = ObjectValidator()
	vr["ves.io.schema.views.aws_tgw_site.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.aws_tgw_site.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.aws_tgw_site.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.aws_tgw_site.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.aws_tgw_site.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.aws_tgw_site.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.aws_tgw_site.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.aws_tgw_site.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.aws_tgw_site.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.aws_tgw_site.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.aws_tgw_site.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.aws_tgw_site.SetTGWInfoRequest"] = SetTGWInfoRequestValidator()
	vr["ves.io.schema.views.aws_tgw_site.SetTGWInfoResponse"] = SetTGWInfoResponseValidator()
	vr["ves.io.schema.views.aws_tgw_site.SetVIPInfoRequest"] = SetVIPInfoRequestValidator()
	vr["ves.io.schema.views.aws_tgw_site.SetVIPInfoResponse"] = SetVIPInfoResponseValidator()
	vr["ves.io.schema.views.aws_tgw_site.SetVPCIpPrefixesRequest"] = SetVPCIpPrefixesRequestValidator()
	vr["ves.io.schema.views.aws_tgw_site.SetVPCIpPrefixesResponse"] = SetVPCIpPrefixesResponseValidator()
	vr["ves.io.schema.views.aws_tgw_site.SetVPNTunnelsRequest"] = SetVPNTunnelsRequestValidator()
	vr["ves.io.schema.views.aws_tgw_site.SetVPNTunnelsResponse"] = SetVPNTunnelsResponseValidator()

	vr["ves.io.schema.views.aws_tgw_site.AWSTGWInfoConfigType"] = AWSTGWInfoConfigTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.AWSTGWResourceShareType"] = AWSTGWResourceShareTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.AWSTGWStatusType"] = AWSTGWStatusTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.AWSVPNTunnelConfigType"] = AWSVPNTunnelConfigTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.ActiveServicePoliciesType"] = ActiveServicePoliciesTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.ExistingTGWType"] = ExistingTGWTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.SecurityConfigType"] = SecurityConfigTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.ServicesVPCReplaceType"] = ServicesVPCReplaceTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.ServicesVPCType"] = ServicesVPCTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.TGWAssignedASNType"] = TGWAssignedASNTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.TGWParamsType"] = TGWParamsTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.VPCAttachmentListType"] = VPCAttachmentListTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.VPCAttachmentType"] = VPCAttachmentTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.VPCIpPrefixesType"] = VPCIpPrefixesTypeValidator()
	vr["ves.io.schema.views.aws_tgw_site.VnConfiguration"] = VnConfigurationValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.aws_tgw_site.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.aws_tgw_site.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.aws_tgw_site.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.aws_tgw_site.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.aws_tgw_site.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.aws_tgw_site.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.aws_tgw_site.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.aws_tgw_site.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.aws_tgw_site.API.Create"] = []string{
		"spec.aws_parameters.assisted",
		"spec.aws_parameters.az_nodes.#.disk_size",
		"spec.vn_config.global_network_list.global_network_connections.#.enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info_internal",
		"spec.vn_config.global_network_list.global_network_connections.#.enable_forward_proxy.tls_intercept.custom_certificate.private_key.secret_encoding_type",
		"spec.vn_config.global_network_list.global_network_connections.#.enable_forward_proxy.tls_intercept.custom_certificate.private_key.vault_secret_info",
		"spec.vn_config.global_network_list.global_network_connections.#.enable_forward_proxy.tls_intercept.custom_certificate.private_key.wingman_secret_info",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.views.aws_tgw_site.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.inside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.outside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.workload_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.custom_dns.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.custom_dns.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.aws_tgw_site.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.inside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.outside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.workload_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.custom_dns.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.custom_dns.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.views.aws_tgw_site.API.Create"] = "ves.io.schema.views.aws_tgw_site.CreateRequest"

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.aws_tgw_site.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.aws_parameters.az_nodes.#.inside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.aws_parameters.az_nodes.#.outside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.aws_parameters.az_nodes.#.workload_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.custom_dns.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "create_form.spec.custom_dns.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.aws_parameters.az_nodes.#.inside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.aws_parameters.az_nodes.#.outside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.aws_parameters.az_nodes.#.workload_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.custom_dns.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.custom_dns.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.inside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.outside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.workload_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.custom_dns.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.custom_dns.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.views.aws_tgw_site.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.aws_parameters.az_nodes.#.inside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.aws_parameters.az_nodes.#.outside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.aws_parameters.az_nodes.#.workload_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.custom_dns.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.custom_dns.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.aws_tgw_site.API.Replace"] = []string{
		"spec.aws_parameters.assisted",
		"spec.aws_parameters.az_nodes.#.disk_size",
		"spec.vn_config.global_network_list.global_network_connections.#.enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info_internal",
		"spec.vn_config.global_network_list.global_network_connections.#.enable_forward_proxy.tls_intercept.custom_certificate.private_key.secret_encoding_type",
		"spec.vn_config.global_network_list.global_network_connections.#.enable_forward_proxy.tls_intercept.custom_certificate.private_key.vault_secret_info",
		"spec.vn_config.global_network_list.global_network_connections.#.enable_forward_proxy.tls_intercept.custom_certificate.private_key.wingman_secret_info",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.views.aws_tgw_site.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.inside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.outside_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_parameters.az_nodes.#.workload_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.custom_dns.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.custom_dns.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.views.aws_tgw_site.API.Replace"] = "ves.io.schema.views.aws_tgw_site.ReplaceRequest"

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.views.aws_tgw_site.CustomAPI.SetVPNTunnels"] = []string{
		"tunnels.#",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.aws_tgw_site.API"] = "config"
	sm["ves.io.schema.views.aws_tgw_site.CustomAPI"] = "config"

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

	customCSR = mdr.PvtCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.views.aws_tgw_site.Object"] = PrivateCustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.views.aws_tgw_site.PrivateCustomAPI"] = NewPrivateCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.views.aws_tgw_site.PrivateCustomAPI"] = NewPrivateCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.views.aws_tgw_site.PrivateCustomAPI"] = RegisterPrivateCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.aws_tgw_site.PrivateCustomAPI"] = RegisterGwPrivateCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.views.aws_tgw_site.PrivateCustomAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewPrivateCustomAPIServer(svc)
		}

	}()

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.views.aws_tgw_site.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.aws_tgw_site.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.aws_tgw_site.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.aws_tgw_site.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.aws_tgw_site.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.aws_tgw_site.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.aws_tgw_site.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.views.aws_tgw_site.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.views.aws_tgw_site.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.views.aws_tgw_site.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.views.aws_tgw_site.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.aws_tgw_site.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.views.aws_tgw_site.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
