// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cloud_credentials

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.cloud_credentials.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.cloud_credentials.Object"] = ObjectValidator()
	vr["ves.io.schema.cloud_credentials.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.cloud_credentials.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.cloud_credentials.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.cloud_credentials.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.cloud_credentials.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.cloud_credentials.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.cloud_credentials.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.cloud_credentials.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.cloud_credentials.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.cloud_credentials.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.cloud_credentials.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.cloud_credentials.AWSAssumeRoleType"] = AWSAssumeRoleTypeValidator()
	vr["ves.io.schema.cloud_credentials.AWSSecretType"] = AWSSecretTypeValidator()
	vr["ves.io.schema.cloud_credentials.AzurePfxType"] = AzurePfxTypeValidator()
	vr["ves.io.schema.cloud_credentials.AzureSecretType"] = AzureSecretTypeValidator()
	vr["ves.io.schema.cloud_credentials.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.cloud_credentials.GCPCredFileType"] = GCPCredFileTypeValidator()
	vr["ves.io.schema.cloud_credentials.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.cloud_credentials.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.cloud_credentials.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.cloud_credentials.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.cloud_credentials.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.cloud_credentials.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.cloud_credentials.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.cloud_credentials.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.cloud_credentials.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.cloud_credentials.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.cloud_credentials.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.cloud_credentials.API.Create"] = []string{
		"spec.aws_secret_key.secret_key.blindfold_secret_info_internal",
		"spec.aws_secret_key.secret_key.secret_encoding_type",
		"spec.aws_secret_key.secret_key.vault_secret_info",
		"spec.aws_secret_key.secret_key.wingman_secret_info",
		"spec.azure_client_secret.client_secret.blindfold_secret_info_internal",
		"spec.azure_client_secret.client_secret.secret_encoding_type",
		"spec.azure_client_secret.client_secret.vault_secret_info",
		"spec.azure_client_secret.client_secret.wingman_secret_info",
		"spec.azure_pfx_certificate.password.blindfold_secret_info_internal",
		"spec.azure_pfx_certificate.password.secret_encoding_type",
		"spec.azure_pfx_certificate.password.vault_secret_info",
		"spec.azure_pfx_certificate.password.wingman_secret_info",
		"spec.gcp_cred_file.credential_file.blindfold_secret_info_internal",
		"spec.gcp_cred_file.credential_file.secret_encoding_type",
		"spec.gcp_cred_file.credential_file.vault_secret_info",
		"spec.gcp_cred_file.credential_file.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.cloud_credentials.API.Create"] = "ves.io.schema.cloud_credentials.CreateRequest"

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.cloud_credentials.API.Replace"] = []string{
		"spec.aws_secret_key.secret_key.blindfold_secret_info_internal",
		"spec.aws_secret_key.secret_key.secret_encoding_type",
		"spec.aws_secret_key.secret_key.vault_secret_info",
		"spec.aws_secret_key.secret_key.wingman_secret_info",
		"spec.azure_client_secret.client_secret.blindfold_secret_info_internal",
		"spec.azure_client_secret.client_secret.secret_encoding_type",
		"spec.azure_client_secret.client_secret.vault_secret_info",
		"spec.azure_client_secret.client_secret.wingman_secret_info",
		"spec.azure_pfx_certificate.password.blindfold_secret_info_internal",
		"spec.azure_pfx_certificate.password.secret_encoding_type",
		"spec.azure_pfx_certificate.password.vault_secret_info",
		"spec.azure_pfx_certificate.password.wingman_secret_info",
		"spec.gcp_cred_file.credential_file.blindfold_secret_info_internal",
		"spec.gcp_cred_file.credential_file.secret_encoding_type",
		"spec.gcp_cred_file.credential_file.vault_secret_info",
		"spec.gcp_cred_file.credential_file.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.cloud_credentials.API.Replace"] = "ves.io.schema.cloud_credentials.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.cloud_credentials.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.cloud_credentials.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.cloud_credentials.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.cloud_credentials.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.cloud_credentials.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.cloud_credentials.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.cloud_credentials.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.cloud_credentials.Object"] = NewCRUDAPIServer

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
