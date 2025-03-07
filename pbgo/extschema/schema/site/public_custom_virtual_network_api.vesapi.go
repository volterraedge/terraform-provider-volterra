// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package site

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

// Create CustomVirtualNetworkListAPI GRPC Client satisfying server.CustomClient
type CustomVirtualNetworkListAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomVirtualNetworkListAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomVirtualNetworkListAPIGrpcClient) doRPCGlobalNetworkList(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &GlobalNetworkListRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.GlobalNetworkListRequest", yamlReq)
	}
	rsp, err := c.grpcClient.GlobalNetworkList(ctx, req, opts...)
	return rsp, err
}

func (c *CustomVirtualNetworkListAPIGrpcClient) doRPCSegmentList(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SegmentListRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.SegmentListRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SegmentList(ctx, req, opts...)
	return rsp, err
}

func (c *CustomVirtualNetworkListAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomVirtualNetworkListAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomVirtualNetworkListAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomVirtualNetworkListAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["GlobalNetworkList"] = ccl.doRPCGlobalNetworkList

	rpcFns["SegmentList"] = ccl.doRPCSegmentList

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomVirtualNetworkListAPI REST Client satisfying server.CustomClient
type CustomVirtualNetworkListAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomVirtualNetworkListAPIRestClient) doRPCGlobalNetworkList(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &GlobalNetworkListRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.GlobalNetworkListRequest: %s", yamlReq, err)
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
		q.Add("site", fmt.Sprintf("%v", req.Site))

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
	pbRsp := &GlobalNetworkListResp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.site.GlobalNetworkListResp", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomVirtualNetworkListAPIRestClient) doRPCSegmentList(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SegmentListRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.SegmentListRequest: %s", yamlReq, err)
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
		q.Add("site", fmt.Sprintf("%v", req.Site))

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
	pbRsp := &SegmentListResp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.site.SegmentListResp", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomVirtualNetworkListAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomVirtualNetworkListAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomVirtualNetworkListAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["GlobalNetworkList"] = ccl.doRPCGlobalNetworkList

	rpcFns["SegmentList"] = ccl.doRPCSegmentList

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customVirtualNetworkListAPIInprocClient

// INPROC Client (satisfying CustomVirtualNetworkListAPIClient interface)
type customVirtualNetworkListAPIInprocClient struct {
	CustomVirtualNetworkListAPIServer
}

func (c *customVirtualNetworkListAPIInprocClient) GlobalNetworkList(ctx context.Context, in *GlobalNetworkListRequest, opts ...grpc.CallOption) (*GlobalNetworkListResp, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.site.CustomVirtualNetworkListAPI.GlobalNetworkList")
	return c.CustomVirtualNetworkListAPIServer.GlobalNetworkList(ctx, in)
}
func (c *customVirtualNetworkListAPIInprocClient) SegmentList(ctx context.Context, in *SegmentListRequest, opts ...grpc.CallOption) (*SegmentListResp, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.site.CustomVirtualNetworkListAPI.SegmentList")
	return c.CustomVirtualNetworkListAPIServer.SegmentList(ctx, in)
}

func NewCustomVirtualNetworkListAPIInprocClient(svc svcfw.Service) CustomVirtualNetworkListAPIClient {
	return &customVirtualNetworkListAPIInprocClient{CustomVirtualNetworkListAPIServer: NewCustomVirtualNetworkListAPIServer(svc)}
}

// RegisterGwCustomVirtualNetworkListAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomVirtualNetworkListAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomVirtualNetworkListAPIHandlerClient(ctx, mux, NewCustomVirtualNetworkListAPIInprocClient(s))
}

// Create customVirtualNetworkListAPISrv

// SERVER (satisfying CustomVirtualNetworkListAPIServer interface)
type customVirtualNetworkListAPISrv struct {
	svc svcfw.Service
}

