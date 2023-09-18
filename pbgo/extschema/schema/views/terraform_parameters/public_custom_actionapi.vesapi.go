// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package terraform_parameters

import (
	"bytes"
	"context"
	"fmt"
	io "io"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

var (
	_ = fmt.Sprintf("dummy for fmt import use")
)

// Create CustomActionAPI GRPC Client satisfying server.CustomClient
type CustomActionAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomActionAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomActionAPIGrpcClient) doRPCForceDelete(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &ForceDeleteRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.terraform_parameters.ForceDeleteRequest", yamlReq)
	}
	rsp, err := c.grpcClient.ForceDelete(ctx, req, opts...)
	return rsp, err
}

func (c *CustomActionAPIGrpcClient) doRPCRun(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &RunRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.terraform_parameters.RunRequest", yamlReq)
	}
	rsp, err := c.grpcClient.Run(ctx, req, opts...)
	return rsp, err
}

func (c *CustomActionAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}
	if cco.YAMLReq == "" {
		return nil, fmt.Errorf("Error, empty request body")
	}
	ctx = client.AddHdrsToCtx(cco.Headers, ctx)

	rsp, err := rpcFn(ctx, cco.YAMLReq, cco.GrpcCallOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using GRPC")
	}
	if cco.OutCallResponse != nil {
		cco.OutCallResponse.ProtoMsg = rsp
	}
	return rsp, nil
}

func NewCustomActionAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomActionAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomActionAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["ForceDelete"] = ccl.doRPCForceDelete

	rpcFns["Run"] = ccl.doRPCRun

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomActionAPI REST Client satisfying server.CustomClient
type CustomActionAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomActionAPIRestClient) doRPCForceDelete(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &ForceDeleteRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.terraform_parameters.ForceDeleteRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post", "put":
		jsn, err := codec.ToJSON(req, codec.ToWithUseProtoFieldName())
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		var op string
		if hm == "post" {
			op = http.MethodPost
		} else {
			op = http.MethodPut
		}
		newReq, err := http.NewRequest(op, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrapf(err, "Creating new HTTP %s request for custom API", op)
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("view_kind", fmt.Sprintf("%v", req.ViewKind))
		q.Add("view_name", fmt.Sprintf("%v", req.ViewName))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	// checking whether the status code is a successful status code (2xx series)
	if rsp.StatusCode < 200 || rsp.StatusCode > 299 {
		body, err := io.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &ForceDeleteResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.views.terraform_parameters.ForceDeleteResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomActionAPIRestClient) doRPCRun(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &RunRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.terraform_parameters.RunRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post", "put":
		jsn, err := codec.ToJSON(req, codec.ToWithUseProtoFieldName())
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		var op string
		if hm == "post" {
			op = http.MethodPost
		} else {
			op = http.MethodPut
		}
		newReq, err := http.NewRequest(op, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrapf(err, "Creating new HTTP %s request for custom API", op)
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("action", fmt.Sprintf("%v", req.Action))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("view_kind", fmt.Sprintf("%v", req.ViewKind))
		q.Add("view_name", fmt.Sprintf("%v", req.ViewName))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	// checking whether the status code is a successful status code (2xx series)
	if rsp.StatusCode < 200 || rsp.StatusCode > 299 {
		body, err := io.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &RunResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.views.terraform_parameters.RunResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomActionAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}

	rsp, err := rpcFn(ctx, cco)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using Rest")
	}
	return rsp, nil
}

func NewCustomActionAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomActionAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["ForceDelete"] = ccl.doRPCForceDelete

	rpcFns["Run"] = ccl.doRPCRun

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customActionAPIInprocClient

// INPROC Client (satisfying CustomActionAPIClient interface)
type customActionAPIInprocClient struct {
	CustomActionAPIServer
}

func (c *customActionAPIInprocClient) ForceDelete(ctx context.Context, in *ForceDeleteRequest, opts ...grpc.CallOption) (*ForceDeleteResponse, error) {
	return c.CustomActionAPIServer.ForceDelete(ctx, in)
}
func (c *customActionAPIInprocClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	return c.CustomActionAPIServer.Run(ctx, in)
}

func NewCustomActionAPIInprocClient(svc svcfw.Service) CustomActionAPIClient {
	return &customActionAPIInprocClient{CustomActionAPIServer: NewCustomActionAPIServer(svc)}
}

// RegisterGwCustomActionAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomActionAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomActionAPIHandlerClient(ctx, mux, NewCustomActionAPIInprocClient(s))
}

// Create customActionAPISrv

// SERVER (satisfying CustomActionAPIServer interface)
type customActionAPISrv struct {
	svc svcfw.Service
}

func (s *customActionAPISrv) ForceDelete(ctx context.Context, in *ForceDeleteRequest) (*ForceDeleteResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.views.terraform_parameters.CustomActionAPI")
	cah, ok := ah.(CustomActionAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomActionAPIServer", ah)
	}

	var (
		rsp *ForceDeleteResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.views.terraform_parameters.ForceDeleteRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomActionAPI.ForceDelete' operation on 'terraform_parameters'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.views.terraform_parameters.CustomActionAPI.ForceDelete"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.ForceDelete(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.views.terraform_parameters.ForceDeleteResponse", rsp)...)

	return rsp, nil
}
func (s *customActionAPISrv) Run(ctx context.Context, in *RunRequest) (*RunResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.views.terraform_parameters.CustomActionAPI")
	cah, ok := ah.(CustomActionAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomActionAPIServer", ah)
	}

	var (
		rsp *RunResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.views.terraform_parameters.RunRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomActionAPI.Run' operation on 'terraform_parameters'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.views.terraform_parameters.CustomActionAPI.Run"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.Run(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.views.terraform_parameters.RunResponse", rsp)...)

	return rsp, nil
}

func NewCustomActionAPIServer(svc svcfw.Service) CustomActionAPIServer {
	return &customActionAPISrv{svc: svc}
}

var CustomActionAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "View Terraform Parameters",
        "description": "View Terraform Parameters is set of parameters that are used by terraform scripts \nto instantiate view objects external to volterra",
        "version": "version not set"
    },
    "schemes": [
        "http",
        "https"
    ],
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "tags": [],
    "paths": {
        "/public/namespaces/{namespace}/terraform/{view_kind}/{view_name}/force-delete": {
            "post": {
                "summary": "Force delete view",
                "description": "force delete view object. This can result in staled objects in cloud provider.",
                "operationId": "ves.io.schema.views.terraform_parameters.CustomActionAPI.ForceDelete",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/terraform_parametersForceDeleteResponse"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "description": "Namespace\n\nx-example: \"value\"\nNamespace for the label to be retrieved",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "view_kind",
                        "description": "Kind of View\n\nx-example: \"value\"\nKind of view of which terraform parameters are requested e.g. aws_vpc_site, azure_vnet_site",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Kind of View"
                    },
                    {
                        "name": "view_name",
                        "description": "Name of view\n\nx-example: \"value\"\nName of the view for which terraform parameters are requested",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Name of view"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/terraform_parametersForceDeleteRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomActionAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-terraform_parameters-customactionapi-forcedelete"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.terraform_parameters.CustomActionAPI.ForceDelete"
            },
            "x-displayname": "View Terraform Parameters Action",
            "x-ves-proto-service": "ves.io.schema.views.terraform_parameters.CustomActionAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/{namespace}/terraform/{view_kind}/{view_name}/run": {
            "post": {
                "summary": "Run Terraform Action for view",
                "description": "perform terraform actions for a given view. Supported actions are apply and plan.",
                "operationId": "ves.io.schema.views.terraform_parameters.CustomActionAPI.Run",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/terraform_parametersRunResponse"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "description": "Namespace\n\nx-example: \"value\"\nNamespace for the label to be retrieved",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "view_kind",
                        "description": "Kind of View\n\nx-example: \"value\"\nKind of view of which terraform parameters are requested e.g. aws_vpc_site, azure_vnet_site",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Kind of View"
                    },
                    {
                        "name": "view_name",
                        "description": "Name of view\n\nx-example: \"value\"\nName of the view for which terraform parameters are requested",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Name of view"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/terraform_parametersRunRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomActionAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-terraform_parameters-customactionapi-run"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.terraform_parameters.CustomActionAPI.Run"
            },
            "x-displayname": "View Terraform Parameters Action",
            "x-ves-proto-service": "ves.io.schema.views.terraform_parameters.CustomActionAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "terraform_parametersForceDeleteRequest": {
            "type": "object",
            "description": "Force delete view request",
            "title": "ForceDeleteRequest",
            "x-displayname": "Force delete view request",
            "x-ves-proto-message": "ves.io.schema.views.terraform_parameters.ForceDeleteRequest",
            "properties": {
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the label to be retrieved\n\nExample: - \"value\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "value"
                },
                "view_kind": {
                    "type": "string",
                    "description": " Kind of view of which terraform parameters are requested e.g. aws_vpc_site, azure_vnet_site\n\nExample: - \"value\"-",
                    "title": "Kind of View",
                    "x-displayname": "Kind of View",
                    "x-ves-example": "value"
                },
                "view_name": {
                    "type": "string",
                    "description": " Name of the view for which terraform parameters are requested\n\nExample: - \"value\"-",
                    "title": "Name of view",
                    "x-displayname": "Name of view",
                    "x-ves-example": "value"
                }
            }
        },
        "terraform_parametersForceDeleteResponse": {
            "type": "object",
            "description": "Force delete view response",
            "title": "ForceDeleteResponse",
            "x-displayname": "Force delete view response",
            "x-ves-proto-message": "ves.io.schema.views.terraform_parameters.ForceDeleteResponse"
        },
        "terraform_parametersRunAction": {
            "type": "string",
            "description": "Terraform action to be performed for a given view e.g. plan, apply\n\nApply action used to apply the changes required to reach the desired state of the configuration\nPlan action is a convenient way to check whether the execution plan for a set of changes matches your expectations without making any changes to real resources or to the state\nDESTROY action is used to destroy the Terraform-managed infrastructure",
            "title": "Terraform Action",
            "enum": [
                "APPLY",
                "PLAN",
                "DESTROY"
            ],
            "default": "APPLY",
            "x-displayname": "Terraform action to be performed",
            "x-ves-proto-enum": "ves.io.schema.views.terraform_parameters.RunAction"
        },
        "terraform_parametersRunRequest": {
            "type": "object",
            "description": "perform terraform actions for a given view. Supported actions are apply and plan.",
            "title": "Run Terraform Action for view",
            "x-displayname": "Run Terraform Action for view",
            "x-ves-proto-message": "ves.io.schema.views.terraform_parameters.RunRequest",
            "properties": {
                "action": {
                    "description": " Terraform action to be performed for a given view e.g. plan, apply",
                    "title": "Terraform action to be performed",
                    "$ref": "#/definitions/terraform_parametersRunAction",
                    "x-displayname": "Terraform action to be performed"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the label to be retrieved\n\nExample: - \"value\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "value"
                },
                "view_kind": {
                    "type": "string",
                    "description": " Kind of view of which terraform parameters are requested e.g. aws_vpc_site, azure_vnet_site\n\nExample: - \"value\"-",
                    "title": "Kind of View",
                    "x-displayname": "Kind of View",
                    "x-ves-example": "value"
                },
                "view_name": {
                    "type": "string",
                    "description": " Name of the view for which terraform parameters are requested\n\nExample: - \"value\"-",
                    "title": "Name of view",
                    "x-displayname": "Name of view",
                    "x-ves-example": "value"
                }
            }
        },
        "terraform_parametersRunResponse": {
            "type": "object",
            "description": "Response for Run API",
            "title": "RunResponse",
            "x-displayname": "Run Response",
            "x-ves-proto-message": "ves.io.schema.views.terraform_parameters.RunResponse"
        }
    },
    "x-displayname": "View Terraform Parameters",
    "x-ves-proto-file": "ves.io/schema/views/terraform_parameters/public_custom_actionapi.proto"
}`
