// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package api_credential

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.api_credential.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.api_credential.Object"] = ObjectValidator()
	vr["ves.io.schema.api_credential.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.api_credential.ValidateTokenRequest"] = ValidateTokenRequestValidator()
	vr["ves.io.schema.api_credential.ValidateTokenResponse"] = ValidateTokenResponseValidator()

	vr["ves.io.schema.api_credential.ApiCertificateType"] = ApiCertificateTypeValidator()
	vr["ves.io.schema.api_credential.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.api_credential.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.api_credential.CreateServiceCredentialsRequest"] = CreateServiceCredentialsRequestValidator()
	vr["ves.io.schema.api_credential.CustomCreateSpecType"] = CustomCreateSpecTypeValidator()
	vr["ves.io.schema.api_credential.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.api_credential.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.api_credential.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.api_credential.GetServiceCredentialsResponse"] = GetServiceCredentialsResponseValidator()
	vr["ves.io.schema.api_credential.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.api_credential.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.api_credential.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.api_credential.RecreateScimTokenRequest"] = RecreateScimTokenRequestValidator()
	vr["ves.io.schema.api_credential.RenewRequest"] = RenewRequestValidator()
	vr["ves.io.schema.api_credential.ReplaceServiceCredentialsRequest"] = ReplaceServiceCredentialsRequestValidator()
	vr["ves.io.schema.api_credential.ReplaceServiceCredentialsResponse"] = ReplaceServiceCredentialsResponseValidator()
	vr["ves.io.schema.api_credential.ScimTokenRequest"] = ScimTokenRequestValidator()
	vr["ves.io.schema.api_credential.SiteKubeconfigType"] = SiteKubeconfigTypeValidator()
	vr["ves.io.schema.api_credential.StatusResponse"] = StatusResponseValidator()
	vr["ves.io.schema.api_credential.Vk8sKubeconfigType"] = Vk8SKubeconfigTypeValidator()

	vr["ves.io.schema.api_credential.GlobalSpecType"] = GlobalSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.api_credential.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.api_credential.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.api_credential.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.api_credential.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.api_credential.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.api_credential.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.api_credential.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.api_credential.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.api_credential.CustomAPI.Create"] = "ves.io.schema.api_credential.CreateRequest"

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.api_credential.CustomAPI.CreateServiceCredentials"] = "ves.io.schema.api_credential.CreateServiceCredentialsRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.api_credential.CustomAPI"] = "web"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

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

		customCSR.SwaggerRegistry["ves.io.schema.api_credential.Object"] = CustomPrivateAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.api_credential.CustomPrivateAPI"] = NewCustomPrivateAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.api_credential.CustomPrivateAPI"] = NewCustomPrivateAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.api_credential.CustomPrivateAPI"] = RegisterCustomPrivateAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.api_credential.CustomPrivateAPI"] = RegisterGwCustomPrivateAPIHandler
		customCSR.ServerRegistry["ves.io.schema.api_credential.CustomPrivateAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomPrivateAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.api_credential.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.api_credential.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.api_credential.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.api_credential.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.api_credential.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.api_credential.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
