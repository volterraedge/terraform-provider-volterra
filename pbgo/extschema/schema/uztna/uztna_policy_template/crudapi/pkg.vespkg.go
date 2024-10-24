// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package crudapi

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectCreateReq"] = ObjectCreateReqValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectCreateRsp"] = ObjectCreateRspValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectDeleteReq"] = ObjectDeleteReqValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectDeleteRsp"] = ObjectDeleteRspValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectGetReq"] = ObjectGetReqValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectGetRsp"] = ObjectGetRspValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectListReq"] = ObjectListReqValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectListRsp"] = ObjectListRspValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectListRspItem"] = ObjectListRspItemValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectReplaceReq"] = ObjectReplaceReqValidator()
	vr["ves.io.schema.uztna.uztna_policy_template.crudapi.ObjectReplaceRsp"] = ObjectReplaceRspValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

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
		csr.CRUDSwaggerRegistry["ves.io.schema.uztna.uztna_policy_template.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.uztna.uztna_policy_template.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.uztna.uztna_policy_template.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.uztna.uztna_policy_template.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.uztna.uztna_policy_template.crudapi.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.uztna.uztna_policy_template.crudapi.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.uztna.uztna_policy_template.Object"] = NewCRUDAPIServer

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
