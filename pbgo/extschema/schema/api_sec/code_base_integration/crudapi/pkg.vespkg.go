// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package crudapi

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectCreateReq"] = ObjectCreateReqValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectCreateRsp"] = ObjectCreateRspValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectDeleteReq"] = ObjectDeleteReqValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectDeleteRsp"] = ObjectDeleteRspValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectGetReq"] = ObjectGetReqValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectGetRsp"] = ObjectGetRspValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectListReq"] = ObjectListReqValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectListRsp"] = ObjectListRspValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectListRspItem"] = ObjectListRspItemValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectReplaceReq"] = ObjectReplaceReqValidator()
	vr["ves.io.schema.api_sec.code_base_integration.crudapi.ObjectReplaceRsp"] = ObjectReplaceRspValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.api_sec.code_base_integration.crudapi.API.Create"] = "ves.io.schema.api_sec.code_base_integration.crudapi.ObjectCreateReq"

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.api_sec.code_base_integration.crudapi.API.Replace"] = "ves.io.schema.api_sec.code_base_integration.crudapi.ObjectReplaceReq"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR
	csr = mdr.PvtCRUDServiceRegistry

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
		mdr.SvcRegisterHandlers["ves.io.schema.api_sec.code_base_integration.crudapi.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.api_sec.code_base_integration.crudapi.API"] = RegisterGwAPIHandler
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