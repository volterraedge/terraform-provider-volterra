// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package xc_saas

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

func (c *CustomAPIGrpcClient) doRPCGetRegistrationDetails(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &GetRegistrationDetailsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.xc_saas.GetRegistrationDetailsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.GetRegistrationDetails(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) doRPCSendSignupEmail(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SendSignupEmailRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.xc_saas.SendSignupEmailRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SendSignupEmail(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) doRPCSignupXCSaaS(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &XCSaaSSignupRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.xc_saas.XCSaaSSignupRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SignupXCSaaS(ctx, req, opts...)
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
	rpcFns["GetRegistrationDetails"] = ccl.doRPCGetRegistrationDetails

	rpcFns["SendSignupEmail"] = ccl.doRPCSendSignupEmail

	rpcFns["SignupXCSaaS"] = ccl.doRPCSignupXCSaaS

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

func (c *CustomAPIRestClient) doRPCGetRegistrationDetails(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &GetRegistrationDetailsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.xc_saas.GetRegistrationDetailsRequest: %s", yamlReq, err)
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
		q.Add("token", fmt.Sprintf("%v", req.Token))

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
	pbRsp := &GetRegistrationDetailsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.marketplace.xc_saas.GetRegistrationDetailsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) doRPCSendSignupEmail(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SendSignupEmailRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.xc_saas.SendSignupEmailRequest: %s", yamlReq, err)
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
		q.Add("token", fmt.Sprintf("%v", req.Token))

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
	pbRsp := &SendEmailResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.marketplace.xc_saas.SendEmailResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) doRPCSignupXCSaaS(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &XCSaaSSignupRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.xc_saas.XCSaaSSignupRequest: %s", yamlReq, err)
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
		q.Add("domain", fmt.Sprintf("%v", req.Domain))
		q.Add("token", fmt.Sprintf("%v", req.Token))

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
	pbRsp := &XCSaaSSignupResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.marketplace.xc_saas.XCSaaSSignupResponse", body)

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
	rpcFns["GetRegistrationDetails"] = ccl.doRPCGetRegistrationDetails

	rpcFns["SendSignupEmail"] = ccl.doRPCSendSignupEmail

	rpcFns["SignupXCSaaS"] = ccl.doRPCSignupXCSaaS

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customAPIInprocClient

// INPROC Client (satisfying CustomAPIClient interface)
type customAPIInprocClient struct {
	CustomAPIServer
}

func (c *customAPIInprocClient) GetRegistrationDetails(ctx context.Context, in *GetRegistrationDetailsRequest, opts ...grpc.CallOption) (*GetRegistrationDetailsResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.marketplace.xc_saas.CustomAPI.GetRegistrationDetails")
	return c.CustomAPIServer.GetRegistrationDetails(ctx, in)
}
func (c *customAPIInprocClient) SendSignupEmail(ctx context.Context, in *SendSignupEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.marketplace.xc_saas.CustomAPI.SendSignupEmail")
	return c.CustomAPIServer.SendSignupEmail(ctx, in)
}
func (c *customAPIInprocClient) SignupXCSaaS(ctx context.Context, in *XCSaaSSignupRequest, opts ...grpc.CallOption) (*XCSaaSSignupResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.marketplace.xc_saas.CustomAPI.SignupXCSaaS")
	return c.CustomAPIServer.SignupXCSaaS(ctx, in)
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

func (s *customAPISrv) GetRegistrationDetails(ctx context.Context, in *GetRegistrationDetailsRequest) (*GetRegistrationDetailsResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.marketplace.xc_saas.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPIServer", ah)
	}

	var (
		rsp *GetRegistrationDetailsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.marketplace.xc_saas.GetRegistrationDetailsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.GetRegistrationDetails' operation on 'xc_saas'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.marketplace.xc_saas.CustomAPI.GetRegistrationDetails"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.GetRegistrationDetails(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.marketplace.xc_saas.GetRegistrationDetailsResponse", rsp)...)

	return rsp, nil
}
func (s *customAPISrv) SendSignupEmail(ctx context.Context, in *SendSignupEmailRequest) (*SendEmailResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.marketplace.xc_saas.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPIServer", ah)
	}

	var (
		rsp *SendEmailResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.marketplace.xc_saas.SendSignupEmailRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.SendSignupEmail' operation on 'xc_saas'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.marketplace.xc_saas.CustomAPI.SendSignupEmail"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SendSignupEmail(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.marketplace.xc_saas.SendEmailResponse", rsp)...)

	return rsp, nil
}
func (s *customAPISrv) SignupXCSaaS(ctx context.Context, in *XCSaaSSignupRequest) (*XCSaaSSignupResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.marketplace.xc_saas.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPIServer", ah)
	}

	var (
		rsp *XCSaaSSignupResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.marketplace.xc_saas.XCSaaSSignupRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.SignupXCSaaS' operation on 'xc_saas'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.marketplace.xc_saas.CustomAPI.SignupXCSaaS"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SignupXCSaaS(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.marketplace.xc_saas.XCSaaSSignupResponse", rsp)...)

	return rsp, nil
}

func NewCustomAPIServer(svc svcfw.Service) CustomAPIServer {
	return &customAPISrv{svc: svc}
}

var CustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "XC SaaS Onboarding API",
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
        "/no_auth/namespaces/system/f5xc-saas/signup": {
            "post": {
                "summary": "Signup XC SaaS",
                "description": "Use this API to signup registered Azure Service Bus (ASB) provisioning message for F5XC service.",
                "operationId": "ves.io.schema.marketplace.xc_saas.CustomAPI.SignupXCSaaS",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/xc_saasXCSaaSSignupResponse"
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
                            "$ref": "#/definitions/xc_saasXCSaaSSignupRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-marketplace-xc_saas-customapi-signupxcsaas"
                },
                "x-ves-proto-rpc": "ves.io.schema.marketplace.xc_saas.CustomAPI.SignupXCSaaS"
            },
            "x-displayname": "XC SaaS Onboarding Custom API",
            "x-ves-proto-service": "ves.io.schema.marketplace.xc_saas.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/no_auth/namespaces/system/f5xc-saas/signup/registration_details": {
            "get": {
                "summary": "Get Registration Details",
                "description": "Use this API to to get registration details (currently limited to email address and domain) associated with a specific asb_message object",
                "operationId": "ves.io.schema.marketplace.xc_saas.CustomAPI.GetRegistrationDetails",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/xc_saasGetRegistrationDetailsResponse"
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
                        "name": "token",
                        "description": "x-example: \"7b22656e7469746c656d656e745f737562736372697074696f6e5f6964223a2261393933666633302d363138392d346231612d623064632d3630623034393063666463302d5445454d2d62343438666663362d34353166222c2276616c69645f74696c6c223a22323032342d31312d31355431373a31393a35382e3334323934352d30383a3030227d:d1bd81da4772d62fb424b805391fc946c611b5c245246e444410bcb0fe548ad7\"\nThe token contains the encrypted data, secured using an HMAC key.",
                        "in": "query",
                        "required": false,
                        "type": "string",
                        "x-displayname": "Token"
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-marketplace-xc_saas-customapi-getregistrationdetails"
                },
                "x-ves-proto-rpc": "ves.io.schema.marketplace.xc_saas.CustomAPI.GetRegistrationDetails"
            },
            "x-displayname": "XC SaaS Onboarding Custom API",
            "x-ves-proto-service": "ves.io.schema.marketplace.xc_saas.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/no_auth/namespaces/system/f5xc-saas/signup/send_email": {
            "post": {
                "summary": "Send Signup Email",
                "description": "Use this API to send a tenant provisioning signup email when the signup URL link in the previously sent email has expired or is no longer accessible",
                "operationId": "ves.io.schema.marketplace.xc_saas.CustomAPI.SendSignupEmail",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/xc_saasSendEmailResponse"
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
                            "$ref": "#/definitions/xc_saasSendSignupEmailRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-marketplace-xc_saas-customapi-sendsignupemail"
                },
                "x-ves-proto-rpc": "ves.io.schema.marketplace.xc_saas.CustomAPI.SendSignupEmail"
            },
            "x-displayname": "XC SaaS Onboarding Custom API",
            "x-ves-proto-service": "ves.io.schema.marketplace.xc_saas.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "xc_saasGetRegistrationDetailsResponse": {
            "type": "object",
            "description": "Response format for returning registration details associated with the ASB message",
            "title": "GetRegistrationDetailsResponse",
            "x-displayname": "Get Registration Details Response",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.GetRegistrationDetailsResponse",
            "properties": {
                "console_url": {
                    "type": "string",
                    "description": " A URL to the console page, returned if the tenant is already signed up and active\n\nExample: - \"console_urlhttps://console.ves.volterra.io/login/start\"-",
                    "title": "console_url",
                    "x-displayname": "Console URL",
                    "x-ves-example": "console_url: https://console.ves.volterra.io/login/start"
                },
                "domain": {
                    "type": "string",
                    "description": " The domain associated with the registered ASB message",
                    "title": "domain",
                    "x-displayname": "Domain"
                },
                "email": {
                    "type": "string",
                    "description": " The email associated with the registered ASB message",
                    "title": "email",
                    "x-displayname": "Email"
                },
                "error_message": {
                    "type": "string",
                    "description": " An error message providing additional details (e.g., tenant already signed up)\n\nExample: - \"error_message\"\"Your account is already active. Please log in to continue.\"-",
                    "title": "error_message",
                    "x-displayname": "Error Message",
                    "x-ves-example": "error_message\": \"Your account is already active. Please log in to continue."
                }
            }
        },
        "xc_saasSendEmailResponse": {
            "type": "object",
            "description": "Response to indicate if the email was sent successfully",
            "title": "SendEmailResponse",
            "x-displayname": "Send Email Response",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.SendEmailResponse",
            "properties": {
                "message": {
                    "type": "string",
                    "description": " Message provides additional context regarding the result of the email send operation",
                    "title": "message",
                    "x-displayname": "Message"
                },
                "success": {
                    "type": "boolean",
                    "description": " Indicates if the email was sent successfully",
                    "title": "success",
                    "format": "boolean",
                    "x-displayname": "Success"
                }
            }
        },
        "xc_saasSendSignupEmailRequest": {
            "type": "object",
            "description": "The request message for SendSignupEmail",
            "title": "SendSignupEmailRequest",
            "x-displayname": "Send Signup Email Request",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.SendSignupEmailRequest",
            "properties": {
                "token": {
                    "type": "string",
                    "description": " The token contains the encrypted data, secured using an HMAC key.\n\nExample: - \"7b22656e7469746c656d656e745f737562736372697074696f6e5f6964223a2261393933666633302d363138392d346231612d623064632d3630623034393063666463302d5445454d2d62343438666663362d34353166222c2276616c69645f74696c6c223a22323032342d31312d31355431373a31393a35382e3334323934352d30383a3030227d:d1bd81da4772d62fb424b805391fc946c611b5c245246e444410bcb0fe548ad7\"-",
                    "title": "token",
                    "x-displayname": "Token",
                    "x-ves-example": "7b22656e7469746c656d656e745f737562736372697074696f6e5f6964223a2261393933666633302d363138392d346231612d623064632d3630623034393063666463302d5445454d2d62343438666663362d34353166222c2276616c69645f74696c6c223a22323032342d31312d31355431373a31393a35382e3334323934352d30383a3030227d:d1bd81da4772d62fb424b805391fc946c611b5c245246e444410bcb0fe548ad7"
                }
            }
        },
        "xc_saasXCSaaSSignupRequest": {
            "type": "object",
            "title": "XCSaaSSignupRequest",
            "x-displayname": "XC SaaS Signup Request",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.XCSaaSSignupRequest",
            "properties": {
                "domain": {
                    "type": "string",
                    "description": " The domain associated with the tenant/account\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 17\n  ves.io.schema.rules.string.min_len: 1\n  ves.io.schema.rules.string.not_in: [\\\"ves\\\",\\\"volterra\\\",\\\"ves-io\\\",\\\"f5xc\\\"]\n  ves.io.schema.rules.string.ves_object_name: true\n",
                    "title": "domain",
                    "minLength": 1,
                    "maxLength": 17,
                    "x-displayname": "Domain",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "17",
                        "ves.io.schema.rules.string.min_len": "1",
                        "ves.io.schema.rules.string.not_in": "[\\\"ves\\\",\\\"volterra\\\",\\\"ves-io\\\",\\\"f5xc\\\"]",
                        "ves.io.schema.rules.string.ves_object_name": "true"
                    }
                },
                "token": {
                    "type": "string",
                    "description": " The token contains the encrypted data, secured using an HMAC key.\n\nExample: - 7b22656e7469746c656d656e745f737562736372697074696f6e5f6964223a2261393933666633302d363138392d346231612d623064632d3630623034393063666463302d5445454d2d62343438666663362d34353166222c2276616c69645f74696c6c223a22323032342d31312d31355431373a31393a35382e3334323934352d30383a3030227d:d1bd81da4772d62fb424b805391fc946c611b5c245246e444410bcb0fe548ad7-",
                    "title": "token",
                    "x-displayname": "Token"
                }
            }
        },
        "xc_saasXCSaaSSignupResponse": {
            "type": "object",
            "title": "XCSaaSSignupResponse",
            "x-displayname": "XC SaaS Signup Response",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.XCSaaSSignupResponse"
        }
    },
    "x-displayname": "XC SaaS Onboarding",
    "x-ves-proto-file": "ves.io/schema/marketplace/xc_saas/public_customapi.proto"
}`
