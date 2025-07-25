// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package uztna_active_sessions

import (
	"bytes"
	"context"
	"fmt"
	io "io"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	google_protobuf "github.com/gogo/protobuf/types"
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

// Create TerminateSessionsAPI GRPC Client satisfying server.CustomClient
type TerminateSessionsAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient TerminateSessionsAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *TerminateSessionsAPIGrpcClient) doRPCTerminateSessions(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &TerminateSessionsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.TerminateSessions(ctx, req, opts...)
	return rsp, err
}

func (c *TerminateSessionsAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewTerminateSessionsAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &TerminateSessionsAPIGrpcClient{
		conn:       cc,
		grpcClient: NewTerminateSessionsAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["TerminateSessions"] = ccl.doRPCTerminateSessions

	ccl.rpcFns = rpcFns

	return ccl
}

// Create TerminateSessionsAPI REST Client satisfying server.CustomClient
type TerminateSessionsAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *TerminateSessionsAPIRestClient) doRPCTerminateSessions(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &TerminateSessionsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsRequest: %s", yamlReq, err)
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
		q.Add("session_id", fmt.Sprintf("%v", req.SessionId))

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
	pbRsp := &google_protobuf.Empty{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *google.protobuf.Empty", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *TerminateSessionsAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewTerminateSessionsAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &TerminateSessionsAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["TerminateSessions"] = ccl.doRPCTerminateSessions

	ccl.rpcFns = rpcFns

	return ccl
}

// Create terminateSessionsAPIInprocClient

// INPROC Client (satisfying TerminateSessionsAPIClient interface)
type terminateSessionsAPIInprocClient struct {
	TerminateSessionsAPIServer
}

func (c *terminateSessionsAPIInprocClient) TerminateSessions(ctx context.Context, in *TerminateSessionsRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsAPI.TerminateSessions")
	return c.TerminateSessionsAPIServer.TerminateSessions(ctx, in)
}

func NewTerminateSessionsAPIInprocClient(svc svcfw.Service) TerminateSessionsAPIClient {
	return &terminateSessionsAPIInprocClient{TerminateSessionsAPIServer: NewTerminateSessionsAPIServer(svc)}
}

// RegisterGwTerminateSessionsAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwTerminateSessionsAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterTerminateSessionsAPIHandlerClient(ctx, mux, NewTerminateSessionsAPIInprocClient(s))
}

// Create terminateSessionsAPISrv

// SERVER (satisfying TerminateSessionsAPIServer interface)
type terminateSessionsAPISrv struct {
	svc svcfw.Service
}

func (s *terminateSessionsAPISrv) TerminateSessions(ctx context.Context, in *TerminateSessionsRequest) (*google_protobuf.Empty, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsAPI")
	cah, ok := ah.(TerminateSessionsAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *TerminateSessionsAPIServer", ah)
	}

	var (
		rsp *google_protobuf.Empty
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'TerminateSessionsAPI.TerminateSessions' operation on 'uztna_active_sessions'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsAPI.TerminateSessions"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.TerminateSessions(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "google.protobuf.Empty", rsp)...)

	return rsp, nil
}

func NewTerminateSessionsAPIServer(svc svcfw.Service) TerminateSessionsAPIServer {
	return &terminateSessionsAPISrv{svc: svc}
}

var TerminateSessionsAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "manage active sessions",
        "description": "APIs to monitor UZTNA active sessions on all applications.",
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
        "/public/namespaces/{namespace}/uztna/delete_access_session/{session_id}": {
            "delete": {
                "summary": "Terminate Access Sessions",
                "description": "Terminate uztna access sessions",
                "operationId": "ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsAPI.TerminateSessions",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {}
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
                        "description": "namespace\n\nx-example: \"ns1\"\nNamespace of the App type for the current request",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "session_id",
                        "description": "session_id\n\nx-required\nx-example: \"9d6d591\"\nID of the session",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Session ID"
                    }
                ],
                "tags": [
                    "TerminateSessionsAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-uztna-uztna_active_sessions-terminatesessionsapi-terminatesessions"
                },
                "x-ves-proto-rpc": "ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsAPI.TerminateSessions"
            },
            "x-displayname": "Terminate Access Session API",
            "x-ves-proto-service": "ves.io.schema.uztna.uztna_active_sessions.TerminateSessionsAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {},
    "x-displayname": "Manage Active Sessions",
    "x-ves-proto-file": "ves.io/schema/uztna/uztna_active_sessions/public_uztna_terminate_sessions.proto"
}`
