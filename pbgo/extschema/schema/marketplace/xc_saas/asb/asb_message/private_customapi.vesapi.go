// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package asb_message

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

func (c *CustomPrivateAPIGrpcClient) doRPCRegisterXCSaaS(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &RegisterXCSaaSRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.xc_saas.asb.asb_message.RegisterXCSaaSRequest", yamlReq)
	}
	rsp, err := c.grpcClient.RegisterXCSaaS(ctx, req, opts...)
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
	rpcFns["RegisterXCSaaS"] = ccl.doRPCRegisterXCSaaS

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

func (c *CustomPrivateAPIRestClient) doRPCRegisterXCSaaS(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &RegisterXCSaaSRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.marketplace.xc_saas.asb.asb_message.RegisterXCSaaSRequest: %s", yamlReq, err)
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
		q.Add("document_type", fmt.Sprintf("%v", req.DocumentType))
		q.Add("document_version", fmt.Sprintf("%v", req.DocumentVersion))
		q.Add("order", fmt.Sprintf("%v", req.Order))
		q.Add("service_order_type", fmt.Sprintf("%v", req.ServiceOrderType))

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
	pbRsp := &RegisterXCSaaSResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.marketplace.xc_saas.asb.asb_message.RegisterXCSaaSResponse", body)

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
	rpcFns["RegisterXCSaaS"] = ccl.doRPCRegisterXCSaaS

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customPrivateAPIInprocClient

// INPROC Client (satisfying CustomPrivateAPIClient interface)
type customPrivateAPIInprocClient struct {
	CustomPrivateAPIServer
}

func (c *customPrivateAPIInprocClient) RegisterXCSaaS(ctx context.Context, in *RegisterXCSaaSRequest, opts ...grpc.CallOption) (*RegisterXCSaaSResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.marketplace.xc_saas.asb.asb_message.CustomPrivateAPI.RegisterXCSaaS")
	return c.CustomPrivateAPIServer.RegisterXCSaaS(ctx, in)
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

func (s *customPrivateAPISrv) RegisterXCSaaS(ctx context.Context, in *RegisterXCSaaSRequest) (*RegisterXCSaaSResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.marketplace.xc_saas.asb.asb_message.CustomPrivateAPI")
	cah, ok := ah.(CustomPrivateAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomPrivateAPIServer", ah)
	}

	var (
		rsp *RegisterXCSaaSResponse
		err error
	)

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.marketplace.xc_saas.asb.asb_message.CustomPrivateAPI.RegisterXCSaaS"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.RegisterXCSaaS(ctx, in)
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
        "title": "ves.io/schema/marketplace/xc_saas/asb/asb_message/private_customapi.proto",
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
        "/ves.io.schema.marketplace.xc_saas.asb.asb_message/register": {
            "post": {
                "summary": "Register XC SaaS",
                "description": "Use this API to register up Azure Service Bus (ASB) provisioning message and further process tenant sign-up for F5XC service.",
                "operationId": "ves.io.schema.marketplace.xc_saas.asb.asb_message.CustomPrivateAPI.RegisterXCSaaS",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/asb_messageRegisterXCSaaSResponse"
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
                            "$ref": "#/definitions/asb_messageRegisterXCSaaSRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomPrivateAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-marketplace-xc_saas-asb-asb_message-customprivateapi-registerxcsaas"
                },
                "x-ves-proto-rpc": "ves.io.schema.marketplace.xc_saas.asb.asb_message.CustomPrivateAPI.RegisterXCSaaS"
            },
            "x-displayname": "ASB Message Custom Private API",
            "x-ves-proto-service": "ves.io.schema.marketplace.xc_saas.asb.asb_message.CustomPrivateAPI",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        }
    },
    "definitions": {
        "asb_messageActionType": {
            "type": "string",
            "description": "Type of action to perform\n\nProvision the requested service for tenant creation\nProvision the requested service for tenant renewal",
            "title": "ActionType",
            "enum": [
                "ENABLE",
                "RENEWAL"
            ],
            "default": "ENABLE",
            "x-displayname": "Action Type",
            "x-ves-proto-enum": "ves.io.schema.marketplace.xc_saas.asb.asb_message.ActionType"
        },
        "asb_messageAdmin": {
            "type": "object",
            "description": "Tenant administrator metadata",
            "title": "Admin",
            "x-displayname": "Admin",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.asb.asb_message.Admin",
            "properties": {
                "email": {
                    "type": "string",
                    "description": " Tenant admin email\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_len: 256\n  ves.io.schema.rules.string.min_len: 3\n",
                    "title": "Email",
                    "minLength": 3,
                    "maxLength": 256,
                    "x-displayname": "Email",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_len": "256",
                        "ves.io.schema.rules.string.min_len": "3"
                    }
                },
                "first_name": {
                    "type": "string",
                    "description": " Tenant admin first name\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_len: 256\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "FirstName",
                    "minLength": 1,
                    "maxLength": 256,
                    "x-displayname": "First Name",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_len": "256",
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                },
                "last_name": {
                    "type": "string",
                    "description": " Tenant admin last name\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_len: 256\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "LastName",
                    "minLength": 1,
                    "maxLength": 256,
                    "x-displayname": "Last Name",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_len": "256",
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                }
            }
        },
        "asb_messageCustomerMetadata": {
            "type": "object",
            "description": "Customer related entitlement data",
            "title": "CustomerMetadata",
            "x-displayname": "Customer Metadata",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.asb.asb_message.CustomerMetadata",
            "properties": {
                "address_one": {
                    "type": "string",
                    "description": " Mailing address one of the company\n\nValidation Rules:\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "AddressOne",
                    "minLength": 1,
                    "x-displayname": "Address One",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                },
                "address_two": {
                    "type": "string",
                    "description": " Mailing address two of the company",
                    "title": "AddressTwo",
                    "x-displayname": "Address Two"
                },
                "city": {
                    "type": "string",
                    "description": " City / town of the company\n\nValidation Rules:\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "City",
                    "minLength": 1,
                    "x-displayname": "City",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                },
                "country": {
                    "type": "string",
                    "description": " Country of the company\n\nValidation Rules:\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "Country",
                    "minLength": 1,
                    "x-displayname": "Country",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                },
                "name": {
                    "type": "string",
                    "description": " Name of the company\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "Name",
                    "minLength": 1,
                    "maxLength": 256,
                    "x-displayname": "Name",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256",
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                },
                "zip_code": {
                    "type": "string",
                    "description": " Zip code of the company\n\nValidation Rules:\n  ves.io.schema.rules.string.min_len: 5\n",
                    "title": "ZipCode",
                    "minLength": 5,
                    "x-displayname": "ZipCode",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.min_len": "5"
                    }
                }
            }
        },
        "asb_messageDocumentType": {
            "type": "string",
            "description": "Identifies the type of message document.\n\nService order generated by TEEM",
            "title": "DocumentType",
            "enum": [
                "TEEM_SERVICE_ORDER"
            ],
            "default": "TEEM_SERVICE_ORDER",
            "x-displayname": "Document Type",
            "x-ves-proto-enum": "ves.io.schema.marketplace.xc_saas.asb.asb_message.DocumentType"
        },
        "asb_messageEntitlement": {
            "type": "object",
            "description": "Entitlement to be applied for the basic provisioning model",
            "title": "Entitlement",
            "x-displayname": "Entitlement",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.asb.asb_message.Entitlement",
            "properties": {
                "customer_metadata": {
                    "description": " Customer related entitlement data",
                    "title": "CustomerMetadata",
                    "$ref": "#/definitions/asb_messageCustomerMetadata",
                    "x-displayname": "Customer Metadata"
                },
                "entitlement_metadata": {
                    "description": " Entitlement data\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "EntitlementMetadata",
                    "$ref": "#/definitions/asb_messageEntitlementMetadata",
                    "x-displayname": "Entitlement Metadata",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "tenant_metadata": {
                    "description": " Tenant related entitlement data\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "TenantMetadata",
                    "$ref": "#/definitions/asb_messageTenantMetadata",
                    "x-displayname": "Tenant Metadata",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                }
            }
        },
        "asb_messageEntitlementMetadata": {
            "type": "object",
            "description": "Entitlement data",
            "title": "EntitlementMetadata",
            "x-displayname": "Entitlement Metadata",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.asb.asb_message.EntitlementMetadata",
            "properties": {
                "entitled_skus": {
                    "type": "array",
                    "description": " SKU information that is used mostly for reporting purposes.\n\nExample: - \"['F5-V-O-ALL-BASE-PK-B','F5-XC-O-ALL-BOT-STD-B','F5-V-O-ADN-MSH-API-B','F5-V-O-ADN-MSH-RLM-B']\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.repeated.min_items: 1\n",
                    "title": "EntitledSKUs",
                    "minItems": 1,
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Entitled SKUs",
                    "x-ves-example": "['F5-V-O-ALL-BASE-PK-B','F5-XC-O-ALL-BOT-STD-B','F5-V-O-ADN-MSH-API-B','F5-V-O-ADN-MSH-RLM-B']",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.repeated.min_items": "1"
                    }
                }
            }
        },
        "asb_messageOrder": {
            "type": "object",
            "description": "Information about the specific service order",
            "title": "Order",
            "x-displayname": "Order",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.asb.asb_message.Order",
            "properties": {
                "account_id": {
                    "type": "string",
                    "description": " The salesforce customer account associated with the service order\n\nExample: - \"SFA-1478257\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_len: 256\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "AccountID",
                    "minLength": 1,
                    "maxLength": 256,
                    "x-displayname": "AccountID",
                    "x-ves-example": "SFA-1478257",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_len": "256",
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                },
                "action_type": {
                    "description": " Type of action to perform\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.enum.defined_only: true\n  ves.io.schema.rules.message.required: true\n",
                    "title": "ActionType",
                    "$ref": "#/definitions/asb_messageActionType",
                    "x-displayname": "Action Type",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.enum.defined_only": "true",
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "entitlement": {
                    "description": " Entitlement to be applied for the basic provisioning model\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "Entitlement",
                    "$ref": "#/definitions/asb_messageEntitlement",
                    "x-displayname": "Entitlement",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "entitlement_id": {
                    "type": "string",
                    "description": " The EntitlementID of the purchased service. The EntitlementID is unique to the AccountID, SubscriptionID.\n In the case of a subscription renewal, the EntitlementID persists although the SubscriptionID changes.\n\nExample: - \"e5712007-0560-4fcc-b8c9-f4ffbeaf3e4e\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_len: 256\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "EntitlementID",
                    "minLength": 1,
                    "maxLength": 256,
                    "x-displayname": "EntitlementID",
                    "x-ves-example": "e5712007-0560-4fcc-b8c9-f4ffbeaf3e4e",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_len": "256",
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                },
                "order_type": {
                    "description": " Order's buying program type",
                    "title": "Order",
                    "$ref": "#/definitions/asb_messageOrderType",
                    "x-displayname": "Order"
                },
                "service": {
                    "description": " Identifies the service that should handle provisioning of the order\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.enum.defined_only: true\n  ves.io.schema.rules.message.required: true\n",
                    "title": "Service",
                    "$ref": "#/definitions/asbasb_messageServiceType",
                    "x-displayname": "Service",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.enum.defined_only": "true",
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "subscription_id": {
                    "type": "string",
                    "description": " The latest zuora subscriptionID associated with the service that triggered the service order\n\nExample: - \"A-S00012023\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_len: 256\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "SubscriptionID",
                    "minLength": 1,
                    "maxLength": 256,
                    "x-displayname": "SubscriptionID",
                    "x-ves-example": "A-S00012023",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_len": "256",
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                }
            }
        },
        "asb_messageOrderType": {
            "type": "string",
            "description": "Order's buying program type\n\nOrder's buying program is a term subscription\nOrder's buying program is a flexible consumption program",
            "title": "OrderType",
            "enum": [
                "PAID",
                "FCP"
            ],
            "default": "PAID",
            "x-displayname": "Order Type",
            "x-ves-proto-enum": "ves.io.schema.marketplace.xc_saas.asb.asb_message.OrderType"
        },
        "asb_messageRegisterXCSaaSRequest": {
            "type": "object",
            "description": "The XC SaaS registration request holds the data required to register and further process signup for F5XC service.\nIf any additional data sent via the azure service bus message(https://docs.f5net.com/pages/viewpage.action?spaceKey=ITCL2PROG\u0026title=Service+Order)\nhas a specific use case in the registration/signup flow, it will be considered as a new requirement.",
            "title": "RegisterXCSaaSRequest",
            "x-displayname": "XC SaaS Registration Request",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.asb.asb_message.RegisterXCSaaSRequest",
            "properties": {
                "document_type": {
                    "description": " Identifies the type of message document\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.enum.defined_only: true\n  ves.io.schema.rules.message.required: true\n",
                    "title": "DocumentType",
                    "$ref": "#/definitions/asb_messageDocumentType",
                    "x-displayname": "Document Type",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.enum.defined_only": "true",
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "document_version": {
                    "type": "string",
                    "description": " The version of the message document format",
                    "title": "DocumentVersion",
                    "x-displayname": "Document Version"
                },
                "order": {
                    "description": " Information about the specific service order\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "Order",
                    "$ref": "#/definitions/asb_messageOrder",
                    "x-displayname": "Order",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "service_order_type": {
                    "description": " Type of order\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.enum.defined_only: true\n  ves.io.schema.rules.message.required: true\n",
                    "title": "ServiceOrderType",
                    "$ref": "#/definitions/asb_messageServiceOrderType",
                    "x-displayname": "Service Order Type",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.enum.defined_only": "true",
                        "ves.io.schema.rules.message.required": "true"
                    }
                }
            }
        },
        "asb_messageRegisterXCSaaSResponse": {
            "type": "object",
            "title": "RegisterXCSaaSResponse",
            "x-displayname": "XC SaaS Registration Response",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.asb.asb_message.RegisterXCSaaSResponse"
        },
        "asb_messageServiceOrderType": {
            "type": "string",
            "description": "Type of order\n\nThe order is related to service provisioning or de-provisioning\nTesting service orders triggered for testing purposes",
            "title": "ServiceOrderType",
            "enum": [
                "PROVISIONING",
                "TESTING"
            ],
            "default": "PROVISIONING",
            "x-displayname": "Service Order Type",
            "x-ves-proto-enum": "ves.io.schema.marketplace.xc_saas.asb.asb_message.ServiceOrderType"
        },
        "asb_messageTenantMetadata": {
            "type": "object",
            "description": "Tenant related entitlement data",
            "title": "TenantMetadata",
            "x-displayname": "Tenant Metadata",
            "x-ves-proto-message": "ves.io.schema.marketplace.xc_saas.asb.asb_message.TenantMetadata",
            "properties": {
                "admins": {
                    "type": "array",
                    "description": " Tenant administrator metadata\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.repeated.min_items: 1\n  ves.io.schema.rules.repeated.unique: true\n",
                    "title": "Admins",
                    "minItems": 1,
                    "items": {
                        "$ref": "#/definitions/asb_messageAdmin"
                    },
                    "x-displayname": "Admins",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.repeated.min_items": "1",
                        "ves.io.schema.rules.repeated.unique": "true"
                    }
                },
                "name": {
                    "type": "string",
                    "description": " Domain name of the account\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 17\n  ves.io.schema.rules.string.min_len: 1\n  ves.io.schema.rules.string.not_in: [\\\"ves\\\",\\\"volterra\\\",\\\"ves-io\\\",\\\"f5xc\\\"]\n  ves.io.schema.rules.string.ves_object_name: true\n",
                    "title": "Name",
                    "minLength": 1,
                    "maxLength": 17,
                    "x-displayname": "Name",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "17",
                        "ves.io.schema.rules.string.min_len": "1",
                        "ves.io.schema.rules.string.not_in": "[\\\"ves\\\",\\\"volterra\\\",\\\"ves-io\\\",\\\"f5xc\\\"]",
                        "ves.io.schema.rules.string.ves_object_name": "true"
                    }
                }
            }
        },
        "asbasb_messageServiceType": {
            "type": "string",
            "description": "Identifies the service that should handle provisioning of the order\n\nImplies that the provisioning is specific to XC service",
            "title": "ServiceType",
            "enum": [
                "XC"
            ],
            "default": "XC",
            "x-displayname": "Service Type",
            "x-ves-proto-enum": "ves.io.schema.marketplace.xc_saas.asb.asb_message.ServiceType"
        }
    },
    "x-displayname": "",
    "x-ves-proto-file": "ves.io/schema/marketplace/xc_saas/asb/asb_message/private_customapi.proto"
}`
