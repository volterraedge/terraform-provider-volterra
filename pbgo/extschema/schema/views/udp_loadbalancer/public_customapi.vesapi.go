// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package udp_loadbalancer

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

func (c *CustomAPIGrpcClient) doRPCGetDnsInfo(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &GetDnsInfoRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.udp_loadbalancer.GetDnsInfoRequest", yamlReq)
	}
	rsp, err := c.grpcClient.GetDnsInfo(ctx, req, opts...)
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
	rpcFns["GetDnsInfo"] = ccl.doRPCGetDnsInfo

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

func (c *CustomAPIRestClient) doRPCGetDnsInfo(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &GetDnsInfoRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.udp_loadbalancer.GetDnsInfoRequest: %s", yamlReq, err)
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
	pbRsp := &GetDnsInfoResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.views.udp_loadbalancer.GetDnsInfoResponse", body)

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
	rpcFns["GetDnsInfo"] = ccl.doRPCGetDnsInfo

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customAPIInprocClient

// INPROC Client (satisfying CustomAPIClient interface)
type customAPIInprocClient struct {
	CustomAPIServer
}

func (c *customAPIInprocClient) GetDnsInfo(ctx context.Context, in *GetDnsInfoRequest, opts ...grpc.CallOption) (*GetDnsInfoResponse, error) {
	ctx = server.ContextWithRpcFQN(ctx, "ves.io.schema.views.udp_loadbalancer.CustomAPI.GetDnsInfo")
	return c.CustomAPIServer.GetDnsInfo(ctx, in)
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

func (s *customAPISrv) GetDnsInfo(ctx context.Context, in *GetDnsInfoRequest) (*GetDnsInfoResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.views.udp_loadbalancer.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPIServer", ah)
	}

	var (
		rsp *GetDnsInfoResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.views.udp_loadbalancer.GetDnsInfoRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.GetDnsInfo' operation on 'udp_loadbalancer'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.views.udp_loadbalancer.CustomAPI.GetDnsInfo"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.GetDnsInfo(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.views.udp_loadbalancer.GetDnsInfoResponse", rsp)...)

	return rsp, nil
}

func NewCustomAPIServer(svc svcfw.Service) CustomAPIServer {
	return &customAPISrv{svc: svc}
}

var CustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "UDP Proxy",
        "description": "UDP Load Balancer view defines a required parameters that can be used in CRUD, to create and manage UDP Load Balancer.\nIt can be used to create UDP Load Balancer and UDP Load Balancer with SNI.\n\nView will create following child objects.\n\n* Virtual-host\n* default route\n* clusters\n* endpoints\n* advertise policy",
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
        "/public/namespaces/{namespace}/udp_loadbalancers/{name}/get-dns-info": {
            "get": {
                "summary": "Get DNS Info",
                "description": "GetDnsInfo is an API to get DNS information for a given UDP load balancer",
                "operationId": "ves.io.schema.views.udp_loadbalancer.CustomAPI.GetDnsInfo",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/udp_loadbalancerGetDnsInfoResponse"
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
                        "description": "Namespace\n\nx-example: \"value\"\nNamespace for the UDP load balancer",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-example: \"value\"\nName of the UDP load balancer",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Name"
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://docs.cloud.f5.com/docs-v2/platform/reference/api-ref/ves-io-schema-views-udp_loadbalancer-customapi-getdnsinfo"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.udp_loadbalancer.CustomAPI.GetDnsInfo"
            },
            "x-displayname": "UDP Load Balancer Custom API",
            "x-ves-proto-service": "ves.io.schema.views.udp_loadbalancer.CustomAPI",
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
        "schemavirtual_host_dns_infoGlobalSpecType": {
            "type": "object",
            "description": "Shape of the virtual host DNS info global specification",
            "title": "Global Specification for Virtual Host DNS info",
            "x-displayname": "Global Spec",
            "x-ves-proto-message": "ves.io.schema.virtual_host_dns_info.GlobalSpecType",
            "properties": {
                "dns_info": {
                    "type": "array",
                    "description": " DNS information for this virtual host",
                    "title": "DNS information",
                    "items": {
                        "$ref": "#/definitions/virtual_host_dns_infoDnsInfo"
                    },
                    "x-displayname": "DNS Information"
                },
                "host_name": {
                    "type": "string",
                    "description": " Host name to be used for the virtual host\n\nExample: - \"value\"-",
                    "title": "host name",
                    "x-displayname": "Host Name",
                    "x-ves-example": "value"
                },
                "virtual_host": {
                    "type": "array",
                    "description": " Reference to virtual host object to which this information applies",
                    "title": "virtual host",
                    "items": {
                        "$ref": "#/definitions/ioschemaObjectRefType"
                    },
                    "x-displayname": "Virtual Host"
                }
            }
        },
        "udp_loadbalancerGetDnsInfoResponse": {
            "type": "object",
            "description": "Response for get-dns-info API",
            "title": "GetDnsInfoResponse",
            "x-displayname": "Get DNS Info Response",
            "x-ves-proto-message": "ves.io.schema.views.udp_loadbalancer.GetDnsInfoResponse",
            "properties": {
                "dns_info": {
                    "description": " DNS information object for this UDP load balancer",
                    "title": "DNS information",
                    "$ref": "#/definitions/schemavirtual_host_dns_infoGlobalSpecType",
                    "x-displayname": "DNS information"
                }
            }
        },
        "virtual_host_dns_infoDnsInfo": {
            "type": "object",
            "description": "A message that contains DNS information for a given IP address",
            "title": "DNS information",
            "x-displayname": "DNS Information",
            "x-ves-proto-message": "ves.io.schema.virtual_host_dns_info.DnsInfo",
            "properties": {
                "ip_address": {
                    "type": "string",
                    "description": " IP address associated with virtual host\n\nValidation Rules:\n  ves.io.schema.rules.string.ip: true\n",
                    "title": "IP address",
                    "x-displayname": "IP Address",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.ip": "true"
                    }
                }
            }
        }
    },
    "x-displayname": "Configure UDP Load Balancer",
    "x-ves-proto-file": "ves.io/schema/views/udp_loadbalancer/public_customapi.proto"
}`
