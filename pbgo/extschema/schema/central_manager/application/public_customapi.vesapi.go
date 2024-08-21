// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package application

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

// Create CustomAPI GRPC Client satisfying server.CustomClient
type CustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomAPIGrpcClient) doRPCGet(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &GetRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.central_manager.application.GetRequest", yamlReq)
	}
	rsp, err := c.grpcClient.Get(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) doRPCList(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &ListRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.central_manager.application.ListRequest", yamlReq)
	}
	rsp, err := c.grpcClient.List(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["Get"] = ccl.doRPCGet

	rpcFns["List"] = ccl.doRPCList

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPI REST Client satisfying server.CustomClient
type CustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomAPIRestClient) doRPCGet(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &GetRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.central_manager.application.GetRequest: %s", yamlReq, err)
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
		q.Add("application_id", fmt.Sprintf("%v", req.ApplicationId))
		q.Add("cm_id", fmt.Sprintf("%v", req.CmId))

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
	pbRsp := &GetResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.central_manager.application.GetResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) doRPCList(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &ListRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.central_manager.application.ListRequest: %s", yamlReq, err)
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
		q.Add("cm_id", fmt.Sprintf("%v", req.CmId))
		q.Add("limit", fmt.Sprintf("%v", req.Limit))
		q.Add("page", fmt.Sprintf("%v", req.Page))

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
	pbRsp := &ListResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.central_manager.application.ListResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["Get"] = ccl.doRPCGet

	rpcFns["List"] = ccl.doRPCList

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customAPIInprocClient

// INPROC Client (satisfying CustomAPIClient interface)
type customAPIInprocClient struct {
	CustomAPIServer
}

func (c *customAPIInprocClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.central_manager.application.CustomAPI.Get")
	return c.CustomAPIServer.Get(ctx, in)
}
func (c *customAPIInprocClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.central_manager.application.CustomAPI.List")
	return c.CustomAPIServer.List(ctx, in)
}

func NewCustomAPIInprocClient(svc svcfw.Service) CustomAPIClient {
	return &customAPIInprocClient{CustomAPIServer: NewCustomAPIServer(svc)}
}

// RegisterGwCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomAPIHandlerClient(ctx, mux, NewCustomAPIInprocClient(s))
}

// Create customAPISrv

// SERVER (satisfying CustomAPIServer interface)
type customAPISrv struct {
	svc svcfw.Service
}

func (s *customAPISrv) Get(ctx context.Context, in *GetRequest) (*GetResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.central_manager.application.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPIServer", ah)
	}

	var (
		rsp *GetResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.central_manager.application.GetRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.Get' operation on 'application'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.central_manager.application.CustomAPI.Get"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.Get(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.central_manager.application.GetResponse", rsp)...)

	return rsp, nil
}
func (s *customAPISrv) List(ctx context.Context, in *ListRequest) (*ListResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.central_manager.application.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPIServer", ah)
	}

	var (
		rsp *ListResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.central_manager.application.ListRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.List' operation on 'application'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.central_manager.application.CustomAPI.List"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.List(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.central_manager.application.ListResponse", rsp)...)

	return rsp, nil
}

func NewCustomAPIServer(svc svcfw.Service) CustomAPIServer {
	return &customAPISrv{svc: svc}
}

var CustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Central Manager Applications",
        "description": "Central Manager Applications APIs enable the fetch of metrics data \nof different BIG-IP Next Central Managers owned by a Tenant",
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
        "/public/central_manager/applications/{cm_id}": {
            "get": {
                "summary": "List Central Manager Applications",
                "description": "List the applications managed by this central manager",
                "operationId": "ves.io.schema.central_manager.application.CustomAPI.List",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/applicationListResponse"
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
                        "name": "cm_id",
                        "description": "Central Manager ID\n\nx-required\nx-example: \"cm01-eastern-us\"\nID of Central Manager",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Central Manager ID"
                    },
                    {
                        "name": "limit",
                        "description": "x-example: \"20\"\nLimits the number of items to return per page",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "x-displayname": "Limit"
                    },
                    {
                        "name": "page",
                        "description": "x-example: \"1\"\nThe number of pages to return",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "x-displayname": "Page"
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-central_manager-application-customapi-list"
                },
                "x-ves-proto-rpc": "ves.io.schema.central_manager.application.CustomAPI.List"
            },
            "x-displayname": "Central Manager Application Custom API",
            "x-ves-proto-service": "ves.io.schema.central_manager.application.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/central_manager/applications/{cm_id}/{application_id}": {
            "get": {
                "summary": "Get Central Manager Application",
                "description": "Get details of a particular application managed by the central manager",
                "operationId": "ves.io.schema.central_manager.application.CustomAPI.Get",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/applicationGetResponse"
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
                        "name": "cm_id",
                        "description": "Central Manager ID\n\nx-required\nx-example: \"cm01-eastern-us\"\nID of Central Manager",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Central Manager ID"
                    },
                    {
                        "name": "application_id",
                        "description": "Central Manager Application ID\n\nx-required\nx-example: \"a0b9fb1b-fcb0-4da5-9046-775cbda7bb2d\"\nID of Central Manager Cloud Application",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Central Manager Application ID"
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-central_manager-application-customapi-get"
                },
                "x-ves-proto-rpc": "ves.io.schema.central_manager.application.CustomAPI.Get"
            },
            "x-displayname": "Central Manager Application Custom API",
            "x-ves-proto-service": "ves.io.schema.central_manager.application.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "applicationAppType": {
            "type": "string",
            "description": "Defines central manager application type\n",
            "title": "Central Manager Application type",
            "enum": [
                "FAST",
                "AS3"
            ],
            "default": "FAST",
            "x-displayname": "Central Manager Application type",
            "x-ves-proto-enum": "ves.io.schema.central_manager.application.AppType"
        },
        "applicationApplication": {
            "type": "object",
            "description": "Application details for a given Central manager",
            "title": "Application",
            "x-displayname": "Application",
            "x-ves-proto-message": "ves.io.schema.central_manager.application.Application",
            "properties": {
                "id": {
                    "type": "string",
                    "description": " Central manager application Id\n\nExample: - \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"-",
                    "title": "id",
                    "x-displayname": "Application ID",
                    "x-ves-example": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
                },
                "name": {
                    "type": "string",
                    "description": " Central manager application name\n\nExample: - \"MyApp1\"-",
                    "title": "name",
                    "x-displayname": "Application Name",
                    "x-ves-example": "MyApp1"
                },
                "type": {
                    "description": " Central manager application type\n\nExample: - \"fast\"-",
                    "title": "type",
                    "$ref": "#/definitions/applicationAppType",
                    "x-displayname": "Application Type",
                    "x-ves-example": "fast"
                }
            }
        },
        "applicationGetResponse": {
            "type": "object",
            "description": "This is the output message of the 'Get' RPC",
            "title": "Get Response",
            "x-displayname": "Get Response",
            "x-ves-proto-message": "ves.io.schema.central_manager.application.GetResponse",
            "properties": {
                "id": {
                    "type": "string",
                    "description": " Central manager application Id\n\nExample: - \"3fa85f64-5717-4562-b3fc-2c963f66afa6\"-",
                    "title": "id",
                    "x-displayname": "Application ID",
                    "x-ves-example": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
                },
                "locations": {
                    "type": "array",
                    "description": " Location of virtual server",
                    "title": "locations",
                    "items": {
                        "$ref": "#/definitions/central_managerapplicationLocation"
                    },
                    "x-displayname": "Location"
                },
                "name": {
                    "type": "string",
                    "description": " Central manager application name\n\nExample: - \"MyApp1\"-",
                    "title": "name",
                    "x-displayname": "Application Name",
                    "x-ves-example": "MyApp1"
                },
                "type": {
                    "description": " Central manager application type",
                    "title": "type",
                    "$ref": "#/definitions/applicationAppType",
                    "x-displayname": "Application Type"
                }
            }
        },
        "applicationLinks": {
            "type": "object",
            "description": "Link is href to next/previous/last/first pages",
            "title": "Links",
            "x-displayname": "Links",
            "x-ves-proto-message": "ves.io.schema.central_manager.application.Links",
            "properties": {
                "first": {
                    "type": "string",
                    "description": " Link that navigates to the first page\n\nExample: - \"?limit=100\u0026page=1\"-",
                    "title": "first",
                    "x-displayname": "First",
                    "x-ves-example": "?limit=100\u0026page=1"
                },
                "last": {
                    "type": "string",
                    "description": " Link that navigates to the last page\n\nExample: - \"?limit=100\u0026page=10\"-",
                    "title": "last",
                    "x-displayname": "Last",
                    "x-ves-example": "?limit=100\u0026page=10"
                },
                "next": {
                    "type": "string",
                    "description": " Link that navigates to the next page\n\nExample: - \"?limit=100\u0026page=2\"-",
                    "title": "next",
                    "x-displayname": "Next",
                    "x-ves-example": "?limit=100\u0026page=2"
                },
                "prev": {
                    "type": "string",
                    "description": " Link that navigates to the previous page\n\nExample: - \"?limit=100\u0026page=1\"-",
                    "title": "prev",
                    "x-displayname": "Prev",
                    "x-ves-example": "?limit=100\u0026page=1"
                },
                "self": {
                    "type": "string",
                    "description": " Link that navigates to the self page\n\nExample: - \"?limit=100\u0026page=1\"-",
                    "title": "self",
                    "x-displayname": "Self",
                    "x-ves-example": "?limit=100\u0026page=1"
                }
            }
        },
        "applicationListResponse": {
            "type": "object",
            "description": "This is the output message of the 'List' RPC",
            "title": "List Response",
            "x-displayname": "List Response",
            "x-ves-proto-message": "ves.io.schema.central_manager.application.ListResponse",
            "properties": {
                "applications": {
                    "type": "array",
                    "description": " Application details for a given Central manager",
                    "title": "Application",
                    "items": {
                        "$ref": "#/definitions/applicationApplication"
                    },
                    "x-displayname": "Application Details"
                },
                "count": {
                    "type": "integer",
                    "description": " Count is number of resources returned in the response",
                    "title": "count",
                    "format": "int32",
                    "x-displayname": "Count"
                },
                "link": {
                    "description": " Link is href to next/previous/last/first pages",
                    "title": "link",
                    "$ref": "#/definitions/applicationLinks",
                    "x-displayname": "Link"
                },
                "total": {
                    "type": "integer",
                    "description": " Total number of applications managed by this central manager",
                    "title": "total",
                    "format": "int32",
                    "x-displayname": "Total"
                }
            }
        },
        "applicationPoolMember": {
            "type": "object",
            "description": "Details of Pool members",
            "title": "pool_members",
            "x-displayname": "Pool Members",
            "x-ves-proto-message": "ves.io.schema.central_manager.application.PoolMember",
            "properties": {
                "ip_address": {
                    "type": "string",
                    "description": "\n\nExample: - \"192.168.1.2\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.ip: true\n",
                    "title": "ip_address",
                    "x-displayname": "Pool Member IP",
                    "x-ves-example": "192.168.1.2",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.ip": "true"
                    }
                },
                "name": {
                    "type": "string",
                    "description": "\n\nExample: - \"PoolMember1\"-",
                    "title": "name",
                    "x-displayname": "Pool Member Name",
                    "x-ves-example": "PoolMember1"
                }
            }
        },
        "applicationVirtualServer": {
            "type": "object",
            "description": "Details of virtual server where application is hosted.",
            "title": "virtual_servers",
            "x-displayname": "Virtual Servers",
            "x-ves-proto-message": "ves.io.schema.central_manager.application.VirtualServer",
            "properties": {
                "ip_address": {
                    "type": "string",
                    "description": "\n\nExample: - 10.20.30.40-\n\nValidation Rules:\n  ves.io.schema.rules.string.ip: true\n",
                    "title": "ip_address",
                    "x-displayname": "Virtual IP",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.ip": "true"
                    }
                },
                "pool_members": {
                    "type": "array",
                    "description": " Details of pool member",
                    "title": "pool_members",
                    "items": {
                        "$ref": "#/definitions/applicationPoolMember"
                    },
                    "x-displayname": "Pool Member"
                },
                "vs_name": {
                    "type": "string",
                    "description": "\n\nExample: - BIGIP-",
                    "title": "vs_name",
                    "x-displayname": "Virtual Server Name"
                }
            }
        },
        "central_managerapplicationLocation": {
            "type": "object",
            "description": "Location of virtual server",
            "title": "Location",
            "x-displayname": "Location",
            "x-ves-proto-message": "ves.io.schema.central_manager.application.Location",
            "properties": {
                "hostname": {
                    "type": "string",
                    "description": " Central manager application domain\n\nExample: - \"bigip-next.asia.com\"-",
                    "title": "hostname",
                    "x-displayname": "HostName",
                    "x-ves-example": "bigip-next.asia.com"
                },
                "ip_address": {
                    "type": "string",
                    "description": " Central manager application IP address\n\nExample: - \"10.240.107.46\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.ip: true\n",
                    "title": "ip_address",
                    "x-displayname": "Application IP",
                    "x-ves-example": "10.240.107.46",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.ip": "true"
                    }
                },
                "virtual_servers": {
                    "type": "array",
                    "description": " Details of virtual server where application is hosted.",
                    "title": "virtual_servers",
                    "items": {
                        "$ref": "#/definitions/applicationVirtualServer"
                    },
                    "x-displayname": "Virtual Server"
                }
            }
        }
    },
    "x-displayname": "Central Manager Application",
    "x-ves-proto-file": "ves.io/schema/central_manager/application/public_customapi.proto"
}`