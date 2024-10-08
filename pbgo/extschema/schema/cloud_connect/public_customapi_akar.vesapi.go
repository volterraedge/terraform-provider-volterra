// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cloud_connect

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

// Create ConfigCustomAPI GRPC Client satisfying server.CustomClient
type ConfigCustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient ConfigCustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *ConfigCustomAPIGrpcClient) doRPCEdgeCredentials(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &CredentialsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.cloud_connect.CredentialsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.EdgeCredentials(ctx, req, opts...)
	return rsp, err
}

func (c *ConfigCustomAPIGrpcClient) doRPCEdgeList(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &EdgeListRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.cloud_connect.EdgeListRequest", yamlReq)
	}
	rsp, err := c.grpcClient.EdgeList(ctx, req, opts...)
	return rsp, err
}

func (c *ConfigCustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewConfigCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &ConfigCustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewConfigCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["EdgeCredentials"] = ccl.doRPCEdgeCredentials

	rpcFns["EdgeList"] = ccl.doRPCEdgeList

	ccl.rpcFns = rpcFns

	return ccl
}

// Create ConfigCustomAPI REST Client satisfying server.CustomClient
type ConfigCustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *ConfigCustomAPIRestClient) doRPCEdgeCredentials(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &CredentialsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.cloud_connect.CredentialsRequest: %s", yamlReq, err)
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
		q.Add("provider", fmt.Sprintf("%v", req.Provider))

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
	pbRsp := &CredentialsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.cloud_connect.CredentialsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *ConfigCustomAPIRestClient) doRPCEdgeList(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &EdgeListRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.cloud_connect.EdgeListRequest: %s", yamlReq, err)
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
	pbRsp := &EdgeListResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.cloud_connect.EdgeListResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *ConfigCustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewConfigCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &ConfigCustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["EdgeCredentials"] = ccl.doRPCEdgeCredentials

	rpcFns["EdgeList"] = ccl.doRPCEdgeList

	ccl.rpcFns = rpcFns

	return ccl
}

// Create configCustomAPIInprocClient

// INPROC Client (satisfying ConfigCustomAPIClient interface)
type configCustomAPIInprocClient struct {
	ConfigCustomAPIServer
}

func (c *configCustomAPIInprocClient) EdgeCredentials(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*CredentialsResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.cloud_connect.ConfigCustomAPI.EdgeCredentials")
	return c.ConfigCustomAPIServer.EdgeCredentials(ctx, in)
}
func (c *configCustomAPIInprocClient) EdgeList(ctx context.Context, in *EdgeListRequest, opts ...grpc.CallOption) (*EdgeListResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.cloud_connect.ConfigCustomAPI.EdgeList")
	return c.ConfigCustomAPIServer.EdgeList(ctx, in)
}

func NewConfigCustomAPIInprocClient(svc svcfw.Service) ConfigCustomAPIClient {
	return &configCustomAPIInprocClient{ConfigCustomAPIServer: NewConfigCustomAPIServer(svc)}
}

// RegisterGwConfigCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwConfigCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterConfigCustomAPIHandlerClient(ctx, mux, NewConfigCustomAPIInprocClient(s))
}

// Create configCustomAPISrv

// SERVER (satisfying ConfigCustomAPIServer interface)
type configCustomAPISrv struct {
	svc svcfw.Service
}

func (s *configCustomAPISrv) EdgeCredentials(ctx context.Context, in *CredentialsRequest) (*CredentialsResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.cloud_connect.ConfigCustomAPI")
	cah, ok := ah.(ConfigCustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *ConfigCustomAPIServer", ah)
	}

	var (
		rsp *CredentialsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.cloud_connect.CredentialsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'ConfigCustomAPI.EdgeCredentials' operation on 'cloud_connect'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.cloud_connect.ConfigCustomAPI.EdgeCredentials"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.EdgeCredentials(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.cloud_connect.CredentialsResponse", rsp)...)

	return rsp, nil
}
func (s *configCustomAPISrv) EdgeList(ctx context.Context, in *EdgeListRequest) (*EdgeListResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.cloud_connect.ConfigCustomAPI")
	cah, ok := ah.(ConfigCustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *ConfigCustomAPIServer", ah)
	}

	var (
		rsp *EdgeListResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.cloud_connect.EdgeListRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'ConfigCustomAPI.EdgeList' operation on 'cloud_connect'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.cloud_connect.ConfigCustomAPI.EdgeList"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.EdgeList(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.cloud_connect.EdgeListResponse", rsp)...)

	return rsp, nil
}

