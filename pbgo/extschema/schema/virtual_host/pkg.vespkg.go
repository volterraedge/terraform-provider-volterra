// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package virtual_host

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.virtual_host.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.virtual_host.Object"] = ObjectValidator()
	vr["ves.io.schema.virtual_host.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.virtual_host.APIEPActivityMetrics"] = APIEPActivityMetricsValidator()
	vr["ves.io.schema.virtual_host.APIEPSummaryFilter"] = APIEPSummaryFilterValidator()
	vr["ves.io.schema.virtual_host.APIEndpoint"] = APIEndpointValidator()
	vr["ves.io.schema.virtual_host.APIEndpointLearntSchemaReq"] = APIEndpointLearntSchemaReqValidator()
	vr["ves.io.schema.virtual_host.APIEndpointLearntSchemaRsp"] = APIEndpointLearntSchemaRspValidator()
	vr["ves.io.schema.virtual_host.APIEndpointPDFReq"] = APIEndpointPDFReqValidator()
	vr["ves.io.schema.virtual_host.APIEndpointPDFRsp"] = APIEndpointPDFRspValidator()
	vr["ves.io.schema.virtual_host.APIEndpointReq"] = APIEndpointReqValidator()
	vr["ves.io.schema.virtual_host.APIEndpointRsp"] = APIEndpointRspValidator()
	vr["ves.io.schema.virtual_host.APIEndpointsReq"] = APIEndpointsReqValidator()
	vr["ves.io.schema.virtual_host.APIEndpointsRsp"] = APIEndpointsRspValidator()
	vr["ves.io.schema.virtual_host.ApiEndpointsStatsReq"] = ApiEndpointsStatsReqValidator()
	vr["ves.io.schema.virtual_host.ApiEndpointsStatsRsp"] = ApiEndpointsStatsRspValidator()
	vr["ves.io.schema.virtual_host.CreateJiraIssueRequest"] = CreateJiraIssueRequestValidator()
	vr["ves.io.schema.virtual_host.CreateTicketRequest"] = CreateTicketRequestValidator()
	vr["ves.io.schema.virtual_host.CreateTicketResponse"] = CreateTicketResponseValidator()
	vr["ves.io.schema.virtual_host.GetAPICallSummaryReq"] = GetAPICallSummaryReqValidator()
	vr["ves.io.schema.virtual_host.GetAPICallSummaryRsp"] = GetAPICallSummaryRspValidator()
	vr["ves.io.schema.virtual_host.GetAPIEndpointsSchemaUpdatesReq"] = GetAPIEndpointsSchemaUpdatesReqValidator()
	vr["ves.io.schema.virtual_host.GetAPIEndpointsSchemaUpdatesResp"] = GetAPIEndpointsSchemaUpdatesRespValidator()
	vr["ves.io.schema.virtual_host.GetTopAPIEndpointsReq"] = GetTopAPIEndpointsReqValidator()
	vr["ves.io.schema.virtual_host.GetTopAPIEndpointsRsp"] = GetTopAPIEndpointsRspValidator()
	vr["ves.io.schema.virtual_host.GetTopSensitiveDataReq"] = GetTopSensitiveDataReqValidator()
	vr["ves.io.schema.virtual_host.GetTopSensitiveDataRsp"] = GetTopSensitiveDataRspValidator()
	vr["ves.io.schema.virtual_host.GetVulnerabilitiesReq"] = GetVulnerabilitiesReqValidator()
	vr["ves.io.schema.virtual_host.GetVulnerabilitiesRsp"] = GetVulnerabilitiesRspValidator()
	vr["ves.io.schema.virtual_host.RequestCountPerResponseClass"] = RequestCountPerResponseClassValidator()
	vr["ves.io.schema.virtual_host.SensitiveDataCount"] = SensitiveDataCountValidator()
	vr["ves.io.schema.virtual_host.SwaggerSpecReq"] = SwaggerSpecReqValidator()
	vr["ves.io.schema.virtual_host.SwaggerSpecRsp"] = SwaggerSpecRspValidator()
	vr["ves.io.schema.virtual_host.TicketDetails"] = TicketDetailsValidator()
	vr["ves.io.schema.virtual_host.UnlinkTicketsRequest"] = UnlinkTicketsRequestValidator()
	vr["ves.io.schema.virtual_host.UnlinkTicketsResponse"] = UnlinkTicketsResponseValidator()
	vr["ves.io.schema.virtual_host.UpdateAPIEndpointsSchemasReq"] = UpdateAPIEndpointsSchemasReqValidator()
	vr["ves.io.schema.virtual_host.UpdateAPIEndpointsSchemasResp"] = UpdateAPIEndpointsSchemasRespValidator()
	vr["ves.io.schema.virtual_host.UpdateVulnerabilitiesStateReq"] = UpdateVulnerabilitiesStateReqValidator()
	vr["ves.io.schema.virtual_host.UpdateVulnerabilitiesStateRsp"] = UpdateVulnerabilitiesStateRspValidator()
	vr["ves.io.schema.virtual_host.VulnEvidence"] = VulnEvidenceValidator()
	vr["ves.io.schema.virtual_host.VulnEvidenceSample"] = VulnEvidenceSampleValidator()
	vr["ves.io.schema.virtual_host.VulnRisk"] = VulnRiskValidator()
	vr["ves.io.schema.virtual_host.Vulnerability"] = VulnerabilityValidator()

	vr["ves.io.schema.virtual_host.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.virtual_host.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.virtual_host.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.virtual_host.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.virtual_host.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.virtual_host.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.virtual_host.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.virtual_host.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.virtual_host.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.virtual_host.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.virtual_host.AssignAPIDefinitionReq"] = AssignAPIDefinitionReqValidator()
	vr["ves.io.schema.virtual_host.AssignAPIDefinitionResp"] = AssignAPIDefinitionRespValidator()
	vr["ves.io.schema.virtual_host.GetDnsInfoRequest"] = GetDnsInfoRequestValidator()
	vr["ves.io.schema.virtual_host.GetDnsInfoResponse"] = GetDnsInfoResponseValidator()

	vr["ves.io.schema.virtual_host.ApiSpec"] = ApiSpecValidator()
	vr["ves.io.schema.virtual_host.AuthenticationDetails"] = AuthenticationDetailsValidator()
	vr["ves.io.schema.virtual_host.AutoCertInfoType"] = AutoCertInfoTypeValidator()
	vr["ves.io.schema.virtual_host.CDNOriginServerPublicIP"] = CDNOriginServerPublicIPValidator()
	vr["ves.io.schema.virtual_host.CDNOriginServerPublicName"] = CDNOriginServerPublicNameValidator()
	vr["ves.io.schema.virtual_host.CDNOriginServerType"] = CDNOriginServerTypeValidator()
	vr["ves.io.schema.virtual_host.CDNUpstreamTlsParameters"] = CDNUpstreamTlsParametersValidator()
	vr["ves.io.schema.virtual_host.CaptchaChallengeType"] = CaptchaChallengeTypeValidator()
	vr["ves.io.schema.virtual_host.CdnOriginPoolType"] = CdnOriginPoolTypeValidator()
	vr["ves.io.schema.virtual_host.CdnServiceType"] = CdnServiceTypeValidator()
	vr["ves.io.schema.virtual_host.ClientIPHeaders"] = ClientIPHeadersValidator()
	vr["ves.io.schema.virtual_host.CompressionType"] = CompressionTypeValidator()
	vr["ves.io.schema.virtual_host.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.virtual_host.DNSCacheProfile"] = DNSCacheProfileValidator()
	vr["ves.io.schema.virtual_host.DNSDDoSProfile"] = DNSDDoSProfileValidator()
	vr["ves.io.schema.virtual_host.DNSProxyConfiguration"] = DNSProxyConfigurationValidator()
	vr["ves.io.schema.virtual_host.DNSRecord"] = DNSRecordValidator()
	vr["ves.io.schema.virtual_host.DNSVHostStatusType"] = DNSVHostStatusTypeValidator()
	vr["ves.io.schema.virtual_host.DefaultCacheAction"] = DefaultCacheActionValidator()
	vr["ves.io.schema.virtual_host.DelegationLocationSelection"] = DelegationLocationSelectionValidator()
	vr["ves.io.schema.virtual_host.DomainCertificates"] = DomainCertificatesValidator()
	vr["ves.io.schema.virtual_host.DynamicReverseProxyType"] = DynamicReverseProxyTypeValidator()
	vr["ves.io.schema.virtual_host.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.virtual_host.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.virtual_host.HeaderControlType"] = HeaderControlTypeValidator()
	vr["ves.io.schema.virtual_host.Http1ProtocolOptions"] = Http1ProtocolOptionsValidator()
	vr["ves.io.schema.virtual_host.HttpProtocolOptions"] = HttpProtocolOptionsValidator()
	vr["ves.io.schema.virtual_host.HttpRedirectOptions"] = HttpRedirectOptionsValidator()
	vr["ves.io.schema.virtual_host.JavaScriptConfigType"] = JavaScriptConfigTypeValidator()
	vr["ves.io.schema.virtual_host.JavascriptChallengeType"] = JavascriptChallengeTypeValidator()
	vr["ves.io.schema.virtual_host.LogHeaderOptions"] = LogHeaderOptionsValidator()
	vr["ves.io.schema.virtual_host.LoggingOptionsType"] = LoggingOptionsTypeValidator()
	vr["ves.io.schema.virtual_host.MaskingConfiguration"] = MaskingConfigurationValidator()
	vr["ves.io.schema.virtual_host.OpenApiValidationSettings"] = OpenApiValidationSettingsValidator()
	vr["ves.io.schema.virtual_host.OriginAdvancedConfiguration"] = OriginAdvancedConfigurationValidator()
	vr["ves.io.schema.virtual_host.OtherSettings"] = OtherSettingsValidator()
	vr["ves.io.schema.virtual_host.PolicyBasedChallenge"] = PolicyBasedChallengeValidator()
	vr["ves.io.schema.virtual_host.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.virtual_host.ServiceDomain"] = ServiceDomainValidator()
	vr["ves.io.schema.virtual_host.ShapeBotDefenseConfigType"] = ShapeBotDefenseConfigTypeValidator()
	vr["ves.io.schema.virtual_host.SlowDDoSMitigation"] = SlowDDoSMitigationValidator()
	vr["ves.io.schema.virtual_host.TemporaryUserBlockingType"] = TemporaryUserBlockingTypeValidator()
	vr["ves.io.schema.virtual_host.VerStatusType"] = VerStatusTypeValidator()
	vr["ves.io.schema.virtual_host.VirtualHostID"] = VirtualHostIDValidator()
	vr["ves.io.schema.virtual_host.ZtnaProxyConfiguration"] = ZtnaProxyConfigurationValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.virtual_host.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.virtual_host.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.virtual_host.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.virtual_host.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.virtual_host.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.virtual_host.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.virtual_host.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.virtual_host.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.virtual_host.ApiepCustomAPI.GetAPIEndpointLearntSchema"] = []string{
		"discovered_openapi_spec",
	}

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.virtual_host.API.Create"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.virtual_host.CreateRequest.spec.api_spec.open_api_validation_choice.enable_open_api_validation",
			AddonServices: []string{"f5xc-waap-advanced"},
		},
		{
			FieldPath:     "ves.io.schema.virtual_host.CreateRequest.spec.masking_config.masking_choice.enable_masking",
			AddonServices: []string{"f5xc-waap-advanced"},
		},
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.virtual_host.API.Create"] = []string{
		"spec.cdn_service.cache_ttl",
		"spec.cdn_service.service_domains.#",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.virtual_host.API.Create"] = []string{
		"spec.api_spec",
		"spec.authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info_internal",
		"spec.authentication.cookie_params.auth_hmac.prim_key.secret_encoding_type",
		"spec.authentication.cookie_params.auth_hmac.prim_key.vault_secret_info",
		"spec.authentication.cookie_params.auth_hmac.prim_key.wingman_secret_info",
		"spec.authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info_internal",
		"spec.authentication.cookie_params.auth_hmac.sec_key.secret_encoding_type",
		"spec.authentication.cookie_params.auth_hmac.sec_key.vault_secret_info",
		"spec.authentication.cookie_params.auth_hmac.sec_key.wingman_secret_info",
		"spec.buffer_policy.max_request_time",
		"spec.cookies_to_modify.#",
		"spec.cors_policy.max_age",
		"spec.dns_proxy_configuration",
		"spec.domain_cert_map",
		"spec.masking_config",
		"spec.max_direct_response_body_size",
		"spec.request_cookies_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.request_cookies_to_add.#.secret_value.secret_encoding_type",
		"spec.request_cookies_to_add.#.secret_value.vault_secret_info",
		"spec.request_cookies_to_add.#.secret_value.wingman_secret_info",
		"spec.request_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.request_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.request_headers_to_add.#.secret_value.vault_secret_info",
		"spec.request_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.response_cookies_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.response_cookies_to_add.#.secret_value.secret_encoding_type",
		"spec.response_cookies_to_add.#.secret_value.vault_secret_info",
		"spec.response_cookies_to_add.#.secret_value.wingman_secret_info",
		"spec.response_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.response_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.response_headers_to_add.#.secret_value.vault_secret_info",
		"spec.response_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.retry_policy.retry_on",
		"spec.temporary_user_blocking",
		"spec.tls_cert_params.crl.#",
		"spec.tls_cert_params.require_client_certificate",
		"spec.tls_cert_params.validation_params.use_volterra_trusted_ca_url",
		"spec.tls_parameters.common_params.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.tls_parameters.common_params.tls_certificates.#.private_key.secret_encoding_type",
		"spec.tls_parameters.common_params.tls_certificates.#.private_key.vault_secret_info",
		"spec.tls_parameters.common_params.tls_certificates.#.private_key.wingman_secret_info",
		"spec.tls_parameters.common_params.trusted_ca_url",
		"spec.tls_parameters.common_params.validation_params.use_volterra_trusted_ca_url",
		"spec.tls_parameters.crl.#",
		"spec.tls_parameters.require_client_certificate",
		"spec.ztna_proxy_configurations",
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.virtual_host.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.cdn_service.cdn_origin_pool.origin_servers.#.public_ip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.cdn_service.content_choice",
			AllowedEnvironments: []string{"demo1"},
		},
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.virtual_host.API.Create"] = "ves.io.schema.virtual_host.CreateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.virtual_host.API.Get"] = []string{
		"spec.cdn_service.cache_ttl",
		"spec.cdn_service.service_domains.#",
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.virtual_host.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.cdn_service.cdn_origin_pool.origin_servers.#.public_ip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.cdn_service.content_choice",
			AllowedEnvironments: []string{"demo1"},
		},
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.virtual_host.API.List"] = []string{
		"items.#.get_spec.cdn_service.cache_ttl",
		"items.#.get_spec.cdn_service.service_domains.#",
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.virtual_host.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.cdn_service.cdn_origin_pool.origin_servers.#.public_ip.ipv6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.cdn_service.content_choice",
			AllowedEnvironments: []string{"demo1"},
		},
	}

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.virtual_host.API.Replace"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.virtual_host.ReplaceRequest.spec.api_spec.open_api_validation_choice.enable_open_api_validation",
			AddonServices: []string{"f5xc-waap-advanced"},
		},
		{
			FieldPath:     "ves.io.schema.virtual_host.ReplaceRequest.spec.masking_config.masking_choice.enable_masking",
			AddonServices: []string{"f5xc-waap-advanced"},
		},
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.virtual_host.API.Replace"] = []string{
		"spec.api_spec",
		"spec.authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info_internal",
		"spec.authentication.cookie_params.auth_hmac.prim_key.secret_encoding_type",
		"spec.authentication.cookie_params.auth_hmac.prim_key.vault_secret_info",
		"spec.authentication.cookie_params.auth_hmac.prim_key.wingman_secret_info",
		"spec.authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info_internal",
		"spec.authentication.cookie_params.auth_hmac.sec_key.secret_encoding_type",
		"spec.authentication.cookie_params.auth_hmac.sec_key.vault_secret_info",
		"spec.authentication.cookie_params.auth_hmac.sec_key.wingman_secret_info",
		"spec.buffer_policy.max_request_time",
		"spec.cookies_to_modify.#",
		"spec.cors_policy.max_age",
		"spec.dns_proxy_configuration",
		"spec.domain_cert_map",
		"spec.masking_config",
		"spec.max_direct_response_body_size",
		"spec.request_cookies_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.request_cookies_to_add.#.secret_value.secret_encoding_type",
		"spec.request_cookies_to_add.#.secret_value.vault_secret_info",
		"spec.request_cookies_to_add.#.secret_value.wingman_secret_info",
		"spec.request_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.request_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.request_headers_to_add.#.secret_value.vault_secret_info",
		"spec.request_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.response_cookies_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.response_cookies_to_add.#.secret_value.secret_encoding_type",
		"spec.response_cookies_to_add.#.secret_value.vault_secret_info",
		"spec.response_cookies_to_add.#.secret_value.wingman_secret_info",
		"spec.response_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.response_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.response_headers_to_add.#.secret_value.vault_secret_info",
		"spec.response_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.retry_policy.retry_on",
		"spec.temporary_user_blocking",
		"spec.tls_cert_params.crl.#",
		"spec.tls_cert_params.require_client_certificate",
		"spec.tls_cert_params.validation_params.use_volterra_trusted_ca_url",
		"spec.tls_parameters.common_params.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.tls_parameters.common_params.tls_certificates.#.private_key.secret_encoding_type",
		"spec.tls_parameters.common_params.tls_certificates.#.private_key.vault_secret_info",
		"spec.tls_parameters.common_params.tls_certificates.#.private_key.wingman_secret_info",
		"spec.tls_parameters.common_params.trusted_ca_url",
		"spec.tls_parameters.common_params.validation_params.use_volterra_trusted_ca_url",
		"spec.tls_parameters.crl.#",
		"spec.tls_parameters.require_client_certificate",
		"spec.ztna_proxy_configurations",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.virtual_host.API.Replace"] = "ves.io.schema.virtual_host.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.virtual_host.ApiepCustomAPI"] = "ml/data"
	sm["ves.io.schema.virtual_host.API"] = "config"
	sm["ves.io.schema.virtual_host.CustomAPI"] = "config"

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

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.virtual_host.Object"] = ApiepCustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.virtual_host.ApiepCustomAPI"] = NewApiepCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.virtual_host.ApiepCustomAPI"] = NewApiepCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.virtual_host.ApiepCustomAPI"] = RegisterApiepCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.virtual_host.ApiepCustomAPI"] = RegisterGwApiepCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.virtual_host.ApiepCustomAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewApiepCustomAPIServer(svc)
		}

	}()

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.virtual_host.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.virtual_host.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.virtual_host.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.virtual_host.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.virtual_host.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.virtual_host.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.virtual_host.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.virtual_host.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.virtual_host.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.virtual_host.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.virtual_host.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.virtual_host.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.virtual_host.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
