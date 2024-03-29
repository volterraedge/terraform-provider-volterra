// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package api_definition

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

// Create PrivateConfigCustomAPI GRPC Client satisfying server.CustomClient
type PrivateConfigCustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient PrivateConfigCustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *PrivateConfigCustomAPIGrpcClient) doRPCUpdateAPIInventoryOpenAPISpecs(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &UpdateAPIInventoryOpenAPISpecsReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.api_definition.UpdateAPIInventoryOpenAPISpecsReq", yamlReq)
	}
	rsp, err := c.grpcClient.UpdateAPIInventoryOpenAPISpecs(ctx, req, opts...)
	return rsp, err
}

func (c *PrivateConfigCustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewPrivateConfigCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &PrivateConfigCustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewPrivateConfigCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["UpdateAPIInventoryOpenAPISpecs"] = ccl.doRPCUpdateAPIInventoryOpenAPISpecs

	ccl.rpcFns = rpcFns

	return ccl
}

// Create PrivateConfigCustomAPI REST Client satisfying server.CustomClient
type PrivateConfigCustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *PrivateConfigCustomAPIRestClient) doRPCUpdateAPIInventoryOpenAPISpecs(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &UpdateAPIInventoryOpenAPISpecsReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.api_definition.UpdateAPIInventoryOpenAPISpecsReq: %s", yamlReq, err)
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
		for _, item := range req.ApiInventoryOpenapiSpecs {
			q.Add("api_inventory_openapi_specs", fmt.Sprintf("%v", item))
		}
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
	pbRsp := &UpdateAPIInventoryOpenAPISpecsResp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.views.api_definition.UpdateAPIInventoryOpenAPISpecsResp", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *PrivateConfigCustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewPrivateConfigCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &PrivateConfigCustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["UpdateAPIInventoryOpenAPISpecs"] = ccl.doRPCUpdateAPIInventoryOpenAPISpecs

	ccl.rpcFns = rpcFns

	return ccl
}

// Create privateConfigCustomAPIInprocClient

// INPROC Client (satisfying PrivateConfigCustomAPIClient interface)
type privateConfigCustomAPIInprocClient struct {
	PrivateConfigCustomAPIServer
}

func (c *privateConfigCustomAPIInprocClient) UpdateAPIInventoryOpenAPISpecs(ctx context.Context, in *UpdateAPIInventoryOpenAPISpecsReq, opts ...grpc.CallOption) (*UpdateAPIInventoryOpenAPISpecsResp, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.views.api_definition.PrivateConfigCustomAPI.UpdateAPIInventoryOpenAPISpecs")
	return c.PrivateConfigCustomAPIServer.UpdateAPIInventoryOpenAPISpecs(ctx, in)
}

func NewPrivateConfigCustomAPIInprocClient(svc svcfw.Service) PrivateConfigCustomAPIClient {
	return &privateConfigCustomAPIInprocClient{PrivateConfigCustomAPIServer: NewPrivateConfigCustomAPIServer(svc)}
}

// RegisterGwPrivateConfigCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwPrivateConfigCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterPrivateConfigCustomAPIHandlerClient(ctx, mux, NewPrivateConfigCustomAPIInprocClient(s))
}

// Create privateConfigCustomAPISrv

// SERVER (satisfying PrivateConfigCustomAPIServer interface)
type privateConfigCustomAPISrv struct {
	svc svcfw.Service
}

func (s *privateConfigCustomAPISrv) UpdateAPIInventoryOpenAPISpecs(ctx context.Context, in *UpdateAPIInventoryOpenAPISpecsReq) (*UpdateAPIInventoryOpenAPISpecsResp, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.views.api_definition.PrivateConfigCustomAPI")
	cah, ok := ah.(PrivateConfigCustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *PrivateConfigCustomAPIServer", ah)
	}

	var (
		rsp *UpdateAPIInventoryOpenAPISpecsResp
		err error
	)

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.views.api_definition.PrivateConfigCustomAPI.UpdateAPIInventoryOpenAPISpecs"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.UpdateAPIInventoryOpenAPISpecs(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	return rsp, nil
}

