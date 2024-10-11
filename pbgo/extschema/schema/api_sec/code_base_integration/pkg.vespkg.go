// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package code_base_integration

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.api_sec.code_base_integration.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.api_sec.code_base_integration.Object"] = ObjectValidator()
	vr["ves.io.schema.api_sec.code_base_integration.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.api_sec.code_base_integration.UpdateCodeBaseIntegrationReq"] = UpdateCodeBaseIntegrationReqValidator()
	vr["ves.io.schema.api_sec.code_base_integration.UpdateCodeBaseIntegrationResp"] = UpdateCodeBaseIntegrationRespValidator()

	vr["ves.io.schema.api_sec.code_base_integration.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.api_sec.code_base_integration.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.api_sec.code_base_integration.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.api_sec.code_base_integration.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.api_sec.code_base_integration.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.api_sec.code_base_integration.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.api_sec.code_base_integration.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.api_sec.code_base_integration.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.api_sec.code_base_integration.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.api_sec.code_base_integration.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.api_sec.code_base_integration.AzureReposIntegration"] = AzureReposIntegrationValidator()
	vr["ves.io.schema.api_sec.code_base_integration.BitBucketCloudIntegration"] = BitBucketCloudIntegrationValidator()
	vr["ves.io.schema.api_sec.code_base_integration.BitBucketServerIntegration"] = BitBucketServerIntegrationValidator()
	vr["ves.io.schema.api_sec.code_base_integration.CodeBaseIntegration"] = CodeBaseIntegrationValidator()
	vr["ves.io.schema.api_sec.code_base_integration.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.api_sec.code_base_integration.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.api_sec.code_base_integration.GithubEnterpriseIntegration"] = GithubEnterpriseIntegrationValidator()
	vr["ves.io.schema.api_sec.code_base_integration.GithubIntegration"] = GithubIntegrationValidator()
	vr["ves.io.schema.api_sec.code_base_integration.GitlabCloudIntegration"] = GitlabCloudIntegrationValidator()
	vr["ves.io.schema.api_sec.code_base_integration.GitlabEnterpriseIntegration"] = GitlabEnterpriseIntegrationValidator()
	vr["ves.io.schema.api_sec.code_base_integration.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.api_sec.code_base_integration.IntegrationHealth"] = IntegrationHealthValidator()
	vr["ves.io.schema.api_sec.code_base_integration.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.api_sec.code_base_integration.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.api_sec.code_base_integration.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.api_sec.code_base_integration.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.api_sec.code_base_integration.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.api_sec.code_base_integration.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.api_sec.code_base_integration.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.api_sec.code_base_integration.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.api_sec.code_base_integration.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.api_sec.code_base_integration.API.Create"] = []string{
		"spec.code_base_integration.azure_repos.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.azure_repos.access_token.secret_encoding_type",
		"spec.code_base_integration.azure_repos.access_token.vault_secret_info",
		"spec.code_base_integration.azure_repos.access_token.wingman_secret_info",
		"spec.code_base_integration.bitbucket.passwd.blindfold_secret_info_internal",
		"spec.code_base_integration.bitbucket.passwd.secret_encoding_type",
		"spec.code_base_integration.bitbucket.passwd.vault_secret_info",
		"spec.code_base_integration.bitbucket.passwd.wingman_secret_info",
		"spec.code_base_integration.bitbucket_server.passwd.blindfold_secret_info_internal",
		"spec.code_base_integration.bitbucket_server.passwd.secret_encoding_type",
		"spec.code_base_integration.bitbucket_server.passwd.vault_secret_info",
		"spec.code_base_integration.bitbucket_server.passwd.wingman_secret_info",
		"spec.code_base_integration.github.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.github.access_token.secret_encoding_type",
		"spec.code_base_integration.github.access_token.vault_secret_info",
		"spec.code_base_integration.github.access_token.wingman_secret_info",
		"spec.code_base_integration.github_enterprise.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.github_enterprise.access_token.secret_encoding_type",
		"spec.code_base_integration.github_enterprise.access_token.vault_secret_info",
		"spec.code_base_integration.github_enterprise.access_token.wingman_secret_info",
		"spec.code_base_integration.gitlab.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.gitlab.access_token.secret_encoding_type",
		"spec.code_base_integration.gitlab.access_token.vault_secret_info",
		"spec.code_base_integration.gitlab.access_token.wingman_secret_info",
		"spec.code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.gitlab_enterprise.access_token.secret_encoding_type",
		"spec.code_base_integration.gitlab_enterprise.access_token.vault_secret_info",
		"spec.code_base_integration.gitlab_enterprise.access_token.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.api_sec.code_base_integration.API.Create"] = "ves.io.schema.api_sec.code_base_integration.CreateRequest"

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.api_sec.code_base_integration.API.Replace"] = []string{
		"spec.code_base_integration.azure_repos.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.azure_repos.access_token.secret_encoding_type",
		"spec.code_base_integration.azure_repos.access_token.vault_secret_info",
		"spec.code_base_integration.azure_repos.access_token.wingman_secret_info",
		"spec.code_base_integration.bitbucket.passwd.blindfold_secret_info_internal",
		"spec.code_base_integration.bitbucket.passwd.secret_encoding_type",
		"spec.code_base_integration.bitbucket.passwd.vault_secret_info",
		"spec.code_base_integration.bitbucket.passwd.wingman_secret_info",
		"spec.code_base_integration.bitbucket_server.passwd.blindfold_secret_info_internal",
		"spec.code_base_integration.bitbucket_server.passwd.secret_encoding_type",
		"spec.code_base_integration.bitbucket_server.passwd.vault_secret_info",
		"spec.code_base_integration.bitbucket_server.passwd.wingman_secret_info",
		"spec.code_base_integration.github.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.github.access_token.secret_encoding_type",
		"spec.code_base_integration.github.access_token.vault_secret_info",
		"spec.code_base_integration.github.access_token.wingman_secret_info",
		"spec.code_base_integration.github_enterprise.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.github_enterprise.access_token.secret_encoding_type",
		"spec.code_base_integration.github_enterprise.access_token.vault_secret_info",
		"spec.code_base_integration.github_enterprise.access_token.wingman_secret_info",
		"spec.code_base_integration.gitlab.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.gitlab.access_token.secret_encoding_type",
		"spec.code_base_integration.gitlab.access_token.vault_secret_info",
		"spec.code_base_integration.gitlab.access_token.wingman_secret_info",
		"spec.code_base_integration.gitlab_enterprise.access_token.blindfold_secret_info_internal",
		"spec.code_base_integration.gitlab_enterprise.access_token.secret_encoding_type",
		"spec.code_base_integration.gitlab_enterprise.access_token.vault_secret_info",
		"spec.code_base_integration.gitlab_enterprise.access_token.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.api_sec.code_base_integration.API.Replace"] = "ves.io.schema.api_sec.code_base_integration.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.api_sec.code_base_integration.API"] = "config"

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

		customCSR.SwaggerRegistry["ves.io.schema.api_sec.code_base_integration.Object"] = PrivateCustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.api_sec.code_base_integration.PrivateCustomAPI"] = NewPrivateCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.api_sec.code_base_integration.PrivateCustomAPI"] = NewPrivateCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.api_sec.code_base_integration.PrivateCustomAPI"] = RegisterPrivateCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.api_sec.code_base_integration.PrivateCustomAPI"] = RegisterGwPrivateCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.api_sec.code_base_integration.PrivateCustomAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewPrivateCustomAPIServer(svc)
		}

	}()

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.api_sec.code_base_integration.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.api_sec.code_base_integration.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.api_sec.code_base_integration.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.api_sec.code_base_integration.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.api_sec.code_base_integration.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.api_sec.code_base_integration.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.api_sec.code_base_integration.Object"] = NewCRUDAPIServer

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