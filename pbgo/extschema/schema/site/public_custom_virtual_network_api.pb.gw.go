// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: ves.io/schema/site/public_custom_virtual_network_api.proto

/*
Package site is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package site

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_CustomVirtualNetworkListAPI_GlobalNetworkList_0(ctx context.Context, marshaler runtime.Marshaler, client CustomVirtualNetworkListAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GlobalNetworkListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "namespace", err)
	}

	val, ok = pathParams["site"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "site")
	}

	protoReq.Site, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "site", err)
	}

	msg, err := client.GlobalNetworkList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CustomVirtualNetworkListAPI_GlobalNetworkList_0(ctx context.Context, marshaler runtime.Marshaler, server CustomVirtualNetworkListAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GlobalNetworkListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "namespace", err)
	}

	val, ok = pathParams["site"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "site")
	}

	protoReq.Site, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "site", err)
	}

	msg, err := server.GlobalNetworkList(ctx, &protoReq)
	return msg, metadata, err

}

func request_CustomVirtualNetworkListAPI_SegmentList_0(ctx context.Context, marshaler runtime.Marshaler, client CustomVirtualNetworkListAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SegmentListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "namespace", err)
	}

	val, ok = pathParams["site"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "site")
	}

	protoReq.Site, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "site", err)
	}

	msg, err := client.SegmentList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CustomVirtualNetworkListAPI_SegmentList_0(ctx context.Context, marshaler runtime.Marshaler, server CustomVirtualNetworkListAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SegmentListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "namespace", err)
	}

	val, ok = pathParams["site"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "site")
	}

	protoReq.Site, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "site", err)
	}

	msg, err := server.SegmentList(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCustomVirtualNetworkListAPIHandlerServer registers the http handlers for service CustomVirtualNetworkListAPI to "mux".
// UnaryRPC     :call CustomVirtualNetworkListAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterCustomVirtualNetworkListAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CustomVirtualNetworkListAPIServer) error {

	mux.Handle("GET", pattern_CustomVirtualNetworkListAPI_GlobalNetworkList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CustomVirtualNetworkListAPI_GlobalNetworkList_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomVirtualNetworkListAPI_GlobalNetworkList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CustomVirtualNetworkListAPI_SegmentList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CustomVirtualNetworkListAPI_SegmentList_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomVirtualNetworkListAPI_SegmentList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCustomVirtualNetworkListAPIHandlerFromEndpoint is same as RegisterCustomVirtualNetworkListAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCustomVirtualNetworkListAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCustomVirtualNetworkListAPIHandler(ctx, mux, conn)
}

// RegisterCustomVirtualNetworkListAPIHandler registers the http handlers for service CustomVirtualNetworkListAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCustomVirtualNetworkListAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCustomVirtualNetworkListAPIHandlerClient(ctx, mux, NewCustomVirtualNetworkListAPIClient(conn))
}

// RegisterCustomVirtualNetworkListAPIHandlerClient registers the http handlers for service CustomVirtualNetworkListAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CustomVirtualNetworkListAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CustomVirtualNetworkListAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CustomVirtualNetworkListAPIClient" to call the correct interceptors.
func RegisterCustomVirtualNetworkListAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CustomVirtualNetworkListAPIClient) error {

	mux.Handle("GET", pattern_CustomVirtualNetworkListAPI_GlobalNetworkList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CustomVirtualNetworkListAPI_GlobalNetworkList_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomVirtualNetworkListAPI_GlobalNetworkList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_CustomVirtualNetworkListAPI_SegmentList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CustomVirtualNetworkListAPI_SegmentList_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomVirtualNetworkListAPI_SegmentList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CustomVirtualNetworkListAPI_GlobalNetworkList_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4, 2, 5}, []string{"public", "namespaces", "namespace", "sites", "site", "global_networks"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_CustomVirtualNetworkListAPI_SegmentList_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4, 2, 5}, []string{"public", "namespaces", "namespace", "sites", "site", "segments"}, "", runtime.AssumeColonVerbOpt(false)))
)

var (
	forward_CustomVirtualNetworkListAPI_GlobalNetworkList_0 = runtime.ForwardResponseMessage

	forward_CustomVirtualNetworkListAPI_SegmentList_0 = runtime.ForwardResponseMessage
)
