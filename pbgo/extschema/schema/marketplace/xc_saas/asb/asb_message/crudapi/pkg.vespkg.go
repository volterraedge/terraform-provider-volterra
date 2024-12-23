// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package crudapi

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectCreateReq"] = ObjectCreateReqValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectCreateRsp"] = ObjectCreateRspValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectDeleteReq"] = ObjectDeleteReqValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectDeleteRsp"] = ObjectDeleteRspValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectGetReq"] = ObjectGetReqValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectGetRsp"] = ObjectGetRspValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectListReq"] = ObjectListReqValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectListRsp"] = ObjectListRspValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectListRspItem"] = ObjectListRspItemValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectReplaceReq"] = ObjectReplaceReqValidator()
	vr["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.ObjectReplaceRsp"] = ObjectReplaceRspValidator()

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
		csr.CRUDSwaggerRegistry["ves.io.schema.marketplace.xc_saas.asb.asb_message.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.marketplace.xc_saas.asb.asb_message.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.marketplace.xc_saas.asb.asb_message.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.marketplace.xc_saas.asb.asb_message.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.marketplace.xc_saas.asb.asb_message.crudapi.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.marketplace.xc_saas.asb.asb_message.Object"] = NewCRUDAPIServer

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
