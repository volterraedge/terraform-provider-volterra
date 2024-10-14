// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package dns_zone

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.dns_zone.DNSSECMode"] = DNSSECModeValidator()
	vr["ves.io.schema.dns_zone.DNSSECModeEnable"] = DNSSECModeEnableValidator()
	vr["ves.io.schema.dns_zone.DNSSECStatus"] = DNSSECStatusValidator()
	vr["ves.io.schema.dns_zone.DSRecord"] = DSRecordValidator()

	vr["ves.io.schema.dns_zone.DNSZoneStatus"] = DNSZoneStatusValidator()
	vr["ves.io.schema.dns_zone.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.dns_zone.Object"] = ObjectValidator()
	vr["ves.io.schema.dns_zone.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.dns_zone.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.dns_zone.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.dns_zone.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.dns_zone.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.dns_zone.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.dns_zone.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.dns_zone.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.dns_zone.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.dns_zone.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.dns_zone.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.dns_zone.DnsZoneMetricsRequest"] = DnsZoneMetricsRequestValidator()
	vr["ves.io.schema.dns_zone.DnsZoneMetricsResponse"] = DnsZoneMetricsResponseValidator()
	vr["ves.io.schema.dns_zone.DnsZoneRequestLogRequest"] = DnsZoneRequestLogRequestValidator()
	vr["ves.io.schema.dns_zone.DnsZoneRequestLogResponse"] = DnsZoneRequestLogResponseValidator()
	vr["ves.io.schema.dns_zone.DnsZoneRequestLogsResponseData"] = DnsZoneRequestLogsResponseDataValidator()
	vr["ves.io.schema.dns_zone.MetricsData"] = MetricsDataValidator()

	vr["ves.io.schema.dns_zone.CloneReq"] = CloneReqValidator()
	vr["ves.io.schema.dns_zone.CloneResp"] = CloneRespValidator()
	vr["ves.io.schema.dns_zone.ExportZoneFileRequest"] = ExportZoneFileRequestValidator()
	vr["ves.io.schema.dns_zone.ExportZoneFileResponse"] = ExportZoneFileResponseValidator()
	vr["ves.io.schema.dns_zone.F5CSDNSZoneConfiguration"] = F5CSDNSZoneConfigurationValidator()
	vr["ves.io.schema.dns_zone.GetLocalZoneFileRequest"] = GetLocalZoneFileRequestValidator()
	vr["ves.io.schema.dns_zone.GetLocalZoneFileResponse"] = GetLocalZoneFileResponseValidator()
	vr["ves.io.schema.dns_zone.GetRemoteZoneFileRequest"] = GetRemoteZoneFileRequestValidator()
	vr["ves.io.schema.dns_zone.GetRemoteZoneFileResponse"] = GetRemoteZoneFileResponseValidator()
	vr["ves.io.schema.dns_zone.ImportAXFRRequest"] = ImportAXFRRequestValidator()
	vr["ves.io.schema.dns_zone.ImportAXFRResponse"] = ImportAXFRResponseValidator()
	vr["ves.io.schema.dns_zone.ImportBINDCreateRequest"] = ImportBINDCreateRequestValidator()
	vr["ves.io.schema.dns_zone.ImportBINDResponse"] = ImportBINDResponseValidator()
	vr["ves.io.schema.dns_zone.ImportBINDValidateRequest"] = ImportBINDValidateRequestValidator()
	vr["ves.io.schema.dns_zone.ImportF5CSZoneRequest"] = ImportF5CSZoneRequestValidator()
	vr["ves.io.schema.dns_zone.ImportF5CSZoneResponse"] = ImportF5CSZoneResponseValidator()
	vr["ves.io.schema.dns_zone.InvalidZone"] = InvalidZoneValidator()
	vr["ves.io.schema.dns_zone.TSIGConfiguration"] = TSIGConfigurationValidator()
	vr["ves.io.schema.dns_zone.ValidZone"] = ValidZoneValidator()

	vr["ves.io.schema.dns_zone.AFSDBRecordValue"] = AFSDBRecordValueValidator()
	vr["ves.io.schema.dns_zone.CERTRecordValue"] = CERTRecordValueValidator()
	vr["ves.io.schema.dns_zone.CERTResourceRecord"] = CERTResourceRecordValidator()
	vr["ves.io.schema.dns_zone.CertificationAuthorityAuthorization"] = CertificationAuthorityAuthorizationValidator()
	vr["ves.io.schema.dns_zone.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.dns_zone.DLVResourceRecord"] = DLVResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSAAAAResourceRecord"] = DNSAAAAResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSAFSDBRecord"] = DNSAFSDBRecordValidator()
	vr["ves.io.schema.dns_zone.DNSAResourceRecord"] = DNSAResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSAliasResourceRecord"] = DNSAliasResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSCAAResourceRecord"] = DNSCAAResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSCDSRecord"] = DNSCDSRecordValidator()
	vr["ves.io.schema.dns_zone.DNSCNAMEResourceRecord"] = DNSCNAMEResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSDSRecord"] = DNSDSRecordValidator()
	vr["ves.io.schema.dns_zone.DNSEUI48ResourceRecord"] = DNSEUI48ResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSEUI64ResourceRecord"] = DNSEUI64ResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSLBResourceRecord"] = DNSLBResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSLOCResourceRecord"] = DNSLOCResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSMXResourceRecord"] = DNSMXResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSNAPTRResourceRecord"] = DNSNAPTRResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSNSResourceRecord"] = DNSNSResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSPTRResourceRecord"] = DNSPTRResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSSRVResourceRecord"] = DNSSRVResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DNSTXTResourceRecord"] = DNSTXTResourceRecordValidator()
	vr["ves.io.schema.dns_zone.DSRecordValue"] = DSRecordValueValidator()
	vr["ves.io.schema.dns_zone.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.dns_zone.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.dns_zone.LOCValue"] = LOCValueValidator()
	vr["ves.io.schema.dns_zone.MailExchanger"] = MailExchangerValidator()
	vr["ves.io.schema.dns_zone.NAPTRValue"] = NAPTRValueValidator()
	vr["ves.io.schema.dns_zone.PrimaryDNSConfig"] = PrimaryDNSConfigValidator()
	vr["ves.io.schema.dns_zone.PrimaryDNSCreateSpecType"] = PrimaryDNSCreateSpecTypeValidator()
	vr["ves.io.schema.dns_zone.PrimaryDNSGetSpecType"] = PrimaryDNSGetSpecTypeValidator()
	vr["ves.io.schema.dns_zone.RRSet"] = RRSetValidator()
	vr["ves.io.schema.dns_zone.RRSetGroup"] = RRSetGroupValidator()
	vr["ves.io.schema.dns_zone.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.dns_zone.SHA1Digest"] = SHA1DigestValidator()
	vr["ves.io.schema.dns_zone.SHA1Fingerprint"] = SHA1FingerprintValidator()
	vr["ves.io.schema.dns_zone.SHA256Digest"] = SHA256DigestValidator()
	vr["ves.io.schema.dns_zone.SHA256Fingerprint"] = SHA256FingerprintValidator()
	vr["ves.io.schema.dns_zone.SHA384Digest"] = SHA384DigestValidator()
	vr["ves.io.schema.dns_zone.SOARecordParameterConfig"] = SOARecordParameterConfigValidator()
	vr["ves.io.schema.dns_zone.SRVService"] = SRVServiceValidator()
	vr["ves.io.schema.dns_zone.SSHFPRecordValue"] = SSHFPRecordValueValidator()
	vr["ves.io.schema.dns_zone.SSHFPResourceRecord"] = SSHFPResourceRecordValidator()
	vr["ves.io.schema.dns_zone.SecondaryDNSConfig"] = SecondaryDNSConfigValidator()
	vr["ves.io.schema.dns_zone.SecondaryDNSCreateSpecType"] = SecondaryDNSCreateSpecTypeValidator()
	vr["ves.io.schema.dns_zone.SecondaryDNSGetSpecType"] = SecondaryDNSGetSpecTypeValidator()
	vr["ves.io.schema.dns_zone.TLSARecordValue"] = TLSARecordValueValidator()
	vr["ves.io.schema.dns_zone.TLSAResourceRecord"] = TLSAResourceRecordValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

	mdr.EntryFactory["ves.io.schema.dns_zone.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.dns_zone.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dns_zone.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.dns_zone.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.dns_zone.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.dns_zone.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.dns_zone.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.dns_zone.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedRequestFieldsRegistry["ves.io.schema.dns_zone.API.Create"] = []string{
		"spec.primary.default_rr_set_group.#.dlv_record",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprint",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprinttype",
		"spec.primary.rr_set_group.#.rr_set.#.dlv_record",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprint",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprinttype",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.dns_zone.API.Create"] = []string{
		"spec.primary.default_rr_set_group.#.dlv_record",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprint",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprinttype",
		"spec.primary.rr_set_group.#.rr_set.#.dlv_record",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprint",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprinttype",
		"spec.secondary.zone_file",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.dns_zone.API.Create"] = []string{
		"spec.primary.default_rr_set_group.#.alias_record.name",
		"spec.primary.rr_set_group.#.metadata.disable",
		"spec.primary.rr_set_group.#.rr_set.#.alias_record.name",
		"spec.secondary.tsig_key_value.blindfold_secret_info_internal",
		"spec.secondary.tsig_key_value.secret_encoding_type",
		"spec.secondary.tsig_key_value.vault_secret_info",
		"spec.secondary.tsig_key_value.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.dns_zone.API.Create"] = "ves.io.schema.dns_zone.CreateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.dns_zone.API.Get"] = []string{
		"create_form.spec.primary.default_rr_set_group.#.dlv_record",
		"create_form.spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprint",
		"create_form.spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprinttype",
		"create_form.spec.primary.rr_set_group.#.rr_set.#.dlv_record",
		"create_form.spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprint",
		"create_form.spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprinttype",
		"replace_form.spec.primary.default_rr_set_group.#.dlv_record",
		"replace_form.spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprint",
		"replace_form.spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprinttype",
		"replace_form.spec.primary.rr_set_group.#.rr_set.#.dlv_record",
		"replace_form.spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprint",
		"replace_form.spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprinttype",
		"spec.primary.default_rr_set_group.#.dlv_record",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprint",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprinttype",
		"spec.primary.rr_set_group.#.rr_set.#.dlv_record",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprint",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprinttype",
		"spec.secondary.zone_file",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.dns_zone.API.List"] = []string{
		"items.#.get_spec.primary.default_rr_set_group.#.dlv_record",
		"items.#.get_spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprint",
		"items.#.get_spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprinttype",
		"items.#.get_spec.primary.rr_set_group.#.rr_set.#.dlv_record",
		"items.#.get_spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprint",
		"items.#.get_spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprinttype",
		"items.#.get_spec.secondary.zone_file",
	}

	mdr.RPCDeprecatedRequestFieldsRegistry["ves.io.schema.dns_zone.API.Replace"] = []string{
		"spec.primary.default_rr_set_group.#.dlv_record",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprint",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprinttype",
		"spec.primary.rr_set_group.#.rr_set.#.dlv_record",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprint",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprinttype",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.dns_zone.API.Replace"] = []string{
		"spec.primary.default_rr_set_group.#.alias_record.name",
		"spec.primary.rr_set_group.#.metadata.disable",
		"spec.primary.rr_set_group.#.rr_set.#.alias_record.name",
		"spec.secondary.tsig_key_value.blindfold_secret_info_internal",
		"spec.secondary.tsig_key_value.secret_encoding_type",
		"spec.secondary.tsig_key_value.vault_secret_info",
		"spec.secondary.tsig_key_value.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.dns_zone.API.Replace"] = "ves.io.schema.dns_zone.ReplaceRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.dns_zone.CustomAPI.ImportAXFR"] = []string{
		"configuration.default_rr_set_group.#.dlv_record",
		"configuration.default_rr_set_group.#.sshfp_record.values.#.fingerprint",
		"configuration.default_rr_set_group.#.sshfp_record.values.#.fingerprinttype",
		"configuration.rr_set_group.#.rr_set.#.dlv_record",
		"configuration.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprint",
		"configuration.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprinttype",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.dns_zone.CustomAPI.ImportAXFR"] = []string{
		"tsig_configuration.tsig_key_value.blindfold_secret_info_internal",
		"tsig_configuration.tsig_key_value.secret_encoding_type",
		"tsig_configuration.tsig_key_value.vault_secret_info",
		"tsig_configuration.tsig_key_value.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.dns_zone.CustomAPI.ImportAXFR"] = "ves.io.schema.dns_zone.ImportAXFRRequest"

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.dns_zone.CustomAPI.ImportBINDCreate"] = "ves.io.schema.dns_zone.ImportBINDCreateRequest"

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.dns_zone.CustomAPI.ImportBINDValidate"] = "ves.io.schema.dns_zone.ImportBINDValidateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.dns_zone.CustomAPI.ImportF5CSZone"] = []string{
		"spec.primary.default_rr_set_group.#.dlv_record",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprint",
		"spec.primary.default_rr_set_group.#.sshfp_record.values.#.fingerprinttype",
		"spec.primary.rr_set_group.#.rr_set.#.dlv_record",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprint",
		"spec.primary.rr_set_group.#.rr_set.#.sshfp_record.values.#.fingerprinttype",
		"spec.secondary.zone_file",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.dns_zone.API"] = "config/dns"
	sm["ves.io.schema.dns_zone.CustomDataAPI"] = "data"
	sm["ves.io.schema.dns_zone.CustomAPI"] = "config/dns"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.dns_zone.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.dns_zone.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.dns_zone.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.dns_zone.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.dns_zone.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.dns_zone.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.dns_zone.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.dns_zone.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.dns_zone.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.dns_zone.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.dns_zone.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.dns_zone.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.dns_zone.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomDataAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.dns_zone.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.dns_zone.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.dns_zone.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.dns_zone.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.dns_zone.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.dns_zone.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
