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
	google_api "google.golang.org/genproto/googleapis/api/httpbody"
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

// Create ConfigKubeConfigAPI GRPC Client satisfying server.CustomClient
type ConfigKubeConfigAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient ConfigKubeConfigAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *ConfigKubeConfigAPIGrpcClient) doRPCCreateLocalKubeConfig(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &CreateKubeConfigReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.CreateKubeConfigReq", yamlReq)
	}
	rsp, err := c.grpcClient.CreateLocalKubeConfig(ctx, req, opts...)
	return rsp, err
}

func (c *ConfigKubeConfigAPIGrpcClient) doRPCListLocalKubeConfig(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &ListKubeConfigReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.ListKubeConfigReq", yamlReq)
	}
	rsp, err := c.grpcClient.ListLocalKubeConfig(ctx, req, opts...)
	return rsp, err
}

func (c *ConfigKubeConfigAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewConfigKubeConfigAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &ConfigKubeConfigAPIGrpcClient{
		conn:       cc,
		grpcClient: NewConfigKubeConfigAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["CreateLocalKubeConfig"] = ccl.doRPCCreateLocalKubeConfig

	rpcFns["ListLocalKubeConfig"] = ccl.doRPCListLocalKubeConfig

	ccl.rpcFns = rpcFns

	return ccl
}

// Create ConfigKubeConfigAPI REST Client satisfying server.CustomClient
type ConfigKubeConfigAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *ConfigKubeConfigAPIRestClient) doRPCCreateLocalKubeConfig(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &CreateKubeConfigReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.CreateKubeConfigReq: %s", yamlReq, err)
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
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))

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
	pbRsp := &google_api.HttpBody{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		// server strips HTTP Body proto message and sends only data, re-build it here
		pbRsp.ContentType = rsp.Header.Get("Content-Type")
		pbRsp.Data = body

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *ConfigKubeConfigAPIRestClient) doRPCListLocalKubeConfig(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &ListKubeConfigReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.ListKubeConfigReq: %s", yamlReq, err)
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
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))

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
	pbRsp := &ListKubeConfigRsp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.site.ListKubeConfigRsp", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *ConfigKubeConfigAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewConfigKubeConfigAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &ConfigKubeConfigAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["CreateLocalKubeConfig"] = ccl.doRPCCreateLocalKubeConfig

	rpcFns["ListLocalKubeConfig"] = ccl.doRPCListLocalKubeConfig

	ccl.rpcFns = rpcFns

	return ccl
}

// Create configKubeConfigAPIInprocClient

// INPROC Client (satisfying ConfigKubeConfigAPIClient interface)
type configKubeConfigAPIInprocClient struct {
	ConfigKubeConfigAPIServer
}

func (c *configKubeConfigAPIInprocClient) CreateLocalKubeConfig(ctx context.Context, in *CreateKubeConfigReq, opts ...grpc.CallOption) (*google_api.HttpBody, error) {
	return c.ConfigKubeConfigAPIServer.CreateLocalKubeConfig(ctx, in)
}
func (c *configKubeConfigAPIInprocClient) ListLocalKubeConfig(ctx context.Context, in *ListKubeConfigReq, opts ...grpc.CallOption) (*ListKubeConfigRsp, error) {
	return c.ConfigKubeConfigAPIServer.ListLocalKubeConfig(ctx, in)
}

func NewConfigKubeConfigAPIInprocClient(svc svcfw.Service) ConfigKubeConfigAPIClient {
	return &configKubeConfigAPIInprocClient{ConfigKubeConfigAPIServer: NewConfigKubeConfigAPIServer(svc)}
}

// RegisterGwConfigKubeConfigAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwConfigKubeConfigAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterConfigKubeConfigAPIHandlerClient(ctx, mux, NewConfigKubeConfigAPIInprocClient(s))
}

// Create configKubeConfigAPISrv

// SERVER (satisfying ConfigKubeConfigAPIServer interface)
type configKubeConfigAPISrv struct {
	svc svcfw.Service
}

