// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package api_crawler

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.api_sec.api_crawler.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.api_sec.api_crawler.Object"] = ObjectValidator()
	vr["ves.io.schema.api_sec.api_crawler.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.api_sec.api_crawler.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.api_sec.api_crawler.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.api_sec.api_crawler.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.api_sec.api_crawler.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.api_sec.api_crawler.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.api_sec.api_crawler.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.api_sec.api_crawler.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.api_sec.api_crawler.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.api_sec.api_crawler.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.api_sec.api_crawler.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.api_sec.api_crawler.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.api_sec.api_crawler.DomainConfiguration"] = DomainConfigurationValidator()
	vr["ves.io.schema.api_sec.api_crawler.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.api_sec.api_crawler.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.api_sec.api_crawler.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.api_sec.api_crawler.SimpleLogin"] = SimpleLoginValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.api_sec.api_crawler.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.api_sec.api_crawler.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.api_sec.api_crawler.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.api_sec.api_crawler.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.api_sec.api_crawler.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.api_sec.api_crawler.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.api_sec.api_crawler.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.api_sec.api_crawler.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.api_sec.api_crawler.API.Create"] = []string{
		"spec.domains.#.simple_login.password.blindfold_secret_info_internal",
		"spec.domains.#.simple_login.password.secret_encoding_type",
		"spec.domains.#.simple_login.password.vault_secret_info",
		"spec.domains.#.simple_login.password.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.api_sec.api_crawler.API.Create"] = "ves.io.schema.api_sec.api_crawler.CreateRequest"

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.api_sec.api_crawler.API.Replace"] = []string{
		"spec.domains.#.simple_login.password.blindfold_secret_info_internal",
		"spec.domains.#.simple_login.password.secret_encoding_type",
		"spec.domains.#.simple_login.password.vault_secret_info",
		"spec.domains.#.simple_login.password.wingman_secret_info",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.api_sec.api_crawler.API.Replace"] = "ves.io.schema.api_sec.api_crawler.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.api_sec.api_crawler.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.api_sec.api_crawler.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.api_sec.api_crawler.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.api_sec.api_crawler.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.api_sec.api_crawler.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.api_sec.api_crawler.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.api_sec.api_crawler.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.api_sec.api_crawler.Object"] = NewCRUDAPIServer

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
