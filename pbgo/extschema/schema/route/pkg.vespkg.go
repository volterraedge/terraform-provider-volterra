// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package route

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.route.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.route.Object"] = ObjectValidator()
	vr["ves.io.schema.route.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.route.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.route.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.route.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.route.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.route.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.route.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.route.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.route.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.route.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.route.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.route.BotDefenseJavascriptInjectionType"] = BotDefenseJavascriptInjectionTypeValidator()
	vr["ves.io.schema.route.ContentRewriteType"] = ContentRewriteTypeValidator()
	vr["ves.io.schema.route.ContextExtensionInfo"] = ContextExtensionInfoValidator()
	vr["ves.io.schema.route.CookieForHashing"] = CookieForHashingValidator()
	vr["ves.io.schema.route.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.route.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.route.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.route.HashPolicyType"] = HashPolicyTypeValidator()
	vr["ves.io.schema.route.JavaScriptTag"] = JavaScriptTagValidator()
	vr["ves.io.schema.route.MirrorPolicyType"] = MirrorPolicyTypeValidator()
	vr["ves.io.schema.route.QueryParamsSimpleRoute"] = QueryParamsSimpleRouteValidator()
	vr["ves.io.schema.route.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.route.RouteDestination"] = RouteDestinationValidator()
	vr["ves.io.schema.route.RouteDestinationList"] = RouteDestinationListValidator()
	vr["ves.io.schema.route.RouteDirectResponse"] = RouteDirectResponseValidator()
	vr["ves.io.schema.route.RouteQueryParams"] = RouteQueryParamsValidator()
	vr["ves.io.schema.route.RouteRedirect"] = RouteRedirectValidator()
	vr["ves.io.schema.route.RouteType"] = RouteTypeValidator()
	vr["ves.io.schema.route.ServicePolicyInfo"] = ServicePolicyInfoValidator()
	vr["ves.io.schema.route.SpdyConfigType"] = SpdyConfigTypeValidator()
	vr["ves.io.schema.route.TagAttribute"] = TagAttributeValidator()
	vr["ves.io.schema.route.VerStatusType"] = VerStatusTypeValidator()
	vr["ves.io.schema.route.WebsocketConfigType"] = WebsocketConfigTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.route.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.route.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.route.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.route.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.route.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.route.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.route.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.route.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.route.API.Create"] = []string{
		"spec.routes.#.bot_defense_javascript_injection_inline_mode",
		"spec.routes.#.disable_custom_script",
		"spec.routes.#.request_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.routes.#.request_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.routes.#.request_headers_to_add.#.secret_value.vault_secret_info",
		"spec.routes.#.request_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.routes.#.response_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.routes.#.response_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.routes.#.response_headers_to_add.#.secret_value.vault_secret_info",
		"spec.routes.#.response_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.routes.#.route_destination.buffer_policy.max_request_time",
		"spec.routes.#.route_destination.cors_policy.max_age",
		"spec.routes.#.route_destination.retry_policy.retry_on",
		"spec.routes.#.route_redirect.all_params",
		"spec.routes.#.route_redirect.port_redirect",
		"spec.routes.#.route_redirect.strip_query_params",
		"spec.routes.#.skip_lb_override",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.route.API.Create"] = "ves.io.schema.route.CreateRequest"

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.route.API.Replace"] = []string{
		"spec.routes.#.bot_defense_javascript_injection_inline_mode",
		"spec.routes.#.disable_custom_script",
		"spec.routes.#.request_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.routes.#.request_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.routes.#.request_headers_to_add.#.secret_value.vault_secret_info",
		"spec.routes.#.request_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.routes.#.response_headers_to_add.#.secret_value.blindfold_secret_info_internal",
		"spec.routes.#.response_headers_to_add.#.secret_value.secret_encoding_type",
		"spec.routes.#.response_headers_to_add.#.secret_value.vault_secret_info",
		"spec.routes.#.response_headers_to_add.#.secret_value.wingman_secret_info",
		"spec.routes.#.route_destination.buffer_policy.max_request_time",
		"spec.routes.#.route_destination.cors_policy.max_age",
		"spec.routes.#.route_destination.retry_policy.retry_on",
		"spec.routes.#.route_redirect.all_params",
		"spec.routes.#.route_redirect.port_redirect",
		"spec.routes.#.route_redirect.strip_query_params",
		"spec.routes.#.skip_lb_override",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.route.API.Replace"] = "ves.io.schema.route.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.route.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.route.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.route.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.route.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.route.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.route.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.route.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.route.Object"] = NewCRUDAPIServer

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