func (s *customVirtualNetworkListAPISrv) GlobalNetworkList(ctx context.Context, in *GlobalNetworkListRequest) (*GlobalNetworkListResp, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.site.CustomVirtualNetworkListAPI")
	cah, ok := ah.(CustomVirtualNetworkListAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomVirtualNetworkListAPIServer", ah)
	}

	var (
		rsp *GlobalNetworkListResp
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.site.GlobalNetworkListRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomVirtualNetworkListAPI.GlobalNetworkList' operation on 'site'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.site.CustomVirtualNetworkListAPI.GlobalNetworkList"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.GlobalNetworkList(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.site.GlobalNetworkListResp", rsp)...)

	return rsp, nil
}
func (s *customVirtualNetworkListAPISrv) SegmentList(ctx context.Context, in *SegmentListRequest) (*SegmentListResp, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.site.CustomVirtualNetworkListAPI")
	cah, ok := ah.(CustomVirtualNetworkListAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomVirtualNetworkListAPIServer", ah)
	}

	var (
		rsp *SegmentListResp
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.site.SegmentListRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomVirtualNetworkListAPI.SegmentList' operation on 'site'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.site.CustomVirtualNetworkListAPI.SegmentList"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SegmentList(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.site.SegmentListResp", rsp)...)

	return rsp, nil
}

func NewCustomVirtualNetworkListAPIServer(svc svcfw.Service) CustomVirtualNetworkListAPIServer {
	return &customVirtualNetworkListAPISrv{svc: svc}
}

var CustomVirtualNetworkListAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "APIs to list segments and global networks for a site",
        "description": "APIs to get list of segments and global networks associated with a site.",
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
        "/public/namespaces/{namespace}/sites/{site}/global_networks": {
            "get": {
                "summary": "Global Network List",
                "description": "API to get list of Global Network in a site.",
                "operationId": "ves.io.schema.site.CustomVirtualNetworkListAPI.GlobalNetworkList",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/siteGlobalNetworkListResp"
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
                        "description": "Namespace\n\nx-example: \"system\"\nSite namespace",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "site",
                        "description": "Site\n\nx-required\nx-example: \"ce398\"\nSite name",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Site Name"
                    }
                ],
                "tags": [
                    "CustomVirtualNetworkListAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-site-customvirtualnetworklistapi-globalnetworklist"
                },
                "x-ves-proto-rpc": "ves.io.schema.site.CustomVirtualNetworkListAPI.GlobalNetworkList"
            },
            "x-displayname": "Custom Virtual Network List API",
            "x-ves-proto-service": "ves.io.schema.site.CustomVirtualNetworkListAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/{namespace}/sites/{site}/segments": {
            "get": {
                "summary": "Segment List",
                "description": "API to get list of segments in a site.",
                "operationId": "ves.io.schema.site.CustomVirtualNetworkListAPI.SegmentList",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/siteSegmentListResp"
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
                        "description": "Namespace\n\nx-example: \"system\"\nSite namespace",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "site",
                        "description": "Site\n\nx-required\nx-example: \"ce398\"\nSite name",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Site Name"
                    }
                ],
                "tags": [
                    "CustomVirtualNetworkListAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-site-customvirtualnetworklistapi-segmentlist"
                },
                "x-ves-proto-rpc": "ves.io.schema.site.CustomVirtualNetworkListAPI.SegmentList"
            },
            "x-displayname": "Custom Virtual Network List API",
            "x-ves-proto-service": "ves.io.schema.site.CustomVirtualNetworkListAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "ioschemaObjectRefType": {
            "type": "object",
            "description": "This type establishes a 'direct reference' from one object(the referrer) to another(the referred).\nSuch a reference is in form of tenant/namespace/name for public API and Uid for private API\nThis type of reference is called direct because the relation is explicit and concrete (as opposed\nto selector reference which builds a group based on labels of selectee objects)",
            "title": "ObjectRefType",
            "x-displayname": "Object reference",
            "x-ves-proto-message": "ves.io.schema.ObjectRefType",
            "properties": {
                "kind": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then kind will hold the referred object's kind (e.g. \"route\")\n\nExample: - \"virtual_site\"-",
                    "title": "kind",
                    "x-displayname": "Kind",
                    "x-ves-example": "virtual_site"
                },
                "name": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then name will hold the referred object's(e.g. route's) name.\n\nExample: - \"contactus-route\"-",
                    "title": "name",
                    "x-displayname": "Name",
                    "x-ves-example": "contactus-route"
                },
                "namespace": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then namespace will hold the referred object's(e.g. route's) namespace.\n\nExample: - \"ns1\"-",
                    "title": "namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1"
                },
                "tenant": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then tenant will hold the referred object's(e.g. route's) tenant.\n\nExample: - \"acmecorp\"-",
                    "title": "tenant",
                    "x-displayname": "Tenant",
                    "x-ves-example": "acmecorp"
                },
                "uid": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then uid will hold the referred object's(e.g. route's) uid.\n\nExample: - \"d15f1fad-4d37-48c0-8706-df1824d76d31\"-",
                    "title": "uid",
                    "x-displayname": "UID",
                    "x-ves-example": "d15f1fad-4d37-48c0-8706-df1824d76d31"
                }
            }
        },
        "siteGlobalNetworkItem": {
            "type": "object",
            "description": "Contains reference to the virtual network, and its optional description",
            "title": "GlobalNetworkItem",
            "x-displayname": "GlobalNetworkItem",
            "x-ves-proto-message": "ves.io.schema.site.GlobalNetworkItem",
            "properties": {
                "description": {
                    "type": "string",
                    "description": " Optional description for the virtual_network",
                    "title": "description",
                    "x-displayname": "Description"
                },
                "ref": {
                    "type": "array",
                    "description": " Reference to the Global Network type virtual_network Object",
                    "title": "Reference",
                    "items": {
                        "$ref": "#/definitions/ioschemaObjectRefType"
                    },
                    "x-displayname": "Reference"
                }
            }
        },
        "siteGlobalNetworkListResp": {
            "type": "object",
            "description": "Global Network List Response",
            "title": "GlobalNetworkListResp",
            "x-displayname": "Global Network List Response",
            "x-ves-proto-message": "ves.io.schema.site.GlobalNetworkListResp",
            "properties": {
                "items": {
                    "type": "array",
                    "description": " List of the global networks connected to site",
                    "title": "Item",
                    "items": {
                        "$ref": "#/definitions/siteGlobalNetworkItem"
                    },
                    "x-displayname": "Items"
                }
            }
        },
        "siteSegmentItem": {
            "type": "object",
            "description": "Contains reference to the segment, and its optional description",
            "title": "SegmentItem",
            "x-displayname": "SegmentItem",
            "x-ves-proto-message": "ves.io.schema.site.SegmentItem",
            "properties": {
                "description": {
                    "type": "string",
                    "description": " Optional description for the segment",
                    "title": "description",
                    "x-displayname": "Description"
                },
                "ref": {
                    "type": "array",
                    "description": " Reference to the segment object",
                    "title": "Reference",
                    "items": {
                        "$ref": "#/definitions/ioschemaObjectRefType"
                    },
                    "x-displayname": "Reference"
                }
            }
        },
        "siteSegmentListResp": {
            "type": "object",
            "description": "Segment List Response",
            "title": "SegmentListResp",
            "x-displayname": "Segment List Response",
            "x-ves-proto-message": "ves.io.schema.site.SegmentListResp",
            "properties": {
                "items": {
                    "type": "array",
                    "description": " List of the segments connected to site",
                    "title": "Item",
                    "items": {
                        "$ref": "#/definitions/siteSegmentItem"
                    },
                    "x-displayname": "Items"
                }
            }
        }
    },
    "x-displayname": "Site",
    "x-ves-proto-file": "ves.io/schema/site/public_custom_virtual_network_api.proto"
}`
