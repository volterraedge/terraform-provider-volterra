// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cloud_link

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

// Create CustomDataAPI GRPC Client satisfying server.CustomClient
type CustomDataAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomDataAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomDataAPIGrpcClient) doRPCReapplyConfig(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &ReapplyConfigRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.cloud_link.ReapplyConfigRequest", yamlReq)
	}
	rsp, err := c.grpcClient.ReapplyConfig(ctx, req, opts...)
	return rsp, err
}

func (c *CustomDataAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomDataAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomDataAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomDataAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["ReapplyConfig"] = ccl.doRPCReapplyConfig

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomDataAPI REST Client satisfying server.CustomClient
type CustomDataAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomDataAPIRestClient) doRPCReapplyConfig(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &ReapplyConfigRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.cloud_link.ReapplyConfigRequest: %s", yamlReq, err)
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
	pbRsp := &ReapplyConfigResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.cloud_link.ReapplyConfigResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomDataAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomDataAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomDataAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["ReapplyConfig"] = ccl.doRPCReapplyConfig

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customDataAPIInprocClient

// INPROC Client (satisfying CustomDataAPIClient interface)
type customDataAPIInprocClient struct {
	CustomDataAPIServer
}

func (c *customDataAPIInprocClient) ReapplyConfig(ctx context.Context, in *ReapplyConfigRequest, opts ...grpc.CallOption) (*ReapplyConfigResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.cloud_link.CustomDataAPI.ReapplyConfig")
	return c.CustomDataAPIServer.ReapplyConfig(ctx, in)
}

func NewCustomDataAPIInprocClient(svc svcfw.Service) CustomDataAPIClient {
	return &customDataAPIInprocClient{CustomDataAPIServer: NewCustomDataAPIServer(svc)}
}

// RegisterGwCustomDataAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomDataAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomDataAPIHandlerClient(ctx, mux, NewCustomDataAPIInprocClient(s))
}

// Create customDataAPISrv

// SERVER (satisfying CustomDataAPIServer interface)
type customDataAPISrv struct {
	svc svcfw.Service
}

func (s *customDataAPISrv) ReapplyConfig(ctx context.Context, in *ReapplyConfigRequest) (*ReapplyConfigResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.cloud_link.CustomDataAPI")
	cah, ok := ah.(CustomDataAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomDataAPIServer", ah)
	}

	var (
		rsp *ReapplyConfigResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.cloud_link.ReapplyConfigRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomDataAPI.ReapplyConfig' operation on 'cloud_link'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.cloud_link.CustomDataAPI.ReapplyConfig"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.ReapplyConfig(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.cloud_link.ReapplyConfigResponse", rsp)...)

	return rsp, nil
}

func NewCustomDataAPIServer(svc svcfw.Service) CustomDataAPIServer {
	return &customDataAPISrv{svc: svc}
}

var CustomDataAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "CloudLink",
        "description": "APIs to manage Cloud Link resources",
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
        "/public/namespaces/system/cloud_links/{name}/reapply_config": {
            "post": {
                "summary": "CloudLink ",
                "description": "Reapply CloudLink Config",
                "operationId": "ves.io.schema.cloud_link.CustomDataAPI.ReapplyConfig",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/cloud_linkReapplyConfigResponse"
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
                        "name": "name",
                        "description": "Name\n\nx-example: \"aws-cloud-link-east\"\nName of Cloud Link.",
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
                            "$ref": "#/definitions/cloud_linkReapplyConfigRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomDataAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-cloud_link-customdataapi-reapplyconfig"
                },
                "x-ves-proto-rpc": "ves.io.schema.cloud_link.CustomDataAPI.ReapplyConfig"
            },
            "x-displayname": "CloudLink Custom Data API",
            "x-ves-proto-service": "ves.io.schema.cloud_link.CustomDataAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "cloud_linkReapplyConfigRequest": {
            "type": "object",
            "description": "Reapply CloudLink Config",
            "title": "CloudLink Reapply Config Request",
            "x-displayname": "CloudLink Reapply Config Request",
            "x-ves-proto-message": "ves.io.schema.cloud_link.ReapplyConfigRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Name of Cloud Link.\n\nExample: - \"aws-cloud-link-east\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "aws-cloud-link-east"
                }
            }
        },
        "cloud_linkReapplyConfigResponse": {
            "type": "object",
            "description": "Reapply CloudLink Config",
            "title": "CloudLink Reapply Config Response",
            "x-displayname": "CloudLink Reapply Config Response",
            "x-ves-proto-message": "ves.io.schema.cloud_link.ReapplyConfigResponse"
        }
    },
    "x-displayname": "CloudLink",
    "x-ves-proto-file": "ves.io/schema/cloud_link/public_custom_api.proto"
}`