func NewPrivateConfigCustomAPIServer(svc svcfw.Service) PrivateConfigCustomAPIServer {
	return &privateConfigCustomAPISrv{svc: svc}
}

var PrivateConfigCustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "api_definition object",
        "description": "The api_definition construct provides a mechanism to create api_groups based on referred OpenAPI specs.\nDefault api_groups, which are built automatically, include a group containing all the operations specified in swaggers; a group defining all possible requests for the given base urls.\nIn addition, we create default api-groups by a predefined OpenAPI extension, for example  x-volterra-apigroup. http_loadbalancer can refer api_definition object and create access policy rules based on its api-groups.",
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
        "/private/namespaces/{namespace}/api_definitions/{name}/update_inventory_specs": {
            "post": {
                "summary": "Update API Inventory OpenAPI Specs",
                "description": "Update the API Definition's API Inventory Openapi Specs with the provided files.",
                "operationId": "ves.io.schema.views.api_definition.PrivateConfigCustomAPI.UpdateAPIInventoryOpenAPISpecs",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/api_definitionUpdateAPIInventoryOpenAPISpecsResp"
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
                        "description": "Namespace\n\nx-example: \"shared\"\nThe namespace of the API Definition for the current request",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-example: \"name\"\nThe name of the API Definition for the current request",
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
                            "$ref": "#/definitions/api_definitionUpdateAPIInventoryOpenAPISpecsReq"
                        }
                    }
                ],
                "tags": [
                    "PrivateConfigCustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-api_definition-privateconfigcustomapi-updateapiinventoryopenapispecs"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.api_definition.PrivateConfigCustomAPI.UpdateAPIInventoryOpenAPISpecs"
            },
            "x-displayname": "Private API Definition Custom API",
            "x-ves-proto-service": "ves.io.schema.views.api_definition.PrivateConfigCustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        }
    },
    "definitions": {
        "api_definitionUpdateAPIInventoryOpenAPISpecsReq": {
            "type": "object",
            "description": "Request shape for Update API Inventory OpenAPI Spec",
            "title": "Update API Inventory OpenAPI Spec Request",
            "x-displayname": "Update API Inventory OpenAPI Spec",
            "x-ves-proto-message": "ves.io.schema.views.api_definition.UpdateAPIInventoryOpenAPISpecsReq",
            "properties": {
                "api_inventory_openapi_specs": {
                    "type": "array",
                    "description": " A stored object link to internally generated OpenAPI specification file.\n\nValidation Rules:\n  ves.io.schema.rules.repeated.max_items: 10\n  ves.io.schema.rules.repeated.unique: true\n",
                    "title": "api_inventory_openapi_spec",
                    "maxItems": 10,
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "API Inventory OpenAPI specification files",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.repeated.max_items": "10",
                        "ves.io.schema.rules.repeated.unique": "true"
                    }
                },
                "name": {
                    "type": "string",
                    "description": " The name of the API Definition for the current request\n\nExample: - \"name\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "name"
                },
                "namespace": {
                    "type": "string",
                    "description": " The namespace of the API Definition for the current request\n\nExample: - \"shared\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "shared"
                }
            }
        },
        "api_definitionUpdateAPIInventoryOpenAPISpecsResp": {
            "type": "object",
            "description": "Response shape for Update API Inventory OpenAPI Specs Schema",
            "title": "Update API Inventory OpenAPI Specs Response",
            "x-displayname": "Update API Inventory OpenAPI Specs Response",
            "x-ves-proto-message": "ves.io.schema.views.api_definition.UpdateAPIInventoryOpenAPISpecsResp"
        }
    },
    "x-displayname": "API Definition",
    "x-ves-proto-file": "ves.io/schema/views/api_definition/private_config_customapi.proto"
}`
