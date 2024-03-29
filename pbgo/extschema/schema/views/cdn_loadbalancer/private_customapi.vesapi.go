// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cdn_loadbalancer

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

// Create PrivateCustomAPI GRPC Client satisfying server.CustomClient
type PrivateCustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient PrivateCustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *PrivateCustomAPIGrpcClient) doRPCUpdateServiceDomains(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &UpdateServiceDomainsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.cdn_loadbalancer.UpdateServiceDomainsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.UpdateServiceDomains(ctx, req, opts...)
	return rsp, err
}

func (c *PrivateCustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewPrivateCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &PrivateCustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewPrivateCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["UpdateServiceDomains"] = ccl.doRPCUpdateServiceDomains

	ccl.rpcFns = rpcFns

	return ccl
}

// Create PrivateCustomAPI REST Client satisfying server.CustomClient
type PrivateCustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *PrivateCustomAPIRestClient) doRPCUpdateServiceDomains(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &UpdateServiceDomainsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.cdn_loadbalancer.UpdateServiceDomainsRequest: %s", yamlReq, err)
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
		for _, item := range req.ServiceDomains {
			q.Add("service_domains", fmt.Sprintf("%v", item))
		}
		q.Add("tenant", fmt.Sprintf("%v", req.Tenant))

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
	pbRsp := &UpdateServiceDomainsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.views.cdn_loadbalancer.UpdateServiceDomainsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *PrivateCustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewPrivateCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &PrivateCustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["UpdateServiceDomains"] = ccl.doRPCUpdateServiceDomains

	ccl.rpcFns = rpcFns

	return ccl
}

// Create privateCustomAPIInprocClient

// INPROC Client (satisfying PrivateCustomAPIClient interface)
type privateCustomAPIInprocClient struct {
	PrivateCustomAPIServer
}

func (c *privateCustomAPIInprocClient) UpdateServiceDomains(ctx context.Context, in *UpdateServiceDomainsRequest, opts ...grpc.CallOption) (*UpdateServiceDomainsResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI.UpdateServiceDomains")
	return c.PrivateCustomAPIServer.UpdateServiceDomains(ctx, in)
}

func NewPrivateCustomAPIInprocClient(svc svcfw.Service) PrivateCustomAPIClient {
	return &privateCustomAPIInprocClient{PrivateCustomAPIServer: NewPrivateCustomAPIServer(svc)}
}

// RegisterGwPrivateCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwPrivateCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterPrivateCustomAPIHandlerClient(ctx, mux, NewPrivateCustomAPIInprocClient(s))
}

// Create privateCustomAPISrv

// SERVER (satisfying PrivateCustomAPIServer interface)
type privateCustomAPISrv struct {
	svc svcfw.Service
}

func (s *privateCustomAPISrv) UpdateServiceDomains(ctx context.Context, in *UpdateServiceDomainsRequest) (*UpdateServiceDomainsResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI")
	cah, ok := ah.(PrivateCustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *PrivateCustomAPIServer", ah)
	}

	var (
		rsp *UpdateServiceDomainsResponse
		err error
	)

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI.UpdateServiceDomains"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.UpdateServiceDomains(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	return rsp, nil
}

func NewPrivateCustomAPIServer(svc svcfw.Service) PrivateCustomAPIServer {
	return &privateCustomAPISrv{svc: svc}
}

var PrivateCustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "CDN",
        "description": "CDN package",
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
        "/private/custom/namespaces/{namespace}/cdn_loadbalancer/{name}/update_service_domains": {
            "post": {
                "summary": "UpdateServiceDomains",
                "description": "Private API to update service domains from service.",
                "operationId": "ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI.UpdateServiceDomains",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/cdn_loadbalancerUpdateServiceDomainsResponse"
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
                        "description": "Namespace\n\nx-example: \"default\"\nx-required\nNamespace scope of the request",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-example: \"default\"\nx-required\nName of the CDN loadbalancer",
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
                            "$ref": "#/definitions/cdn_loadbalancerUpdateServiceDomainsRequest"
                        }
                    }
                ],
                "tags": [
                    "PrivateCustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-cdn_loadbalancer-privatecustomapi-updateservicedomains"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI.UpdateServiceDomains"
            },
            "x-displayname": "Private Custom API for CDN LoadBalancer",
            "x-ves-proto-service": "ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        },
        "/ves.io.schema/introspect/write/custom/namespaces/{namespace}/cdn_loadbalancer/{name}/update_service_domains": {
            "post": {
                "summary": "UpdateServiceDomains",
                "description": "Private API to update service domains from service.",
                "operationId": "ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI.UpdateServiceDomains",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/cdn_loadbalancerUpdateServiceDomainsResponse"
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
                        "description": "Namespace\n\nx-example: \"default\"\nx-required\nNamespace scope of the request",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-example: \"default\"\nx-required\nName of the CDN loadbalancer",
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
                            "$ref": "#/definitions/cdn_loadbalancerUpdateServiceDomainsRequest"
                        }
                    }
                ],
                "tags": [
                    "PrivateCustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-cdn_loadbalancer-privatecustomapi-updateservicedomains"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI.UpdateServiceDomains"
            },
            "x-displayname": "Private Custom API for CDN LoadBalancer",
            "x-ves-proto-service": "ves.io.schema.views.cdn_loadbalancer.PrivateCustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        }
    },
    "definitions": {
        "cdn_loadbalancerUpdateServiceDomainsRequest": {
            "type": "object",
            "description": "UpdateServiceDomainsRequest",
            "title": "UpdateServiceDomainsRequest",
            "x-displayname": "UpdateServiceDomainsRequest",
            "x-ves-proto-message": "ves.io.schema.views.cdn_loadbalancer.UpdateServiceDomainsRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Name of the CDN loadbalancer\n\nExample: - \"default\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "default",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace scope of the request\n\nExample: - \"default\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "service_domains": {
                    "type": "array",
                    "description": " CNAME provided from service per domain ",
                    "title": "Service Domains",
                    "items": {
                        "$ref": "#/definitions/virtual_hostServiceDomain"
                    },
                    "x-displayname": "Service Domains"
                },
                "tenant": {
                    "type": "string",
                    "description": " Tenant name \n\nExample: - \"default\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "Tenant",
                    "x-displayname": "Tenant",
                    "x-ves-example": "default",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                }
            }
        },
        "cdn_loadbalancerUpdateServiceDomainsResponse": {
            "type": "object",
            "description": "UpdateServiceDomainsResponse",
            "title": "UpdateServiceDomainsResponse",
            "x-displayname": "UpdateServiceDomainsResponse",
            "x-ves-proto-message": "ves.io.schema.views.cdn_loadbalancer.UpdateServiceDomainsResponse"
        },
        "virtual_hostServiceDomain": {
            "type": "object",
            "x-ves-proto-message": "ves.io.schema.virtual_host.ServiceDomain",
            "properties": {
                "domain": {
                    "type": "string",
                    "description": " Domain Name\n\nExample: - \"cdn.acmecorp.com\"-",
                    "title": "Domain Name",
                    "x-displayname": "Domain Name",
                    "x-ves-example": "cdn.acmecorp.com"
                },
                "service_domain": {
                    "type": "string",
                    "description": " Service Domain\n\nExample: - \"ves-io-cdn-cdn-acmecorp-com.demo1.ac.vh.volterra.us\"-",
                    "title": "Service Domain",
                    "x-displayname": "Service Domain",
                    "x-ves-example": "ves-io-cdn-cdn-acmecorp-com.demo1.ac.vh.volterra.us"
                }
            }
        }
    },
    "x-displayname": "CDN Loadbalancer",
    "x-ves-proto-file": "ves.io/schema/views/cdn_loadbalancer/private_customapi.proto"
}`
