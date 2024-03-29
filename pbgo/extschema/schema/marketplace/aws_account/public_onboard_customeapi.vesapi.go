// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package aws_account

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

// Create OnboardCustomAPI GRPC Client satisfying server.CustomClient
type OnboardCustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient OnboardCustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *OnboardCustomAPIGrpcClient) doRPCRegisterNewAWSAccount(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &RegistrationRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.aws_account.RegistrationRequest", yamlReq)
	}
	rsp, err := c.grpcClient.RegisterNewAWSAccount(ctx, req, opts...)
	return rsp, err
}

func (c *OnboardCustomAPIGrpcClient) doRPCSignupAWSAccount(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &AWSAccountSignupRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.aws_account.AWSAccountSignupRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SignupAWSAccount(ctx, req, opts...)
	return rsp, err
}

func (c *OnboardCustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewOnboardCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &OnboardCustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewOnboardCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["RegisterNewAWSAccount"] = ccl.doRPCRegisterNewAWSAccount

	rpcFns["SignupAWSAccount"] = ccl.doRPCSignupAWSAccount

	ccl.rpcFns = rpcFns

	return ccl
}

// Create OnboardCustomAPI REST Client satisfying server.CustomClient
type OnboardCustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *OnboardCustomAPIRestClient) doRPCRegisterNewAWSAccount(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &RegistrationRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.aws_account.RegistrationRequest: %s", yamlReq, err)
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
		q.Add("x_amzn_marketplace_token", fmt.Sprintf("%v", req.XAmznMarketplaceToken))

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
	pbRsp := &RegistrationResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.marketplace.aws_account.RegistrationResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *OnboardCustomAPIRestClient) doRPCSignupAWSAccount(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &AWSAccountSignupRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.aws_account.AWSAccountSignupRequest: %s", yamlReq, err)
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
		q.Add("account_details", fmt.Sprintf("%v", req.AccountDetails))
		q.Add("company_details", fmt.Sprintf("%v", req.CompanyDetails))
		q.Add("crm_details", fmt.Sprintf("%v", req.CrmDetails))
		q.Add("user_details", fmt.Sprintf("%v", req.UserDetails))

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
	pbRsp := &AWSAccountSignupResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.marketplace.aws_account.AWSAccountSignupResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *OnboardCustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewOnboardCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &OnboardCustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["RegisterNewAWSAccount"] = ccl.doRPCRegisterNewAWSAccount

	rpcFns["SignupAWSAccount"] = ccl.doRPCSignupAWSAccount

	ccl.rpcFns = rpcFns

	return ccl
}

// Create onboardCustomAPIInprocClient

// INPROC Client (satisfying OnboardCustomAPIClient interface)
type onboardCustomAPIInprocClient struct {
	OnboardCustomAPIServer
}

func (c *onboardCustomAPIInprocClient) RegisterNewAWSAccount(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.marketplace.aws_account.OnboardCustomAPI.RegisterNewAWSAccount")
	return c.OnboardCustomAPIServer.RegisterNewAWSAccount(ctx, in)
}
func (c *onboardCustomAPIInprocClient) SignupAWSAccount(ctx context.Context, in *AWSAccountSignupRequest, opts ...grpc.CallOption) (*AWSAccountSignupResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.marketplace.aws_account.OnboardCustomAPI.SignupAWSAccount")
	return c.OnboardCustomAPIServer.SignupAWSAccount(ctx, in)
}

func NewOnboardCustomAPIInprocClient(svc svcfw.Service) OnboardCustomAPIClient {
	return &onboardCustomAPIInprocClient{OnboardCustomAPIServer: NewOnboardCustomAPIServer(svc)}
}

// RegisterGwOnboardCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwOnboardCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterOnboardCustomAPIHandlerClient(ctx, mux, NewOnboardCustomAPIInprocClient(s))
}

// Create onboardCustomAPISrv

// SERVER (satisfying OnboardCustomAPIServer interface)
type onboardCustomAPISrv struct {
	svc svcfw.Service
}

