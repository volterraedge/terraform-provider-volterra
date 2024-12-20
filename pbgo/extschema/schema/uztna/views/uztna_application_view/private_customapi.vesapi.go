// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package uztna_application_view

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

// Create CustomPrivateAPI GRPC Client satisfying server.CustomClient
type CustomPrivateAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomPrivateAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomPrivateAPIGrpcClient) doRPCDeleteExtendedTags(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &DeleteExtendedTagsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.uztna.views.uztna_application_view.DeleteExtendedTagsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.DeleteExtendedTags(ctx, req, opts...)
	return rsp, err
}

func (c *CustomPrivateAPIGrpcClient) doRPCSetExtendedTags(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SetExtendedTagsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.uztna.views.uztna_application_view.SetExtendedTagsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SetExtendedTags(ctx, req, opts...)
	return rsp, err
}

func (c *CustomPrivateAPIGrpcClient) doRPCUpdateExtendedTags(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &UpdateExtendedTagsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.uztna.views.uztna_application_view.UpdateExtendedTagsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.UpdateExtendedTags(ctx, req, opts...)
	return rsp, err
}

func (c *CustomPrivateAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomPrivateAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomPrivateAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomPrivateAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["DeleteExtendedTags"] = ccl.doRPCDeleteExtendedTags

	rpcFns["SetExtendedTags"] = ccl.doRPCSetExtendedTags

	rpcFns["UpdateExtendedTags"] = ccl.doRPCUpdateExtendedTags

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomPrivateAPI REST Client satisfying server.CustomClient
type CustomPrivateAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomPrivateAPIRestClient) doRPCDeleteExtendedTags(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &DeleteExtendedTagsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.uztna.views.uztna_application_view.DeleteExtendedTagsRequest: %s", yamlReq, err)
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
		for _, item := range req.ExtendedTags {
			q.Add("extended_tags", fmt.Sprintf("%v", item))
		}
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("objname", fmt.Sprintf("%v", req.Objname))

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
	pbRsp := &DeleteExtendedTagsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.uztna.views.uztna_application_view.DeleteExtendedTagsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomPrivateAPIRestClient) doRPCSetExtendedTags(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SetExtendedTagsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.uztna.views.uztna_application_view.SetExtendedTagsRequest: %s", yamlReq, err)
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
		for _, item := range req.ExtendedTags {
			q.Add("extended_tags", fmt.Sprintf("%v", item))
		}
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("objname", fmt.Sprintf("%v", req.Objname))

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
	pbRsp := &SetExtendedTagsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.uztna.views.uztna_application_view.SetExtendedTagsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomPrivateAPIRestClient) doRPCUpdateExtendedTags(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &UpdateExtendedTagsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.uztna.views.uztna_application_view.UpdateExtendedTagsRequest: %s", yamlReq, err)
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
		for _, item := range req.AddExtendedTags {
			q.Add("add_extended_tags", fmt.Sprintf("%v", item))
		}
		for _, item := range req.DeleteExtendedTags {
			q.Add("delete_extended_tags", fmt.Sprintf("%v", item))
		}
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("objname", fmt.Sprintf("%v", req.Objname))

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
	pbRsp := &UpdateExtendedTagsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.uztna.views.uztna_application_view.UpdateExtendedTagsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomPrivateAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomPrivateAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomPrivateAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["DeleteExtendedTags"] = ccl.doRPCDeleteExtendedTags

	rpcFns["SetExtendedTags"] = ccl.doRPCSetExtendedTags

	rpcFns["UpdateExtendedTags"] = ccl.doRPCUpdateExtendedTags

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customPrivateAPIInprocClient

// INPROC Client (satisfying CustomPrivateAPIClient interface)
type customPrivateAPIInprocClient struct {
	CustomPrivateAPIServer
}

func (c *customPrivateAPIInprocClient) DeleteExtendedTags(ctx context.Context, in *DeleteExtendedTagsRequest, opts ...grpc.CallOption) (*DeleteExtendedTagsResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.DeleteExtendedTags")
	return c.CustomPrivateAPIServer.DeleteExtendedTags(ctx, in)
}
func (c *customPrivateAPIInprocClient) SetExtendedTags(ctx context.Context, in *SetExtendedTagsRequest, opts ...grpc.CallOption) (*SetExtendedTagsResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.SetExtendedTags")
	return c.CustomPrivateAPIServer.SetExtendedTags(ctx, in)
}
func (c *customPrivateAPIInprocClient) UpdateExtendedTags(ctx context.Context, in *UpdateExtendedTagsRequest, opts ...grpc.CallOption) (*UpdateExtendedTagsResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.UpdateExtendedTags")
	return c.CustomPrivateAPIServer.UpdateExtendedTags(ctx, in)
}

func NewCustomPrivateAPIInprocClient(svc svcfw.Service) CustomPrivateAPIClient {
	return &customPrivateAPIInprocClient{CustomPrivateAPIServer: NewCustomPrivateAPIServer(svc)}
}

// RegisterGwCustomPrivateAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomPrivateAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomPrivateAPIHandlerClient(ctx, mux, NewCustomPrivateAPIInprocClient(s))
}

