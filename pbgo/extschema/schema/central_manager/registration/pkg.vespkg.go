// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package registration

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.central_manager.registration.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.central_manager.registration.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.central_manager.registration.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.central_manager.registration.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.central_manager.registration.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.central_manager.registration.UpdateStatusReq"] = UpdateStatusReqValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.central_manager.registration.CustomAPI"] = "clearview"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		customCSR.SwaggerRegistry["ves.io.schema.central_manager.registration.CustomAPI"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.central_manager.registration.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.central_manager.registration.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.central_manager.registration.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.central_manager.registration.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.central_manager.registration.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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