func (s *onboardCustomAPISrv) RegisterNewAWSAccount(ctx context.Context, in *RegistrationRequest) (*RegistrationResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.marketplace.aws_account.OnboardCustomAPI")
	cah, ok := ah.(OnboardCustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *OnboardCustomAPIServer", ah)
	}

	var (
		rsp *RegistrationResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.marketplace.aws_account.RegistrationRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'OnboardCustomAPI.RegisterNewAWSAccount' operation on 'aws_account'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.marketplace.aws_account.OnboardCustomAPI.RegisterNewAWSAccount"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.RegisterNewAWSAccount(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.marketplace.aws_account.RegistrationResponse", rsp)...)

	return rsp, nil
}
func (s *onboardCustomAPISrv) SignupAWSAccount(ctx context.Context, in *AWSAccountSignupRequest) (*AWSAccountSignupResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.marketplace.aws_account.OnboardCustomAPI")
	cah, ok := ah.(OnboardCustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *OnboardCustomAPIServer", ah)
	}

	var (
		rsp *AWSAccountSignupResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.marketplace.aws_account.AWSAccountSignupRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'OnboardCustomAPI.SignupAWSAccount' operation on 'aws_account'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.marketplace.aws_account.OnboardCustomAPI.SignupAWSAccount"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SignupAWSAccount(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.marketplace.aws_account.AWSAccountSignupResponse", rsp)...)

	return rsp, nil
}

func NewOnboardCustomAPIServer(svc svcfw.Service) OnboardCustomAPIServer {
	return &onboardCustomAPISrv{svc: svc}
}

var OnboardCustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Onboard AWS SaaS PAYG account",
        "description": "APIs to manage AWS Account resources.",
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
        "/no_auth/namespaces/system/aws/f5xc-saas/register": {
            "post": {
                "summary": "Register New AWS Account",
                "description": "Use this API to register F5XC AWS marketplace product for F5XC service.",
                "operationId": "ves.io.schema.marketplace.aws_account.OnboardCustomAPI.RegisterNewAWSAccount",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/aws_accountRegistrationResponse"
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
                            "$ref": "#/definitions/aws_accountRegistrationRequest"
                        }
                    }
                ],
                "tags": [
                    "OnboardCustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-marketplace-aws_account-onboardcustomapi-registernewawsaccount"
                },
                "x-ves-proto-rpc": "ves.io.schema.marketplace.aws_account.OnboardCustomAPI.RegisterNewAWSAccount"
            },
            "x-displayname": "AWS SaaS PAYG account",
            "x-ves-proto-service": "ves.io.schema.marketplace.aws_account.OnboardCustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/no_auth/namespaces/system/aws/f5xc-saas/signup": {
            "post": {
                "summary": "Signup AWS Account",
                "description": "Use this API to signup AWS account for F5XC service.",
                "operationId": "ves.io.schema.marketplace.aws_account.OnboardCustomAPI.SignupAWSAccount",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/aws_accountAWSAccountSignupResponse"
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
                            "$ref": "#/definitions/aws_accountAWSAccountSignupRequest"
                        }
                    }
                ],
                "tags": [
                    "OnboardCustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-marketplace-aws_account-onboardcustomapi-signupawsaccount"
                },
                "x-ves-proto-rpc": "ves.io.schema.marketplace.aws_account.OnboardCustomAPI.SignupAWSAccount"
            },
            "x-displayname": "AWS SaaS PAYG account",
            "x-ves-proto-service": "ves.io.schema.marketplace.aws_account.OnboardCustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "aws_accountAWSAccountSignupRequest": {
            "type": "object",
            "title": "AWSAccountSignupRequest",
            "x-displayname": "AWS Account Signup Request",
            "x-ves-proto-message": "ves.io.schema.marketplace.aws_account.AWSAccountSignupRequest",
            "properties": {
                "account_details": {
                    "description": " Details of the new F5XC account to be created\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "Account Details",
                    "$ref": "#/definitions/signupAccountMeta",
                    "x-displayname": "Account Details",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "company_details": {
                    "description": " Details of the company\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "Company Details",
                    "$ref": "#/definitions/signupCompanyMeta",
                    "x-displayname": "Company Details",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "crm_details": {
                    "description": " This field holds CRM information",
                    "title": "CRM Details",
                    "$ref": "#/definitions/schemaCRMInfo",
                    "x-displayname": "CRM Details"
                },
                "user_details": {
                    "description": " Details of the user\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "User Details",
                    "$ref": "#/definitions/signupUserMeta",
                    "x-displayname": "User Details",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                }
            }
        },
        "aws_accountAWSAccountSignupResponse": {
            "type": "object",
            "title": "AWSAccountSignupResponse",
            "x-displayname": "AWS Account Signup Response",
            "x-ves-proto-message": "ves.io.schema.marketplace.aws_account.AWSAccountSignupResponse"
        },
        "aws_accountRegistrationRequest": {
            "type": "object",
            "description": "Request to register F5XC AWS marketplace product for F5XC service.",
            "title": "RegistrationRequest",
            "x-displayname": "AWS SaaS PAYG Registration Request",
            "x-ves-proto-message": "ves.io.schema.marketplace.aws_account.RegistrationRequest",
            "properties": {
                "x_amzn_marketplace_token": {
                    "type": "string",
                    "description": " AWS customer’s registration token\n\nExample: - \"x-amzn-marketplace-token=\u003ctoken val\u003e\"-",
                    "title": "x-amzn-marketplace-token",
                    "x-displayname": "x-amzn-marketplace-token",
                    "x-ves-example": "x-amzn-marketplace-token=\u003ctoken val\u003e"
                }
            }
        },
        "aws_accountRegistrationResponse": {
            "type": "object",
            "description": "Response to register F5XC AWS marketplace product",
            "title": "RegistrationResponse",
            "x-displayname": "AWS SaaS PAYG Registration Response",
            "x-ves-proto-message": "ves.io.schema.marketplace.aws_account.RegistrationResponse",
            "properties": {
                "redirect_url": {
                    "type": "string",
                    "description": " Registration redirect URL for redirecting the AWS customer to login/registration page\n\nExample: - \"redirect_urlhttps://console.ves.volterra.io/login/start\"-",
                    "title": "redirect_url",
                    "x-displayname": "Redirect URL",
                    "x-ves-example": "redirect_url: https://console.ves.volterra.io/login/start"
                }
            }
        },
        "schemaCRMInfo": {
            "type": "object",
            "description": "CRM Information",
            "title": "CRM Information",
            "x-displayname": "CRM Information",
            "x-ves-proto-message": "ves.io.schema.CRMInfo"
        },
        "signupAccountMeta": {
            "type": "object",
            "title": "Account Meta",
            "x-displayname": "Account Meta",
            "x-ves-proto-message": "ves.io.schema.signup.AccountMeta",
            "properties": {
                "domain": {
                    "type": "string",
                    "description": " domain of the account\n\nExample: - \"john-deer\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 17\n  ves.io.schema.rules.string.not_in: [\\\"ves\\\",\\\"volterra\\\"]\n  ves.io.schema.rules.string.ves_object_name: true\n",
                    "title": "Domain",
                    "maxLength": 17,
                    "x-displayname": "Domain",
                    "x-ves-example": "john-deer",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "17",
                        "ves.io.schema.rules.string.not_in": "[\\\"ves\\\",\\\"volterra\\\"]",
                        "ves.io.schema.rules.string.ves_object_name": "true"
                    }
                },
                "locale": {
                    "type": "string",
                    "description": " locale of the account\n\nExample: - \"en-US\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 16\n",
                    "title": "Locale",
                    "maxLength": 16,
                    "x-displayname": "Locale",
                    "x-ves-example": "en-US",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "16"
                    }
                },
                "tos_accepted_at": {
                    "type": "string",
                    "description": " terms of services accepted timestamp\n\nExample: - \"2020-04-20T12:32:51.341959216Z\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.timestamp.lt_now: true\n",
                    "title": "TOS Accepted Timestamp",
                    "format": "date-time",
                    "x-displayname": "TOS Accepted Timestamp",
                    "x-ves-example": "2020-04-20T12:32:51.341959216Z",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.timestamp.lt_now": "true"
                    }
                },
                "tos_version": {
                    "type": "string",
                    "description": " terms of services version\n\nExample: - \"v2022.3.14\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_len: 12\n  ves.io.schema.rules.string.min_len: 2\n",
                    "title": "TOS Version",
                    "minLength": 2,
                    "maxLength": 12,
                    "x-displayname": "TOS Version",
                    "x-ves-example": "v2022.3.14",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_len": "12",
                        "ves.io.schema.rules.string.min_len": "2"
                    }
                }
            }
        },
        "signupCompanyMeta": {
            "type": "object",
            "title": "Company Meta",
            "x-displayname": "Company Meta",
            "x-ves-proto-message": "ves.io.schema.signup.CompanyMeta",
            "properties": {
                "mailing_address": {
                    "description": " mailing address of the company",
                    "title": "Mailing Address",
                    "$ref": "#/definitions/signupContactMeta",
                    "x-displayname": "Mailing Address"
                },
                "name": {
                    "type": "string",
                    "description": " name of the company\n\nExample: - \"F5 Networks, Inc\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n",
                    "title": "Name",
                    "maxLength": 256,
                    "x-displayname": "Name",
                    "x-ves-example": "F5 Networks, Inc",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256"
                    }
                }
            }
        },
        "signupContactMeta": {
            "type": "object",
            "description": "Instance of one single contact that can be used to communicate with customers.\nDepending on contact type we use these details to send general communication (regular, physical mail) or invoices.",
            "title": "ContactMeta",
            "x-displayname": "ContactMeta",
            "x-ves-proto-message": "ves.io.schema.signup.ContactMeta",
            "properties": {
                "address1": {
                    "type": "string",
                    "description": " address line 1\n\nExample: - \"1234 Main road\"-",
                    "title": "address1",
                    "x-displayname": "Address Line 1",
                    "x-ves-example": "1234 Main road"
                },
                "address2": {
                    "type": "string",
                    "description": " address line 2\n\nExample: - \"P.O BOX 56\"-",
                    "title": "address2",
                    "x-displayname": "Address Line 2",
                    "x-ves-example": "P.O BOX 56"
                },
                "city": {
                    "type": "string",
                    "description": " city / town of the contact\n\nExample: - \"Sunnyvale\"-",
                    "title": "city",
                    "x-displayname": "City",
                    "x-ves-example": "Sunnyvale"
                },
                "country": {
                    "type": "string",
                    "description": " country of contact (e.g. USA). refer to https://en.wikipedia.org/wiki/ISO_3166-1, column alpha-2\n\nExample: - \"US\"-",
                    "title": "country",
                    "x-displayname": "Country",
                    "x-ves-example": "US"
                },
                "county": {
                    "type": "string",
                    "description": " county (optional, for countries where they have counties)\n\nExample: - \"Santa Clara\"-",
                    "title": "county",
                    "x-displayname": "County",
                    "x-ves-example": "Santa Clara"
                },
                "phone_number": {
                    "type": "string",
                    "description": " phone number of the contact\n\nExample: - \"+11234567890\"-",
                    "title": "phone_number",
                    "x-displayname": "Phone Number",
                    "x-ves-example": "+11234567890"
                },
                "state": {
                    "type": "string",
                    "description": " state (optional, for countries where they have states)\n\nExample: - \"California\"-",
                    "title": "state",
                    "x-displayname": "State",
                    "x-ves-example": "California"
                },
                "state_code": {
                    "type": "string",
                    "description": " state code (optional, for countries where they have states)\n\nExample: - \"CA\"-",
                    "title": "state code",
                    "x-displayname": "State Code",
                    "x-ves-example": "CA"
                },
                "zip_code": {
                    "type": "string",
                    "description": " zip or postal code\n\nExample: - \"95054\"-",
                    "title": "zip_code",
                    "x-displayname": "ZIP code",
                    "x-ves-example": "95054"
                }
            }
        },
        "signupUserMeta": {
            "type": "object",
            "title": "User Meta",
            "x-displayname": "User Meta",
            "x-ves-proto-message": "ves.io.schema.signup.UserMeta",
            "properties": {
                "contact_number": {
                    "type": "string",
                    "description": " contact number of the user\n\nExample: - \"+14084004001\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n",
                    "title": "Contact Number",
                    "maxLength": 256,
                    "x-displayname": "Contact Number",
                    "x-ves-example": "+14084004001",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256"
                    }
                },
                "email": {
                    "type": "string",
                    "description": " email of the user\n\nExample: - \"jane.doe@gmail.com\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n",
                    "title": "E-mail",
                    "maxLength": 256,
                    "x-displayname": "E-mail",
                    "x-ves-example": "jane.doe@gmail.com",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256"
                    }
                },
                "first_name": {
                    "type": "string",
                    "description": " first name of the user\n\nExample: - \"Jane\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n",
                    "title": "First Name",
                    "maxLength": 256,
                    "x-displayname": "First Name",
                    "x-ves-example": "Jane",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256"
                    }
                },
                "last_name": {
                    "type": "string",
                    "description": " last name of the user\n\nExample: - \"Doe\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n",
                    "title": "Last Name",
                    "maxLength": 256,
                    "x-displayname": "Last Name",
                    "x-ves-example": "Doe",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256"
                    }
                }
            }
        }
    },
    "x-displayname": "Onboard AWS SaaS PAYG account",
    "x-ves-proto-file": "ves.io/schema/marketplace/aws_account/public_onboard_customeapi.proto"
}`