func (s *configKubeConfigAPISrv) CreateLocalKubeConfig(ctx context.Context, in *CreateKubeConfigReq) (*google_api.HttpBody, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.site.ConfigKubeConfigAPI")
	cah, ok := ah.(ConfigKubeConfigAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *ConfigKubeConfigAPIServer", ah)
	}

	var (
		rsp *google_api.HttpBody
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.site.CreateKubeConfigReq", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'ConfigKubeConfigAPI.CreateLocalKubeConfig' operation on 'site'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.site.ConfigKubeConfigAPI.CreateLocalKubeConfig"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.CreateLocalKubeConfig(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "google.api.HttpBody", rsp)...)

	return rsp, nil
}
func (s *configKubeConfigAPISrv) ListLocalKubeConfig(ctx context.Context, in *ListKubeConfigReq) (*ListKubeConfigRsp, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.site.ConfigKubeConfigAPI")
	cah, ok := ah.(ConfigKubeConfigAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *ConfigKubeConfigAPIServer", ah)
	}

	var (
		rsp *ListKubeConfigRsp
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.site.ListKubeConfigReq", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'ConfigKubeConfigAPI.ListLocalKubeConfig' operation on 'site'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.site.ConfigKubeConfigAPI.ListLocalKubeConfig"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.ListLocalKubeConfig(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.site.ListKubeConfigRsp", rsp)...)

	return rsp, nil
}

func NewConfigKubeConfigAPIServer(svc svcfw.Service) ConfigKubeConfigAPIServer {
	return &configKubeConfigAPISrv{svc: svc}
}

var ConfigKubeConfigAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Site KubeConfig custom API",
        "description": "API for manage kube configs.",
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
        "/public/namespaces/{namespace}/sites/{name}/local-kubeconfig": {
            "post": {
                "summary": "Create K8s Cluster Local Kube Config",
                "description": "Down load kube config for local k8s cluster access",
                "operationId": "ves.io.schema.site.ConfigKubeConfigAPI.CreateLocalKubeConfig",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/apiHttpBody"
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
                        "description": "Namespace\n\nx-required\nx-example: \"system\"\nK8s Cluster namespace",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-required\nx-example: \"ce398\"\nK8s Cluster name",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Name"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/siteCreateKubeConfigReq"
                        }
                    }
                ],
                "tags": [
                    "ConfigKubeConfigAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-site-configkubeconfigapi-createlocalkubeconfig"
                },
                "x-ves-proto-rpc": "ves.io.schema.site.ConfigKubeConfigAPI.CreateLocalKubeConfig"
            },
            "x-displayname": "Local KubeConfig API",
            "x-ves-proto-service": "ves.io.schema.site.ConfigKubeConfigAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/{namespace}/sites/{name}/local-kubeconfigs": {
            "get": {
                "summary": "List Local Kube Configs",
                "description": "Returns list of all local active kubeconfig minted for this site",
                "operationId": "ves.io.schema.site.ConfigKubeConfigAPI.ListLocalKubeConfig",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/siteListKubeConfigRsp"
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
                        "description": "Namespace\n\nx-required\nx-example: \"system\"\nK8s Cluster namespace",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-required\nx-example: \"ce398\"\nK8s Cluster name",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Name"
                    }
                ],
                "tags": [
                    "ConfigKubeConfigAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-site-configkubeconfigapi-listlocalkubeconfig"
                },
                "x-ves-proto-rpc": "ves.io.schema.site.ConfigKubeConfigAPI.ListLocalKubeConfig"
            },
            "x-displayname": "Local KubeConfig API",
            "x-ves-proto-service": "ves.io.schema.site.ConfigKubeConfigAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "apiHttpBody": {
            "type": "object",
            "description": "Message that represents an arbitrary HTTP body. It should only be used for\npayload formats that can't be represented as JSON, such as raw binary or\nan HTML page.\n\n\nThis message can be used both in streaming and non-streaming API methods in\nthe request as well as the response.\n\nIt can be used as a top-level request field, which is convenient if one\nwants to extract parameters from either the URL or HTTP template into the\nrequest fields and also want access to the raw HTTP body.\n\nExample:\n\n    message GetResourceRequest {\n      // A unique request id.\n      string request_id = 1;\n\n      // The raw HTTP body is bound to this field.\n      google.api.HttpBody http_body = 2;\n    }\n\n    service ResourceService {\n      rpc GetResource(GetResourceRequest) returns (google.api.HttpBody);\n      rpc UpdateResource(google.api.HttpBody) returns\n      (google.protobuf.Empty);\n    }\n\nExample with streaming methods:\n\n    service CaldavService {\n      rpc GetCalendar(stream google.api.HttpBody)\n        returns (stream google.api.HttpBody);\n      rpc UpdateCalendar(stream google.api.HttpBody)\n        returns (stream google.api.HttpBody);\n    }\n\nUse of this type only changes how the request and response bodies are\nhandled, all other features will continue to work unchanged.",
            "properties": {
                "content_type": {
                    "type": "string",
                    "description": "The HTTP Content-Type header value specifying the content type of the body."
                },
                "data": {
                    "type": "string",
                    "description": "The HTTP request/response body as raw binary.",
                    "format": "byte"
                },
                "extensions": {
                    "type": "array",
                    "description": "Application specific response metadata. Must be set in the first response\nfor streaming APIs.",
                    "items": {
                        "$ref": "#/definitions/protobufAny"
                    }
                }
            }
        },
        "protobufAny": {
            "type": "object",
            "description": "-Any- contains an arbitrary serialized protocol buffer message along with a\nURL that describes the type of the serialized message.\n\nProtobuf library provides support to pack/unpack Any values in the form\nof utility functions or additional generated methods of the Any type.\n\nExample 1: Pack and unpack a message in C++.\n\n    Foo foo = ...;\n    Any any;\n    any.PackFrom(foo);\n    ...\n    if (any.UnpackTo(\u0026foo)) {\n      ...\n    }\n\nExample 2: Pack and unpack a message in Java.\n\n    Foo foo = ...;\n    Any any = Any.pack(foo);\n    ...\n    if (any.is(Foo.class)) {\n      foo = any.unpack(Foo.class);\n    }\n\n Example 3: Pack and unpack a message in Python.\n\n    foo = Foo(...)\n    any = Any()\n    any.Pack(foo)\n    ...\n    if any.Is(Foo.DESCRIPTOR):\n      any.Unpack(foo)\n      ...\n\n Example 4: Pack and unpack a message in Go\n\n     foo := \u0026pb.Foo{...}\n     any, err := ptypes.MarshalAny(foo)\n     ...\n     foo := \u0026pb.Foo{}\n     if err := ptypes.UnmarshalAny(any, foo); err != nil {\n       ...\n     }\n\nThe pack methods provided by protobuf library will by default use\n'type.googleapis.com/full.type.name' as the type URL and the unpack\nmethods only use the fully qualified type name after the last '/'\nin the type URL, for example \"foo.bar.com/x/y.z\" will yield type\nname \"y.z\".\n\n\nJSON\n====\nThe JSON representation of an -Any- value uses the regular\nrepresentation of the deserialized, embedded message, with an\nadditional field -@type- which contains the type URL. Example:\n\n    package google.profile;\n    message Person {\n      string first_name = 1;\n      string last_name = 2;\n    }\n\n    {\n      \"@type\": \"type.googleapis.com/google.profile.Person\",\n      \"firstName\": \u003cstring\u003e,\n      \"lastName\": \u003cstring\u003e\n    }\n\nIf the embedded message type is well-known and has a custom JSON\nrepresentation, that representation will be embedded adding a field\n-value- which holds the custom JSON in addition to the -@type-\nfield. Example (for message [google.protobuf.Duration][]):\n\n    {\n      \"@type\": \"type.googleapis.com/google.protobuf.Duration\",\n      \"value\": \"1.212s\"\n    }",
            "properties": {
                "type_url": {
                    "type": "string",
                    "description": "A URL/resource name that uniquely identifies the type of the serialized\nprotocol buffer message. This string must contain at least\none \"/\" character. The last segment of the URL's path must represent\nthe fully qualified name of the type (as in\n-path/google.protobuf.Duration-). The name should be in a canonical form\n(e.g., leading \".\" is not accepted).\n\nIn practice, teams usually precompile into the binary all types that they\nexpect it to use in the context of Any. However, for URLs which use the\nscheme -http-, -https-, or no scheme, one can optionally set up a type\nserver that maps type URLs to message definitions as follows:\n\n* If no scheme is provided, -https- is assumed.\n* An HTTP GET on the URL must yield a [google.protobuf.Type][]\n  value in binary format, or produce an error.\n* Applications are allowed to cache lookup results based on the\n  URL, or have them precompiled into a binary to avoid any\n  lookup. Therefore, binary compatibility needs to be preserved\n  on changes to types. (Use versioned type names to manage\n  breaking changes.)\n\nNote: this functionality is not currently available in the official\nprotobuf release, and it is not used for type URLs beginning with\ntype.googleapis.com.\n\nSchemes other than -http-, -https- (or the empty scheme) might be\nused with implementation specific semantics."
                },
                "value": {
                    "type": "string",
                    "description": "Must be a valid serialized protocol buffer of the above specified type.",
                    "format": "byte"
                }
            }
        },
        "siteCreateKubeConfigReq": {
            "type": "object",
            "description": "Create kubeconfig request parameters",
            "title": "Create Kube Config Request",
            "x-displayname": "Create Kube Config Request",
            "x-ves-proto-message": "ves.io.schema.site.CreateKubeConfigReq",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " K8s Cluster name\n\nExample: - \"ce398\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_bytes: 64\n  ves.io.schema.rules.string.min_bytes: 1\n",
                    "title": "Name",
                    "minLength": 1,
                    "maxLength": 64,
                    "x-displayname": "Name",
                    "x-ves-example": "ce398",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_bytes": "64",
                        "ves.io.schema.rules.string.min_bytes": "1"
                    }
                },
                "namespace": {
                    "type": "string",
                    "description": " K8s Cluster namespace\n\nExample: - \"system\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_bytes: 64\n  ves.io.schema.rules.string.min_bytes: 1\n",
                    "title": "Namespace",
                    "minLength": 1,
                    "maxLength": 64,
                    "x-displayname": "Namespace",
                    "x-ves-example": "system",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_bytes": "64",
                        "ves.io.schema.rules.string.min_bytes": "1"
                    }
                }
            }
        },
        "siteListKubeConfigRsp": {
            "type": "object",
            "description": "List kubeconfig response",
            "title": "List Kube Config Response",
            "x-displayname": "List Kube Config Response",
            "x-ves-proto-message": "ves.io.schema.site.ListKubeConfigRsp",
            "properties": {
                "items": {
                    "type": "array",
                    "description": " List of kubeconfig items.",
                    "title": "List of the kubeconfigs",
                    "items": {
                        "$ref": "#/definitions/siteListKubeConfigRspItem"
                    },
                    "x-displayname": "List of Kube Configs"
                }
            }
        },
        "siteListKubeConfigRspItem": {
            "type": "object",
            "description": "Each item of credential list request.",
            "title": "List Kube Config item",
            "x-displayname": "List Kube Config item",
            "x-ves-proto-message": "ves.io.schema.site.ListKubeConfigRspItem",
            "properties": {
                "create_timestamp": {
                    "type": "string",
                    "description": " Create time of API credential.",
                    "title": "Create timestamp",
                    "format": "date-time",
                    "x-displayname": "Creation Time"
                },
                "expiry_timestamp": {
                    "type": "string",
                    "description": " Expiry time of credential.",
                    "title": "Expiry time",
                    "format": "date-time",
                    "x-displayname": "Expiry Time"
                },
                "name": {
                    "type": "string",
                    "description": " Name of this credential\n\nExample: - \"api-cred-x89sf\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_bytes: 64\n  ves.io.schema.rules.string.min_bytes: 1\n",
                    "title": "Name",
                    "minLength": 1,
                    "maxLength": 64,
                    "x-displayname": "Name",
                    "x-ves-example": "api-cred-x89sf",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_bytes": "64",
                        "ves.io.schema.rules.string.min_bytes": "1"
                    }
                },
                "uid": {
                    "type": "string",
                    "description": " UUID of API credential object.\n\nExample: - \"fa45979f-4e41-4f4b-8b0b-c3ab844ab0aa\"-",
                    "title": "uuid of the record",
                    "x-displayname": "UUID",
                    "x-ves-example": "fa45979f-4e41-4f4b-8b0b-c3ab844ab0aa"
                },
                "user_email": {
                    "type": "string",
                    "description": " User email of user that requested credential .\n\nExample: - \"admin@acmecorp.com\"-",
                    "title": "Email of user",
                    "x-displayname": "User",
                    "x-ves-example": "admin@acmecorp.com"
                }
            }
        }
    },
    "x-displayname": "Site",
    "x-ves-proto-file": "ves.io/schema/site/public_config_kubeconfig.proto"
}`
