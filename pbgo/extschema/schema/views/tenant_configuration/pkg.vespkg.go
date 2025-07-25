// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package tenant_configuration

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.views.tenant_configuration.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.views.tenant_configuration.Object"] = ObjectValidator()
	vr["ves.io.schema.views.tenant_configuration.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.views.tenant_configuration.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.views.tenant_configuration.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.views.tenant_configuration.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.views.tenant_configuration.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.views.tenant_configuration.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.views.tenant_configuration.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.views.tenant_configuration.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.views.tenant_configuration.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.views.tenant_configuration.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.views.tenant_configuration.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.views.tenant_configuration.BasicConfiguration"] = BasicConfigurationValidator()
	vr["ves.io.schema.views.tenant_configuration.BruteForceDetectionSettings"] = BruteForceDetectionSettingsValidator()
	vr["ves.io.schema.views.tenant_configuration.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.views.tenant_configuration.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.views.tenant_configuration.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.views.tenant_configuration.PasswordPolicy"] = PasswordPolicyValidator()
	vr["ves.io.schema.views.tenant_configuration.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.views.tenant_configuration.SessionManagement"] = SessionManagementValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.views.tenant_configuration.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.views.tenant_configuration.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.tenant_configuration.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.views.tenant_configuration.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.views.tenant_configuration.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.views.tenant_configuration.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.views.tenant_configuration.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.views.tenant_configuration.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.views.tenant_configuration.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.views.tenant_configuration.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.views.tenant_configuration.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.views.tenant_configuration.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.views.tenant_configuration.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.views.tenant_configuration.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.views.tenant_configuration.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.views.tenant_configuration.Object"] = NewCRUDAPIServer

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