// Create customPrivateAPISrv

// SERVER (satisfying CustomPrivateAPIServer interface)
type customPrivateAPISrv struct {
	svc svcfw.Service
}

func (s *customPrivateAPISrv) DeleteExtendedTags(ctx context.Context, in *DeleteExtendedTagsRequest) (*DeleteExtendedTagsResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI")
	cah, ok := ah.(CustomPrivateAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomPrivateAPIServer", ah)
	}

	var (
		rsp *DeleteExtendedTagsResponse
		err error
	)

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.DeleteExtendedTags"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.DeleteExtendedTags(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	return rsp, nil
}
func (s *customPrivateAPISrv) SetExtendedTags(ctx context.Context, in *SetExtendedTagsRequest) (*SetExtendedTagsResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI")
	cah, ok := ah.(CustomPrivateAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomPrivateAPIServer", ah)
	}

	var (
		rsp *SetExtendedTagsResponse
		err error
	)

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.SetExtendedTags"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SetExtendedTags(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	return rsp, nil
}
func (s *customPrivateAPISrv) UpdateExtendedTags(ctx context.Context, in *UpdateExtendedTagsRequest) (*UpdateExtendedTagsResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI")
	cah, ok := ah.(CustomPrivateAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomPrivateAPIServer", ah)
	}

	var (
		rsp *UpdateExtendedTagsResponse
		err error
	)

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.UpdateExtendedTags"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.UpdateExtendedTags(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	return rsp, nil
}

func NewCustomPrivateAPIServer(svc svcfw.Service) CustomPrivateAPIServer {
	return &customPrivateAPISrv{svc: svc}
}

var CustomPrivateAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Onboarded Application Object",
        "description": " \nApplications are services that run within your organization, \nand end users are granted access to them after the ZTNA Access Policy \nis successfully evaluated. Access to these applications is subject to user authorization.\nThis section allows you to onboard your private applications to the ZTNA platform.",
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
        "/private/namespaces/{namespace}/uztna_application": {
            "delete": {
                "summary": "Delete Extended Tags",
                "description": "Delete Extended Tags from application group",
                "operationId": "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.DeleteExtendedTags",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/uztna_application_viewDeleteExtendedTagsResponse"
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
                        "description": "Namespace\n\nx-example: \"default\"\nNamespace for the object to be configured",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/uztna_application_viewDeleteExtendedTagsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomPrivateAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-uztna-views-uztna_application_view-customprivateapi-deleteextendedtags"
                },
                "x-ves-proto-rpc": "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.DeleteExtendedTags"
            },
            "post": {
                "summary": "Configure Extended Tags",
                "description": "Configure Extended Tags from application group",
                "operationId": "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.SetExtendedTags",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/uztna_application_viewSetExtendedTagsResponse"
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
                        "description": "Namespace\n\nx-example: \"default\"\nNamespace for the object to be configured",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/uztna_application_viewSetExtendedTagsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomPrivateAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-uztna-views-uztna_application_view-customprivateapi-setextendedtags"
                },
                "x-ves-proto-rpc": "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.SetExtendedTags"
            },
            "put": {
                "summary": "Update Extended Tags",
                "description": "Update Extended Tags from application group",
                "operationId": "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.UpdateExtendedTags",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/uztna_application_viewUpdateExtendedTagsResponse"
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
                        "description": "Namespace\n\nx-example: \"default\"\nNamespace for the object to be configured",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/uztna_application_viewUpdateExtendedTagsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomPrivateAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-uztna-views-uztna_application_view-customprivateapi-updateextendedtags"
                },
                "x-ves-proto-rpc": "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI.UpdateExtendedTags"
            },
            "x-displayname": "Custom API for UZTNA Application",
            "x-ves-proto-service": "ves.io.schema.uztna.views.uztna_application_view.CustomPrivateAPI",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        }
    },
    "definitions": {
        "uztna_application_viewDeleteExtendedTagsRequest": {
            "type": "object",
            "description": "Request to Delete Extended Tags from application group",
            "title": "Request to Delete Extended Tags",
            "x-displayname": "Request to Delete Extended Tags",
            "x-ves-proto-message": "ves.io.schema.uztna.views.uztna_application_view.DeleteExtendedTagsRequest",
            "properties": {
                "extended_tags": {
                    "type": "array",
                    "description": " Extended Tags of Application from application group\n\nExample: - \"[\"HR\",\"Eng\"]\"-",
                    "title": "Extended Tags of Application",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Extended Tags of Application",
                    "x-ves-example": "[\"HR\",\"Eng\"]"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                },
                "objname": {
                    "type": "string",
                    "description": " Object Name\n\nExample: - \"example\"-",
                    "title": "Object Name",
                    "x-displayname": "Object Name",
                    "x-ves-example": "example"
                }
            }
        },
        "uztna_application_viewDeleteExtendedTagsResponse": {
            "type": "object",
            "title": "Response to Delete Extended Tags",
            "x-displayname": "Response to Delete Extended Tags",
            "x-ves-proto-message": "ves.io.schema.uztna.views.uztna_application_view.DeleteExtendedTagsResponse"
        },
        "uztna_application_viewSetExtendedTagsRequest": {
            "type": "object",
            "description": "Request to Configure Extended Tags from application group",
            "title": "Request to Configure Extended Tags",
            "x-displayname": "Request to Configure Extended Tags",
            "x-ves-proto-message": "ves.io.schema.uztna.views.uztna_application_view.SetExtendedTagsRequest",
            "properties": {
                "extended_tags": {
                    "type": "array",
                    "description": " Extended Tags of Application from application group\n\nExample: - \"[\"HR\",\"Eng\"]\"-",
                    "title": "Extended Tags of Application",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Extended Tags of Application",
                    "x-ves-example": "[\"HR\",\"Eng\"]"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                },
                "objname": {
                    "type": "string",
                    "description": " Object Name\n\nExample: - \"example\"-",
                    "title": "Object Name",
                    "x-displayname": "Object Name",
                    "x-ves-example": "example"
                }
            }
        },
        "uztna_application_viewSetExtendedTagsResponse": {
            "type": "object",
            "title": "Response to configure Configure Extended Tags",
            "x-displayname": "Response to Configure Extended Tags ",
            "x-ves-proto-message": "ves.io.schema.uztna.views.uztna_application_view.SetExtendedTagsResponse"
        },
        "uztna_application_viewUpdateExtendedTagsRequest": {
            "type": "object",
            "description": "Request to Update Extended Tags from application group",
            "title": "Request to Update Extended Tags",
            "x-displayname": "Request to Update Extended Tags",
            "x-ves-proto-message": "ves.io.schema.uztna.views.uztna_application_view.UpdateExtendedTagsRequest",
            "properties": {
                "add_extended_tags": {
                    "type": "array",
                    "description": " Extended Tags of Application to be Added from application group\n\nExample: - \"[\"HR\",\"Eng\"]\"-",
                    "title": "Extended Tags of Application to be Added",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Extended Tags of Application to be Added",
                    "x-ves-example": "[\"HR\",\"Eng\"]"
                },
                "delete_extended_tags": {
                    "type": "array",
                    "description": " Extended Tags of Application to be Deleted from application group\n\nExample: - \"[\"HR\",\"Eng\"]\"-",
                    "title": "Extended Tags of Application to be Deleted",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Extended Tags of Application to be Deleted",
                    "x-ves-example": "[\"HR\",\"Eng\"]"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                },
                "objname": {
                    "type": "string",
                    "description": " Object Name\n\nExample: - \"example\"-",
                    "title": "Object Name",
                    "x-displayname": "Object Name",
                    "x-ves-example": "example"
                }
            }
        },
        "uztna_application_viewUpdateExtendedTagsResponse": {
            "type": "object",
            "title": "Response to Update Extended Tags",
            "x-displayname": "Response to Update Extended Tags",
            "x-ves-proto-message": "ves.io.schema.uztna.views.uztna_application_view.UpdateExtendedTagsResponse"
        }
    },
    "x-displayname": "Onboarded Application",
    "x-ves-proto-file": "ves.io/schema/uztna/views/uztna_application_view/private_customapi.proto"
}`