func NewConfigCustomAPIServer(svc svcfw.Service) ConfigCustomAPIServer {
	return &configCustomAPISrv{svc: svc}
}

var ConfigCustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Cloud Connect",
        "description": "Cloud Connect related public APIs served by akar.\nIt is always scoped by system namespace.",
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
        "/public/namespaces/system/edge_credentials": {
            "post": {
                "summary": "Cloud Credential",
                "description": "Returns the cloud credential for the matching edge type",
                "operationId": "ves.io.schema.cloud_connect.ConfigCustomAPI.EdgeCredentials",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/cloud_connectCredentialsResponse"
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
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cloud_connectCredentialsRequest"
                        }
                    }
                ],
                "tags": [
                    "ConfigCustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-cloud_connect-configcustomapi-edgecredentials"
                },
                "x-ves-proto-rpc": "ves.io.schema.cloud_connect.ConfigCustomAPI.EdgeCredentials"
            },
            "x-displayname": "Cloud Connect Akar CustomAPI",
            "x-ves-proto-service": "ves.io.schema.cloud_connect.ConfigCustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/system/edge_list": {
            "get": {
                "summary": "Edge List",
                "description": "Returns the online edge sites (Both Customer Edge and Cloud Edge)",
                "operationId": "ves.io.schema.cloud_connect.ConfigCustomAPI.EdgeList",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/cloud_connectEdgeListResponse"
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
                "tags": [
                    "ConfigCustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-cloud_connect-configcustomapi-edgelist"
                },
                "x-ves-proto-rpc": "ves.io.schema.cloud_connect.ConfigCustomAPI.EdgeList"
            },
            "x-displayname": "Cloud Connect Akar CustomAPI",
            "x-ves-proto-service": "ves.io.schema.cloud_connect.ConfigCustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "cloud_connectCloudConnectProviderType": {
            "type": "string",
            "description": "Cloud Connect Provider Type\n\n - AWS: AWS\n\nCloud connects backed by AWS cloud\n - AZURE: AZURE\n\nCloud connects backed by Azure cloud\n - GCP: GCP\n\nCloud connects backed by GCP cloud",
            "title": "Cloud Connect Provider",
            "enum": [
                "AWS",
                "AZURE",
                "GCP"
            ],
            "default": "AWS",
            "x-displayname": "Cloud Connect Provider",
            "x-ves-proto-enum": "ves.io.schema.cloud_connect.CloudConnectProviderType"
        },
        "cloud_connectCredentialsRequest": {
            "type": "object",
            "description": "Request to return all the credentials for the matching cloud site type.",
            "title": "CredentialsRequest",
            "x-displayname": "Edge Credentials Request",
            "x-ves-proto-message": "ves.io.schema.cloud_connect.CredentialsRequest",
            "properties": {
                "provider": {
                    "description": " Edge site provider",
                    "title": "Provider",
                    "$ref": "#/definitions/cloud_connectCloudConnectProviderType",
                    "x-displayname": "Provider"
                }
            }
        },
        "cloud_connectCredentialsResponse": {
            "type": "object",
            "description": "Response that returns all the credentials for the matching provider.",
            "title": "CredentialsResponse",
            "x-displayname": "CredentialsResponse",
            "x-ves-proto-message": "ves.io.schema.cloud_connect.CredentialsResponse",
            "properties": {
                "cred": {
                    "type": "array",
                    "description": " Cloud credentials",
                    "title": "Credentials",
                    "items": {
                        "$ref": "#/definitions/schemaviewsObjectRefType"
                    },
                    "x-displayname": "Cloud Credentials"
                }
            }
        },
        "cloud_connectEdgeListResponse": {
            "type": "object",
            "description": "Response body that returns online edge sites both customer edge and cloud edge.",
            "title": "EdgeListResponse",
            "x-displayname": "Edge List Response",
            "x-ves-proto-message": "ves.io.schema.cloud_connect.EdgeListResponse",
            "properties": {
                "edge_site": {
                    "type": "array",
                    "description": " Reference to a edge site",
                    "title": "Edge Site",
                    "items": {
                        "$ref": "#/definitions/cloud_connectEdgeSite"
                    },
                    "x-displayname": "Edge Site"
                }
            }
        },
        "cloud_connectEdgeSite": {
            "type": "object",
            "description": "Reference to a edge site",
            "title": "Edge Site",
            "x-displayname": "Edge Site",
            "x-ves-proto-message": "ves.io.schema.cloud_connect.EdgeSite",
            "properties": {
                "coordinates": {
                    "description": " Edge Coordinates",
                    "title": "Edge Coordinates",
                    "$ref": "#/definitions/siteCoordinates",
                    "x-displayname": "Edge Coordinates"
                },
                "edge_site": {
                    "description": " Reference to a edge site",
                    "title": "Edge Site",
                    "$ref": "#/definitions/ioschemaObjectRefType",
                    "x-displayname": "Egde Site Reference"
                },
                "provider": {
                    "description": " Edge site provider",
                    "title": "Provider",
                    "$ref": "#/definitions/cloud_connectCloudConnectProviderType",
                    "x-displayname": "Provider"
                },
                "region": {
                    "type": "string",
                    "description": " Edge site region",
                    "title": "Region",
                    "x-displayname": "Region"
                }
            }
        },
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
        "schemaviewsObjectRefType": {
            "type": "object",
            "description": "This type establishes a direct reference from one object(the referrer) to another(the referred).\nSuch a reference is in form of tenant/namespace/name",
            "title": "ObjectRefType",
            "x-displayname": "Object reference",
            "x-ves-proto-message": "ves.io.schema.views.ObjectRefType",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then name will hold the referred object's(e.g. route's) name.\n\nExample: - \"contacts-route\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_bytes: 128\n  ves.io.schema.rules.string.min_bytes: 1\n",
                    "title": "name",
                    "minLength": 1,
                    "maxLength": 128,
                    "x-displayname": "Name",
                    "x-ves-example": "contacts-route",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_bytes": "128",
                        "ves.io.schema.rules.string.min_bytes": "1"
                    }
                },
                "namespace": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then namespace will hold the referred object's(e.g. route's) namespace.\n\nExample: - \"ns1\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_bytes: 64\n",
                    "title": "namespace",
                    "maxLength": 64,
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_bytes": "64"
                    }
                },
                "tenant": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then tenant will hold the referred object's(e.g. route's) tenant.\n\nExample: - \"acmecorp\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_bytes: 64\n",
                    "title": "tenant",
                    "maxLength": 64,
                    "x-displayname": "Tenant",
                    "x-ves-example": "acmecorp",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_bytes": "64"
                    }
                }
            }
        },
        "siteCoordinates": {
            "type": "object",
            "description": "Coordinates of the site which provides the site physical location",
            "title": "Site Coordinates",
            "x-displayname": "Site Coordinates",
            "x-ves-proto-message": "ves.io.schema.site.Coordinates",
            "properties": {
                "latitude": {
                    "type": "number",
                    "description": " Latitude of the site location\n\nExample: - \"10.0\"-\n\nValidation Rules:\n  ves.io.schema.rules.float.gte: -90.0\n  ves.io.schema.rules.float.lte: 90.0\n",
                    "title": "latitude",
                    "format": "float",
                    "x-displayname": "Latitude",
                    "x-ves-example": "10.0",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.float.gte": "-90.0",
                        "ves.io.schema.rules.float.lte": "90.0"
                    }
                },
                "longitude": {
                    "type": "number",
                    "description": " longitude of site location\n\nExample: - \"20.0\"-\n\nValidation Rules:\n  ves.io.schema.rules.float.gte: -180.0\n  ves.io.schema.rules.float.lte: 180.0\n",
                    "title": "longitude",
                    "format": "float",
                    "x-displayname": "Longitude",
                    "x-ves-example": "20.0",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.float.gte": "-180.0",
                        "ves.io.schema.rules.float.lte": "180.0"
                    }
                }
            }
        }
    },
    "x-displayname": "Cloud Connect",
    "x-ves-proto-file": "ves.io/schema/cloud_connect/public_customapi_akar.proto"
}`
