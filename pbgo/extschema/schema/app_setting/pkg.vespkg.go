// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package app_setting

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.app_setting.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.app_setting.Object"] = ObjectValidator()
	vr["ves.io.schema.app_setting.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.app_setting.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.app_setting.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.app_setting.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.app_setting.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.app_setting.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.app_setting.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.app_setting.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.app_setting.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.app_setting.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.app_setting.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.app_setting.SuspiciousUser"] = SuspiciousUserValidator()
	vr["ves.io.schema.app_setting.SuspiciousUserStatusReq"] = SuspiciousUserStatusReqValidator()
	vr["ves.io.schema.app_setting.SuspiciousUserStatusRsp"] = SuspiciousUserStatusRspValidator()

	vr["ves.io.schema.app_setting.AppTypeSettings"] = AppTypeSettingsValidator()
	vr["ves.io.schema.app_setting.BolaDetectionManualSettings"] = BolaDetectionManualSettingsValidator()
	vr["ves.io.schema.app_setting.BusinessLogicMarkupSetting"] = BusinessLogicMarkupSettingValidator()
	vr["ves.io.schema.app_setting.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.app_setting.FailedLoginActivitySetting"] = FailedLoginActivitySettingValidator()
	vr["ves.io.schema.app_setting.ForbiddenActivitySetting"] = ForbiddenActivitySettingValidator()
	vr["ves.io.schema.app_setting.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.app_setting.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.app_setting.MaliciousUserDetectionSetting"] = MaliciousUserDetectionSettingValidator()
	vr["ves.io.schema.app_setting.MetricSelector"] = MetricSelectorValidator()
	vr["ves.io.schema.app_setting.NonexistentUrlAutomaticActivitySetting"] = NonexistentUrlAutomaticActivitySettingValidator()
	vr["ves.io.schema.app_setting.NonexistentUrlCustomActivitySetting"] = NonexistentUrlCustomActivitySettingValidator()
	vr["ves.io.schema.app_setting.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.app_setting.TimeseriesAnalysesSetting"] = TimeseriesAnalysesSettingValidator()
	vr["ves.io.schema.app_setting.UserBehaviorAnalysisSetting"] = UserBehaviorAnalysisSettingValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.app_setting.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.app_setting.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.app_setting.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.app_setting.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.app_setting.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.app_setting.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.app_setting.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.app_setting.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.app_setting.API.Create"] = []string{
		"spec.anomaly_types.#",
		"spec.app_type_refs.#",
		"spec.app_type_settings.#.user_behavior_analysis_setting.enable_detection.bola_detection_manual",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.app_setting.API.Replace"] = []string{
		"spec.anomaly_types.#",
		"spec.app_type_refs.#",
		"spec.app_type_settings.#.user_behavior_analysis_setting.enable_detection.bola_detection_manual",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.app_setting.API"] = "config"
	sm["ves.io.schema.app_setting.CustomAPI"] = "ml/data"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.app_setting.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.app_setting.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.app_setting.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.app_setting.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.app_setting.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.app_setting.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.app_setting.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.app_setting.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.app_setting.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.app_setting.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.app_setting.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.app_setting.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.app_setting.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
