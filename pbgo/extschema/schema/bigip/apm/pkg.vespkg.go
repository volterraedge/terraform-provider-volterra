// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package apm

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.bigip.apm.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.bigip.apm.Object"] = ObjectValidator()
	vr["ves.io.schema.bigip.apm.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.bigip.apm.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.bigip.apm.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.bigip.apm.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.bigip.apm.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.bigip.apm.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.bigip.apm.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.bigip.apm.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.bigip.apm.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.bigip.apm.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.bigip.apm.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.bigip.apm.MetricData"] = MetricDataValidator()
	vr["ves.io.schema.bigip.apm.MetricTypeData"] = MetricTypeDataValidator()
	vr["ves.io.schema.bigip.apm.MetricsRequest"] = MetricsRequestValidator()
	vr["ves.io.schema.bigip.apm.MetricsResponse"] = MetricsResponseValidator()

	vr["ves.io.schema.bigip.apm.APMBigIpAWSReplaceType"] = APMBigIpAWSReplaceTypeValidator()
	vr["ves.io.schema.bigip.apm.APMBigIpAWSType"] = APMBigIpAWSTypeValidator()
	vr["ves.io.schema.bigip.apm.AWSMarketPlaceImageTypeAPMaaS"] = AWSMarketPlaceImageTypeAPMaaSValidator()
	vr["ves.io.schema.bigip.apm.AWSSiteTypeChoice"] = AWSSiteTypeChoiceValidator()
	vr["ves.io.schema.bigip.apm.AWSSiteTypeChoiceReplaceType"] = AWSSiteTypeChoiceReplaceTypeValidator()
	vr["ves.io.schema.bigip.apm.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.bigip.apm.EndpointServiceReplaceType"] = EndpointServiceReplaceTypeValidator()
	vr["ves.io.schema.bigip.apm.EndpointServiceType"] = EndpointServiceTypeValidator()
	vr["ves.io.schema.bigip.apm.F5BigIpAppStackBareMetalChoiceReplaceType"] = F5BigIpAppStackBareMetalChoiceReplaceTypeValidator()
	vr["ves.io.schema.bigip.apm.F5BigIpAppStackBareMetalTypeChoice"] = F5BigIpAppStackBareMetalTypeChoiceValidator()
	vr["ves.io.schema.bigip.apm.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.bigip.apm.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.bigip.apm.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.bigip.apm.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.bigip.apm.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.bigip.apm.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.bigip.apm.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.bigip.apm.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.bigip.apm.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.bigip.apm.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.bigip.apm.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.bigip.apm.API.Create"] = []string{
		"spec.aws_site_type_choice.apm_aws_site.admin_password.blindfold_secret_info_internal",
		"spec.aws_site_type_choice.apm_aws_site.admin_password.secret_encoding_type",
		"spec.aws_site_type_choice.apm_aws_site.admin_password.vault_secret_info",
		"spec.aws_site_type_choice.apm_aws_site.admin_password.wingman_secret_info",
		"spec.baremetal_site_type_choice.f5_bare_metal_site.admin_password.blindfold_secret_info_internal",
		"spec.baremetal_site_type_choice.f5_bare_metal_site.admin_password.secret_encoding_type",
		"spec.baremetal_site_type_choice.f5_bare_metal_site.admin_password.vault_secret_info",
		"spec.baremetal_site_type_choice.f5_bare_metal_site.admin_password.wingman_secret_info",
		"spec.baremetal_site_type_choice.f5_bare_metal_site.bigiq_instance.password.blindfold_secret_info_internal",
		"spec.baremetal_site_type_choice.f5_bare_metal_site.bigiq_instance.password.secret_encoding_type",
		"spec.baremetal_site_type_choice.f5_bare_metal_site.bigiq_instance.password.vault_secret_info",
		"spec.baremetal_site_type_choice.f5_bare_metal_site.bigiq_instance.password.wingman_secret_info",
		"spec.https_management.advertise_on_public",
		"spec.https_management.advertise_on_public_default_vip",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.disable_local",
		"spec.https_management.do_not_advertise",
		"spec.https_management.do_not_advertise_on_internet",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.bigip.apm.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.aws_site_type_choice.apm_aws_site.nodes.#.mgmt_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.bigip.apm.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.aws_site_type_choice.apm_aws_site.nodes.#.mgmt_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.bigip.apm.API.Create"] = "ves.io.schema.bigip.apm.CreateRequest"

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.bigip.apm.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.aws_site_type_choice.apm_aws_site.nodes.#.mgmt_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.aws_site_type_choice.apm_aws_site.nodes.#.mgmt_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.bigip.apm.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.aws_site_type_choice.apm_aws_site.nodes.#.mgmt_subnet.subnet_param.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.bigip.apm.API.Replace"] = []string{
		"spec.https_management.advertise_on_public",
		"spec.https_management.advertise_on_public_default_vip",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.disable_local",
		"spec.https_management.do_not_advertise",
		"spec.https_management.do_not_advertise_on_internet",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.bigip.apm.API.Replace"] = "ves.io.schema.bigip.apm.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.bigip.apm.API"] = "config"
	sm["ves.io.schema.bigip.apm.CustomDataAPI"] = "data"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.bigip.apm.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.bigip.apm.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.bigip.apm.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.bigip.apm.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.bigip.apm.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.bigip.apm.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.bigip.apm.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.bigip.apm.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.bigip.apm.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.bigip.apm.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.bigip.apm.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.bigip.apm.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.bigip.apm.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
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
