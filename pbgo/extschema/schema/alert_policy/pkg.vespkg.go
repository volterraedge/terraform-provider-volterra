// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package alert_policy

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.alert_policy.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.alert_policy.Object"] = ObjectValidator()
	vr["ves.io.schema.alert_policy.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.alert_policy.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.alert_policy.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.alert_policy.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.alert_policy.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.alert_policy.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.alert_policy.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.alert_policy.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.alert_policy.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.alert_policy.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.alert_policy.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.alert_policy.AlertPolicyMatch"] = AlertPolicyMatchValidator()
	vr["ves.io.schema.alert_policy.AlertPolicyMatchRequest"] = AlertPolicyMatchRequestValidator()
	vr["ves.io.schema.alert_policy.AlertPolicyMatchResponse"] = AlertPolicyMatchResponseValidator()

	vr["ves.io.schema.alert_policy.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.alert_policy.CustomGroupBy"] = CustomGroupByValidator()
	vr["ves.io.schema.alert_policy.CustomMatcher"] = CustomMatcherValidator()
	vr["ves.io.schema.alert_policy.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.alert_policy.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.alert_policy.GroupMatcher"] = GroupMatcherValidator()
	vr["ves.io.schema.alert_policy.LabelMatcher"] = LabelMatcherValidator()
	vr["ves.io.schema.alert_policy.NotificationParameters"] = NotificationParametersValidator()
	vr["ves.io.schema.alert_policy.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.alert_policy.Route"] = RouteValidator()
	vr["ves.io.schema.alert_policy.SeverityMatcher"] = SeverityMatcherValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.alert_policy.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.alert_policy.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.alert_policy.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.alert_policy.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.alert_policy.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.alert_policy.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.alert_policy.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.alert_policy.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.alert_policy.API"] = "config"
	sm["ves.io.schema.alert_policy.CustomAPI"] = "alert"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.alert_policy.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.alert_policy.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.alert_policy.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.alert_policy.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.alert_policy.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.alert_policy.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.alert_policy.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.alert_policy.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.alert_policy.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.alert_policy.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.alert_policy.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.alert_policy.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.alert_policy.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
