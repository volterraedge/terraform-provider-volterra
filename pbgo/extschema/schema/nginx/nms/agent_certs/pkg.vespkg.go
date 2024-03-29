// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package agent_certs

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {

	vr["ves.io.schema.nginx.nms.agent_certs.GenerateAgentCertRequest"] = GenerateAgentCertRequestValidator()
	vr["ves.io.schema.nginx.nms.agent_certs.GenerateAgentCertResponse"] = GenerateAgentCertResponseValidator()
	vr["ves.io.schema.nginx.nms.agent_certs.RevokeAgentCertRequest"] = RevokeAgentCertRequestValidator()
	vr["ves.io.schema.nginx.nms.agent_certs.RevokeAgentCertResponse"] = RevokeAgentCertResponseValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.nginx.nms.agent_certs.CustomAPI"] = "nginx/nms"

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
		customCSR.SwaggerRegistry["ves.io.schema.nginx.nms.agent_certs.CustomAPI"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.nginx.nms.agent_certs.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.nginx.nms.agent_certs.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.nginx.nms.agent_certs.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.nginx.nms.agent_certs.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.nginx.nms.agent_certs.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
