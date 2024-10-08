// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package known_label_key

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.known_label_key.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.known_label_key.Object"] = ObjectValidator()

	vr["ves.io.schema.known_label_key.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.known_label_key.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.known_label_key.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.known_label_key.DeleteResponse"] = DeleteResponseValidator()
	vr["ves.io.schema.known_label_key.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.known_label_key.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.known_label_key.LabelKeyType"] = LabelKeyTypeValidator()

	vr["ves.io.schema.known_label_key.GlobalSpecType"] = GlobalSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.known_label_key.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.known_label_key.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.known_label_key.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.known_label_key.Object"] = GetObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.known_label_key.CustomAPI"] = "config"

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

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.known_label_key.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.known_label_key.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.known_label_key.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.known_label_key.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.known_label_key.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.known_label_key.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
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
